import React, { useState, useEffect } from 'react';
import { LayoutMain } from '../../components/Layout';
import { useRouter } from 'next/router';
import API from '../../service/APIService';
import ErrorMessage from '../../components/ErrorMessage';
import SuccessMessage from '../../components/SuccessMessage';
import ForgotPasswordForm from '../../components/user/forms/ForgotPasswordForm';
import useAsync from '../../hooks/useAsync';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';

export default function ForgotPassword(): JSX.Element {
  const [respError, setRespError] = useState();
  const [success, setSuccess] = useState();

  async function HandleSubmit(body: string) {
    const res = await API.ForgotPasswordEmail(body);
    if (res.status == 200) {
      setSuccess(true);
      setRespError(false);
    } else {
      setRespError(res);
      setSuccess(false);
    }
  }

  return (
    <LayoutMain>
      <Container>
        <Row>
          <Col />
          <Col sm="12" md="6">
            <ForgotPasswordForm success={success} error={respError} HandleSubmit={HandleSubmit} />
          </Col>
          <Col />
        </Row>
      </Container>
    </LayoutMain>
  );
}

ForgotPassword.getInitialProps = async () => {
  return {};
};
