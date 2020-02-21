import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Spinner from 'react-bootstrap/Spinner';

const Loading = () => (
  <div>
    <Container>
      <Row>
        <Col>
          <div className="spinner-div">
            <Spinner className="loading-spinner" animation="border" variant="primary" />
          </div>
        </Col>
      </Row>
    </Container>
  </div>
);

export default Loading;
