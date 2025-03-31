import { building } from '$app/environment';

if (!building) {
	process.env.NODE_TLS_REJECT_UNAUTHORIZED = '0';
}

export const handle = async ({ event, resolve }) => {
	return resolve(event);
};
