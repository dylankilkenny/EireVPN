import React from 'react';
import styled from 'styled-components';
import { FiArrowLeftCircle } from 'react-icons/fi';

import { MdSettings } from 'react-icons/md';
import { IconContext } from 'react-icons';
import PropTypes from 'prop-types';

const Cont = styled.div`
  padding-top: 10px; 
  height 70px;
`;

const BrandingCont = styled.div`
  float: left;
  width: 85%;
  color: #5a6268;
`;

const IconCont = styled.div`
  margin-left: 85%;
`;

const LoginHeader = () => (
  <Cont>
    <BrandingCont>
      <h2>ÉireVPN</h2>
    </BrandingCont>
  </Cont>
);

const MainHeader = ({ renderSettings }) => (
  <Cont>
    <BrandingCont>
      <h2>EireVPN</h2>
    </BrandingCont>
    <IconCont>
      <IconContext.Provider value={{ size: '2em', color: '#5a6268' }}>
        <div>
          <MdSettings style={{ cursor: 'pointer' }} onClick={renderSettings} />
        </div>
      </IconContext.Provider>
    </IconCont>
  </Cont>
);

const SettingsHeader = ({ renderMain }) => (
  <Cont>
    <IconContext.Provider value={{ size: '2em', color: '#5a6268' }}>
      <div>
        <FiArrowLeftCircle style={{ cursor: 'pointer' }} onClick={renderMain} />
      </div>
    </IconContext.Provider>
  </Cont>
);

const Header = ({ view, renderSettings, renderMain }) => {
  switch (view) {
    case 0:
      return <LoginHeader />;
    case 1:
      return <MainHeader renderSettings={renderSettings} />;
    case 2:
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
