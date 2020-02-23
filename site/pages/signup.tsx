import React, { useState } from 'react';
import { LayoutLogin } from '../components/Layout';
import SignupForm from '../components/admin/forms/SignupForm';
import Router from 'next/router';
import API from '../service/APIService';

export default function SignupPage(): JSX.Element {
  const [error, setError] = useState();

  async function HandleSignup(body: string) {
    const res = await API.Signup(body);
    if (res.status == 200) {
      Router.push('/downloads?signedup=1');
    } else {
      setError(res);
    }
  }

  return (
    <LayoutLogin>
      <div className="signup-cont">
        <SignupForm error={error} HandleSignup={HandleSignup} />
      </div>
    </LayoutLogin>
  );
}
