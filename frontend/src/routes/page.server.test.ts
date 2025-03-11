import { describe, test, expect, vi } from 'vitest';
import { load } from './+page.server';
import { getAlbums } from '$lib/db/index.js';

vi.mock('$lib/db/index.js', () => {
	return {
		getAlbums: vi.fn()
	};
});

describe('+page.server.ts', () => {
	test('should load albums data', async () => {
		const mockAlbums = [
			{ id: '1', title: 'Album 1', artist: 'Artist 1', price: 9.99, imageUrl: '/album1.jpg' }
		];
		vi.mocked(getAlbums).mockResolvedValue(mockAlbums);

		const mockFetch = vi.fn();

		const result = await load({ fetch: mockFetch } as never);

		expect(getAlbums).toHaveBeenCalledWith(mockFetch);

		expect(result).toHaveProperty('albums');

		const albums = await result.albums;
		expect(albums).toEqual(mockAlbums);
	});
});
