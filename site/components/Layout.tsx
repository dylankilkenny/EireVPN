import { HeaderMain, HeaderDashboard } from './Header';
import Head from 'next/head';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { initGA, logPageView } from '../service/Analytics';
import React, { useEffect } from 'react';

interface Props {
  children: React.ReactNode;
}

const HtmlHead = (): JSX.Element => {
  useEffect(() => {
    initGA();
    logPageView();
  });
  return (
    <Head>
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      <meta
        name="description"
        content="ÉireVPN is an irish based VPN solution with ad blocking technology baked right in. Why use 2 extensions when you can have one?"
      />
      <meta property="og:title" content="ÉireVPN - Free VPN & Ad Blocker" />
      <meta
        property="og:description"
        content="ÉireVPN is an irish based VPN solution with ad blocking technology baked right in. Why use 2 extensions when you can have one?"
      />
      <meta property="og:image" content="../static/images/links-image.png" />
      <meta property="og:site_name" content="ÉireVPN" />
      <meta property="twitter:title" content="ÉireVPN - Free VPN & Ad Blocker" />
      <meta
        property="twitter:description"
        content="ÉireVPN is an irish based VPN solution with ad blocking technology baked right in. Why use 2 extensions when you can have one?"
      />
      <meta property="twitter:image" content="../static/images/links-image.png" />
      <meta name="twitter:site" content="@eirevpn" />
      <meta name="twitter:creator" content="@eirevpn" />
      <meta name="robots" content="index, follow" />
      <link
        href="https://fonts.googleapis.com/css?family=Fredoka+One&display=swap"
        rel="stylesheet"
      />
      <link rel="stylesheet" type="text/css" href="//wpcc.io/lib/1.0.2/cookieconsent.min.css" />
      <link rel="canonical" href="https://eirevpn.ie/" />
      <script src="../static/js/cookie-policy.js"></script>
      <script src="//wpcc.io/lib/1.0.2/cookieconsent.min.js"></script>
      <link rel="icon" type="image/png" href="../static/images/shield.png"></link>
      <title>ÉireVPN - Free VPN & Ad Blocker</title>
    </Head>
  );
};

const LayoutLanding: React.FC<Props> = props => (
  <div>
    <HtmlHead />
    <div>
      <HeaderMain />
      {props.children}
    </div>
  </div>
);

const LayoutLogin: React.FC<Props> = props => (
  <div>
    <HtmlHead />
    <HeaderMain />
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
    <HeaderMain />
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
