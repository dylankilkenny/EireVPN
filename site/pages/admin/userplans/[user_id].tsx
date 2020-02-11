import React, { useState } from 'react';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import ErrorMessage from '../../../components/ErrorMessage';
import API from '../../../service/APIService';
import useAsync from '../../../hooks/useAsync';
import UserPlanEditForm from '../../../components/admin/forms/UserPlanEditForm';
import { useRouter } from 'next/router';

export default function UserPlanEdit(): JSX.Element {
  const router = useRouter();
  const userID = router.query.user_id.toString();
  const { data, loading, error } = useAsync(() => API.GetUserPlanByUserID(userID));
  const [respError, setRespError] = useState();
  const [success, setSuccess] = useState();
  const hasError = !!error;

  async function HandleSave(body: string) {
    console.log(body);
    const res = await API.UpdateUserPlan(userID, body);
    if (res.status == 200) {
      setSuccess(true);
      setRespError(false);
    } else {
      setRespError(res);
      setSuccess(false);
    }
  }

  async function HandleDelete() {
    const res = await API.DeleteUserPlan(userID);
    if (res.status == 200) {
      router.push('/admin/userplans');
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
        <UserPlanEditForm
          success={success}
          HandleDelete={HandleDelete}
          HandleSave={HandleSave}
          error={respError}
          userplan={data.userplan}
        />
      ) : (
        <ErrorMessage show={hasError} error={error} />
      )}
    </LayoutAdminDash>
  );
}

UserPlanEdit.getInitialProps = async () => {
  return {};
};
