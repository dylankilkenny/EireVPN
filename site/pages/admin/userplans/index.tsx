import React from 'react';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import ErrorMessage from '../../../components/ErrorMessage';
import { LayoutAdminDash } from '../../../components/Layout';
import UserPlansTable from '../../../components/admin/tables/UserPlansTable';
import useAsync from '../../../hooks/useAsync';
import API from '../../../service/APIService';
import CrudToolbar from '../../../components/CrudToolbar';
import Router from 'next/router';

export default function Users(): JSX.Element {
  const { data, loading, error } = useAsync(API.GetUserPlansList);
  const hasError = !!error;

  if (loading) {
    return <div></div>;
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <UserPlansTable show={!hasError} userplans={data?.userplans} />
      <ErrorMessage show={hasError} error={error} />
    </LayoutAdminDash>
  );
}
