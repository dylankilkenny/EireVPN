import React from 'react';
import Table from 'react-bootstrap/Table';
import Connection from '../../../interfaces/connection';
import dayjs from 'dayjs';
import Router from 'next/router';

interface ConnTableProps {
  connections: Connection[];
  show: boolean;
}

const ConnectionsTable: React.FC<ConnTableProps> = ({ connections, show }) => {
  if (!show) {
    return <div />;
  }
  return (
    <Table striped bordered hover responsive size="sm">
      <thead>
        <tr>
          <th>#</th>
          <th>Created</th>
          <th>User ID</th>
          <th>Server ID</th>
          <th>Country</th>
        </tr>
      </thead>
      <tbody className="table-admin-list">
        {connections.map(c => (
          <tr key={c.id}>
            <td>{c.id}</td>
            <td>
              {dayjs(c.createdAt)
                .format('DD-MM-YYYY H:mm')
                .toString()}
            </td>
            <td>{c.user_id}</td>
            <td>{c.server_id}</td>
            <td>{c.server_country}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
};

export default ConnectionsTable;
