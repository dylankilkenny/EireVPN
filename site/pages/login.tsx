import React, { useState } from 'react';
import { LayoutLogin } from '../components/Layout';
import LoginForm from '../components/admin/forms/LoginForm';
import { useRouter } from 'next/router';
import API from '../service/APIService';

export default function LoginPage(): JSX.Element {
  const router = useRouter();
  const signedup = router.query.signedup;
  const [error, setError] = useState();

  async function HandleLogin(body: string) {
    const res = await API.Login(body);
    if (res.status == 200) {
      router.push('/account');
    } else {
      setError(res);
    }
  }

  return (
    <LayoutLogin>
      <div className="signup-cont">
        <LoginForm
          signedup={signedup == '1' ? true : false}
          error={error}
          HandleLogin={HandleLogin}
        />
      </div>
    </LayoutLogin>
  );
}
