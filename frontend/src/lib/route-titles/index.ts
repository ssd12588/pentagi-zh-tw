import type { ComponentType } from 'react';

import {
    FlowDocument,
    FlowTemplateDocument,
    KnowledgeDocumentDocument,
    SettingsProvidersDocument,
} from '@/graphql/types';

import { apolloTitle } from './apollo-title';
import { formatPromptId } from './format-prompt-id';
import { type RouteParams } from './render-title';

export interface RouteTitleHandle {
    title: TitleResolver;
}

/**
 * A `handle.title` value can be one of three forms:
 *   - `string` — fully static, known at build time.
 *   - `(params) => string` — derived synchronously from URL params.
 *   - `ComponentType<{ params }>` — reactive (e.g. subscribes to Apollo
 *     cache for resource-driven titles). Must be produced by `apolloTitle()`
 *     so the marker it attaches lets `DocumentTitle` distinguish a component
 *     from a `(params) => string` resolver at runtime. A hand-rolled component
 *     function will be misdetected as a resolver and called with raw params —
 *     always route reactive titles through `apolloTitle()`.
 */
export type TitleResolver = ((params: RouteParams) => string) | ComponentType<{ params: RouteParams }> | string;

/**
 * Single source of truth for every route's document `<title>`. `app.tsx`
 * imports nothing from Apollo for title purposes — it only spreads handles
 * from this registry onto the matching <Route>.
 */
export const routeTitles = {
    account: { title: '帳號' },
    apiTokens: { title: 'API Token' },
    dashboard: { title: '儀表板' },
    flow: {
        title: apolloTitle({
            document: FlowDocument,
            select: (data, { flowId }) =>
                data?.flow?.title && flowId ? `Flow #${flowId} — ${data.flow.title}` : 'Flow',
            variables: ({ flowId }) => (flowId ? { id: flowId } : null),
        }),
    },
    flowReport: { title: '流程報告' },
    flows: { title: '流程' },
    knowledge: {
        title: apolloTitle({
            document: KnowledgeDocumentDocument,
            select: (data, { knowledgeId }) =>
                knowledgeId === 'new' ? 'New knowledge' : data?.knowledgeDocument?.question || 'Knowledge',
            variables: ({ knowledgeId }) => (!knowledgeId || knowledgeId === 'new' ? null : { id: knowledgeId }),
        }),
    },
    knowledges: { title: '知識庫' },
    login: { title: '登入' },
    newFlow: { title: '新增流程' },
    oauth: { title: 'OAuth' },
    prompt: {
        title: (params: RouteParams) => (params.promptId ? formatPromptId(params.promptId) : 'Prompt'),
    },
    prompts: { title: '提示詞' },

    provider: {
        title: apolloTitle({
            document: SettingsProvidersDocument,
            select: (data, { providerId }) => {
                if (providerId === 'new') {
                    return 'New provider';
                }

                const provider = data?.settingsProviders.userDefined?.find(
                    (candidate) => String(candidate.id) === providerId,
                );

                return provider?.name || 'Provider';
            },
            variables: ({ providerId }) => (providerId === 'new' ? null : {}),
        }),
    },

    providers: { title: '供應商' },

    resources: { title: '資源' },

    template: {
        title: apolloTitle({
            document: FlowTemplateDocument,
            select: (data, { templateId }) =>
                templateId === 'new' ? 'New template' : data?.flowTemplate?.title || 'Template',
            variables: ({ templateId }) => (!templateId || templateId === 'new' ? null : { templateId }),
        }),
    },

    templates: { title: '範本' },
} as const satisfies Record<string, RouteTitleHandle>;
