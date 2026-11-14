import { zodResolver } from '@hookform/resolvers/zod';
import { ChevronDown, Database, ListFilter, Search, X } from 'lucide-react';
import { useEffect, useMemo, useState } from 'react';
import { useForm } from 'react-hook-form';
import { useDebouncedCallback } from 'use-debounce';
import { z } from 'zod';

import { Button } from '@/components/ui/button';
import { Empty, EmptyContent, EmptyDescription, EmptyHeader, EmptyMedia, EmptyTitle } from '@/components/ui/empty';
import { Form, FormControl, FormField } from '@/components/ui/form';
import { InputGroup, InputGroupAddon, InputGroupButton, InputGroupInput } from '@/components/ui/input-group';
import { useAutoScroll } from '@/hooks/use-auto-scroll';
import { useFlow } from '@/providers/flow-provider';

import FlowTasksDropdown from '../flow-tasks-dropdown';
import FlowVectorStore from './flow-vector-store';

const searchFormSchema = z.object({
    filter: z
        .object({
            subtaskIds: z.array(z.string()),
            taskIds: z.array(z.string()),
        })
        .optional(),
    search: z.string(),
});

function FlowVectorStores() {
    const { flowData, flowId } = useFlow();

    const logs = useMemo(() => flowData?.vectorStoreLogs ?? [], [flowData?.vectorStoreLogs]);
    const [debouncedSearchValue, setDebouncedSearchValue] = useState('');

    const { containerRef, endRef, hasNewMessages, isScrolledToBottom, scrollToEnd } = useAutoScroll(logs, flowId);

    const form = useForm<z.infer<typeof searchFormSchema>>({
        defaultValues: {
            filter: {
                subtaskIds: [],
                taskIds: [],
            },
            search: '',
        },
        resolver: zodResolver(searchFormSchema),
    });

    const searchValue = form.watch('search');
    const filter = form.watch('filter');

    const debouncedUpdateSearch = useDebouncedCallback((value: string) => {
        setDebouncedSearchValue(value);
    }, 500);

    useEffect(() => {
        debouncedUpdateSearch(searchValue);

        return () => {
            debouncedUpdateSearch.cancel();
        };
    }, [searchValue, debouncedUpdateSearch]);

    useEffect(() => {
        return () => {
            debouncedUpdateSearch.cancel();
        };
    }, [debouncedUpdateSearch]);

    useEffect(() => {
        form.reset({
            filter: {
                subtaskIds: [],
                taskIds: [],
            },
            search: '',
        });
        setDebouncedSearchValue('');
        debouncedUpdateSearch.cancel();
    }, [flowId, form, debouncedUpdateSearch]);

    const hasActiveFilters = useMemo(() => {
        const hasSearch = !!searchValue.trim();
        const hasTaskFilters = !!(filter?.taskIds?.length || filter?.subtaskIds?.length);

        return hasSearch || hasTaskFilters;
    }, [searchValue, filter]);

    const filteredLogs = useMemo(() => {
        const search = debouncedSearchValue.toLowerCase().trim();

        let filtered = logs || [];

        if (search) {
            filtered = filtered.filter((log) => {
                return (
                    log.query.toLowerCase().includes(search) ||
                    log.result?.toLowerCase().includes(search) ||
                    log.filter.toLowerCase().includes(search) ||
                    log.action.toLowerCase().includes(search) ||
                    log.executor.toLowerCase().includes(search) ||
                    log.initiator.toLowerCase().includes(search)
                );
            });
        }

        if (filter?.taskIds?.length || filter?.subtaskIds?.length) {
            const selectedTaskIds = new Set(filter.taskIds ?? []);
            const selectedSubtaskIds = new Set(filter.subtaskIds ?? []);

            filtered = filtered.filter((log) => {
                if (log.taskId && selectedTaskIds.has(log.taskId)) {
                    return true;
                }

                if (log.subtaskId && selectedSubtaskIds.has(log.subtaskId)) {
                    return true;
                }

                return false;
            });
        }

        return filtered;
    }, [logs, debouncedSearchValue, filter]);

    const hasLogs = filteredLogs && filteredLogs.length > 0;

    const handleResetFilters = () => {
        form.reset({
            filter: {
                subtaskIds: [],
                taskIds: [],
            },
            search: '',
        });
        setDebouncedSearchValue('');
        debouncedUpdateSearch.cancel();
    };

    return (
        <div className="flex h-full flex-col">
            <div className="bg-background sticky top-0 z-10 pb-4">
                <Form {...form}>
                    <div className="flex gap-2 p-px">
                        <FormField
                            control={form.control}
                            name="search"
                            render={({ field }) => (
                                <FormControl>
                                    <InputGroup className="flex-1">
                                        <InputGroupAddon>
                                            <Search />
                                        </InputGroupAddon>
                                        <InputGroupInput
                                            {...field}
                                            autoComplete="off"
                                            placeholder="搜尋向量儲存日誌…"
                                            type="text"
                                        />
                                        {field.value && (
                                            <InputGroupAddon align="inline-end">
                                                <InputGroupButton
                                                    aria-label="清除向量儲存搜尋"
                                                    onClick={() => {
                                                        form.reset({ search: '' });
                                                        setDebouncedSearchValue('');
                                                        debouncedUpdateSearch.cancel();
                                                    }}
                                                    type="button"
                                                >
                                                    <X />
                                                </InputGroupButton>
                                            </InputGroupAddon>
                                        )}
                                    </InputGroup>
                                </FormControl>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="filter"
                            render={({ field }) => (
                                <FormControl>
                                    <FlowTasksDropdown
                                        onChange={field.onChange}
                                        value={field.value}
                                    />
                                </FormControl>
                            )}
                        />
                    </div>
                </Form>
            </div>
            {hasLogs ? (
                <div className="relative flex-1 overflow-y-hidden">
                    <div
                        className="flex h-full flex-col gap-4 overflow-y-auto"
                        ref={containerRef}
                    >
                        {filteredLogs.map((log) => (
                            <FlowVectorStore
                                key={log.id}
                                log={log}
                                searchValue={debouncedSearchValue}
                            />
                        ))}
                        <div ref={endRef} />
                    </div>

                    {!isScrolledToBottom && (
                        <Button
                            aria-label="捲動到最新向量儲存日誌"
                            className="absolute right-4 bottom-4 z-10 shadow-md hover:shadow-lg"
                            onClick={() => scrollToEnd()}
                            size="icon-sm"
                            type="button"
                            variant="outline"
                        >
                            <ChevronDown />
                            {hasNewMessages && (
                                <span className="bg-primary absolute -top-1.5 -right-1.5 size-3 rounded-full" />
                            )}
                        </Button>
                    )}
                </div>
            ) : hasActiveFilters ? (
                <Empty>
                    <EmptyHeader>
                        <EmptyMedia variant="icon">
                            <ListFilter />
                        </EmptyMedia>
                        <EmptyTitle>找不到向量儲存日誌</EmptyTitle>
                        <EmptyDescription>請調整搜尋關鍵字或篩選條件</EmptyDescription>
                    </EmptyHeader>
                    <EmptyContent>
                        <Button
                            onClick={handleResetFilters}
                            variant="outline"
                        >
                            <X />
                            重設篩選條件
                        </Button>
                    </EmptyContent>
                </Empty>
            ) : (
                <Empty>
                    <EmptyHeader>
                        <EmptyMedia variant="icon">
                            <Database />
                        </EmptyMedia>
                        <EmptyTitle>沒有可用的向量儲存日誌</EmptyTitle>
                        <EmptyDescription>
                            代理使用知識庫後，向量儲存日誌會顯示在這裡
                        </EmptyDescription>
                    </EmptyHeader>
                </Empty>
            )}
        </div>
    );
}

export default FlowVectorStores;
