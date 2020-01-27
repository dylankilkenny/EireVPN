import React, { useState } from 'react';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import ErrorMessage from '../../../components/ErrorMessage';
import API from '../../../service/APIService';
import Auth from '../../../service/Auth';
import useAsync from '../../../hooks/useAsync';
import UserForm from '../../../components/admin/forms/UserEditForm';
import { useRouter } from 'next/router';

export default function UserEdit(): JSX.Element {
  const router = useRouter();
  const userID = router.query.id.toString();
  const { data, loading, error } = useAsync(() => API.GetUserByID(userID));
  const [respError, setRespError] = useState();
  const [success, setSuccess] = useState();
  const hasError = !!error;

  async function HandleSave(body: string) {
    const res = await API.UpdateUser(userID, body);
    if (res.status == 200) {
      setSuccess(true);
      setRespError(false);
    } else {
      setRespError(res);
      setSuccess(false);
    }
  }

  async function HandleDelete() {
    const res = await API.DeleteUser(userID);
    if (res.status == 200) {
      router.push('/admin/users');
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
        <UserForm
          success={success}
          HandleDelete={HandleDelete}
          HandleSave={HandleSave}
          error={respError}
          user={data.user}
        />
      ) : (
        <ErrorMessage show={hasError} error={error} />
      )}
    </LayoutAdminDash>
  );
}

UserEdit.getInitialProps = async () => {
  return {};
};
