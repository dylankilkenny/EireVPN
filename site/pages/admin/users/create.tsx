import React, { useState } from 'react';
import { useRouter } from 'next/router';
import { LayoutAdminDash } from '../../../components/Layout';
import AdminSidePanel from '../../../components/admin/AdminSidePanel';
import API from '../../../service/APIService';
import UserCreateForm from '../../../components/admin/forms/UserCreateForm';

export default function UserCreate(): JSX.Element {
  const router = useRouter();
  const [error, setError] = useState();

  async function HandleSave(body: string) {
    const res = await API.CreateUser(body);
    if (res.status == 200) {
      router.push('/admin/users');
    } else {
      setError(res);
    }
  }

  return (
    <LayoutAdminDash AdminSidePanel={<AdminSidePanel />}>
      <UserCreateForm error={error} HandleSave={HandleSave} />
    </LayoutAdminDash>
  );
}

UserCreate.getInitialProps = async () => {
  return {};
};
