import type { Album } from '$types';
import { getAlbum } from '$lib/db';

export type PageLoad = {
	album: Promise<Album>;
};

export const load = async ({ fetch, params }): Promise<PageLoad> => {
	return {
		album: getAlbum(fetch, params.albumId)
	};
};
