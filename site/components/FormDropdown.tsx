import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';

interface FormDropdownProps {
  name: string;
  label: string;
  value: string;
  className?: string;
  options?: string[];
  optionsKV?: { value: string; name: string }[];
  onChange?: (value: React.SetStateAction<string>) => void;
}

type Event = React.ChangeEvent<HTMLSelectElement>;

const FormGroup: React.FC<FormDropdownProps> = ({
  name,
  label,
  value,
  options,
  optionsKV,
  className,
  onChange
}): JSX.Element => {
  let DropdownOptions;
  if (options) {
    options.unshift('...');
    DropdownOptions = options.map((val, i) => (
      <option key={i} value={val}>
        {val}
      </option>
    ));
  } else if (optionsKV) {
    optionsKV.unshift({ value: '', name: '...' });
    DropdownOptions = optionsKV.map((val, i) => (
      <option key={i} value={val.value}>
        {val.name}
      </option>
    ));
  } else {
    DropdownOptions = <div />;
  }
  return (
    <Form.Group as={Col} controlId={name}>
      <Form.Label sm="2">{label}</Form.Label>
      <Form.Control
        className={`${className}`}
        as="select"
        value={value}
        name={name}
        onChange={(e: Event) => {
          if (onChange) {
            onChange(e.target.value);
          }
        }}
      >
        {DropdownOptions}
      </Form.Control>
    </Form.Group>
  );
};

export default FormGroup;
