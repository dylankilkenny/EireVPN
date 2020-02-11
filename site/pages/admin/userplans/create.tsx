import React, { useState } from 'react';
import { useRouter } from 'next/router';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import API from '../../../service/APIService';
import UserPlanCreateForm from '../../../components/admin/forms/UserPlanCreateForm';
import useAsync from '../../../hooks/useAsync';
import ErrorMessage from '../../../components/ErrorMessage';

export default function UserPlanCreate(): JSX.Element {
  const router = useRouter();
  const userid = router.query.userid;
  const { data, loading, error } = useAsync(() => API.GetPlansList());
  const [respError, setRespError] = useState();

  async function HandleSave(body: string) {
    const res = await API.CreateUserPlan(body);
    if (res.status == 200) {
      router.push('/admin/userplans');
    } else {
      setRespError(res);
    }
  }

  if (loading) {
    return <div></div>;
  }
  console.log(data);
  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      {!error ? (
        <UserPlanCreateForm
          userid={userid.toString()}
          planlist={data.plans}
          error={respError}
          HandleSave={HandleSave}
        />
      ) : (
        <ErrorMessage show={true} error={error} />
      )}
    </LayoutAdminDash>
  );
}

UserPlanCreate.getInitialProps = async () => {
  return {};
};
