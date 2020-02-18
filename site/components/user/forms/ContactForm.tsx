import React, { useState, useEffect } from 'react';
import Form from 'react-bootstrap/Form';
import ButtonMain from '../../ButtonMain';
import Card from 'react-bootstrap/Card';
import ErrorMessage from '../../ErrorMessage';
import FormInput from '../../admin/forms/FormInput';

interface ContactFormProps {
  success: boolean;
  HandleSend: (body: string) => Promise<void>;
}

type TFormEvent = React.FormEvent<HTMLFormElement>;

const ContactForm: React.FC<ContactFormProps> = ({ success, HandleSend }) => {
  const [email, setEmail] = useState();
  const [subject, setSubject] = useState();
  const [message, setMessage] = useState();
  const [validated, setValidated] = useState(false);

  // if saved successfully clear password fields
  useEffect(() => {
    if (success) {
      setEmail('');
      setSubject('');
      setMessage('');
      setValidated(false);
    }
  });

  const handleSubmit = (event: TFormEvent) => {
    event.stopPropagation();
    event.preventDefault();
    const form = event.currentTarget;
    // this sets form validation feedback to visible
    setValidated(true);
    if (form.checkValidity() === true) {
      HandleSend(JSON.stringify({ email, subject, message }));
    }
  };

  return (
    <div className="contact-form">
      <Card>
        <Card.Body>
          <Form noValidate validated={validated} onSubmit={(e: TFormEvent) => handleSubmit(e)}>
            <Card.Title className="card-title-form">
              Contact Support
              <div className="button-toolbar">
                <ButtonMain type="submit" value="Send" />
              </div>
            </Card.Title>
            <Form.Row>
              <FormInput
                required
                type="email"
                name="email"
                label="Email"
                value={email}
                onChange={setEmail}
                feebackType="invalid"
                feebackValue="Required"
              />
            </Form.Row>
            <Form.Row>
              <FormInput
                required
                name="subject"
                label="Subject"
                value={subject}
                onChange={setSubject}
                feebackType="invalid"
                feebackValue="Required"
              />
            </Form.Row>
            <Form.Row>
              <FormInput
                textarea
                required
                name="message"
                label="Message"
                value={message}
                onChange={setMessage}
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

export default ContactForm;
