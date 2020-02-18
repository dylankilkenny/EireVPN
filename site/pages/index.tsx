import { LayoutLanding } from '../components/Layout';
import Image from 'react-bootstrap/Image';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Button from 'react-bootstrap/Button';

export default function LandingPage() {
  return (
    <LayoutLanding>
      <div className="landing-page">
        <Container>
          <Row>
            <Col xs={12} md={6}>
              <div className="headline">High speed Irish VPN and Ad Blocker.</div>
              <p className="p-landing">
                Safely secure your browsing and enjoy unrestricted access worldwide.
              </p>
              <Button className="btn-landing" variant="primary">
                Try Free
              </Button>
            </Col>
            <Col xs={12} md={6}>
              <Image
                style={{ marginTop: 90 }}
                fluid
                src="../static/images/undraw_security_o890.png"
              />
            </Col>
          </Row>
        </Container>
      </div>
    </LayoutLanding>
  );
}
