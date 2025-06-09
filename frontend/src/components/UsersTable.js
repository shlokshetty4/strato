import React, { useEffect, useState } from 'react';

export default function UsersTable() {
  const [users, setUsers] = useState([]);
  const [filterMfa, setFilterMfa] = useState('');

  useEffect(() => {
    fetch('/api/users')
      .then((res) => res.json())
      .then(setUsers)
      .catch(console.error);
  }, []);

  const filtered = users.filter(u => {
    if (filterMfa === 'enabled') return u.mfaEnabled;
    if (filterMfa === 'disabled') return !u.mfaEnabled;
    return true;
  });

  return (
    <div>
      <div className="mb-4">
        <label className="mr-2">Filter MFA:</label>
        <select
          value={filterMfa}
          onChange={e => setFilterMfa(e.target.value)}
          className="border px-2 py-1"
        >
          <option value="">All</option>
          <option value="enabled">Enabled</option>
          <option value="disabled">Disabled</option>
        </select>
      </div>

      <table className="min-w-full bg-white border">
        <thead>
          <tr>
            <th className="px-4 py-2 border">Human User</th>
            <th className="px-4 py-2 border">Create Date</th>
            <th className="px-4 py-2 border">Password Changed Date</th>
            <th className="px-4 py-2 border">Days since last password change</th>
            <th className="px-4 py-2 border">Last Access Date</th>
            <th className="px-4 py-2 border">Days since Last Access</th>
            <th className="px-4 py-2 border">MFA Enabled</th>
          </tr>
        </thead>
        <tbody>
          {filtered.map(u => {
            const warnPwd = u.daysSincePasswordChange > 365;
            const warnAccess = u.daysSinceLastAccess > 90;
            return (
              <tr key={u.id} className={warnPwd || warnAccess ? 'bg-white-100' : ''}>
                <td className="px-4 py-2 border">{u.name}</td>
                <td className="px-4 py-2 border">{u.createDate}</td>
                <td className="px-4 py-2 border">{u.passwordChangedDate}</td>
                <td className="px-4 py-2 border">{u.daysSincePasswordChange}</td>
                <td className="px-4 py-2 border">{u.lastAccessDate}</td>
                <td className="px-4 py-2 border">{u.daysSinceLastAccess}</td>
                <td className="px-4 py-2 border">{u.mfaEnabled ? 'Yes' : 'No'}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}
