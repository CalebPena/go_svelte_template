import { applyAction } from '$app/forms';
import { goto } from '$app/navigation';
import { DEFUALT_ERROR_MESSAGE } from '$lib/constants';
import type { ActionResult } from '@sveltejs/kit';
import type { z } from 'zod';

export function formInputError<T>(schema: z.Schema<T>, value: unknown): string {
	const data = schema.safeParse(value);

	if (!data.success) {
		if (data.error.issues.length === 0) {
			return DEFUALT_ERROR_MESSAGE;
		}

		return data.error.issues[0].message;
	}

	return '';
}

type ActionWrapper<T extends Record<string, unknown>> = {
	formData: FormData;
	formElement: HTMLFormElement;
	action: URL;
	result: ActionResult<T>;
	update: (options?: { reset?: boolean; invalidateAll?: boolean }) => Promise<void>;
};

export function handleActionWrapper<T extends Record<string, unknown>>(
	func?: (arg: ActionWrapper<T>) => void,
) {
	return async (arg: ActionWrapper<T>) => {
		if (arg.result.type === 'redirect') {
			goto(arg.result.location);
			return;
		}

		if (func) {
			func(arg);
		}

		await applyAction(arg.result);
	};
}
