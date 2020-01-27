import Header from './Header';
import Head from 'next/head';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

interface Props {
  children: React.ReactNode;
}

const HtmlHead = (): JSX.Element => {
  return (
    <Head>
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    </Head>
  );
};

const Layout: React.FC<Props> = props => (
  <div>
    <HtmlHead />
    <div>
      <Header />
      {props.children}
    </div>
  </div>
);

const LayoutLogin: React.FC<Props> = props => (
  <div>
    <HtmlHead />
    <Header />
    <Container>
      <Row>
        <Col sm={12} md={2} lg={3}></Col>
        <Col sm={12} md={8} lg={6}>
          {props.children}
        </Col>
        <Col sm={12} md={2} lg={6}></Col>
      </Row>
    </Container>
  </div>
);

interface AdminDashProps {
  AdminSidePanel: React.ReactNode;
  children: React.ReactNode;
}

const LayoutAdminDash: React.FC<AdminDashProps> = ({
  AdminSidePanel,
  children
}) => (
  <div>
    <HtmlHead />
    <Header />
    <Container fluid>
      <Row>
        <Col sm={12} md={3} lg={2}>
          {AdminSidePanel}
        </Col>
        <Col>{children}</Col>
      </Row>
    </Container>
  </div>
);

export { Layout, LayoutLogin, LayoutAdminDash };
