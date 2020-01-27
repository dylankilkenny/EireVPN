import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';

interface FormDropdownProps {
  name: string;
  label: string;
  value: string;
  options: string[];
  onChange?: (value: React.SetStateAction<string>) => void;
}

type Event = React.ChangeEvent<HTMLSelectElement>;

const FormGroup: React.FC<FormDropdownProps> = ({
  name,
  label,
  value,
  options,
  onChange
}): JSX.Element => {
  return (
    <Form.Group as={Col} controlId={name}>
      <Form.Label sm="2">{label}</Form.Label>
      <Form.Control
        as="select"
        value={value}
        name={name}
        onChange={(e: Event) => {
          if (onChange) {
            onChange(e.target.value);
          }
        }}
      >
        {options.map((val, i) => (
          <option key={i} value={val}>
            {val}
          </option>
        ))}
      </Form.Control>
    </Form.Group>
  );
};

export default FormGroup;
