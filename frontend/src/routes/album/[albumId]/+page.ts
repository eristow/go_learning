import type { Album } from '$types';

export type PageLoad = {
	album: Promise<Album>;
};

export const load = async ({ fetch, params }): Promise<PageLoad> => {
	const fetchAlbum = async (): Promise<Album> => {
		try {
			// TODO: change the backend URL to an env var
			const albumRes = await fetch(`http://localhost:8080/albums/${params.albumId}`);
			const albumData = await albumRes.json();
			return albumData;
		} catch (e) {
			console.error(e);
			return {} as Album;
		}
	};

	return {
		album: fetchAlbum()
	};
};
