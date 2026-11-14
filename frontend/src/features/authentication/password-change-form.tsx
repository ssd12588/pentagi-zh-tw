import { type ComponentProps, useState } from 'react';
import { toast } from 'sonner';
import * as z from 'zod';

import { Button } from '@/components/ui/button';
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { FormSubmitButton } from '@/components/ui/form-submit-button';
import { InputPassword } from '@/components/ui/input-password';
import { useAppForm } from '@/hooks/use-app-form';
import { api, resolveApiErrorMessage } from '@/lib/axios';
import { cn } from '@/lib/utils';

const passwordChangeFormSchema = z
    .object({
        confirmPassword: z.string().min(1, { message: '確認密碼' }),
        currentPassword: z.string().min(1, { message: '必須輸入目前密碼' }),
        newPassword: z
            .string()
            .min(8, { message: '密碼長度至少 8 個字元' })
            // bcrypt, which hashes it server-side, refuses anything longer than 72 bytes.
            .refine((password) => new TextEncoder().encode(password).length <= 72, {
                message: '密碼不得超過 72 個字元',
            })
            .refine(
                (password) => {
                    const length = [...password].length;

                    if (length > 15) {
                        return true;
                    }

                    return (
                        length >= 8 &&
                        /[0-9]/.test(password) &&
                        /[a-z]/.test(password) &&
                        /[A-Z]/.test(password) &&
                        /[!@#$&*]/.test(password)
                    );
                },
                {
                    message:
                        '密碼長度須超過 15 個字元，或至少 8 個字元且包含數字、小寫、大寫與特殊字元（!@#$&*）',
                },
            ),
    })
    .refine((data) => data.newPassword === data.confirmPassword, {
        message: "兩次輸入的密碼不一致",
        path: ['confirmPassword'],
    })
    .refine((data) => data.currentPassword !== data.newPassword, {
        message: '新密碼必須與目前密碼不同',
        path: ['newPassword'],
    });

const ERROR_BY_CODE: Record<string, string> = {
    'Users.ChangePasswordCurrentUser.InvalidCurrentPassword': 'Current password is incorrect',
    'Users.ChangePasswordCurrentUser.InvalidNewPassword': 'New password does not meet requirements',
    'Users.ChangePasswordCurrentUser.InvalidPassword': 'Password validation failed',
    'Users.NotFound': 'User not found',
};

interface PasswordChangeFormProps {
    buttonSize?: ComponentProps<typeof Button>['size'];
    layout?: 'horizontal' | 'vertical';
    onCancel?: () => void;
    onSkip?: () => void;
    onSuccess?: () => void;
}

type PasswordChangeFormValues = z.infer<typeof passwordChangeFormSchema>;

export function PasswordChangeForm({
    buttonSize = 'default',
    layout = 'horizontal',
    onCancel,
    onSkip,
    onSuccess,
}: PasswordChangeFormProps) {
    const [error, setError] = useState<null | string>(null);

    const form = useAppForm<PasswordChangeFormValues>({
        defaultValues: {
            confirmPassword: '',
            currentPassword: '',
            newPassword: '',
        },
        schema: passwordChangeFormSchema,
    });

    const handleSubmit = async (values: PasswordChangeFormValues) => {
        setError(null);

        try {
            await api.put('/user/password', {
                confirm_password: values.confirmPassword,
                current_password: values.currentPassword,
                password: values.newPassword,
            });

            form.reset();
            toast.success('密碼變更成功');

            onSuccess?.();
        } catch (err: unknown) {
            setError(resolveApiErrorMessage(err, ERROR_BY_CODE, 'Failed to change password'));
        }
    };

    const isVertical = layout === 'vertical';

    const skipButton = onSkip && (
        <Button
            className={cn('text-muted-foreground', isVertical && 'w-full')}
            onClick={onSkip}
            size={buttonSize}
            type="button"
            variant="ghost"
        >
            暫時略過
        </Button>
    );
    const cancelButton = onCancel && (
        <Button
            className={cn(isVertical && 'w-full')}
            onClick={onCancel}
            size={buttonSize}
            type="button"
            variant="outline"
        >
            取消
        </Button>
    );
    const submitButton = (
        <FormSubmitButton
            className={cn(isVertical && 'w-full')}
            size={buttonSize}
        >
            <span>更新密碼</span>
        </FormSubmitButton>
    );

    return (
        <Form {...form}>
            <form
                className="flex flex-col gap-4"
                noValidate
                onSubmit={form.handleSubmit(handleSubmit)}
            >
                <FormField
                    control={form.control}
                    name="currentPassword"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>目前密碼</FormLabel>
                            <FormControl>
                                <InputPassword
                                    {...field}
                                    placeholder="輸入目前的密碼"
                                />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />

                <FormField
                    control={form.control}
                    name="newPassword"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>新密碼</FormLabel>
                            <FormControl>
                                <InputPassword
                                    {...field}
                                    placeholder="輸入新密碼"
                                />
                            </FormControl>
                            <FormDescription className="text-xs">
                                Must be 16+ characters, or 8+ with number, lowercase, uppercase, and special character
                                (!@#$&*)
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />

                <FormField
                    control={form.control}
                    name="confirmPassword"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>確認新密碼</FormLabel>
                            <FormControl>
                                <InputPassword
                                    {...field}
                                    placeholder="再次輸入新密碼"
                                />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />

                {error && <div className="text-destructive text-sm">{error}</div>}

                {isVertical ? (
                    <div className="flex flex-col gap-2 pt-2">
                        {submitButton}
                        {cancelButton}
                        {skipButton}
                    </div>
                ) : (
                    <div className="flex justify-end gap-2 pt-2">
                        {skipButton}
                        {cancelButton}
                        {submitButton}
                    </div>
                )}
            </form>
        </Form>
    );
}
