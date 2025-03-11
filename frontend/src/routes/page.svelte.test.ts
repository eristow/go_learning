import { describe, test, expect } from 'vitest';
import '@testing-library/jest-dom/vitest';
import { render, screen } from '@testing-library/svelte';
import Page from './+page.svelte';

describe('/+page.svelte', () => {
	test('should render correctly with albums', async () => {
		const mockAlbums = Promise.resolve([
			{
				id: '1',
				title: 'Test Album',
				artist: 'Test Artist',
				price: 9.99,
				imageUrl: '/test.jpg'
			}
		]);

		render(Page, { props: { data: { albums: mockAlbums } } });

		expect(screen.getByText('Welcome to The Music Store!')).toBeInTheDocument();
		expect(screen.getByText('Loading...')).toBeInTheDocument();

		await screen.findByText('Test Album');
		expect(screen.getByText('Test Artist')).toBeInTheDocument();
		expect(screen.getByText('$9.99')).toBeInTheDocument();
	});

	test('should handle error state', async () => {
		const mockError = Promise.reject(new Error('Failed to load albums'));

		render(Page, { props: { data: { albums: mockError } } });

		await screen.findByText('Failed to load albums');
	});
});
