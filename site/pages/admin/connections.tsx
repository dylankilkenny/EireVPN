import React, { useState, useEffect } from 'react';
import { LayoutAdminDash } from '../../components/Layout';
import AdminSidePanel from '../../components/admin/AdminSidePanel';
import ConnectionsTable from '../../components/admin/tables/ConnectionsTable';
import ErrorMessage from '../../components/ErrorMessage';
import Pagination from '../../components/Pagination';
import CrudToolbar from '../../components/CrudToolbar';
import API from '../../service/APIService';
import useAsync from '../../hooks/useAsync';
import Router from 'next/router';

export default function Connections(): JSX.Element {
  const [offset, setOffset] = useState(0);
  const { data, loading, error } = useAsync(() => API.GetConnectionsList(offset), [offset]);
  const hasError = !!error;
  const pageLimit = 20;

  const handlePagination = (page_number: number) => {
    setOffset((page_number - 1) * pageLimit);
  };

  if (loading) {
    return <div></div>;
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <ConnectionsTable show={!hasError} connections={data?.connections} />
      <Pagination count={data?.count} handlePagination={handlePagination} pageLimit={pageLimit} />
      <ErrorMessage show={hasError} error={error} />
    </LayoutAdminDash>
  );
}
