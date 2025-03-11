import { render, fireEvent } from '@testing-library/svelte';
import { describe, expect, vi, test } from 'vitest';
import Button from './Button.svelte';

describe('Button', () => {
	test('renders with correct text', () => {
		const { getByRole } = render(Button, { props: { children: () => 'Click Me', type: 'button' } });

		const button = getByRole('button');

		expect(button).toBeInTheDocument();
	});

	test('calls onClick when clicked', async () => {
		const mock = vi.fn();
		const { getByRole } = render(Button, {
			props: {
				children: () => 'Click Me',
				type: 'button',
				onClick: mock
			}
		});

		const button = getByRole('button');
		await fireEvent.click(button);
		expect(mock).toHaveBeenCalledTimes(1);
	});

	test('renders with correct type', () => {
		const { getByRole } = render(Button, { props: { children: () => 'Click Me', type: 'submit' } });
		const button = getByRole('button');
		expect(button).toHaveAttribute('type', 'submit');
	});
});
