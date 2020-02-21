import React, { useState, useEffect } from 'react';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import SuccessMessage from '../../SuccessMessage';
import APIError from '../../../interfaces/error';
import ButtonMain from '../../ButtonMain';
import FormInput from '../../FormInput';
import { IconContext } from 'react-icons';
import { MdAccountCircle } from 'react-icons/md';
import Link from 'next/link';

interface LoginFormProps {
  signedup?: boolean;
  error: APIError;
  HandleLogin: (body: string) => Promise<void>;
}

type TFormEvent = React.FormEvent<HTMLFormElement>;

const LoginForm: React.FC<LoginFormProps> = ({ HandleLogin, signedup, error }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [validated, setValidated] = useState(false);

  const handleSubmit = (event: TFormEvent) => {
    event.stopPropagation();
    event.preventDefault();
    const form = event.currentTarget;
    if (form.checkValidity() === true) {
      HandleLogin(JSON.stringify({ email, password }));
    }
    // this sets form validation feedback to visible
    setValidated(true);
  };

  return (
    <Card style={{ padding: 20 }}>
      <Card.Body>
        <Card.Title className="card-title-form">
          <h2 className="center">Login</h2>
        </Card.Title>
        <SuccessMessage show={!!signedup} message="Sign up complete, please log in" />
        <Form noValidate validated={validated} onSubmit={(e: TFormEvent) => handleSubmit(e)}>
          <Form.Row>
            <FormInput
              required
              type="email"
              name="email"
              label="Email"
              value={email}
              onChange={setEmail}
              feebackType="invalid"
              feebackValue="Not a valid email"
            />
          </Form.Row>
          <Form.Row>
            <FormInput
              required
              type="password"
              name="password"
              label="Password"
              value={password}
              onChange={setPassword}
              feebackType="invalid"
              feebackValue="Required"
            />
          </Form.Row>
          <ErrorMessage show={!!error} error={error} />
          <ButtonMain className="w-100" type="submit" value="Submit" />
        </Form>

        <div className="login-links-cont">
          <Link href="/forgotpass">
            <a className="forgot-pass-link">Forgot Password?</a>
          </Link>
          <Link href="/signup">
            <a className="create-account-link">
              <IconContext.Provider value={{ size: '1.5em' }}>
                <MdAccountCircle />
              </IconContext.Provider>
              Create Account
            </a>
          </Link>
        </div>
      </Card.Body>
    </Card>
  );
};

export default LoginForm;
