import React, { useState } from 'react';
import { LayoutAdminDash } from '../../components/Layout';
import AdminSidePanel from '../../components/admin/AdminSidePanel';
import ErrorMessage from '../../components/ErrorMessage';
import API from '../../service/APIService';
import useAsync from '../../hooks/useAsync';
import SettingsForm from '../../components/admin/forms/SettingsForm';

export default function SettingsPage(): JSX.Element {
  const { data, loading, error } = useAsync(() => API.GetSettings());
  const [respError, setRespError] = useState();
  const [success, setSuccess] = useState();
  const hasError = !!error;

  async function HandleSave(body: string) {
    const res = await API.UpdateSettings(body);
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

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      {!error ? (
        <SettingsForm
          success={success}
          HandleSave={HandleSave}
          error={respError}
          settings={data.settings}
        />
      ) : (
        <ErrorMessage show={hasError} error={error} />
      )}
    </LayoutAdminDash>
  );
}

SettingsPage.getInitialProps = async () => {
  return {};
};
