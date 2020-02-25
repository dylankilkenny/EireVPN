import { LayoutMain } from '../components/Layout';
import Image from 'react-bootstrap/Image';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Button from 'react-bootstrap/Button';
import SuccessMessage from '../components/SuccessMessage';
import { useRouter } from 'next/router';

export default function DownloadsPage() {
  const router = useRouter();
  const signedup = router.query.signedup;

  const chromeLink =
    'https://chrome.google.com/webstore/detail/%C3%A9irevpn-beta-free-vpn-ad/fmhlhjegfkgemampnpomkajhdpjnhmie?hl=en-GB';
  const firefoxLink =
    'https://addons.mozilla.org/en-US/firefox/addon/%C3%A9irevpn-free-vpn-ad-blocker/';
  return (
    <LayoutMain>
      <div className="downloads-cont">
        <Container>
          <Row>
            <Col xs={12} md={6}>
              <Image style={{ marginTop: 20 }} fluid src="../static/images/download.png" />
            </Col>
            <Col xs={12} md={6}>
              <Row>
                <Col>
                  <SuccessMessage
                    show={!!signedup}
                    message="Sign up complete, download an extension for your preferred browser."
                  />
                  <h2 className="downloads-heading">Downloads</h2>
                </Col>
              </Row>
              <Row>
                <Col>
                  <Image
                    className="img-browser"
                    fluid
                    src="../static/images/icons8-chrome-144.png"
                  />
                </Col>
                <Col>
                  <a target="_blank" href={chromeLink}>
                    <Button className="btn-downloads" variant="primary">
                      Add to Chrome
                    </Button>
                  </a>
                </Col>
              </Row>
              <hr></hr>
              <Row>
                <Col>
                  <Image
                    className="img-browser"
                    fluid
                    src="../static/images/icons8-firefox-144.png"
                  />
                </Col>
                <Col>
                  <a target="_blank" href={firefoxLink}>
                    <Button className="btn-downloads" variant="primary">
                      Add to Firefox
                    </Button>
                  </a>
                </Col>
              </Row>
            </Col>
          </Row>
        </Container>
      </div>
    </LayoutMain>
  );
}
