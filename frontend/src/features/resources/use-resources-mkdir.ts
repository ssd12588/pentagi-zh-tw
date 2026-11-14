import { useCallback, useState } from 'react';
import { toast } from 'sonner';
import { z } from 'zod';

import { api, getApiErrorMessage } from '@/lib/axios';

import { RESOURCES_MKDIR_API_PATH } from './resources-constants';

export const resourcesMkdirFormSchema = z.object({
    path: z
        .string()
        .trim()
        .min(1, { message: '路徑不能為空' })
        .refine((value) => !value.startsWith('/'), { message: '路徑必須是相對路徑（開頭不能是「/」）' })
        .refine((value) => !value.split('/').includes('..'), { message: '路徑不能包含「..」' }),
});

export type ResourcesMkdirFormValues = z.infer<typeof resourcesMkdirFormSchema>;

interface MkdirRequestBody {
    path: string;
}

interface UseResourcesMkdirResult {
    isCreating: boolean;
    mkdir: (values: ResourcesMkdirFormValues) => Promise<boolean>;
}

/** Wraps `POST /resources/mkdir` (idempotent — returns existing dir on hit). */
export function useResourcesMkdir(): UseResourcesMkdirResult {
    const [isCreating, setIsCreating] = useState(false);

    const mkdir = useCallback(async ({ path }: ResourcesMkdirFormValues): Promise<boolean> => {
        setIsCreating(true);

        try {
            await api.post<void, MkdirRequestBody>(RESOURCES_MKDIR_API_PATH, { path: path.trim() });

            toast.success('目錄已建立', { description: `Created at /${path.trim()}` });

            return true;
        } catch (error) {
            const description = getApiErrorMessage(error, 'Failed to create directory');

            toast.error('建立目錄失敗', { description });

            return false;
        } finally {
            setIsCreating(false);
        }
    }, []);

    return {
        isCreating,
        mkdir,
    };
}
