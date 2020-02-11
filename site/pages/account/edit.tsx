import React, { useState, useEffect } from 'react';
import { LayoutUserDash } from '../../components/Layout';
import EditDetailsDashboard from '../../components/user/EditDetailsDashboard';
import API from '../../service/APIService';
import useAsync from '../../hooks/useAsync';
import { useRouter } from 'next/router';
import Cookies from 'js-cookie';

export default function PlansEdit(): JSX.Element {
  const router = useRouter();
  let userid = Cookies.get('uid')!;
  useEffect(() => {
    if (!userid) {
      router.push('/login');
    }
  });
  const { data, loading, error } = useAsync(() => API.GetUserByID(userid));
  const [respError, setRespError] = useState();
  const [success, setSuccess] = useState();

  async function HandleDetailsSave(body: string) {
    const res = await API.UpdateUser(userid, body);
    if (res.status == 200) {
      setSuccess(true);
      setRespError(false);
    } else {
      setRespError(res);
      setSuccess(false);
    }
  }

  async function HandlePasswordSave(body: string) {
    const res = await API.UpdatePassword(body);
    if (res.status == 200) {
      setSuccess(true);
      setRespError(false);
    } else {
      setRespError(res);
      setSuccess(false);
    }
  }

  if (loading) {
    return <div></div>;
  }

  if (error) {
    console.log(error);
  }

  return (
    <LayoutUserDash>
      <EditDetailsDashboard
        error={respError}
        success={success}
        user={data.user}
        HandleDetailsSave={HandleDetailsSave}
        HandlePasswordSave={HandlePasswordSave}
      />
    </LayoutUserDash>
  );
}

PlansEdit.getInitialProps = async () => {
  return {};
};
