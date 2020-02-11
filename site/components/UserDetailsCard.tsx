import React, { useState } from 'react';
import { IconContext } from 'react-icons';
import { FaRegEdit } from 'react-icons/fa';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import useAsync from '../hooks/useAsync';
import API from '../service/APIService';
import { useRouter } from 'next/router';

interface UserDetailsCardProps {
  userid: string;
}

const UserDetailsCard: React.FC<UserDetailsCardProps> = ({ userid }) => {
  const router = useRouter();
  const { data, loading, error } = useAsync(() => API.GetUserByID(userid));
  const handleEditClick = () => {
    router.push('/account/edit');
  };

  if (loading) {
    return <div></div>;
  }

  const { user } = data;

  return (
    <Card className="dash-card">
      <Card.Body>
        <Card.Title>
          My Details
          <IconContext.Provider value={{ color: '#004643' }}>
            <div onClick={handleEditClick} className="details-edit">
              <FaRegEdit />
              <div style={{ paddingLeft: 10 }}>Edit</div>
            </div>
          </IconContext.Provider>
        </Card.Title>
        <hr></hr>
        <div className="details-card">
          <Row>
            <Col>
              <label htmlFor="firstname">First Name</label>

              {user.lastname ? (
                <div id="firstname">{user.firstname}</div>
              ) : (
                <div className="empty-field" id="firstname">
                  Click edit to fill...
                </div>
              )}
            </Col>
            <Col>
              <label htmlFor="lastname">Last Name</label>
              {user.lastname ? (
                <div id="lastname">{user.lastname}</div>
              ) : (
                <div className="empty-field" id="lastname">
                  Click edit to fill...
                </div>
              )}
            </Col>
          </Row>
          <Row>
            <Col>
              <label htmlFor="email">Email</label>
              <div id="email">{user.email}</div>
            </Col>
            <Col>
              <label htmlFor="lastname">Password</label>
              <div id="lastname">• • • • • • • • • </div>
            </Col>
          </Row>
        </div>
      </Card.Body>
    </Card>
  );
};

export default UserDetailsCard;
