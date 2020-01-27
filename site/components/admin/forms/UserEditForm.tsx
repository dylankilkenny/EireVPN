import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import User from '../../../interfaces/user';
import Button from 'react-bootstrap/Button';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';
import dayjs from 'dayjs';
import FormGroup from './FormGroup';
import SuccessMessage from '../../SuccessMessage';

interface UserEditFormProps {
  user: User;
  error: APIError;
  success: boolean;
  HandleSave: (body: string) => Promise<void>;
  HandleDelete: () => Promise<void>;
}

const UserForm: React.FC<UserEditFormProps> = ({
  user,
  error,
  success,
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

  return (
    <div>
      <div className="button-toolbar">
        <Button onClick={handleSaveClick} className="btn-save" variant="outline-secondary">
          Save
        </Button>
        <Button onClick={HandleDelete} variant="outline-secondary">
          Delete
        </Button>
      </div>
      <ErrorMessage show={hasError} error={error} />
      <SuccessMessage show={success} message="Server Updated Successfully" />
      <Form>
        <Form.Row>
          <FormGroup
            textOnly
            name="createdAt"
            label="Created At"
            value={dayjs(user.createdAt)
              .format('DD-MM-YYYY H:mm')
              .toString()}
          />
          <FormGroup
            textOnly
            name="updateAt"
            label="Updated At"
            value={dayjs(user.updatedAt)
              .format('DD-MM-YYYY H:mm')
              .toString()}
          />
        </Form.Row>
        <Form.Row>
          <FormGroup textOnly name="stripe_id" label="Stripe ID" value={user.stripe_customer_id} />
          <FormGroup name="email" label="Email" value={email} onChange={setEmail} />
        </Form.Row>
        <Form.Row>
          <FormGroup name="firstname" label="Firstname" value={firstname} onChange={setFirstname} />
          <FormGroup name="lastname" label="Lastname" value={lastname} onChange={setLastname} />
        </Form.Row>
      </Form>
    </div>
  );
};

export default UserForm;
