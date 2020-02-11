import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import User from '../../../interfaces/user';
import ButtonMain from '../../ButtonMain';
import Card from 'react-bootstrap/Card';
import FormInput from '../../admin/forms/FormInput';

interface UserEditFormProps {
  user: User;
  HandleSave: (body: string) => Promise<void>;
}

const UserForm: React.FC<UserEditFormProps> = ({ user, HandleSave }) => {
  const [firstname, setFirstname] = useState(user.firstname);
  const [lastname, setLastname] = useState(user.lastname);
  const [email, setEmail] = useState(user.email);

  const handleSaveClick = () => {
    HandleSave(JSON.stringify({ firstname, lastname, email }));
  };

  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title className="card-title-form">
            Details
            <div className="button-toolbar">
              <ButtonMain onClick={handleSaveClick} value="Save" />
            </div>
          </Card.Title>
          <Form>
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
              <FormInput
                className="w-75"
                name="email"
                label="Email"
                value={email}
                onChange={setEmail}
              />
            </Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default UserForm;
