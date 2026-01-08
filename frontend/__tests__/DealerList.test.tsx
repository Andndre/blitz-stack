import { render, screen } from '@testing-library/react';
import DealerList from '../components/DealerList';
import { Dealer } from '../types';

describe('DealerList Component', () => {
  it('renders "Tidak ada data" message when list is empty', () => {
    render(<DealerList dealers={[]} />);
    const message = screen.getByText(/Tidak ada data dealer ditemukan/i);
    expect(message).toBeInTheDocument();
  });

  it('renders list of dealers correctly', () => {
    const mockDealers: Dealer[] = [
      { id: 1, name: 'Dealer Alpha', city: 'Jakarta' },
      { id: 2, name: 'Dealer Beta', city: 'Bandung' },
    ];

    render(<DealerList dealers={mockDealers} />);

    // Check if names are rendered
    expect(screen.getByText('Dealer Alpha')).toBeInTheDocument();
    expect(screen.getByText('Dealer Beta')).toBeInTheDocument();

    // Check if cities are rendered
    expect(screen.getByText('Jakarta')).toBeInTheDocument();
    expect(screen.getByText('Bandung')).toBeInTheDocument();

    // Check if correct number of items rendered
    const items = screen.getAllByTestId('dealer-item');
    expect(items).toHaveLength(2);
  });
});
