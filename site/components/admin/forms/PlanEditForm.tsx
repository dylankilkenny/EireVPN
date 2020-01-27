import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Plan from '../../../interfaces/plan';
import Button from 'react-bootstrap/Button';
import dayjs from 'dayjs';
import APIError from '../../../interfaces/error';
import FormGroup from './FormGroup';
import SuccessMessage from '../../SuccessMessage';
import ErrorMessage from '../../ErrorMessage';

interface PlanEditFormProps {
  plan: Plan;
  error: APIError;
  success: boolean;
  HandleSave: (body: string) => Promise<void>;
  HandleDelete: () => Promise<void>;
}

const UserForm: React.FC<PlanEditFormProps> = ({
  plan,
  error,
  success,
  HandleSave,
  HandleDelete
}) => {
  const hasError = !!error;
  const [name, setName] = useState(plan.name);

  const handleSaveClick = () => {
    HandleSave(JSON.stringify({ name }));
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
            value={dayjs(plan.createdAt)
              .format('DD-MM-YYYY H:mm')
              .toString()}
          />
          <FormGroup
            textOnly
            name="updateAt"
            label="Updated At"
            value={dayjs(plan.updatedAt)
              .format('DD-MM-YYYY H:mm')
              .toString()}
          />
        </Form.Row>
        <Form.Row>
          <FormGroup name="name" label="Name" value={name} onChange={setName} />
          <FormGroup textOnly name="amount" label="Amount" value={plan.amount.toString()} />
        </Form.Row>
        <Form.Row>
          <FormGroup textOnly name="interval" label="Interval" value={plan.interval} />
          <FormGroup
            textOnly
            name="interval_count"
            label="Interval Count"
            value={plan.interval_count.toString()}
          />
        </Form.Row>
        <Form.Row>
          <FormGroup textOnly name="plan_type" label="Plan Type" value={plan.plan_type} />
          <FormGroup textOnly name="currency" label="Currency" value={plan.currency} />
        </Form.Row>
        <Form.Row>
          <FormGroup
            textOnly
            name="stripe_plan_id"
            label="Stripe Plan ID"
            value={plan.stripe_plan_id}
          />
          <FormGroup
            textOnly
            name="stripe_product_id"
            label="Stripe Product ID"
            value={plan.stripe_product_id}
          />
        </Form.Row>
      </Form>
    </div>
  );
};

export default UserForm;
