import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Image from 'react-bootstrap/Image';
import Server from '../../../interfaces/server';
import Button from 'react-bootstrap/Button';
import dayjs from 'dayjs';
import FormGroup from './FormGroup';
import ErrorMessage from '../../ErrorMessage';
import SuccessMessage from '../../SuccessMessage';
import APIError from '../../../interfaces/error';

interface ServerEditFormProps {
  server: Server;
  error: APIError;
  success: boolean;
  HandleSave: (body: string) => Promise<void>;
  HandleDelete: () => Promise<void>;
}

const ServerEditForm: React.FC<ServerEditFormProps> = ({
  server,
  error,
  success,
  HandleSave,
  HandleDelete
}) => {
  const hasError = !!error;
  const [ip, setIp] = useState(server.ip);
  const [portString, setPortString] = useState(server.port.toString());
  const [username, setUsername] = useState(server.username);
  const [password, setPassword] = useState(server.password);

  const handleSaveClick = () => {
    const port = parseInt(portString);
    HandleSave(JSON.stringify({ ip, port, username, password }));
  };

  return (
    <div>
      <div className="button-toolbar">
        <Button onClick={handleSaveClick} className="btn-save" variant="outline-secondary">
          Save
        </Button>
        <Button onClick={HandleDelete} variant="outline-secondary">
          Delete
        </Button>
      </div>
      <ErrorMessage show={hasError} error={error} />
      <SuccessMessage show={success} message="Server Updated Successfully" />
      <Form>
        <Form.Row>
          <FormGroup
            textOnly
            name="createdAt"
            label="Created At"
            value={dayjs(server.createdAt)
              .format('DD-MM-YYYY H:mm')
              .toString()}
          />
          <FormGroup
            textOnly
            name="updateAt"
            label="Updated At"
            value={dayjs(server.updatedAt)
              .format('DD-MM-YYYY H:mm')
              .toString()}
          />
        </Form.Row>
        <Form.Row>
          <FormGroup textOnly name="Country" label="Country" value={server.country} />
          <FormGroup
            textOnly
            name="country_code"
            label="Country Code"
            value={server.country_code.toString()}
          />
        </Form.Row>
        <Form.Row>
          <FormGroup name="ip" label="Ip" value={ip} onChange={setIp} />
          <FormGroup
            name="Port"
            label="port"
            value={portString.toString()}
            onChange={setPortString}
          />
        </Form.Row>
        <Form.Row>
          <FormGroup name="username" label="Username" value={username} onChange={setUsername} />
          <FormGroup name="password" label="Password" value={password} onChange={setPassword} />
        </Form.Row>
        <Form.Row>
          <Form.Group as={Col} controlId="image_path">
            <Form.Label column sm="2">
              Image
            </Form.Label>
            <Col sm="5">
              <Image src={`http://localhost:3000/${server.image_path}`} rounded />
            </Col>
          </Form.Group>
        </Form.Row>
      </Form>
    </div>
  );
};

export default ServerEditForm;
