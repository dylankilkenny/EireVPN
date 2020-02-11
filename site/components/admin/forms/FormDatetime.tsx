import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import DatePicker from 'react-datepicker';
import dayjs from 'dayjs';

interface FormDatetimeProps {
  name: string;
  label: string;
  value: string;
  onChange: (value: React.SetStateAction<string>) => void;
}

const FormGroup: React.FC<FormDatetimeProps> = ({ name, label, value, onChange }): JSX.Element => {
  return (
    <Form.Group as={Col} controlId={name}>
      <Form.Label sm="2">{label}</Form.Label>
      <div>
        <DatePicker
          className="form-control"
          selected={dayjs(value).toDate()}
          onChange={(date: Date) => onChange(dayjs(date).toString())}
          showTimeSelect
          timeFormat="HH:mm"
          timeIntervals={15}
          timeCaption="time"
          dateFormat="MMMM d, yyyy h:mm aa"
        />
      </div>
    </Form.Group>
  );
};

export default FormGroup;
