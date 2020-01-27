import React from 'react';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import PlansTable from '../../../components/admin/tables/PlansTable';
import ErrorMessage from '../../../components/ErrorMessage';
import API from '../../../service/APIService';
import useAsync from '../../../hooks/useAsync';
import CrudToolbar from '../../../components/CrudToolbar';
import Router from 'next/router';

export default function Plans(): JSX.Element {
  const { data, loading, error } = useAsync(API.GetPlansList);
  const hasError = !!error;

  const HandleCreate = () => {
    Router.push('/admin/plans/create');
  };

  if (loading) {
    return <div></div>;
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <CrudToolbar HandleCreate={HandleCreate} />
      <PlansTable show={!hasError} plans={data?.plans} />
      <ErrorMessage show={hasError} error={error} />
    </LayoutAdminDash>
  );
}
