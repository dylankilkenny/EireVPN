import React from 'react';
import Table from 'react-bootstrap/Table';
import Plan from '../../../interfaces/plan';
import dayjs from 'dayjs';
import Router from 'next/router';

interface PlansTableProps {
  plans: Plan[];
  show: boolean;
}

const UsersTable: React.FC<PlansTableProps> = ({ plans, show }) => {
  if (!show) {
    return <div />;
  }
  return (
    <Table className="table-admin-list" striped bordered hover responsive size="sm">
      <thead>
        <tr>
          <th>#</th>
          <th>Created</th>
          <th>Name</th>
          <th>Cost</th>
          <th>Interval</th>
          <th>Interval Count</th>
          <th>Stripe Plan ID</th>
          <th>Stripe Product ID</th>
        </tr>
      </thead>
      <tbody className="table-admin-list">
        {plans.map(plan => (
          <tr
            key={plan.id}
            onClick={() => Router.push('/admin/plans/[id]', '/admin/plans/' + plan.id)}
          >
            <td>{plan.id}</td>
            <td>
              {dayjs(plan.createdAt)
                .format('DD-MM-YYYY H:mm')
                .toString()}
            </td>
            <td>{plan.name}</td>
            <td>{plan.amount}</td>
            <td>{plan.interval}</td>
            <td>{plan.interval_count}</td>
            <td>{plan.stripe_plan_id}</td>
            <td>{plan.stripe_product_id}</td>
          </tr>
        ))}
      </tbody>
    </Table>
  );
};

export default UsersTable;
