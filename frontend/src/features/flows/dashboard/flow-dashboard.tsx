import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { FlowDashboardOverview } from '@/features/flows/dashboard/flow-dashboard-overview';
import { useFlow } from '@/providers/flow-provider';

function FlowDashboard() {
    const { flowId } = useFlow();

    if (!flowId) {
        return (
            <div className="text-muted-foreground flex items-center justify-center py-12">
                選擇流程以檢視儀表板
            </div>
        );
    }

    return (
        <Tabs defaultValue="overview">
            <TabsList className="hidden">
                <TabsTrigger value="overview">總覽</TabsTrigger>
            </TabsList>

            <TabsContent value="overview">
                <FlowDashboardOverview flowId={flowId} />
            </TabsContent>
        </Tabs>
    );
}

export default FlowDashboard;
