import React, { useState, useEffect } from 'react';
import { LayoutLanding } from '../components/Layout';
import { useRouter } from 'next/router';
import API from '../service/APIService';
import ErrorMessage from '../components/ErrorMessage';
import SuccessMessage from '../components/SuccessMessage';
import useAsync from '../hooks/useAsync';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Container from 'react-bootstrap/Container';

export default function ConfirmEmailPage(): JSX.Element {
  const router = useRouter();
  const token = router.query.token;
  const { data, loading, error } = useAsync(() => API.ConfirmEmail(token.toString()));

  useEffect(() => {
    if (!token) {
      router.push('/');
    }
  });

  if (loading) {
    return <div></div>;
  }

  return (
    <LayoutLanding>
      <Container>
        <Row>
          <Col />
          <Col>
            <ErrorMessage show={!!error} error={error} />
            <SuccessMessage show={!error} message="Thank you for confirming your email" />
          </Col>
          <Col />
        </Row>
      </Container>
    </LayoutLanding>
  );
}

ConfirmEmailPage.getInitialProps = async () => {
  return {};
};
