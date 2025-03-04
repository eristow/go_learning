import type { Album } from '$types';
import type { AlbumDTO } from '$types/Album';
import { PUBLIC_BACKEND_URL } from '$env/static/public';

type Fetch = (input: RequestInfo, init?: RequestInit) => Promise<Response>;

const backendUrl = `${PUBLIC_BACKEND_URL}/albums`;

export const getAlbums = async (fetch: Fetch): Promise<Album[]> => {
	try {
		const albumsRes = await fetch(backendUrl);

		if (albumsRes.status !== 200) {
			throw new Error('Failed to get albums');
		}

		const albumsData: Album[] = await albumsRes.json();

		return albumsData;
	} catch (e) {
		console.error(e);
		return [];
	}
};

export const getAlbum = async (fetch: Fetch, albumId: string): Promise<Album> => {
	try {
		const albumRes = await fetch(`${backendUrl}/${albumId}`);
		console.log(albumRes);

		if (albumRes.status !== 200) {
			throw new Error('Failed to get album');
		}

		const albumData: Album = await albumRes.json();

		return albumData;
	} catch (e) {
		console.error(e);
		return {} as Album;
	}
};

export const createAlbum = async (fetch: Fetch, album: AlbumDTO): Promise<Album> => {
	try {
		const albumRes = await fetch(backendUrl, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(album)
		});

		if (albumRes.status !== 201) {
			throw new Error('Failed to create album');
		}

		const albumData: Album = await albumRes.json();

		return albumData;
	} catch (e) {
		console.error(e);
		return {} as Album;
	}
};

export const editAlbum = async (fetch: Fetch, albumId: string, album: AlbumDTO): Promise<Album> => {
	try {
		const albumRes = await fetch(`${backendUrl}/${albumId}`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(album)
		});

		if (albumRes.status !== 200) {
			throw new Error('Failed to edit album');
		}

		const albumData: Album = await albumRes.json();

		return albumData;
	} catch (e) {
		console.error(e);
		return {} as Album;
	}
};

export const deleteAlbum = async (fetch: Fetch, albumId: string): Promise<Album> => {
	try {
		const albumRes = await fetch(`${backendUrl}/${albumId}`, {
			method: 'DELETE'
		});

		if (albumRes.status !== 200) {
			throw new Error('Failed to delete album');
		}

		const albumData: Album = await albumRes.json();

		return albumData;
	} catch (e) {
		console.error(e);
		return {} as Album;
	}
};
