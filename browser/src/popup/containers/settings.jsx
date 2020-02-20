import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import PropTypes from 'prop-types';
import { IoIosMail, IoIosLogOut } from 'react-icons/io';
import { IconContext } from 'react-icons';
import ext from '../../utils/ext';

const Settings = ({ logout }) => {
  const forgotPass = () => {
    const newURL = `${process.env.DOMAIN}/contact`;
    ext.tabs.create({ url: newURL });
  };
  return (
    <Container>
      <Row>
        <Col />
        <Col>
          <div className="btn-custom-logout" onClick={logout}>
            Logout
            <IconContext.Provider value={{ size: '1.5em' }}>
              <IoIosLogOut style={{ cursor: 'pointer', marginLeft: 7 }} />
            </IconContext.Provider>
          </div>
        </Col>
        <Col />
      </Row>
      <Row>
        <Col />
        <Col>
          <div className="create-acc-cont">
            <div onClick={() => forgotPass()} className="create-account-link">
              <IconContext.Provider value={{ size: '1.5em' }}>
                <IoIosMail />
              </IconContext.Provider>
              Contact Support
            </div>
          </div>
        </Col>
        <Col />
      </Row>
    </Container>
  );
};

Settings.propTypes = {
  logout: PropTypes.func.isRequired
};

export default Settings;
