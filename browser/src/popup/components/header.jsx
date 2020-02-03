import React from 'react';
import styled from 'styled-components';
import { FaArrowLeft } from 'react-icons/fa';
import Image from 'react-bootstrap/Image';
import { MdSettings } from 'react-icons/md';
import { IconContext } from 'react-icons';
import PropTypes from 'prop-types';

const Cont = styled.div`
  padding-top: 5px; 
  height 60px;
  padding-left: 14px;
  background-color: #004643;
`;

const BrandingCont = styled.div`
  float: left;
  width: 85%;
  color: #5a6268;
`;

const IconCont = styled.div`
  margin-left: 85%;
`;

const HeaderBranding = () => (
  <div>
    <Image className="shield" src="../../assets/icons/logo-shield-48.png" />
    <h2 className="branding">ÉireVPN</h2>
  </div>
);

const LoginHeader = () => (
  <Cont>
    <BrandingCont>
      <HeaderBranding />
    </BrandingCont>
  </Cont>
);

const MainHeader = ({ renderSettings }) => (
  <Cont>
    <BrandingCont>
      <HeaderBranding />
    </BrandingCont>
    <IconCont>
      <IconContext.Provider value={{ size: '1.5em', color: 'white', className: 'settings-icon' }}>
        <MdSettings onClick={renderSettings} />
      </IconContext.Provider>
    </IconCont>
  </Cont>
);

const SettingsHeader = ({ renderMain }) => (
  <Cont>
    <IconContext.Provider value={{ size: '1.5em', color: 'white', className: 'back-arrow-icon' }}>
      <FaArrowLeft onClick={renderMain} />
    </IconContext.Provider>
  </Cont>
);

const Header = ({ view, renderSettings, renderMain }) => {
  switch (view) {
    case 0:
      return <LoginHeader />;
    case 1:
      return <MainHeader renderSettings={renderSettings} />;
    case 3:
      return <SettingsHeader renderMain={renderMain} />;
    default:
      return <LoginHeader />;
  }
};

MainHeader.propTypes = {
  renderSettings: PropTypes.func.isRequired
};

SettingsHeader.propTypes = {
  renderMain: PropTypes.func.isRequired
};

Header.propTypes = {
  renderSettings: PropTypes.func,
  renderMain: PropTypes.func,
  view: PropTypes.number.isRequired
};

export default Header;
