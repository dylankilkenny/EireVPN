import React from 'react';
import Alert from 'react-bootstrap/Alert';

const LargeAlert = ({ variant, heading, body }) => (
  <Alert variant={variant}>
    <Alert.Heading>{heading}</Alert.Heading>
    <p>{body}</p>
  </Alert>
);

export default LargeAlert;
