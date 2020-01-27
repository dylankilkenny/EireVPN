import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import FormGroup from './FormGroup';
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
    <div className="server-create">
      <Form onSubmit={(e: TFormEvent) => handleSubmit(e)}>
        <div className="button-toolbar">
          <Button type="submit" variant="outline-secondary">
            Save
          </Button>
        </div>
        <ErrorMessage show={hasError} error={error} />
        <Form.Row>
          <FormGroup name="country" label="Country" value={country} onChange={setCountry} />
          <FormGroup
            name="country_code"
            label="Country Code"
            value={countryCode}
            onChange={setCountryCode}
          />
        </Form.Row>
        <Form.Row>
          <FormGroup name="ip" label="IP" value={ip} onChange={setIp} />
          <FormGroup name="port" label="Port" value={port} onChange={setPort} />
        </Form.Row>
        <Form.Row>
          <FormGroup name="username" label="Username" value={username} onChange={setUsername} />
          <FormGroup name="password" label="Password" value={password} onChange={setPassword} />
        </Form.Row>
        <Form.Row>
          <Form.Group as={Row} controlId="img">
            <Form.Label column sm="2">
              Image
            </Form.Label>
            <Col sm="5">
              <input type="file" id="img" name="img" />
            </Col>
          </Form.Group>
        </Form.Row>
      </Form>
    </div>
  );
};

export default ServerCreateForm;
