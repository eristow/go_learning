import type { Album } from '$types';

export type PageLoad = {
	albums: Promise<Album[]>;
};

export const load = async ({ fetch }): Promise<PageLoad> => {
	const fetchAlbums = async (): Promise<Album[]> => {
		try {
			// TODO: change the backend URL to an env var
			const albumsRes = await fetch('http://localhost:8080/albums');
			const albumsData = await albumsRes.json();
			return albumsData;
		} catch (e) {
			console.error(e);
			return [];
		}
	};

	return {
		albums: fetchAlbums()
	};
};
