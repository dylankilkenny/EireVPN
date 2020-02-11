import React from 'react';
import Table from 'react-bootstrap/Table';
import UserPlan from '../../../interfaces/userplan';
import dayjs from 'dayjs';
import Router from 'next/router';

interface UserPlansTableProps {
  userplans: UserPlan[];
  show: boolean;
}

const UsersTable: React.FC<UserPlansTableProps> = ({ userplans, show }) => {
  if (!show) {
    return <div />;
  }
  return (
    <Table className="table-admin-list" striped bordered hover responsive size="sm">
      <thead>
        <tr>
          <th>#</th>
          <th>Created</th>
          <th>User ID</th>
          <th>Plan ID</th>
          <th>Active</th>
          <th>Start Date</th>
          <th>End Date</th>
        </tr>
      </thead>
      <tbody className="table-admin-list">
        {userplans.map(up => (
          <tr
            key={up.id}
            onClick={() =>
              Router.push('/admin/userplans/[user_id]', '/admin/userplans/' + up.user_id)
            }
          >
            <td>{up.id}</td>
            <td>
              {dayjs(up.createdAt)
                .format('DD-MM-YYYY H:mm')
                .toString()}
            </td>
            <td>{up.user_id}</td>
            <td>{up.plan_id}</td>
            <td>{up.active.toString()}</td>
            <td>
              {dayjs(up.start_date)
                .format('DD-MM-YYYY H:mm')
                .toString()}
            </td>
            <td>
              {dayjs(up.expiry_date)
                .format('DD-MM-YYYY H:mm')
                .toString()}
            </td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
};

export default UsersTable;
