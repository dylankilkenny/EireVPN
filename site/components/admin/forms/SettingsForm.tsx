import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Settings from '../../../interfaces/settings';
import Button from 'react-bootstrap/Button';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';
import FormDropdown from './FormDropdown';
import SuccessMessage from '../../SuccessMessage';

interface SettingsFormProps {
  settings: Settings;
  error: APIError;
  success: boolean;
  HandleSave: (body: string) => Promise<void>;
}

const SettingsForm: React.FC<SettingsFormProps> = ({ settings, error, success, HandleSave }) => {
  const hasError = !!error;
  const [enableCsrf, setEnableCsrf] = useState(settings.enableCsrf);
  const [enableSubscriptions, setEnableSubs] = useState(settings.enableSubscriptions);
  const [enableAuth, setEnableAuth] = useState(settings.enableAuth);
  const [enableStripe, setEnableStripe] = useState(settings.enableStripe);

  const handleSaveClick = () => {
    HandleSave(
      JSON.stringify({
        enableCsrf,
        enableSubscriptions,
        enableAuth,
        enableStripe
      })
    );
  };

  return (
    <div>
      <div className="button-toolbar">
        <Button onClick={handleSaveClick} className="btn-save" variant="outline-secondary">
          Save
        </Button>
      </div>
      <ErrorMessage show={hasError} error={error} />
      <SuccessMessage show={success} message="Settings Updated Successfully" />
      <Form>
        <Form.Row>
          <FormDropdown
            name="enableCsrf"
            label="Enable CSRF"
            value={enableCsrf}
            options={['true', 'false']}
            onChange={setEnableCsrf}
          />
          <FormDropdown
            name="enableSubscriptions"
            label="Enable Subscriptions"
            value={enableSubscriptions}
            options={['true', 'false']}
            onChange={setEnableSubs}
          />
        </Form.Row>
        <Form.Row>
          <FormDropdown
            name="enableAuth"
            label="Enable Auth"
            value={enableAuth}
            options={['true', 'false']}
            onChange={setEnableAuth}
          />
          <FormDropdown
            name="enableStripe"
            label="Enable Stripe"
            value={enableStripe}
            options={['true', 'false']}
            onChange={setEnableStripe}
          />
        </Form.Row>
      </Form>
    </div>
  );
};

export default SettingsForm;
