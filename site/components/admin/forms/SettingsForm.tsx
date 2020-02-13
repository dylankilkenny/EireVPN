import React, { useState } from 'react';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import Settings from '../../../interfaces/settings';
import ButtonMain from '../../ButtonMain';
import ErrorMessage from '../../ErrorMessage';
import APIError from '../../../interfaces/error';
import FormDropdown from './FormDropdown';
import FormInput from './FormInput';
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
  const [allowedOrigins, setAllowedOrigins] = useState(settings.allowedOrigins.join(',\n'));

  const handleSaveClick = () => {
    HandleSave(
      JSON.stringify({
        enableCsrf,
        enableSubscriptions,
        enableAuth,
        enableStripe,
        allowedOrigins: allowedOrigins.replace(/\r?\n|\r/g, '').split(',')
      })
    );
  };

  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title className="card-title-form">
            Edit Settings
            <div className="button-toolbar">
              <ButtonMain onClick={handleSaveClick} value="Save" />
            </div>
          </Card.Title>
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
            <Form.Row>
              <FormInput
                textarea
                name="allowed_origins"
                label="Allowed Origins"
                value={allowedOrigins}
                onChange={setAllowedOrigins}
                feebackType="invalid"
                feebackValue="Required"
              />
            </Form.Row>
          </Form>
        </Card.Body>
      </Card>
    </div>
  );
};

export default SettingsForm;
