import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import FormGroup from './FormGroup';
import Button from 'react-bootstrap/Button';
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
    <div className="server-create">
      <Form>
        <div className="button-toolbar">
          <Button onClick={handleSaveClick} variant="outline-secondary">
            Save
          </Button>
        </div>
        <ErrorMessage show={hasError} error={error} />
        <Form.Row>
          <FormGroup name="name" label="Name" value={name} onChange={setName} />
          <FormGroup name="amount" label="Amount" value={amountString} onChange={setAmount} />
        </Form.Row>
        <Form.Row>
          <FormGroup name="interval" label="Interval" value={interval} onChange={setInterval} />
          <FormGroup
            name="intervalCount"
            label="Interval Count"
            value={intervalCountString}
            onChange={setIntervalCount}
          />
        </Form.Row>
        <Form.Row>
          <FormGroup name="plantType" label="Plant Type" value={plan_type} onChange={setPlanType} />
        </Form.Row>
      </Form>
    </div>
  );
};

export default UserCreateForm;
