import React, { useState } from 'react';
import User from '../../interfaces/user';
import APIError from '../../interfaces/error';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { IconContext } from 'react-icons';
import { IoMdPerson, IoIosArrowBack } from 'react-icons/io';
import EditDetailsForm from '../../components/user/forms/EditDetailsForm';
import ChangePasswordForm from '../../components/user/forms/ChangePasswordForm';
import { useRouter } from 'next/router';
import SuccessMessage from '../SuccessMessage';
import ErrorMessage from '../ErrorMessage';

interface EditDetailsDashbaordProps {
  user: User;
  success: boolean;
  error: APIError;
  HandleDetailsSave: (body: string) => Promise<void>;
  HandlePasswordSave: (body: string) => Promise<void>;
}

const EditDetailsDashbaord: React.FC<EditDetailsDashbaordProps> = ({
  user,
  success,
  error,
  HandleDetailsSave,
  HandlePasswordSave
}) => {
  const router = useRouter();
  const handleBackClick = () => {
    router.push('/account');
  };
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
          <div className="button-toolbar">
            <IconContext.Provider value={{ color: '#004643' }}>
              <h4 onClick={handleBackClick} className="details-back-btn">
                <IoIosArrowBack />
                <div>Back</div>
              </h4>
            </IconContext.Provider>
          </div>
        </Col>
      </Row>
      <hr></hr>
      <ErrorMessage show={!!error} error={error} />
      <SuccessMessage show={success} message="Saved Successfully" />
      <Row>
        <Col sm={12} md={12} lg={8}>
          <EditDetailsForm user={user} HandleSave={HandleDetailsSave} />
        </Col>
        <Col sm={12} md={12} lg={4}>
          <ChangePasswordForm success={success} HandleSave={HandlePasswordSave} />
        </Col>
      </Row>
    </Container>
  );
};

export default EditDetailsDashbaord;
