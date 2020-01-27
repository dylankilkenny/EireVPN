import React, { useState } from 'react';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import ErrorMessage from '../../../components/ErrorMessage';
import API from '../../../service/APIService';
import useAsync from '../../../hooks/useAsync';
import PlanForm from '../../../components/admin/forms/PlanEditForm';
import { useRouter } from 'next/router';

export default function PlansEdit(): JSX.Element {
  const router = useRouter();
  const planID = router.query.id.toString();
  const { data, loading, error } = useAsync(() => API.GetPlanByID(planID));
  const [respError, setRespError] = useState();
  const [success, setSuccess] = useState();
  const hasError = !!error;

  async function HandleSave(body: string) {
    const res = await API.UpdatePlan(planID, body);
    if (res.status == 200) {
      setSuccess(true);
      setRespError(false);
    } else {
      setRespError(res);
      setSuccess(false);
    }
  }

  async function HandleDelete() {
    const res = await API.DeletePlan(planID);
    if (res.status == 200) {
      router.push('/admin/plans');
    } else {
      setRespError(res);
    }
  }

  if (loading) {
    return <div></div>;
  }
  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      {!error ? (
        <PlanForm
          plan={data.plan}
          success={success}
          error={respError}
          HandleDelete={HandleDelete}
          HandleSave={HandleSave}
        />
      ) : (
        <ErrorMessage show={hasError} error={error} />
      )}
    </LayoutAdminDash>
  );
}

PlansEdit.getInitialProps = async () => {
  return {};
};
