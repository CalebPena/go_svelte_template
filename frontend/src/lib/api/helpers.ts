import { DEFUALT_ERROR_MESSAGE } from '$lib/constants';
import { errorResponseSchema } from '$lib/schemas';
import { error, fail, isHttpError, type Action } from '@sveltejs/kit';
import type { z } from 'zod';

type ApiRequestReturn<T> = {
	data: T;
	res: Response;
};

export async function checkResponse(res: Response) {
	if (res.status >= 200 && res.status < 300) {
		return null;
	}

	let rawData;
	try {
		rawData = await res.json();
	} catch {
		error(500, DEFUALT_ERROR_MESSAGE);
	}

	const data = errorResponseSchema.safeParse(rawData);

	if (!data.success) {
		error(500, DEFUALT_ERROR_MESSAGE);
	}

	error(res.status, data.data.error);
}

export async function requestApi<T>(
	url: string,
	responseSchema: z.Schema<T>,
	method: 'GET' | 'OPTIONS' | 'POST' | 'PUT' | 'PATCH' | 'DELETE',
	data?: unknown,
	token = '',
): Promise<ApiRequestReturn<T>> {
	const headers = new Headers();

	headers.set('Content-Type', 'application/json');

	if (token !== '') {
		headers.set('Authorization', `Bearer ${token}`);
	}

	let res: Response;
	try {
		res = await fetch(url, {
			headers: headers,
			method: method,
			body: JSON.stringify(data),
		});
	} catch {
		error(500, DEFUALT_ERROR_MESSAGE);
	}

	const err = await checkResponse(res);

	if (err !== null) {
		return err;
	}

	try {
		const data = await res.json();
		return { data: responseSchema.parse(data), res: res };
	} catch {
		error(500, DEFUALT_ERROR_MESSAGE);
	}
}

export function actionWrapper(handler: Action): Action {
	return async (...args) => {
		try {
			return await handler(...args);
		} catch (err) {
			if (isHttpError(err)) {
				return fail(err.status, { message: err.body.message });
			}

			throw err;
		}
	};
}
