import { z } from 'zod';
import { DEFUALT_ERROR_MESSAGE, MIN_PASSWORD_LENGTH } from './constants';
import { error } from '@sveltejs/kit';

/**
 * Validate the value with a zod schema
 * Throw a 400 response if the value is invalid
 * Return the validated data
 */
export function validateApiInput<T>(schema: z.Schema<T>, value: unknown): T {
	const data = schema.safeParse(value);

	if (!data.success) {
		if (data.error.issues.length === 0) {
			throw error(500, DEFUALT_ERROR_MESSAGE);
		}

		throw error(400, data.error.issues[0].message);
	}

	return data.data;
}

export const errorResponseSchema = z.object({ error: z.string() });
export function responseSchemaWrapper<T>(schema: z.Schema<T>) {
	return z.object({ response: schema });
}

export const emailSchema = z
	.string()
	.trim()
	.email('Please enter a valid email');
export const passwordSchema = z
	.string()
	.min(MIN_PASSWORD_LENGTH, 'Password must be at least 8 characters long');
export const loginPasswordSchema = z.string().min(1, 'Please enter a password');
