import React, { useState, useEffect } from 'react';
import { LayoutLanding } from '../../components/Layout';
import { useRouter } from 'next/router';
import API from '../../service/APIService';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';
import UpdatePasswordForm from '../../components/user/forms/UpdatePasswordForm';

export default function UpdatePassword(): JSX.Element {
  const router = useRouter();
  const token = router.query.token.toString();
  const [respError, setRespError] = useState();
  const [success, setSuccess] = useState();

  async function HandleSubmit(body: string) {
    const res = await API.UpdatePassword(body, token);
    if (res.status == 200) {
      setSuccess(true);
      setRespError(false);
    } else {
      setRespError(res);
      setSuccess(false);
    }
  }

  useEffect(() => {
    if (!token) {
      router.push('/');
    }
  });

  return (
    <LayoutLanding>
      <Container>
        <Row>
          <Col />
          <Col sm="12" md="6">
            <UpdatePasswordForm success={success} error={respError} HandleSubmit={HandleSubmit} />
          </Col>
          <Col />
        </Row>
      </Container>
    </LayoutLanding>
  );
}

UpdatePassword.getInitialProps = async () => {
  return {};
};
