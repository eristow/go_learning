import type { Album, AlbumDTO } from '$types/Album';
import { redirect } from '@sveltejs/kit';
import type { Actions } from './$types';
import { createAlbum } from '$lib/db';

export const actions = {
	create: async ({ request, fetch }) => {
		const data = await request.formData();

		const album: AlbumDTO = {
			title: (data.get('albumTitle') || '').toString(),
			artist: (data.get('albumArtist') || '').toString(),
			price: parseFloat((data.get('albumPrice') || '0').toString()),
			imageUrl: (data.get('albumArt') || '').toString()
		};

		const newAlbum: Album = await createAlbum(fetch, album);

		redirect(303, `/album/${newAlbum.id}`);
	}
} satisfies Actions;
