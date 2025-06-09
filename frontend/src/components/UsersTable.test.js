import React from 'react';
import { render, screen, act } from '@testing-library/react';
import UsersTable from './UsersTable';

describe('UsersTable component with sample data', () => {
  beforeEach(() => {
    global.fetch = jest.fn(() =>
      Promise.resolve({
        json: () => Promise.resolve([
          {
            name: 'Foo Bar1',
            createDate: 'Oct 1 2020',
            passwordChangedDate: 'Oct 1 2021',
            daysSincePasswordChange: 1325,
            lastAccessDate: 'Jan 4 2025',
            daysSinceLastAccess: 134,
            mfaEnabled: true,
          },
          {
            name: 'Foo1 Bar1',
            createDate: 'Sep 20 2019',
            passwordChangedDate: 'Sep 22 2019',
            daysSincePasswordChange: 2064,
            lastAccessDate: 'Feb 8 2025',
            daysSinceLastAccess: 99,
            mfaEnabled: false,
          },
          {
            name: 'Foo2 Bar2',
            createDate: 'Feb 3 2022',
            passwordChangedDate: 'Feb 3 2022',
            daysSincePasswordChange: 1200,
            lastAccessDate: 'Feb 12 2025',
            daysSinceLastAccess: 95,
            mfaEnabled: false,
          },
          {
            name: 'Foo3 Bar3',
            createDate: 'Mar 7 2023',
            passwordChangedDate: 'Mar 10 2023',
            daysSincePasswordChange: 801,
            lastAccessDate: 'Jan 3 2022',
            daysSinceLastAccess: 1232,
            mfaEnabled: true,
          },
          {
            name: 'Foo Bar4',
            createDate: 'Apr 8 2018',
            passwordChangedDate: 'Apr 12 2020',
            daysSincePasswordChange: 1862,
            lastAccessDate: 'Oct 4 2022',
            daysSinceLastAccess: 956,
            mfaEnabled: false,
          },
        ]),
      })
    );
  });

  it('renders all five sample users', async () => {
    await act(async () => render(<UsersTable />));

    expect(screen.getByText('Foo Bar1')).toBeInTheDocument();
    expect(screen.getByText('Oct 1 2020')).toBeInTheDocument();
    expect(screen.getByText('1325')).toBeInTheDocument();

    expect(screen.getByText('Foo Bar4')).toBeInTheDocument();
    expect(screen.getByText('Apr 12 2020')).toBeInTheDocument();
    expect(screen.getByText('956')).toBeInTheDocument();
  });
});