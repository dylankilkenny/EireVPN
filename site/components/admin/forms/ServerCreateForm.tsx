import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Card from 'react-bootstrap/Card';
import FormInput from '../../FormInput';
import Button from 'react-bootstrap/Button';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';

interface ServerCreateFormProps {
  HandleSave: (body: FormData) => Promise<void>;
  error: APIError;
}

type TFormEvent = React.FormEvent<HTMLFormElement>;

const ServerCreateForm: React.FC<ServerCreateFormProps> = ({ HandleSave, error }) => {
  const [country, setCountry] = useState('');
  const [countryCode, setCountryCode] = useState('');
  const [ip, setIp] = useState('');
  const [port, setPort] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const hasError = !!error;
  const handleSubmit = (event: TFormEvent) => {
    event.preventDefault();
    console.log(event.target);
    const data = new FormData(event.target as HTMLFormElement);
    for (var pair of data.entries()) {
      console.log(pair[0] + ', ' + pair[1]);
    }
    HandleSave(data);
  };

  return (
    <div>
      <Form onSubmit={(e: TFormEvent) => handleSubmit(e)}>
        <Card>
          <Card.Body>
            <Card.Title className="card-title-form">
              Create Server
              <div className="button-toolbar">
                <Button type="submit" className="btn-main" variant="outline-secondary">
                  Save
                </Button>
              </div>
            </Card.Title>

            <ErrorMessage show={hasError} error={error} />
            <Form.Row>
              <FormInput name="country" label="Country" value={country} onChange={setCountry} />
              <FormInput
                name="country_code"
                label="Country Code"
                value={countryCode}
                onChange={setCountryCode}
              />
            </Form.Row>
            <Form.Row>
              <FormInput name="ip" label="IP" value={ip} onChange={setIp} />
              <FormInput name="port" label="Port" value={port} onChange={setPort} />
            </Form.Row>
            <Form.Row>
              <FormInput name="username" label="Username" value={username} onChange={setUsername} />
              <FormInput name="password" label="Password" value={password} onChange={setPassword} />
            </Form.Row>
            <Form.Row>
              <Form.Group as={Row} controlId="img">
                <Form.Label column sm="3">
                  Image
                </Form.Label>
                <Col sm="5">
                  <input type="file" id="img" name="img" />
                </Col>
              </Form.Group>
            </Form.Row>
          </Card.Body>
        </Card>
      </Form>
    </div>
  );
};

export default ServerCreateForm;
