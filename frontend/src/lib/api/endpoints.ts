import { BACKEND_DOMAIN } from '$env/static/private';

export function loginEndpoint() {
	return `${BACKEND_DOMAIN}/auth/login`;
}
export function registerEndpoint() {
	return `${BACKEND_DOMAIN}/auth/register`;
}
