import React from 'react';
import Link from 'next/link';

export default function AdminSidePanel(): JSX.Element {
  return (
    <div className="admin-side-panel">
      <ul>
        <li>
          <Link href="/admin/users">
            <a>Users</a>
          </Link>
        </li>
        <li>
          <Link href="/admin/userplans">
            <a>User Plans</a>
          </Link>
        </li>
        <li>
          <Link href="/admin/plans">
            <a>Plans</a>
          </Link>
        </li>
        <li>
          <Link href="/admin/servers">
            <a>Servers</a>
          </Link>
        </li>
        <li>
          <Link href="/admin/connections">
            <a>Connections</a>
          </Link>
        </li>
        <li>
          <Link href="/admin/settings">
            <a>Settings</a>
          </Link>
        </li>
        <style jsx>{`
          ul {
            list-style-type: none;
          }
          a {
            text-decoration: none;
            color: #5f6368;
          }
          a:hover {
            opacity: 0.6;
          }
        `}</style>
      </ul>
    </div>
  );
}
