import React, { useState, useEffect } from 'react';
import Form from 'react-bootstrap/Form';
import ButtonMain from '../../ButtonMain';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import FormInput from '../../FormInput';

interface ChangePasswordFormProps {
  success: boolean;
  HandleSave: (body: string) => Promise<void>;
}

type TFormEvent = React.FormEvent<HTMLFormElement>;

const ChangePasswordForm: React.FC<ChangePasswordFormProps> = ({ success, HandleSave }) => {
  const [err, setError] = useState();
  const [current_password, setCurrentPassword] = useState();
  const [new_password, setNewPassword] = useState();
  const [confirmNewPassword, setConfirmNewPassword] = useState();
  const [validated, setValidated] = useState(false);

  // if saved successfully clear password fields
  useEffect(() => {
    if (success) {
      setCurrentPassword('');
      setNewPassword('');
      setConfirmNewPassword('');
      setValidated(false);
    }
  });

  const handleSubmit = (event: TFormEvent) => {
    event.stopPropagation();
    event.preventDefault();
    const form = event.currentTarget;
    // this sets form validation feedback to visible
    setValidated(true);
    if (form.checkValidity() === true) {
      if (new_password == confirmNewPassword) {
        setError(undefined);
        HandleSave(JSON.stringify({ current_password, new_password }));
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

  return (
    <div>
      <Card>
        <Card.Body>
          <Form noValidate validated={validated} onSubmit={(e: TFormEvent) => handleSubmit(e)}>
            <Card.Title className="card-title-form">
              Change Password
              <div className="button-toolbar">
                <ButtonMain type="submit" value="Save" />
              </div>
            </Card.Title>
            <ErrorMessage show={!!err} error={err} />
            <Form.Row>
              <FormInput
                required
                type="password"
                name="current_password"
                label="Current Password"
                value={current_password}
                onChange={setCurrentPassword}
                feebackType="invalid"
                feebackValue="Required"
              />
            </Form.Row>
            <Form.Row>
              <FormInput
                required
                type="password"
                name="new_password"
                label="New Password"
                value={new_password}
                onChange={setNewPassword}
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
                value={confirmNewPassword}
                onChange={setConfirmNewPassword}
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

export default ChangePasswordForm;
