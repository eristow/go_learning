import { describe, test, expect, vi } from 'vitest';
import { load } from './+page.server';
import { getAlbum } from '$lib/db/index.js';

vi.mock('$lib/db/index.js', () => {
	return {
		getAlbum: vi.fn()
	};
});

describe('+page.server.ts', () => {
	test('should load album data', async () => {
		const mockAlbum = {
			id: '1',
			title: 'Album 1',
			artist: 'Artist 1',
			price: 9.99,
			imageUrl: '/album1.jpg'
		};
		vi.mocked(getAlbum).mockResolvedValue(mockAlbum);

		const mockFetch = vi.fn();

		const result = await load({
			fetch: mockFetch,
			params: {
				albumId: '1'
			}
		} as never);

		expect(getAlbum).toHaveBeenCalledWith(mockFetch, '1');

		expect(result).toHaveProperty('album');

		const albums = await result.album;
		expect(albums).toEqual(mockAlbum);
	});
});
