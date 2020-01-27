import React, { useState } from 'react';
import { useRouter } from 'next/router';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import API from '../../../service/APIService';
import ServerCreateForm from '../../../components/admin/forms/ServerCreateForm';

export default function ServerCreate(): JSX.Element {
  const router = useRouter();
  const [error, setError] = useState();

  async function HandleSave(body: FormData) {
    const res = await API.CreateServer(body);
    if (res.status == 200) {
      router.push('/admin/servers');
    } else {
      setError(res);
    }
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <ServerCreateForm error={error} HandleSave={HandleSave} />
    </LayoutAdminDash>
  );
}

ServerCreate.getInitialProps = async () => {
  return {};
};
