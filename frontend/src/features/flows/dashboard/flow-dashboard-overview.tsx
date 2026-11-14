import { useQuery } from '@apollo/client/react';
import { Activity, CircleDollarSign, Cpu, GitFork } from 'lucide-react';
import { useMemo } from 'react';

import { MetricCard, UsageStatsRow } from '@/components/dashboard';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip';
import FlowAgentIcon from '@/features/flows/agents/flow-agent-icon';
import {
    AgentType,
    FlowStatsByFlowDocument,
    ToolcallsStatsByFlowDocument,
    ToolcallsStatsByFunctionForFlowDocument,
    UsageStatsByAgentTypeForFlowDocument,
    UsageStatsByFlowDocument,
    UsageStatsByModelAgentsForFlowDocument,
} from '@/graphql/types';
import { formatCost, formatDuration, formatNumber, formatTokenCount } from '@/lib/utils/format';

export function FlowDashboardOverview({ flowId }: { flowId: string }) {
    const {
        data: usageData,
        error: usageError,
        loading: usageLoading,
    } = useQuery(UsageStatsByFlowDocument, {
        variables: { flowId },
    });
    const { data: usageByAgentData } = useQuery(UsageStatsByAgentTypeForFlowDocument, {
        variables: { flowId },
    });
    const { data: usageByModelAgentsData } = useQuery(UsageStatsByModelAgentsForFlowDocument, {
        variables: { flowId },
    });
    const {
        data: toolcallsData,
        error: toolcallsError,
        loading: toolcallsLoading,
    } = useQuery(ToolcallsStatsByFlowDocument, {
        variables: { flowId },
    });
    const { data: toolcallsByFunctionData } = useQuery(ToolcallsStatsByFunctionForFlowDocument, {
        variables: { flowId },
    });
    const {
        data: flowStatsData,
        error: flowStatsError,
        loading: flowStatsLoading,
    } = useQuery(FlowStatsByFlowDocument, {
        variables: { flowId },
    });

    const usage = usageData?.usageStatsByFlow;
    const toolcalls = toolcallsData?.toolcallsStatsByFlow;
    const flowStats = flowStatsData?.flowStatsByFlow;

    const totalCost = usage ? usage.totalUsageCostIn + usage.totalUsageCostOut : 0;
    const totalTokens = usage ? usage.totalUsageIn + usage.totalUsageOut : 0;

    const agentTypeRows = useMemo(() => {
        const seen = new Set<string>();

        return (usageByAgentData?.usageStatsByAgentTypeForFlow ?? [])
            .filter((item) => {
                if (seen.has(item.agentType)) {
                    return false;
                }

                seen.add(item.agentType);

                return true;
            })
            .map((item) => ({
                label: item.agentType,
                stats: item.stats,
            }));
    }, [usageByAgentData]);

    const modelAgentRows = useMemo(() => {
        const seen = new Set<string>();

        return (usageByModelAgentsData?.usageStatsByModelAgentsForFlow ?? []).filter((item) => {
            const key = `${item.model}|${item.provider}`;

            if (seen.has(key)) {
                return false;
            }

            seen.add(key);

            return true;
        });
    }, [usageByModelAgentsData]);

    const toolcallsByFunction = useMemo(() => {
        const seen = new Set<string>();

        return [...(toolcallsByFunctionData?.toolcallsStatsByFunctionForFlow ?? [])]
            .filter((item) => {
                if (seen.has(item.functionName)) {
                    return false;
                }

                seen.add(item.functionName);

                return true;
            })
            .sort((a, b) => b.totalCount - a.totalCount);
    }, [toolcallsByFunctionData]);

    const anyLoading =
        (usageLoading && !usageData) || (toolcallsLoading && !toolcallsData) || (flowStatsLoading && !flowStatsData);

    return (
        <div className="flex flex-col gap-6">
            <div className="grid grid-cols-2 gap-4 md:grid-cols-4">
                <MetricCard
                    description={`Subtasks: ${flowStats?.totalSubtasksCount ?? 0} · Assistants: ${flowStats?.totalAssistantsCount ?? 0}`}
                    error={!!flowStatsError && !flowStatsData}
                    icon={<GitFork className="text-muted-foreground size-4" />}
                    loading={anyLoading}
                    title="任務"
                    value={flowStats ? formatNumber(flowStats.totalTasksCount) : '0'}
                />
                <MetricCard
                    description={`Duration: ${toolcalls ? formatDuration(toolcalls.totalDurationSeconds) : '—'}`}
                    error={!!toolcallsError && !toolcallsData}
                    icon={<Activity className="text-muted-foreground size-4" />}
                    loading={anyLoading}
                    title="工具呼叫"
                    value={toolcalls ? formatNumber(toolcalls.totalCount) : '0'}
                />
                <MetricCard
                    description="輸入 + 輸出 Token"
                    error={!!usageError && !usageData}
                    icon={<Cpu className="text-muted-foreground size-4" />}
                    loading={anyLoading}
                    title="Token"
                    value={formatTokenCount(totalTokens)}
                />
                <MetricCard
                    description="此流程的 LLM 花費"
                    error={!!usageError && !usageData}
                    icon={<CircleDollarSign className="text-muted-foreground size-4" />}
                    loading={anyLoading}
                    title="成本"
                    value={formatCost(totalCost)}
                />
            </div>

            {!!modelAgentRows.length && (
                <Card>
                    <CardHeader>
                        <CardTitle>各模型與供應商的用量</CardTitle>
                        <CardDescription>
                            依模型與供應商分組的 LLM Token 用量與成本（含使用的代理類型）
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <Table aria-label="各模型與供應商的用量">
                            <TableHeader>
                                <TableRow>
                                    <TableHead className="w-40 whitespace-nowrap">模型</TableHead>
                                    <TableHead className="whitespace-nowrap">供應商</TableHead>
                                    <TableHead className="w-36 whitespace-nowrap">代理</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">輸入 Token</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">輸出 Token</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">快取讀取</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">快取寫入</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">輸入成本</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">輸出成本</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">總成本</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {modelAgentRows.map((row) => (
                                    <TableRow key={`${row.model}|${row.provider}`}>
                                        <TableCell className="font-medium">
                                            <Tooltip>
                                                <TooltipTrigger asChild>
                                                    <span className="block w-40 truncate">{row.model}</span>
                                                </TooltipTrigger>
                                                <TooltipContent>{row.model}</TooltipContent>
                                            </Tooltip>
                                        </TableCell>
                                        <TableCell>{row.provider}</TableCell>
                                        <TableCell>
                                            <div className="flex w-36 flex-wrap gap-1">
                                                {row.agentTypes.map((agentType) => (
                                                    <FlowAgentIcon
                                                        className="size-3.5"
                                                        key={agentType}
                                                        tooltip={agentType}
                                                        type={agentType as AgentType}
                                                    />
                                                ))}
                                            </div>
                                        </TableCell>
                                        <TableCell className="text-right">
                                            {formatTokenCount(row.stats.totalUsageIn)}
                                        </TableCell>
                                        <TableCell className="text-right">
                                            {formatTokenCount(row.stats.totalUsageOut)}
                                        </TableCell>
                                        <TableCell className="text-right">
                                            {formatTokenCount(row.stats.totalUsageCacheIn)}
                                        </TableCell>
                                        <TableCell className="text-right">
                                            {formatTokenCount(row.stats.totalUsageCacheOut)}
                                        </TableCell>
                                        <TableCell className="text-right">
                                            {formatCost(row.stats.totalUsageCostIn)}
                                        </TableCell>
                                        <TableCell className="text-right">
                                            {formatCost(row.stats.totalUsageCostOut)}
                                        </TableCell>
                                        <TableCell className="text-right font-semibold">
                                            {formatCost(row.stats.totalUsageCostIn + row.stats.totalUsageCostOut)}
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </CardContent>
                </Card>
            )}

            {!!agentTypeRows.length && (
                <Card>
                    <CardHeader>
                        <CardTitle>各代理類型的用量</CardTitle>
                        <CardDescription>此流程中各代理類型的 LLM Token 用量與成本</CardDescription>
                    </CardHeader>
                    <CardContent>
                        <Table aria-label="各代理類型的用量">
                            <TableHeader>
                                <TableRow>
                                    <TableHead>代理類型</TableHead>
                                    <TableHead className="text-right">輸入 Token</TableHead>
                                    <TableHead className="text-right">輸出 Token</TableHead>
                                    <TableHead className="text-right">快取讀取</TableHead>
                                    <TableHead className="text-right">快取寫入</TableHead>
                                    <TableHead className="text-right">輸入成本</TableHead>
                                    <TableHead className="text-right">輸出成本</TableHead>
                                    <TableHead className="text-right">總成本</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {agentTypeRows.map((row) => (
                                    <UsageStatsRow
                                        key={row.label}
                                        label={row.label}
                                        stats={row.stats}
                                    />
                                ))}
                            </TableBody>
                        </Table>
                    </CardContent>
                </Card>
            )}

            {!!toolcallsByFunction.length && (
                <Card>
                    <CardHeader>
                        <CardTitle>各函式的工具呼叫次數</CardTitle>
                        <CardDescription>此流程中各工具函式的執行統計</CardDescription>
                    </CardHeader>
                    <CardContent>
                        <Table aria-label="各函式的工具呼叫次數">
                            <TableHeader>
                                <TableRow>
                                    <TableHead>函式</TableHead>
                                    <TableHead>類型</TableHead>
                                    <TableHead className="text-right">數量</TableHead>
                                    <TableHead className="text-right">總耗時</TableHead>
                                    <TableHead className="text-right">平均耗時</TableHead>
                                </TableRow>
                            </TableHeader>
                            <TableBody>
                                {toolcallsByFunction.map((item) => (
                                    <TableRow key={item.functionName}>
                                        <TableCell className="font-medium">{item.functionName}</TableCell>
                                        <TableCell>
                                            <Badge variant={item.isAgent ? 'secondary' : 'outline'}>
                                                {item.isAgent ? '代理' : '工具'}
                                            </Badge>
                                        </TableCell>
                                        <TableCell className="text-right">{formatNumber(item.totalCount)}</TableCell>
                                        <TableCell className="text-right">
                                            {formatDuration(item.totalDurationSeconds)}
                                        </TableCell>
                                        <TableCell className="text-right">
                                            {formatDuration(item.avgDurationSeconds)}
                                        </TableCell>
                                    </TableRow>
                                ))}
                            </TableBody>
                        </Table>
                    </CardContent>
                </Card>
            )}
        </div>
    );
}
