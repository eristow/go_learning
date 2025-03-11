import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render } from '@testing-library/svelte';
import '@testing-library/jest-dom/vitest';
import AlbumDetail from './+page.svelte';
import type { Album } from '$types';

const mocks = vi.hoisted(() => {
	return {
		mockPage: {
			params: {
				albumId: '123'
			}
		}
	};
});

vi.mock('$app/state', () => ({
	page: mocks.mockPage
}));

describe('Album Detail Page', () => {
	describe('Loading state', () => {
		it('should show loading indicator when album is pending', async () => {
			const albumPromise = new Promise<Album>(() => {});

			document.body.innerHTML = '<div id="app"></div>';

			render(AlbumDetail, {
				props: {
					data: {
						album: albumPromise
					}
				},
				target: document.getElementById('app')!
			});

			await new Promise((resolve) => setTimeout(resolve, 0));

			expect(document.body.textContent).toContain('Loading...');
		});
	});

	describe('Loaded state', () => {
		const mockAlbum = {
			id: '123',
			title: 'Test Album',
			artist: 'Test Artist',
			price: 12.99,
			imageUrl: 'http://example.com/image.jpg'
		};

		beforeEach(async () => {
			document.body.innerHTML = '<div id="app"></div>';

			const albumPromise = Promise.resolve(mockAlbum);

			render(AlbumDetail, {
				props: {
					data: {
						album: albumPromise
					}
				},
				target: document.getElementById('app')!
			});

			await albumPromise;

			await new Promise((resolve) => setTimeout(resolve, 0));
		});

		it('should display album details when loaded', () => {
			const content = document.body.textContent || '';
			expect(content).toContain('Test Album');
			expect(content).toContain('Test Artist');
			expect(content).toContain('$12.99');
		});

		it('should have edit and delete buttons', () => {
			const content = document.body.innerHTML || '';
			expect(content).toContain('Edit');
			expect(content).toContain('Delete');
			expect(content).toContain('/album/123/edit');
			expect(content).toContain('/album/123?/delete');
			expect(content).toContain('_method');
			expect(content).toContain('DELETE');
		});
	});

	describe('Error state', () => {
		it('should display error message when album loading fails', async () => {
			document.body.innerHTML = '<div id="app"></div>';

			const errorMessage = 'Failed to load album';
			const errorPromise = Promise.reject(new Error(errorMessage));

			render(AlbumDetail, {
				props: {
					data: {
						album: errorPromise
					}
				},
				target: document.getElementById('app')!
			});

			errorPromise.catch(() => {});

			await new Promise((resolve) => setTimeout(resolve, 0));

			expect(document.body.textContent).toContain(errorMessage);
		});
	});
});
