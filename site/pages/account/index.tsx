import React, { useEffect } from 'react';
import { LayoutUserDash } from '../../components/Layout';
import UserDashboard from '../../components/UserDashboard';
import { useRouter } from 'next/router';
import Cookies from 'js-cookie';

export default function Account(): JSX.Element {
  const router = useRouter();
  let userid = Cookies.get('uid')!;
  useEffect(() => {
    if (!userid) {
      router.push('/login');
    }
  });

  return (
    <LayoutUserDash>
      <UserDashboard userid={userid} />
    </LayoutUserDash>
  );
}

Account.getInitialProps = async () => {
  return {};
};
