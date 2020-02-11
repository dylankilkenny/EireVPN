import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import UserPlan from '../../../interfaces/userplan';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';
import dayjs from 'dayjs';
import FormInput from './FormInput';
import FormDropdown from './FormDropdown';
import FormDatetime from './FormDatetime';
import SuccessMessage from '../../SuccessMessage';
import ButtonMain from '../../ButtonMain';

interface UserPlanEditFormProps {
  userplan: UserPlan;
  error: APIError;
  success: boolean;
  HandleSave: (body: string) => Promise<void>;
  HandleDelete: () => Promise<void>;
}

const UserPlanEditForm: React.FC<UserPlanEditFormProps> = ({
  userplan,
  error,
  success,
  HandleSave,
  HandleDelete
}) => {
  const hasError = !!error;
  const [active, setActive] = useState(userplan.active.toString());
  const [start_date, setStartDate] = useState(userplan.start_date);
  const [expiry_date, setExpiryDate] = useState(userplan.expiry_date);

  const handleSaveClick = () => {
    HandleSave(
      JSON.stringify({
        active,
        // datetime should be return in this format to be parsed
        // correctly on the server
        start_date: dayjs(start_date).format('YYYY-MM-DD H:mm'),
        expiry_date: dayjs(expiry_date).format('YYYY-MM-DD H:mm')
      })
    );
  };

  const handleDeleteClick = () => {
    HandleDelete();
  };

  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title className="card-title-form">
            User Plan Details
            <div className="button-toolbar">
              <ButtonMain onClick={handleSaveClick} value="Save" />
              <ButtonMain onClick={handleDeleteClick} value="Delete" />
            </div>
          </Card.Title>

          <ErrorMessage show={hasError} error={error} />
          <SuccessMessage show={success} message="User Plan Updated Successfully" />
          <Form>
            <Form.Row>
              <FormInput
                textOnly
                name="createdAt"
                label="Created At"
                value={dayjs(userplan.createdAt)
                  .format('DD-MM-YYYY H:mm')
                  .toString()}
              />
              <FormInput
                textOnly
                name="updateAt"
                label="Updated At"
                value={dayjs(userplan.updatedAt)
                  .format('DD-MM-YYYY H:mm')
                  .toString()}
              />
            </Form.Row>
            <Form.Row>
              <FormInput
                textOnly
                name="user_id"
                label="User ID"
                value={userplan.user_id.toString()}
              />
              <FormInput
                textOnly
                name="plan_id"
                label="Plan ID"
                value={userplan.plan_id.toString()}
              />
            </Form.Row>
            <Form.Row>
              <FormDatetime
                value={start_date}
                name="start_date"
                label="Start Date"
                onChange={setStartDate}
              />
              <FormDatetime
                value={expiry_date}
                name="expiry_date"
                label="End Date"
                onChange={setExpiryDate}
              />
              <FormDropdown
                name="active"
                label="Active"
                value={active}
                options={['true', 'false']}
                onChange={setActive}
              />
            </Form.Row>
            <Form.Row></Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default UserPlanEditForm;
