import React from 'react';
import Container from 'react-bootstrap/Container';
import PropTypes from 'prop-types';

const PopupContainer = ({ children }) => (
  <div className="popup">
    <Container className="main-container">{children}</Container>
  </div>
);

PopupContainer.propTypes = {
  children: PropTypes.any.isRequired
};

export default PopupContainer;
