import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';

interface FormGroupProps {
  name: string;
  label: string;
  value: string;
  textOnly?: boolean;
  onChange?: (value: React.SetStateAction<string>) => void;
}

type Event = React.ChangeEvent<HTMLInputElement>;

const FormGroup: React.FC<FormGroupProps> = ({
  name,
  label,
  value,
  textOnly,
  onChange
}): JSX.Element => {
  const plainText = textOnly ? true : false;
  return (
    <Form.Group as={Col} controlId={name}>
      <Form.Label sm="2">{label}</Form.Label>
      <Form.Control
        readOnly={plainText}
        size="sm"
        type="text"
        placeholder={label}
        name={name}
        value={value}
        onChange={(e: Event) => {
          if (onChange) {
            onChange(e.target.value);
          }
        }}
      />
    </Form.Group>
  );
};

export default FormGroup;
