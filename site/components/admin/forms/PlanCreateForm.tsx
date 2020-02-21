import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import FormInput from '../../FormInput';
import FormDropdown from '../../FormDropdown';
import ButtonMain from '../../ButtonMain';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';

interface PlanCreateFormProps {
  error: APIError;
  HandleSave: (body: string) => Promise<void>;
}

type Event = React.ChangeEvent<HTMLInputElement>;

const UserCreateForm: React.FC<PlanCreateFormProps> = ({ error, HandleSave }) => {
  const [name, setName] = useState('');
  const [amountString, setAmount] = useState('');
  const [interval, setInterval] = useState('');
  const [intervalCountString, setIntervalCount] = useState('');
  const [plan_type, setPlanType] = useState('');

  const hasError = !!error;

  const handleSaveClick = () => {
    const amount = parseInt(amountString);
    const interval_count = parseInt(intervalCountString);
    const currency = 'EUR';
    HandleSave(JSON.stringify({ name, amount, interval, interval_count, plan_type, currency }));
  };

  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title className="card-title-form">
            Create Plan
            <div className="button-toolbar">
              <ButtonMain onClick={handleSaveClick} value="Save" />
            </div>
          </Card.Title>
          <Form>
            <ErrorMessage show={hasError} error={error} />
            <Form.Row>
              <FormInput name="name" label="Name" value={name} onChange={setName} />
              <FormInput name="amount" label="Amount" value={amountString} onChange={setAmount} />
            </Form.Row>
            <Form.Row>
              <FormInput name="interval" label="Interval" value={interval} onChange={setInterval} />
              <FormInput
                name="intervalCount"
                label="Interval Count"
                value={intervalCountString}
                onChange={setIntervalCount}
              />
            </Form.Row>
            <Form.Row>
              <FormDropdown
                name="plantType"
                label="Plant Type"
                value={plan_type}
                options={['PAYG', 'SUB', 'FREE']}
                onChange={setPlanType}
              />
            </Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default UserCreateForm;
