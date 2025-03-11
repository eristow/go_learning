import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';
import type { Album } from '$types';
import { deleteAlbum, editAlbum, getAlbum } from '$lib/db';
import type { AlbumDTO } from '$types/Album';

export type PageLoad = {
	album: Promise<Album>;
};

export const load = async ({ fetch, params }): Promise<PageLoad> => {
	return {
		album: getAlbum(fetch, params.albumId)
	};
};

export const actions = {
	edit: async ({ request, fetch, params }) => {
		const data = await request.formData();

		const album: AlbumDTO = {
			title: (data.get('albumTitle') || '').toString(),
			artist: (data.get('albumArtist') || '').toString(),
			price: parseFloat((data.get('albumPrice') || '0').toString()),
			imageUrl: (data.get('albumArt') || '').toString()
		};

		const newAlbum: Album = await editAlbum(fetch, params.albumId, album);

		redirect(303, `/album/${newAlbum.id}`);
	},
	delete: async ({ fetch, params }) => {
		await deleteAlbum(fetch, params.albumId);

		redirect(303, '/');
	}
} satisfies Actions;
