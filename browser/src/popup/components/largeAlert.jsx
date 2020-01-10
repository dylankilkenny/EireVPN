import React from 'react';
import Alert from 'react-bootstrap/Alert';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const LargeAlert = ({ variant, heading, body }) => (
  <Alert variant={variant}>
    <Alert.Heading>{heading}</Alert.Heading>
    <p>{body}</p>
  </Alert>
);

// PopupContainer.propTypes = {
//   children: PropTypes.any.isRequired
// };

export default LargeAlert;
