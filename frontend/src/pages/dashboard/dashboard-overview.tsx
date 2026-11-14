import { useQuery } from '@apollo/client/react';
import { Activity, CircleDollarSign, Cpu, GitFork } from 'lucide-react';

import type { UsageStatsFragmentFragment } from '@/graphql/types';

import { DashboardError, LoadingTable, MetricCard, UsageStatsRow } from '@/components/dashboard';
import { Badge } from '@/components/ui/badge';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import {
    FlowsStatsTotalDocument,
    ToolcallsStatsByFunctionDocument,
    ToolcallsStatsTotalDocument,
    UsageStatsByAgentTypeDocument,
    UsageStatsByModelDocument,
    UsageStatsByProviderDocument,
    UsageStatsTotalDocument,
} from '@/graphql/types';
import { formatCost, formatDuration, formatNumber, formatTokenCount } from '@/lib/utils/format';

export function DashboardOverview() {
    const {
        data: usageTotalData,
        error: usageTotalError,
        loading: usageTotalLoading,
    } = useQuery(UsageStatsTotalDocument);
    const {
        data: usageByProviderData,
        error: usageByProviderError,
        loading: usageByProviderLoading,
    } = useQuery(UsageStatsByProviderDocument);
    const {
        data: usageByModelData,
        error: usageByModelError,
        loading: usageByModelLoading,
    } = useQuery(UsageStatsByModelDocument);
    const {
        data: usageByAgentTypeData,
        error: usageByAgentTypeError,
        loading: usageByAgentTypeLoading,
    } = useQuery(UsageStatsByAgentTypeDocument);
    const {
        data: toolcallsTotalData,
        error: toolcallsTotalError,
        loading: toolcallsTotalLoading,
    } = useQuery(ToolcallsStatsTotalDocument);
    const {
        data: toolcallsByFunctionData,
        error: toolcallsByFunctionError,
        loading: toolcallsByFunctionLoading,
    } = useQuery(ToolcallsStatsByFunctionDocument);
    const {
        data: flowsTotalData,
        error: flowsTotalError,
        loading: flowsTotalLoading,
    } = useQuery(FlowsStatsTotalDocument);

    const usageTotal = usageTotalData?.usageStatsTotal;
    const toolcallsTotal = toolcallsTotalData?.toolcallsStatsTotal;
    const flowsTotal = flowsTotalData?.flowsStatsTotal;

    const totalCost = usageTotal ? usageTotal.totalUsageCostIn + usageTotal.totalUsageCostOut : 0;
    const totalTokens = usageTotal ? usageTotal.totalUsageIn + usageTotal.totalUsageOut : 0;

    const providerRows = (usageByProviderData?.usageStatsByProvider ?? []).map((item) => ({
        label: item.provider,
        stats: item.stats,
    }));
    const modelRows = (usageByModelData?.usageStatsByModel ?? []).map((item) => ({
        label: `${item.model} (${item.provider})`,
        stats: item.stats,
    }));
    const agentTypeRows = (usageByAgentTypeData?.usageStatsByAgentType ?? []).map((item) => ({
        label: item.agentType,
        stats: item.stats,
    }));

    const toolcallsByFunction = [...(toolcallsByFunctionData?.toolcallsStatsByFunction ?? [])].sort(
        (a, b) => b.totalCount - a.totalCount,
    );

    return (
        <div className="flex flex-col gap-6">
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
                <MetricCard
                    description={`Tasks: ${flowsTotal?.totalTasksCount ?? 0} · Subtasks: ${flowsTotal?.totalSubtasksCount ?? 0} · Assistants: ${flowsTotal?.totalAssistantsCount ?? 0}`}
                    error={!!flowsTotalError && !flowsTotalData}
                    icon={<GitFork className="text-muted-foreground size-4" />}
                    loading={flowsTotalLoading && !flowsTotalData}
                    title="流程總數"
                    value={flowsTotal ? formatNumber(flowsTotal.totalFlowsCount) : '0'}
                />
                <MetricCard
                    description={`Total duration: ${toolcallsTotal ? formatDuration(toolcallsTotal.totalDurationSeconds) : '—'}`}
                    error={!!toolcallsTotalError && !toolcallsTotalData}
                    icon={<Activity className="text-muted-foreground size-4" />}
                    loading={toolcallsTotalLoading && !toolcallsTotalData}
                    title="工具呼叫"
                    value={toolcallsTotal ? formatNumber(toolcallsTotal.totalCount) : '0'}
                />
                <MetricCard
                    description="已處理的輸入 + 輸出 Token"
                    error={!!usageTotalError && !usageTotalData}
                    icon={<Cpu className="text-muted-foreground size-4" />}
                    loading={usageTotalLoading && !usageTotalData}
                    title="Token 總數"
                    value={formatTokenCount(totalTokens)}
                />
                <MetricCard
                    description="所有供應商的 LLM 總花費"
                    error={!!usageTotalError && !usageTotalData}
                    icon={<CircleDollarSign className="text-muted-foreground size-4" />}
                    loading={usageTotalLoading && !usageTotalData}
                    title="總成本"
                    value={formatCost(totalCost)}
                />
            </div>

            <Card>
                <CardHeader>
                    <CardTitle>各供應商的用量</CardTitle>
                    <CardDescription>依供應商分組的 LLM Token 用量與成本</CardDescription>
                </CardHeader>
                <CardContent>
                    {usageByProviderLoading && !usageByProviderData ? (
                        <LoadingTable />
                    ) : usageByProviderError && !usageByProviderData ? (
                        <ErrorTable />
                    ) : (
                        <UsageStatsTable
                            label="各供應商的用量"
                            rows={providerRows}
                        />
                    )}
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>各模型的用量</CardTitle>
                    <CardDescription>依模型分組的 LLM Token 用量與成本</CardDescription>
                </CardHeader>
                <CardContent>
                    {usageByModelLoading && !usageByModelData ? (
                        <LoadingTable />
                    ) : usageByModelError && !usageByModelData ? (
                        <ErrorTable />
                    ) : (
                        <UsageStatsTable
                            label="各模型的用量"
                            rows={modelRows}
                        />
                    )}
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>各代理類型的用量</CardTitle>
                    <CardDescription>依代理類型分組的 LLM Token 用量與成本</CardDescription>
                </CardHeader>
                <CardContent>
                    {usageByAgentTypeLoading && !usageByAgentTypeData ? (
                        <LoadingTable />
                    ) : usageByAgentTypeError && !usageByAgentTypeData ? (
                        <ErrorTable />
                    ) : (
                        <UsageStatsTable
                            label="各代理類型的用量"
                            rows={agentTypeRows}
                        />
                    )}
                </CardContent>
            </Card>

            <Card>
                <CardHeader>
                    <CardTitle>各函式的工具呼叫次數</CardTitle>
                    <CardDescription>各工具函式的執行統計</CardDescription>
                </CardHeader>
                <CardContent>
                    {toolcallsByFunctionLoading && !toolcallsByFunctionData ? (
                        <LoadingTable />
                    ) : toolcallsByFunctionError && !toolcallsByFunctionData ? (
                        <ErrorTable />
                    ) : (
                        <Table aria-label="各函式的工具呼叫次數">
                            <TableHeader>
                                <TableRow>
                                    <TableHead className="whitespace-nowrap">函式</TableHead>
                                    <TableHead className="whitespace-nowrap">類型</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">數量</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">總耗時</TableHead>
                                    <TableHead className="text-right whitespace-nowrap">平均耗時</TableHead>
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
                    )}
                </CardContent>
            </Card>
        </div>
    );
}

function ErrorTable() {
    return <DashboardError className="py-8" />;
}

function UsageStatsTable({
    label,
    rows,
}: {
    label: string;
    rows: Array<{ label: string; stats: UsageStatsFragmentFragment }>;
}) {
    return (
        <Table aria-label={label}>
            <TableHeader>
                <TableRow>
                    <TableHead className="whitespace-nowrap">名稱</TableHead>
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
                {rows.map((row) => (
                    <UsageStatsRow
                        key={row.label}
                        label={row.label}
                        stats={row.stats}
                    />
                ))}
            </TableBody>
        </Table>
    );
}
