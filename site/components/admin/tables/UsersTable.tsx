import React from 'react';
import Table from 'react-bootstrap/Table';
import User from '../../../interfaces/user';
import dayjs from 'dayjs';
import Router from 'next/router';

interface UsersTableProps {
  users: User[];
  show: boolean;
}

const UsersTable: React.FC<UsersTableProps> = ({ users, show }) => {
  if (!show) {
    return <div />;
  }
  return (
    <Table striped bordered hover responsive size="sm">
      <thead>
        <tr>
          <th>#</th>
          <th>Created</th>
          <th>First Name</th>
          <th>Last Name</th>
          <th>Email</th>
          <th>Stripe ID</th>
        </tr>
      </thead>
      <tbody className="table-admin-list">
        {users?.map(user => (
          <tr
            key={user.id}
            onClick={() => Router.push('/admin/users/[id]', '/admin/users/' + user.id)}
          >
            <td>{user.id}</td>
            <td>
              {dayjs(user.createdAt)
                .format('DD-MM-YYYY H:mm')
                .toString()}
            </td>
            <td>{user.firstname}</td>
            <td>{user.lastname}</td>
            <td>{user.email}</td>
            <td>{user.stripe_customer_id}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
};

export default UsersTable;
