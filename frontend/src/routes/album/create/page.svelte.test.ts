import { describe, it, expect, beforeEach } from 'vitest';
import { render, screen } from '@testing-library/svelte';
import '@testing-library/jest-dom/vitest';
import AlbumCreate from './+page.svelte';

describe('Album Create Page', () => {
	beforeEach(() => {
		render(AlbumCreate);
	});

	it('should render the page title', () => {
		expect(screen.getByText('Album Create')).toBeInTheDocument();
	});

	it('should have a form with correct action and method', () => {
		const form = screen.getByTestId('form');
		expect(form).toHaveAttribute('action', '/album?/create');
		expect(form).toHaveAttribute('method', 'POST');
	});

	it('should have album title input field', () => {
		const titleLabel = screen.getByText('Album Title:');
		expect(titleLabel).toBeInTheDocument();

		const titleInput = screen.getByLabelText('Album Title:');
		expect(titleInput).toHaveAttribute('type', 'text');
		expect(titleInput).toHaveAttribute('id', 'albumTitle');
		expect(titleInput).toHaveAttribute('name', 'albumTitle');
		expect(titleInput).toHaveAttribute('required');
	});

	it('should have artist input field', () => {
		const artistLabel = screen.getByText('Artist:');
		expect(artistLabel).toBeInTheDocument();

		const artistInput = screen.getByLabelText('Artist:');
		expect(artistInput).toHaveAttribute('type', 'text');
		expect(artistInput).toHaveAttribute('id', 'albumArtist');
		expect(artistInput).toHaveAttribute('name', 'albumArtist');
		expect(artistInput).toHaveAttribute('required');
	});

	it('should have price input field', () => {
		const priceLabel = screen.getByText('Price:');
		expect(priceLabel).toBeInTheDocument();

		const priceInput = screen.getByLabelText('Price:');
		expect(priceInput).toHaveAttribute('type', 'number');
		expect(priceInput).toHaveAttribute('step', '0.01');
		expect(priceInput).toHaveAttribute('id', 'albumPrice');
		expect(priceInput).toHaveAttribute('name', 'albumPrice');
		expect(priceInput).toHaveAttribute('required');
	});

	it('should have album art URL input field', () => {
		const artLabel = screen.getByText('Album Art URL:');
		expect(artLabel).toBeInTheDocument();

		const artInput = screen.getByLabelText('Album Art URL:');
		expect(artInput).toHaveAttribute('type', 'text');
		expect(artInput).toHaveAttribute('id', 'albumArt');
		expect(artInput).toHaveAttribute('name', 'albumArt');
		expect(artInput).toHaveAttribute('required');
	});

	it('should have a submit button', () => {
		const submitButton = screen.getByRole('button', { name: 'Create Album' });
		expect(submitButton).toBeInTheDocument();
		expect(submitButton).toHaveAttribute('type', 'submit');
	});
});
