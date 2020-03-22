import React, { useState } from 'react';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import ErrorMessage from '../../../components/ErrorMessage';
import { LayoutAdminDash } from '../../../components/Layout';
import UsersTable from '../../../components/admin/tables/UsersTable';
import useAsync from '../../../hooks/useAsync';
import API from '../../../service/APIService';
import CrudToolbar from '../../../components/CrudToolbar';
import Pagination from '../../../components/Pagination';
import Router from 'next/router';

export default function Users(): JSX.Element {
  const [offset, setOffset] = useState(0);
  const { data, loading, error } = useAsync(() => API.GetUsersList(offset), [offset]);
  const hasError = !!error;
  const pageLimit = 20;

  const HandleCreate = () => {
    Router.push('/admin/users/create');
  };

  const handlePagination = (page_number: number) => {
    setOffset((page_number - 1) * pageLimit);
  };

  if (loading) {
    return <div></div>;
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <CrudToolbar HandleCreate={HandleCreate} />
      <UsersTable show={!hasError} users={data?.users} />
      <Pagination count={data?.count} handlePagination={handlePagination} pageLimit={pageLimit} />
      <ErrorMessage show={hasError} error={error} />
    </LayoutAdminDash>
  );
}
