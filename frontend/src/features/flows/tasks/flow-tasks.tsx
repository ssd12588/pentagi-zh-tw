import { zodResolver } from '@hookform/resolvers/zod';
import { ChevronDown, ListTodo, Search, X } from 'lucide-react';
import { useEffect, useMemo, useState } from 'react';
import { useForm } from 'react-hook-form';
import { useDebouncedCallback } from 'use-debounce';
import { z } from 'zod';

import { Button } from '@/components/ui/button';
import { Empty, EmptyDescription, EmptyHeader, EmptyMedia, EmptyTitle } from '@/components/ui/empty';
import { Form, FormControl, FormField } from '@/components/ui/form';
import { InputGroup, InputGroupAddon, InputGroupButton, InputGroupInput } from '@/components/ui/input-group';
import { useAutoScroll } from '@/hooks/use-auto-scroll';
import { matchesSearchTerm } from '@/lib/text-filter';
import { useFlow } from '@/providers/flow-provider';

import FlowTask from './flow-task';

const searchFormSchema = z.object({
    search: z.string(),
});

function FlowTasks() {
    const { flowData, flowId } = useFlow();

    const tasks = useMemo(() => flowData?.tasks ?? [], [flowData?.tasks]);
    const [debouncedSearchValue, setDebouncedSearchValue] = useState('');

    const { containerRef, endRef, hasNewMessages, isScrolledToBottom, scrollToEnd } = useAutoScroll(tasks, flowId);

    const form = useForm<z.infer<typeof searchFormSchema>>({
        defaultValues: {
            search: '',
        },
        resolver: zodResolver(searchFormSchema),
    });

    const searchValue = form.watch('search');

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
        form.reset({ search: '' });
        setDebouncedSearchValue('');
        debouncedUpdateSearch.cancel();
    }, [flowId, form, debouncedUpdateSearch]);

    const filteredTasks = useMemo(() => {
        const search = debouncedSearchValue.toLowerCase().trim();

        if (!search || !tasks) {
            return tasks || [];
        }

        return tasks.filter((task) => {
            const taskMatches = matchesSearchTerm(task.title, search) || matchesSearchTerm(task.result, search);

            const subtaskMatches =
                task.subtasks?.some(
                    (subtask) =>
                        matchesSearchTerm(subtask.title, search) ||
                        matchesSearchTerm(subtask.description, search) ||
                        matchesSearchTerm(subtask.result, search),
                ) || false;

            return taskMatches || subtaskMatches;
        });
    }, [tasks, debouncedSearchValue]);

    const sortedTasks = [...(filteredTasks || [])].sort((a, b) => +a.id - +b.id);
    const hasTasks = filteredTasks && filteredTasks.length > 0;

    return (
        <div className="flex h-full flex-col">
            <div className="bg-background sticky top-0 z-10 pb-4">
                <Form {...form}>
                    <div className="p-px">
                        <FormField
                            control={form.control}
                            name="search"
                            render={({ field }) => (
                                <FormControl>
                                    <InputGroup>
                                        <InputGroupAddon>
                                            <Search />
                                        </InputGroupAddon>
                                        <InputGroupInput
                                            {...field}
                                            autoComplete="off"
                                            placeholder="搜尋任務與子任務…"
                                            type="text"
                                        />
                                        {field.value && (
                                            <InputGroupAddon align="inline-end">
                                                <InputGroupButton
                                                    aria-label="清除任務搜尋"
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
                    </div>
                </Form>
            </div>

            {hasTasks ? (
                <div className="relative flex-1 overflow-y-hidden">
                    <div
                        className="flex h-full flex-col gap-4 overflow-y-auto"
                        ref={containerRef}
                    >
                        {sortedTasks.map((task) => (
                            <FlowTask
                                key={task.id}
                                searchValue={debouncedSearchValue}
                                task={task}
                            />
                        ))}
                        <div ref={endRef} />
                    </div>

                    {!isScrolledToBottom && (
                        <Button
                            aria-label="捲動到最新任務"
                            className="absolute right-4 bottom-4 z-10 shadow-md hover:shadow-lg"
                            onClick={() => scrollToEnd()}
                            size="icon-sm"
                            type="button"
                            variant="outline"
                        >
                            <ChevronDown />
                            {hasNewMessages && (
                                <span className="bg-primary absolute -top-1 -right-1 size-3 rounded-full" />
                            )}
                        </Button>
                    )}
                </div>
            ) : (
                <Empty>
                    <EmptyHeader>
                        <EmptyMedia variant="icon">
                            <ListTodo />
                        </EmptyMedia>
                        <EmptyTitle>此流程沒有任務</EmptyTitle>
                        <EmptyDescription>代理開始運作後，任務會顯示在這裡</EmptyDescription>
                    </EmptyHeader>
                </Empty>
            )}
        </div>
    );
}

export default FlowTasks;
