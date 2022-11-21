import { env } from '$env/dynamic/public';

export default {
	url: env.PUBLIC_API_URL ?? 'http://localhost',
};
