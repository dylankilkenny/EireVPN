import React from 'react';
import Table from 'react-bootstrap/Table';
import Server from '../../../interfaces/server';
import dayjs from 'dayjs';
import Router from 'next/router';

interface ServersTableProps {
  servers: Server[];
  show: boolean;
}

const ServersTable: React.FC<ServersTableProps> = ({ servers, show }) => {
  if (!show) {
    return <div />;
  }
  return (
    <Table striped bordered hover responsive size="sm">
      <thead>
        <tr>
          <th>#</th>
          <th>Created</th>
          <th>Country</th>
          <th>Code</th>
          <th>Type</th>
          <th>IP</th>
          <th>Port</th>
          <th>Username</th>
          <th>Password</th>
        </tr>
      </thead>
      <tbody className="table-admin-list">
        {servers.map(server => (
          <tr
            key={server.id}
            onClick={() => Router.push('/admin/servers/[id]', '/admin/servers/' + server.id)}
          >
            <td>{server.id}</td>
            <td>
              {dayjs(server.createdAt)
                .format('DD-MM-YYYY H:mm')
                .toString()}
            </td>
            <td>{server.country}</td>
            <td>{server.country_code}</td>
            <td>{server.type}</td>
            <td>{server.ip}</td>
            <td>{server.port}</td>
            <td>{server.username}</td>
            <td>{server.password}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
};

export default ServersTable;
