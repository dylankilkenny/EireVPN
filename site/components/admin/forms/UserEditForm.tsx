import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import User from '../../../interfaces/user';
import ButtonMain from '../../ButtonMain';
import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';
import dayjs from 'dayjs';
import FormInput from './FormInput';
import SuccessMessage from '../../SuccessMessage';
import Router from 'next/router';

interface UserEditFormProps {
  user: User;
  error: APIError;
  success: boolean;
  showCreateUserPlan: boolean;
  HandleSave: (body: string) => Promise<void>;
  HandleDelete: () => Promise<void>;
}

const UserForm: React.FC<UserEditFormProps> = ({
  user,
  error,
  success,
  showCreateUserPlan,
  HandleSave,
  HandleDelete
}) => {
  const hasError = !!error;
  const [firstname, setFirstname] = useState(user.firstname);
  const [lastname, setLastname] = useState(user.lastname);
  const [email, setEmail] = useState(user.email);

  const handleSaveClick = () => {
    HandleSave(JSON.stringify({ firstname, lastname, email }));
  };

  const handleDeleteClick = () => {
    HandleDelete();
  };

  const goToUserPlan = () => {
    Router.push('/admin/userplans/[userid]', '/admin/userplans/' + user.id);
  };

  const goToCreateUserPlan = () => {
    Router.push({
      pathname: '/admin/userplans/create',
      query: { userid: user.id }
    });
  };

  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title className="card-title-form">
            User Details
            <div className="button-toolbar">
              <ButtonMain onClick={handleSaveClick} value="Save" />
              <ButtonMain onClick={handleDeleteClick} value="Delete" />
            </div>
          </Card.Title>
          <ErrorMessage show={hasError} error={error} />
          <SuccessMessage show={success} message="Server Updated Successfully" />
          <Form>
            <Form.Row>
              <FormInput
                textOnly
                name="createdAt"
                label="Created At"
                value={dayjs(user.createdAt)
                  .format('DD-MM-YYYY H:mm')
                  .toString()}
              />
              <FormInput
                textOnly
                name="updateAt"
                label="Updated At"
                value={dayjs(user.updatedAt)
                  .format('DD-MM-YYYY H:mm')
                  .toString()}
              />
            </Form.Row>
            <Form.Row>
              <FormInput
                textOnly
                name="stripe_id"
                label="Stripe ID"
                value={user.stripe_customer_id}
              />
              <FormInput name="email" label="Email" value={email} onChange={setEmail} />
            </Form.Row>
            <Form.Row>
              <FormInput
                name="firstname"
                label="Firstname"
                value={firstname}
                onChange={setFirstname}
              />
              <FormInput name="lastname" label="Lastname" value={lastname} onChange={setLastname} />
            </Form.Row>
            <Form.Row>
              {showCreateUserPlan ? (
                <ButtonMain onClick={goToCreateUserPlan} value="Create User Plan" />
              ) : (
                <ButtonMain onClick={goToUserPlan} value="User Plan" />
              )}
            </Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default UserForm;
