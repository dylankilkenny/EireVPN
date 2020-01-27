import React, { useState } from 'react';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import ErrorMessage from '../../../components/ErrorMessage';
import API from '../../../service/APIService';
import useAsync from '../../../hooks/useAsync';
import ServerEditForm from '../../../components/admin/forms/ServerEditForm';
import { useRouter } from 'next/router';

export default function ServerEdit(): JSX.Element {
  const router = useRouter();
  const serverID = router.query.id.toString();
  const { data, loading, error } = useAsync(() => API.GetServerByID(serverID));
  const [respError, setRespError] = useState();
  const [success, setSuccess] = useState();
  const hasError = !!error;

  async function HandleSave(body: string) {
    const res = await API.UpdateServer(serverID, body);
    if (res.status == 200) {
      setSuccess(true);
      setRespError(false);
    } else {
      setRespError(res);
      setSuccess(false);
    }
  }

  async function HandleDelete() {
    const res = await API.DeleteServer(serverID);
    if (res.status == 200) {
      router.push('/admin/servers');
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
        <ServerEditForm
          HandleSave={HandleSave}
          HandleDelete={HandleDelete}
          server={data.server}
          error={respError}
          success={success}
        />
      ) : (
        <ErrorMessage show={hasError} error={error} />
      )}
    </LayoutAdminDash>
  );
}

ServerEdit.getInitialProps = async () => {
  return {};
};
