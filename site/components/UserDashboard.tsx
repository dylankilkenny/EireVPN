import React, { useState } from 'react';
import { IconContext } from 'react-icons';
import { IoMdPerson, IoIosCalendar } from 'react-icons/io';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import UserDetailsCard from './UserDetailsCard';
import SubscriptionCard from './SubscriptionCard';

interface UserDashboardProps {
  userid: string;
}

const UserDashboard: React.FC<UserDashboardProps> = ({ userid }) => {
  return (
    <Container>
      <Row>
        <Col>
          <IconContext.Provider value={{ color: '#626262' }}>
            <h2 className="dashboard-heading">
              <IoMdPerson />
              <div style={{ paddingLeft: 10 }}>My Account</div>
            </h2>
          </IconContext.Provider>
        </Col>
      </Row>
      <hr></hr>
      <Row>
        <Col sm={12} md={12} lg={4}>
          <SubscriptionCard userid={userid} />
        </Col>
        <Col sm={12} md={12} lg={8}>
          <UserDetailsCard userid={userid} />
        </Col>
      </Row>
    </Container>
  );
};

export default UserDashboard;
