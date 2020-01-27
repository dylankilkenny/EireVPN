import React from 'react';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import ServersTable from '../../../components/admin/tables/ServersTable';
import ErrorMessage from '../../../components/ErrorMessage';
import CrudToolbar from '../../../components/CrudToolbar';
import API from '../../../service/APIService';
import useAsync from '../../../hooks/useAsync';
import Router from 'next/router';

export default function Servers(): JSX.Element {
  const { data, loading, error } = useAsync(API.GetServersList);
  const hasError = !!error;

  const HandleCreate = () => {
    Router.push('/admin/servers/create');
  };

  if (loading) {
    return <div></div>;
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <CrudToolbar HandleCreate={HandleCreate} />
      <ServersTable show={!hasError} servers={data?.servers} />
      <ErrorMessage show={hasError} error={error} />
    </LayoutAdminDash>
  );
}
