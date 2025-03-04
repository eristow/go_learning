import { getAlbums } from '$lib/db/index.js';
import type { Album } from '$types';

export type PageLoad = {
	albums: Promise<Album[]>;
};

export const load = async ({ fetch }): Promise<PageLoad> => {
	return {
		albums: getAlbums(fetch)
	};
};
