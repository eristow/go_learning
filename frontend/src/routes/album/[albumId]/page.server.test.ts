import { describe, test, expect, vi, beforeEach } from 'vitest';
import { actions } from './+page.server';
import { deleteAlbum, editAlbum } from '$lib/db';
import { redirect } from '@sveltejs/kit';

// Mock dependencies
vi.mock('$lib/db', () => ({
	editAlbum: vi.fn(),
	deleteAlbum: vi.fn()
}));

vi.mock('@sveltejs/kit', () => ({
	redirect: vi.fn()
}));

describe('+page.server.ts', () => {
	beforeEach(() => {
		vi.clearAllMocks();
	});

	test('should edit album with valid form data', async () => {
		const formData = new FormData();
		formData.append('albumTitle', 'Test Album Edited');
		formData.append('albumArtist', 'Test Artist Edited');
		formData.append('albumPrice', '15.99');
		formData.append('albumArt', 'http://example.com/image-edited.jpg');

		const mockRequest = {
			formData: () => Promise.resolve(formData)
		};

		const mockFetch = vi.fn();

		const mockNewAlbum = {
			id: '123',
			title: 'Test Album Edited',
			artist: 'Test Artist Edited',
			price: 15.99,
			imageUrl: 'http://example.com/image-edited.jpg'
		};
		vi.mocked(editAlbum).mockResolvedValue(mockNewAlbum);

		await actions.edit({
			request: mockRequest,
			params: { albumId: '123' },
			fetch: mockFetch
		} as never);

		expect(redirect).toHaveBeenCalledWith(303, '/album/123');
	});

	test('should handle missing form values with defaults', async () => {
		const formData = new FormData();
		formData.append('albumTitle', 'Test Album');

		const mockRequest = {
			formData: () => Promise.resolve(formData)
		};

		const mockFetch = vi.fn();

		const mockNewAlbum = { id: '456', title: 'Test Album', artist: '', price: 0, imageUrl: '' };
		vi.mocked(editAlbum).mockResolvedValue(mockNewAlbum);

		await actions.edit({
			request: mockRequest,
			params: { albumId: '456' },
			fetch: mockFetch
		} as never);

		expect(editAlbum).toHaveBeenCalledWith(mockFetch, '456', {
			title: 'Test Album',
			artist: '',
			price: 0,
			imageUrl: ''
		});

		expect(redirect).toHaveBeenCalledWith(303, '/album/456');
	});

	test('should handle invalid price input', async () => {
		const formData = new FormData();
		formData.append('albumTitle', 'Test Album');
		formData.append('albumArtist', 'Test Artist');
		formData.append('albumPrice', 'not-a-number');
		formData.append('albumArt', 'http://example.com/image.jpg');

		const mockRequest = {
			formData: () => Promise.resolve(formData)
		};

		const mockFetch = vi.fn();

		const mockNewAlbum = {
			id: '789',
			title: 'Test Album',
			artist: 'Test Artist',
			price: NaN,
			imageUrl: 'http://example.com/image.jpg'
		};
		vi.mocked(editAlbum).mockResolvedValue(mockNewAlbum);

		await actions.edit({
			request: mockRequest,
			params: { albumId: '789' },
			fetch: mockFetch
		} as never);

		const calledArgument = vi.mocked(editAlbum).mock.calls[0][2];
		expect(calledArgument.title).toBe('Test Album');
		expect(calledArgument.artist).toBe('Test Artist');
		expect(Number.isNaN(calledArgument.price)).toBe(true);
		expect(calledArgument.imageUrl).toBe('http://example.com/image.jpg');

		expect(redirect).toHaveBeenCalledWith(303, '/album/789');
	});

	test('should delete album', async () => {
		const mockFetch = vi.fn();

		await actions.delete({
			params: { albumId: '123' },
			fetch: mockFetch
		} as never);

		expect(deleteAlbum).toHaveBeenCalledWith(mockFetch, '123');
		expect(redirect).toHaveBeenCalledWith(303, '/');
	});
});
