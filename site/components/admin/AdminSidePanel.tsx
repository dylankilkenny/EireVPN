import React from 'react';
import Link from 'next/link';

export default function AdminSidePanel(): JSX.Element {
  return (
    <div>
      <ul>
        <li>
          <Link href="/admin/users">
            <a>Users</a>
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
          <Link href="/admin/settings">
            <a>Settings</a>
          </Link>
        </li>
        <style jsx>{`
          ul {
            list-style-type: none;
          }
        `}</style>
      </ul>
    </div>
  );
}
