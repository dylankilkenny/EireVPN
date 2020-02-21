import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Image from 'react-bootstrap/Image';
import Server from '../../../interfaces/server';
import Card from 'react-bootstrap/Card';
import dayjs from 'dayjs';
import FormInput from '../../FormInput';
import ErrorMessage from '../../ErrorMessage';
import SuccessMessage from '../../SuccessMessage';
import APIError from '../../../interfaces/error';
import ButtonMain from '../../ButtonMain';

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

  const handleDeleteClick = () => {
    HandleDelete();
  };

  return (
    <div>
      <Form>
        <Card>
          <Card.Body>
            <Card.Title className="card-title-form">
              Edit Server
              <div className="button-toolbar">
                <ButtonMain onClick={handleSaveClick} value="Save" />
                <ButtonMain onClick={handleDeleteClick} value="Delete" />
              </div>
            </Card.Title>
            <ErrorMessage show={hasError} error={error} />
            <SuccessMessage show={success} message="Server Updated Successfully" />
            <Form.Row>
              <FormInput
                textOnly
                name="createdAt"
                label="Created At"
                value={dayjs(server.createdAt)
                  .format('DD-MM-YYYY H:mm')
                  .toString()}
              />
              <FormInput
                textOnly
                name="updateAt"
                label="Updated At"
                value={dayjs(server.updatedAt)
                  .format('DD-MM-YYYY H:mm')
                  .toString()}
              />
            </Form.Row>
            <Form.Row>
              <FormInput textOnly name="Country" label="Country" value={server.country} />
              <FormInput
                textOnly
                name="country_code"
                label="Country Code"
                value={server.country_code.toString()}
              />
            </Form.Row>
            <Form.Row>
              <FormInput name="ip" label="Ip" value={ip} onChange={setIp} />
              <FormInput
                name="Port"
                label="port"
                value={portString.toString()}
                onChange={setPortString}
              />
            </Form.Row>
            <Form.Row>
              <FormInput name="username" label="Username" value={username} onChange={setUsername} />
              <FormInput name="password" label="Password" value={password} onChange={setPassword} />
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
          </Card.Body>
        </Card>
      </Form>
    </div>
  );
};

export default ServerEditForm;
