import { describe, test, expect, vi, beforeEach } from 'vitest';
import '@testing-library/jest-dom/vitest';
import { render, screen } from '@testing-library/svelte';
import Layout from './+layout.svelte';
import { createRawSnippet } from 'svelte';

const mockChild = createRawSnippet(() => {
	return {
		render: () => '<div>Child content</div>',
		setup: () => {}
	};
});

const mocks = vi.hoisted(() => {
	return {
		mockPage: {
			url: {
				pathname: '/'
			}
		}
	};
});

vi.mock('$app/state', () => ({
	page: mocks.mockPage
}));

describe('/+layout.svelte', () => {
	beforeEach(() => {
		mocks.mockPage.url.pathname = '/';
	});

	test('should render header and children', () => {
		render(Layout, {
			props: {
				children: mockChild
			}
		});

		expect(screen.getByText('The Music Store')).toBeInTheDocument();
		expect(screen.getByText('Home')).toBeInTheDocument();
		expect(screen.getByText('Add Album')).toBeInTheDocument();
		expect(screen.getByText('Child content')).toBeInTheDocument();
	});

	test('should hide "Add Album" button on non-home pages', () => {
		mocks.mockPage.url.pathname = '/album/123';

		render(Layout, {
			props: {
				children: mockChild
			}
		});

		expect(screen.getByText('The Music Store')).toBeInTheDocument();
		expect(screen.getByText('Home')).toBeInTheDocument();
		expect(screen.queryByText('Add Album')).not.toBeInTheDocument();
	});
});
