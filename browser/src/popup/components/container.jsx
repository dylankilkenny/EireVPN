import React from 'react';
import Container from 'react-bootstrap/Container';
import PropTypes from 'prop-types';
import styled from 'styled-components';

const Wrapper = styled.div`
  height: 400px;
  width: 300px;
`;

const PopupContainer = ({ children }) => (
  <Wrapper>
    <Container>{children}</Container>
  </Wrapper>
);

PopupContainer.propTypes = {
  children: PropTypes.any.isRequired,
};

export default PopupContainer;
