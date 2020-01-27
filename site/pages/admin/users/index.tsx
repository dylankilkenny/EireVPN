import React from 'react';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import ErrorMessage from '../../../components/ErrorMessage';
import { LayoutAdminDash } from '../../../components/Layout';
import UsersTable from '../../../components/admin/tables/UsersTable';
import useAsync from '../../../hooks/useAsync';
import API from '../../../service/APIService';
import CrudToolbar from '../../../components/CrudToolbar';
import Router from 'next/router';

export default function Users(): JSX.Element {
  const { data, loading, error } = useAsync(API.GetUsersList);
  const hasError = !!error;

  const HandleCreate = () => {
    Router.push('/admin/users/create');
  };

  if (loading) {
    return <div></div>;
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <CrudToolbar HandleCreate={HandleCreate} />
      <UsersTable show={!hasError} users={data?.users} />
      <ErrorMessage show={hasError} error={error} />
    </LayoutAdminDash>
  );
}
