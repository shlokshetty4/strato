import React from 'react';
import UsersTable from './components/UsersTable';

function App() {
  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Users</h1>
      <UsersTable />
    </div>
  );
}

export default App;
