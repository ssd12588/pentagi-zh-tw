import { AlertCircle } from 'lucide-react';

import { ProviderIcon } from '@/components/icons/provider-icon';
import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip';
import { cn } from '@/lib/utils';
import { isProviderMissing, type Provider } from '@/models/provider';

interface FlowProviderLabelProps {
    isOwnFlow: boolean;
    provider: Provider;
    providers: Provider[];
}

export function FlowProviderLabel({ isOwnFlow, provider, providers }: FlowProviderLabelProps) {
    const isMissing = isOwnFlow && isProviderMissing(provider, providers);

    return (
        <div
            className={cn('flex min-w-0 items-center gap-2', isMissing && 'opacity-50')}
            data-slot="flow-provider-label"
        >
            <ProviderIcon
                className="size-4 shrink-0"
                provider={provider}
            />
            <span
                className="truncate text-sm"
                title={provider?.name}
            >
                {provider?.name || '不適用'}
            </span>
            {isMissing && (
                <Tooltip>
                    <TooltipTrigger asChild>
                        <AlertCircle
                            aria-label="無法使用"
                            className="text-destructive size-4 shrink-0 cursor-pointer"
                            role="img"
                            tabIndex={0}
                        />
                    </TooltipTrigger>
                    <TooltipContent>無法使用</TooltipContent>
                </Tooltip>
            )}
        </div>
    );
}
