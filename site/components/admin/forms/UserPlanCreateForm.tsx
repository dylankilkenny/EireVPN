import React, { useState } from 'react';
import dayjs from 'dayjs';
import Form from 'react-bootstrap/Form';
import FormInput from '../../FormInput';
import FormDropdown from '../../FormDropdown';
import FormDatetime from '../../FormDatetime';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';
import Plan from '../../../interfaces/plan';
import ButtonMain from '../../ButtonMain';

interface UserPlanCreateFormProps {
  userid: string;
  planlist: Plan[];
  error: APIError;
  HandleSave: (body: string) => Promise<void>;
}

const UserPlanCreateForm: React.FC<UserPlanCreateFormProps> = ({
  userid,
  planlist,
  error,
  HandleSave
}) => {
  const [planid, setPlanID] = useState();
  const [active, setActive] = useState();
  const [start_date, setStartDate] = useState(new Date().toString());
  const [expiry_date, setExpiryDate] = useState(new Date().toString());

  const hasError = !!error;

  let plans = planlist.map((p, i) => {
    return {
      value: p.id,
      name: p.name
    };
  });

  const handleSaveClick = () => {
    HandleSave(
      JSON.stringify({
        user_id: parseInt(userid),
        plan_id: parseInt(planid),
        active,
        start_date: dayjs(start_date).format('YYYY-MM-DD H:mm'),
        expiry_date: dayjs(expiry_date).format('YYYY-MM-DD H:mm')
      })
    );
  };

  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title className="card-title-form">
            Create User Plan
            <div className="button-toolbar">
              <ButtonMain onClick={handleSaveClick} value="Save" />
            </div>
          </Card.Title>
          <Form>
            <ErrorMessage show={hasError} error={error} />
            <Form.Row>
              <FormInput textOnly name="user_id" label="User ID" value={userid} />
            </Form.Row>
            <Form.Row>
              <FormDropdown
                name="plan"
                label="Plan"
                value={planid}
                optionsKV={plans}
                onChange={setPlanID}
                className="active-dropdown"
              />
            </Form.Row>
            <Form.Row>
              <FormDropdown
                name="active"
                label="Active"
                value={active}
                options={['true', 'false']}
                onChange={setActive}
                className="active-dropdown"
              />
            </Form.Row>
            <Form.Row>
              <FormDatetime
                value={start_date}
                name="start_date"
                label="Start Date"
                onChange={setStartDate}
              />
            </Form.Row>
            <Form.Row>
              <FormDatetime
                value={expiry_date}
                name="expiry_date"
                label="End Date"
                onChange={setExpiryDate}
              />
            </Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default UserPlanCreateForm;
