import React, { useState, useEffect } from 'react';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';
import ButtonMain from '../../ButtonMain';
import FormInput from './FormInput';
import Link from 'next/link';

interface SignupFormProps {
  error: APIError | undefined;
  HandleSignup: (body: string) => Promise<void>;
}

type TFormEvent = React.FormEvent<HTMLFormElement>;

const LoginForm: React.FC<SignupFormProps> = ({ error, HandleSignup }) => {
  const [err, setError] = useState(error);
  const [email, setEmail] = useState();
  const [password, setPassword] = useState();
  const [confirmPassword, setConfirmPassword] = useState();
  const [validated, setValidated] = useState(false);

  // update err with the response from the server if there is any
  useEffect(() => {
    setError(error);
  }, [error]);

  const handleSubmit = (event: TFormEvent) => {
    event.stopPropagation();
    event.preventDefault();
    const form = event.currentTarget;
    if (form.checkValidity() === true) {
      if (password == confirmPassword) {
        setError(undefined);
        HandleSignup(JSON.stringify({ email, password }));
      } else {
        setError({
          status: 0,
          code: '',
          title: '',
          detail: 'Passwords do not match'
        });
      }
    }
    // this sets form validation feedback to visible
    setValidated(true);
  };

  return (
    <Card style={{ padding: 20 }}>
      <Card.Body>
        <Card.Title className="card-title-form">
          <h2 className="center">Create an account</h2>
        </Card.Title>
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
            <FormInput
              required
              type="password"
              name="password"
              label="Confirm Password"
              value={confirmPassword}
              onChange={setConfirmPassword}
              feebackType="invalid"
              feebackValue="Required"
            />
          </Form.Row>
          <ErrorMessage show={!!err} error={err} />
          <ButtonMain className="w-100" type="submit" value="Submit" />
        </Form>
        <div className="signup-links-cont">
          Already have an account?
          <Link href="/login">
            <a className="login-link">Login</a>
          </Link>
        </div>
      </Card.Body>
    </Card>
  );
};

export default LoginForm;
