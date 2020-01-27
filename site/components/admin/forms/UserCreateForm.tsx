import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import FormGroup from './FormGroup';
import Button from 'react-bootstrap/Button';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';

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

  const handleSave = () => {
    HandleSave(JSON.stringify({ firstname, lastname, email, password }));
  };

  return (
    <div className="server-create">
      <Form>
        <div className="button-toolbar">
          <Button onClick={handleSave} variant="outline-secondary">
            Save
          </Button>
        </div>
        <ErrorMessage show={hasError} error={error} />
        <Form.Row>
          <FormGroup name="firstname" label="Firstname" value={firstname} onChange={setFirstname} />
          <FormGroup name="lastname" label="Lastname" value={lastname} onChange={setLastname} />
        </Form.Row>
        <Form.Row>
          <FormGroup name="email" label="Email" value={email} onChange={setEmail} />
          <FormGroup name="password" label="Password" value={password} onChange={setPassword} />
        </Form.Row>
      </Form>
    </div>
  );
};

export default UserCreateForm;
