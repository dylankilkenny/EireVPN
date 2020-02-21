import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import FormInput from '../../FormInput';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';
import ButtonMain from '../../ButtonMain';

interface UserCreateFormProps {
  HandleSave: (body: string) => Promise<void>;
  error: APIError;
}

const UserCreateForm: React.FC<UserCreateFormProps> = ({ HandleSave, error }) => {
  const [firstname, setFirstname] = useState('');
  const [lastname, setLastname] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const hasError = !!error;

  const handleSaveClick = () => {
    HandleSave(JSON.stringify({ firstname, lastname, email, password }));
  };

  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title className="card-title-form">
            Create User
            <div className="button-toolbar">
              <ButtonMain onClick={handleSaveClick} value="Save" />
            </div>
          </Card.Title>
          <Form>
            <ErrorMessage show={hasError} error={error} />
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
              <FormInput name="email" label="Email" value={email} onChange={setEmail} />
              <FormInput
                type="password"
                name="password"
                label="Password"
                value={password}
                onChange={setPassword}
              />
            </Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default UserCreateForm;
