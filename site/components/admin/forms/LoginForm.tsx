import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';

interface LoginFormProps {
  HandleLogin: (email: string, pass: string) => Promise<void>;
  error: APIError;
}

type ChangeEvent = React.ChangeEvent<HTMLInputElement>;
type KeyEvent = React.KeyboardEvent<HTMLInputElement>;

const LoginForm: React.FC<LoginFormProps> = ({ HandleLogin, error }) => {
  const hasError = !!error;
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleKeyPress = (target: KeyEvent) => {
    if (target.charCode == 13) {
      HandleLogin(email, password);
    }
  };

  return (
    <Card style={{ padding: 20 }}>
      <Form>
        <Form.Group controlId="formGroupEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
            onChange={(e: ChangeEvent) => setEmail(e.target.value)}
            type="email"
            placeholder="Enter email"
          />
        </Form.Group>
        <Form.Group controlId="formGroupPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            onKeyPress={handleKeyPress}
            onChange={(e: ChangeEvent) => setPassword(e.target.value)}
            type="password"
            placeholder="Password"
          />
        </Form.Group>
      </Form>
      <ErrorMessage show={hasError} error={error} />
      <Button
        onClick={() => HandleLogin(email, password)}
        style={{ width: '5em' }}
        variant="primary"
      >
        Submit
      </Button>
    </Card>
  );
};

export default LoginForm;
