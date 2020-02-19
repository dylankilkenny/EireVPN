import React, { useState, useEffect } from 'react';
import Form from 'react-bootstrap/Form';
import ButtonMain from '../../ButtonMain';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import FormInput from '../../admin/forms/FormInput';
import APIError from '../../../interfaces/error';
import SuccessMessage from '../../SuccessMessage';

interface ForgotPasswordProps {
  success: boolean;
  error: APIError;
  HandleSubmit: (body: string) => Promise<void>;
}

type TFormEvent = React.FormEvent<HTMLFormElement>;

const ForgotPasswordForm: React.FC<ForgotPasswordProps> = ({ success, error, HandleSubmit }) => {
  const [email, setEmail] = useState();
  const [validated, setValidated] = useState(false);

  const handleSubmit = (event: TFormEvent) => {
    event.stopPropagation();
    event.preventDefault();
    const form = event.currentTarget;
    // this sets form validation feedback to visible
    setValidated(true);
    if (form.checkValidity() === true) {
      HandleSubmit(JSON.stringify({ email }));
    }
  };

  return (
    <div>
      <Card className="password-reset-card">
        <Card.Body>
          {success ? (
            <div>
              <p>
                If a matching account with that email has been found we will send you a reset link.
              </p>
              <p>
                <strong>If you do not recieve a email, please check your spam folder</strong>
              </p>
            </div>
          ) : (
            <Form noValidate validated={validated} onSubmit={(e: TFormEvent) => handleSubmit(e)}>
              <Card.Title className="card-title-form">
                Forgotten Password
                <div className="button-toolbar">
                  <ButtonMain type="submit" value="Save" />
                </div>
              </Card.Title>
              <p>Please enter your email and we will send you a reset link</p>
              <ErrorMessage show={!!error} error={error} />
              <Form.Row>
                <FormInput
                  required
                  type="email"
                  name="email"
                  label="Email"
                  value={email}
                  onChange={setEmail}
                  feebackType="invalid"
                  feebackValue="Required"
                />
              </Form.Row>
            </Form>
          )}
        </Card.Body>
      </Card>
    </div>
  );
};

export default ForgotPasswordForm;
