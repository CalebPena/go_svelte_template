import { dev } from '$app/environment';
import { loginEndpoint } from '$lib/api/endpoints';
import { actionWrapper, requestApi } from '$lib/api/helpers';
import { HOME_PAGE_PATH, REDIRECT_PARAM_NAME, TOKEN_HEADER } from '$lib/constants';
import {
	emailSchema,
	loginPasswordSchema,
	responseSchemaWrapper,
	validateApiInput,
} from '$lib/schemas';
import { redirect, type Actions } from '@sveltejs/kit';
import { z } from 'zod';

export const actions: Actions = {
	default: actionWrapper(async ({ request, cookies, url }) => {
		const formData = await request.formData();

		const email = validateApiInput(emailSchema, formData.get('email'));
		const password = validateApiInput(loginPasswordSchema, formData.get('password'));

		const { data } = await requestApi(
			loginEndpoint(),
			responseSchemaWrapper(z.object({ token: z.string() })),
			'POST',
			{ email: email, password: password },
		);

		cookies.set(TOKEN_HEADER, data.response.token, {
			httpOnly: true,
			secure: !dev,
			sameSite: 'strict',
			path: '/',
		});

		const redirectUrl = url.searchParams.get(REDIRECT_PARAM_NAME);

		if (redirectUrl !== null) {
			throw redirect(303, redirectUrl);
		}

		throw redirect(303, HOME_PAGE_PATH);
	}),
};
