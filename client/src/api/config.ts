import { env } from '$env/dynamic/public';

export default {
	origin: env.PUBLIC_DOMAIN ?? 'localhost',
	port: env.PUBLIC_API_PORT
};
