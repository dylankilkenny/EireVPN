import Image from 'react-bootstrap/Image';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Link from 'next/link';

export default function Footer() {
  return (
    <div className="footer">
      <Container>
        <Row>
          <Col xs={12} md={5}>
            <Row>
              <Link href="/">
                <a>
                  <Image className="branding-shield" fluid src="../static/images/shield.png" />
                  <h4 className="branding">ÉireVPN</h4>
                </a>
              </Link>
            </Row>
            <Row>
              <span style={{ fontSize: 10 }}>Copyright © ÉireVPN 2020 - All rights reserved.</span>
            </Row>
          </Col>
          <Col xs={12} md={6}>
            <Row>
              <Col sm={3}>
                <Link href="/downloads">
                  <a>Downloads</a>
                </Link>
              </Col>
              <Col sm={3}>
                <Link href="/contact">
                  <a>Support</a>
                </Link>
              </Col>
              <Col sm={3}>
                <Link href="/login">
                  <a>Login</a>
                </Link>
              </Col>
              <Col sm={3}>
                <Link href="/signup">
                  <a>Sign Up</a>
                </Link>
              </Col>
            </Row>
            <Row>
              <Col sm={3}>
                <Link href="/policies/privacy-policy">
                  <a>Privacy</a>
                </Link>
              </Col>
              <Col sm={3}>
                <Link href="/policies/cookie-policy">
                  <a>Cookies</a>
                </Link>
              </Col>
              <Col sm={3}>
                <Link href="/policies/terms-conditions">
                  <a>Terms</a>
                </Link>
              </Col>
            </Row>
            <style jsx>{`
              a {
                text-decoration: none;
                color: #abd1c6;
              }
              a:hover {
                opacity: 0.6;
              }
            `}</style>
          </Col>
        </Row>
      </Container>
    </div>
  );
}
