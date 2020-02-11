import React, { useState } from 'react';
import { useRouter } from 'next/router';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import API from '../../../service/APIService';
import PlanCreateForm from '../../../components/admin/forms/PlanCreateForm';

export default function PlanCreate(): JSX.Element {
  const router = useRouter();
  const [error, setError] = useState();

  async function HandleSave(body: string) {
    console.log(body);
    const res = await API.CreatePlan(body);
    if (res.status == 200) {
      router.push('/admin/plans');
    } else {
      setError(res);
    }
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <PlanCreateForm error={error} HandleSave={HandleSave} />
    </LayoutAdminDash>
  );
}

PlanCreate.getInitialProps = async () => {
  return {};
};
