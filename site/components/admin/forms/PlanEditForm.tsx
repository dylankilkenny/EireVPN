import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Plan from '../../../interfaces/plan';
import Card from 'react-bootstrap/Card';
import dayjs from 'dayjs';
import APIError from '../../../interfaces/error';
import FormInput from '../../FormInput';
import SuccessMessage from '../../SuccessMessage';
import ErrorMessage from '../../ErrorMessage';
import ButtonMain from '../../ButtonMain';

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

  const handleDeleteClick = () => {
    HandleDelete();
  };

  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title className="card-title-form">
            Edit Plan
            <div className="button-toolbar">
              <ButtonMain onClick={handleSaveClick} value="Save" />
              <ButtonMain onClick={handleDeleteClick} value="Delete" />
            </div>
          </Card.Title>

          <ErrorMessage show={hasError} error={error} />
          <SuccessMessage show={success} message="Server Updated Successfully" />
          <Form>
            <Form.Row>
              <FormInput
                textOnly
                name="createdAt"
                label="Created At"
                value={dayjs(plan.createdAt)
                  .format('DD-MM-YYYY H:mm')
                  .toString()}
              />
              <FormInput
                textOnly
                name="updateAt"
                label="Updated At"
                value={dayjs(plan.updatedAt)
                  .format('DD-MM-YYYY H:mm')
                  .toString()}
              />
            </Form.Row>
            <Form.Row>
              <FormInput name="name" label="Name" value={name} onChange={setName} />
              <FormInput textOnly name="amount" label="Amount" value={plan.amount.toString()} />
            </Form.Row>
            <Form.Row>
              <FormInput textOnly name="interval" label="Interval" value={plan.interval} />
              <FormInput
                textOnly
                name="interval_count"
                label="Interval Count"
                value={plan.interval_count.toString()}
              />
            </Form.Row>
            <Form.Row>
              <FormInput textOnly name="plan_type" label="Plan Type" value={plan.plan_type} />
              <FormInput textOnly name="currency" label="Currency" value={plan.currency} />
            </Form.Row>
            <Form.Row>
              <FormInput
                textOnly
                name="stripe_plan_id"
                label="Stripe Plan ID"
                value={plan.stripe_plan_id}
              />
              <FormInput
                textOnly
                name="stripe_product_id"
                label="Stripe Product ID"
                value={plan.stripe_product_id}
              />
            </Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default UserForm;
