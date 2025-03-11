import { describe, test, expect, vi, beforeEach } from 'vitest';
import { actions } from './+page.server';
import { createAlbum } from '$lib/db';
import { redirect } from '@sveltejs/kit';

// Mock dependencies
vi.mock('$lib/db', () => ({
	createAlbum: vi.fn()
}));

vi.mock('@sveltejs/kit', () => ({
	redirect: vi.fn()
}));

describe('+page.server.ts', () => {
	beforeEach(() => {
		vi.clearAllMocks();
	});

	test('should create album with valid form data', async () => {
		// Setup form data
		const formData = new FormData();
		formData.append('albumTitle', 'Test Album');
		formData.append('albumArtist', 'Test Artist');
		formData.append('albumPrice', '12.99');
		formData.append('albumArt', 'http://example.com/image.jpg');

		const mockRequest = {
			formData: () => Promise.resolve(formData)
		};

		const mockFetch = vi.fn();

		const mockNewAlbum = {
			id: '123',
			title: 'Test Album',
			artist: 'Test Artist',
			price: 12.99,
			imageUrl: 'http://example.com/image.jpg'
		};
		vi.mocked(createAlbum).mockResolvedValue(mockNewAlbum);

		await actions.create({ request: mockRequest, fetch: mockFetch } as never);

		expect(createAlbum).toHaveBeenCalledWith(mockFetch, {
			title: 'Test Album',
			artist: 'Test Artist',
			price: 12.99,
			imageUrl: 'http://example.com/image.jpg'
		});

		expect(redirect).toHaveBeenCalledWith(303, '/album/123');
	});

	test('should handle missing form values with defaults', async () => {
		const formData = new FormData();
		formData.append('albumTitle', 'Test Album');
		// Omitting other fields to test defaults

		const mockRequest = {
			formData: () => Promise.resolve(formData)
		};

		const mockFetch = vi.fn();

		const mockNewAlbum = { id: '456', title: 'Test Album', artist: '', price: 0, imageUrl: '' };
		vi.mocked(createAlbum).mockResolvedValue(mockNewAlbum);

		await actions.create({ request: mockRequest, fetch: mockFetch } as never);

		expect(createAlbum).toHaveBeenCalledWith(mockFetch, {
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
		vi.mocked(createAlbum).mockResolvedValue(mockNewAlbum);

		await actions.create({ request: mockRequest, fetch: mockFetch } as never);

		const calledArgument = vi.mocked(createAlbum).mock.calls[0][1];
		expect(calledArgument.title).toBe('Test Album');
		expect(calledArgument.artist).toBe('Test Artist');
		expect(Number.isNaN(calledArgument.price)).toBe(true);
		expect(calledArgument.imageUrl).toBe('http://example.com/image.jpg');

		expect(redirect).toHaveBeenCalledWith(303, '/album/789');
	});
});
