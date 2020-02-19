import React, { useState, useEffect } from 'react';
import Form from 'react-bootstrap/Form';
import ButtonMain from '../../ButtonMain';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import FormInput from '../../admin/forms/FormInput';
import APIError from '../../../interfaces/error';
import SuccessMessage from '../../SuccessMessage';

interface UpdatePasswordProps {
  success: boolean;
  error: APIError;
  HandleSubmit: (body: string) => Promise<void>;
}

type TFormEvent = React.FormEvent<HTMLFormElement>;

const UpdatePasswordForm: React.FC<UpdatePasswordProps> = ({ success, error, HandleSubmit }) => {
  const [err, setError] = useState();
  const [password, setPassword] = useState();
  const [confirmPassword, setConfirmPassword] = useState();
  const [validated, setValidated] = useState(false);

  const handleSubmit = (event: TFormEvent) => {
    event.stopPropagation();
    event.preventDefault();
    const form = event.currentTarget;
    // this sets form validation feedback to visible
    setValidated(true);
    if (form.checkValidity() === true) {
      if (password == confirmPassword) {
        setError(undefined);
        HandleSubmit(JSON.stringify({ password }));
      } else {
        setError({
          status: 0,
          code: '',
          title: '',
          detail: 'Passwords do not match'
        });
        setValidated(false);
      }
    }
  };

  if (!success) {
    return (
      <Card className="password-reset-card">
        <SuccessMessage show={true} message="Password updated successfully." />
      </Card>
    );
  }

  if (error) {
    return (
      <Card className="password-reset-card">
        <ErrorMessage show={true} error={error} />
      </Card>
    );
  }

  return (
    <div>
      <Card className="password-reset-card">
        <Card.Body>
          <Form noValidate validated={validated} onSubmit={(e: TFormEvent) => handleSubmit(e)}>
            <Card.Title className="card-title-form">
              Reset Password
              <div className="button-toolbar">
                <ButtonMain type="submit" value="Save" />
              </div>
            </Card.Title>
            <p>Please enter your new password</p>
            <ErrorMessage show={!!err} error={err} />
            <Form.Row>
              <FormInput
                required
                type="password"
                name="new_password"
                label="New Password"
                value={password}
                onChange={setPassword}
                feebackType="invalid"
                feebackValue="Required"
              />
            </Form.Row>
            <Form.Row>
              <FormInput
                required
                type="password"
                name="confirm_password"
                label="Confirm Password"
                value={confirmPassword}
                onChange={setConfirmPassword}
                feebackType="invalid"
                feebackValue="Required"
              />
            </Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default UpdatePasswordForm;
