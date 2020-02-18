import { HeaderUser, HeaderDashboard, HeaderLogin } from './Header';
import Head from 'next/head';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { logPageView } from '../service/Analytics';
import React, { useEffect } from 'react';

interface Props {
  children: React.ReactNode;
}

const HtmlHead = (): JSX.Element => {
  useEffect(() => {
    logPageView();
  });
  return (
    <Head>
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      <link
        href="https://fonts.googleapis.com/css?family=Fredoka+One&display=swap"
        rel="stylesheet"
      />
    </Head>
  );
};

const LayoutLanding: React.FC<Props> = props => (
  <div>
    <HtmlHead />
    <div>
      <HeaderUser />
      {props.children}
    </div>
  </div>
);

const LayoutLogin: React.FC<Props> = props => (
  <div>
    <HtmlHead />
    <HeaderLogin />
    <Container>
      <Row>
        <Col sm={12} md={2} lg={3}></Col>
        <Col sm={12} md={8} lg={6}>
          {props.children}
        </Col>
        <Col sm={12} md={2} lg={3}></Col>
      </Row>
    </Container>
  </div>
);

const LayoutContact: React.FC<Props> = props => (
  <div>
    <HtmlHead />
    <HeaderUser />
    <Container fluid>
      <Row>
        <Col>
          <div className="account-dash">{props.children}</div>
        </Col>
      </Row>
    </Container>
  </div>
);

interface UserDashProps {
  children: React.ReactNode;
}

const LayoutUserDash: React.FC<UserDashProps> = ({ children }) => (
  <div>
    <HtmlHead />
    <HeaderDashboard />
    <div className="account-layout">
      <Container fluid>
        <Row>
          <Col sm={12} md={1} lg={1} />
          <Col sm={12} md={10} lg={10}>
            <div className="account-dash">{children}</div>
          </Col>
          <Col sm={12} md={1} lg={1} />
        </Row>
      </Container>
    </div>
  </div>
);

interface AdminDashProps {
  AdminSidePanel: React.ReactNode;
  children: React.ReactNode;
}

const LayoutAdminDash: React.FC<AdminDashProps> = ({ AdminSidePanel, children }) => (
  <div>
    <HtmlHead />
    <HeaderDashboard />
    <div className="admin-layout">
      <Container fluid>
        <Row>
          <Col sm={12} md={3} lg={2}>
            {AdminSidePanel}
          </Col>
          <Col>
            <div className="admin-dash">{children}</div>
          </Col>
        </Row>
      </Container>
    </div>
  </div>
);

export { LayoutLanding, LayoutLogin, LayoutAdminDash, LayoutUserDash, LayoutContact };
