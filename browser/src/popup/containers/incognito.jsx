import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Alert from 'react-bootstrap/Alert';
import Image from 'react-bootstrap/Image';

const Incognito = () => (
  <div>
    <Alert className="alert-incognito" variant="warning">
      Permission to run in private windows is required by Firefox.
    </Alert>
    <Container>
      <Row>
        <Col>1. Right click on the extensions icon and click manage</Col>
      </Row>
      <Row>
        <Col>
          <Image className="image-incognito" src="../../assets/images/manage.png" />
        </Col>
      </Row>
      <Row>
        <Col>2. Scroll down to Run in Private Windows and choose Allow.</Col>
      </Row>
      <Row>
        <Col>
          <Image className="image-incognito allow-img" src="../../assets/images/allow.png" />
        </Col>
      </Row>
    </Container>
  </div>
);

export default Incognito;
