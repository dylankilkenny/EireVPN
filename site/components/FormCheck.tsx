import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import DatePicker from 'react-datepicker';
import dayjs from 'dayjs';

interface FormDatetimeProps {
  name: string;
  label?: string;
  labelEl?: JSX.Element;
  feedback: string;
  className?: string;
}

const FormGroup: React.FC<FormDatetimeProps> = ({
  name,
  label,
  labelEl,
  feedback,
  className
}): JSX.Element => {
  if (!!className) className = '';
  return (
    <Form.Group as={Col} controlId={name}>
      <Form.Check
        className={`${className}`}
        required
        name={name}
        label={label ? label : labelEl}
        feedback={feedback}
        id="name"
      />
    </Form.Group>
  );
};

export default FormGroup;
