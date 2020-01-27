import React, { useState } from 'react';
import { LayoutLogin } from '../../components/Layout';
import LoginForm from '../../components/admin/forms/LoginForm';
import Router from 'next/router';
import API from '../../service/APIService';

export default function LoginPage(): JSX.Element {
  const [error, setError] = useState();

  async function HandleLogin(email: string, password: string) {
    const res = await API.Login(JSON.stringify({ email, password }));
    if (res.status == 200) {
      Router.push('/admin/users');
    } else {
      setError(res);
    }
  }

  return (
    <LayoutLogin>
      <h2>Login</h2>
      <LoginForm error={error} HandleLogin={HandleLogin} />
    </LayoutLogin>
  );
}
