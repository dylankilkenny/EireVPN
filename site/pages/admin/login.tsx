import React, { useState } from 'react';
import { LayoutLogin } from '../../components/Layout';
import LoginForm from '../../components/admin/forms/LoginForm';
import Router from 'next/router';
import API from '../../service/APIService';

export default function LoginPage(): JSX.Element {
  const [error, setError] = useState();

  async function HandleLogin(body: string) {
    const res = await API.Login(body);
    if (res.status == 200) {
      Router.push('/admin/users');
    } else {
      setError(res);
    }
  }

  return (
    <LayoutLogin>
      <div className="login-cont">
        <LoginForm error={error} HandleLogin={HandleLogin} />
      </div>
    </LayoutLogin>
  );
}
