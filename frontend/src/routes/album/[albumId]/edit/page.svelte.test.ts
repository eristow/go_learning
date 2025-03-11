import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render } from '@testing-library/svelte';
import '@testing-library/jest-dom/vitest';
import EditAlbum from './+page.svelte';
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

describe('Album Edit Page', () => {
	describe('Loading state', () => {
		it('should show loading indicator when album is pending', async () => {
			const albumPromise = new Promise<Album>(() => {});

			document.body.innerHTML = '<div id="app"></div>';

			render(EditAlbum, {
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

			render(EditAlbum, {
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
			const titleInput = document.getElementById('albumTitle') as HTMLInputElement;
			const artistInput = document.getElementById('albumArtist') as HTMLInputElement;
			const priceInput = document.getElementById('albumPrice') as HTMLInputElement;
			const imageInput = document.getElementById('albumArt') as HTMLInputElement;

			expect(titleInput).not.toBeNull();
			expect(artistInput).not.toBeNull();
			expect(priceInput).not.toBeNull();
			expect(imageInput).not.toBeNull();

			expect(titleInput.value).toBe(mockAlbum.title);
			expect(artistInput.value).toBe(mockAlbum.artist);
			expect(priceInput.value).toBe(mockAlbum.price.toString());
			expect(imageInput.value).toBe(mockAlbum.imageUrl);
		});

		it('should have edit button', () => {
			const content = document.body.innerHTML || '';
			expect(content).toContain('Edit Album');
		});
	});

	describe('Error state', () => {
		it('should display error message when album loading fails', async () => {
			document.body.innerHTML = '<div id="app"></div>';

			const errorMessage = 'Failed to load album';
			const errorPromise = Promise.reject(new Error(errorMessage));

			render(EditAlbum, {
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
