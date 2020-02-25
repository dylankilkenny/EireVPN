import { LayoutMain } from '../components/Layout';
import Image from 'react-bootstrap/Image';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Button from 'react-bootstrap/Button';
import Link from 'next/link';

export default function LandingPage() {
  return (
    <LayoutMain>
      <div className="landing-page">
        <Container>
          <Row>
            <Col xs={12} md={6}>
              <div className="headline">High speed Irish VPN and Ad Blocker.</div>
              <p className="p-landing">
                Safely secure your browsing and enjoy unrestricted access worldwide.
              </p>
              <Link href="/signup">
                <Button className="btn-landing" variant="primary">
                  Try Free
                </Button>
              </Link>
            </Col>
            <Col xs={12} md={6}>
              <div className="landing-image-cont">
                <Image fluid src="../static/images/undraw_security_o890.png" />
              </div>
            </Col>
          </Row>
        </Container>
      </div>
    </LayoutMain>
  );
}
