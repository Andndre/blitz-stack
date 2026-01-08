import { render, screen } from '@testing-library/react';
import ItemList from '../components/ItemList';
import { Item } from '../types';

describe('ItemList Component', () => {
  it('renders "No items found" message when list is empty', () => {
    render(<ItemList items={[]} />);
    const message = screen.getByText(/No items found/i);
    expect(message).toBeInTheDocument();
  });

  it('renders list of items correctly', () => {
    const mockItems: Item[] = [
      { id: 1, title: 'Item Alpha', description: 'Desc A' },
      { id: 2, title: 'Item Beta', description: 'Desc B' },
    ];

    render(<ItemList items={mockItems} />);

    expect(screen.getByText('Item Alpha')).toBeInTheDocument();
    expect(screen.getByText('Item Beta')).toBeInTheDocument();
    expect(screen.getByText('Desc A')).toBeInTheDocument();
    expect(screen.getByText('Desc B')).toBeInTheDocument();

    const items = screen.getAllByTestId('item-row');
    expect(items).toHaveLength(2);
  });
});
