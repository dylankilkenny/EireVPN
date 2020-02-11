import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';

type Event = React.ChangeEvent<HTMLInputElement>;
type KeyEvent = React.KeyboardEvent<HTMLInputElement>;

interface FormGroupProps {
  name: string;
  label: string;
  value: string;
  required?: boolean;
  type?: string;
  className?: string;
  textOnly?: boolean;
  isInvalid?: boolean;
  feebackType?: 'valid' | 'invalid';
  feebackValue?: string;
  onKeyPress?: (target: KeyEvent) => void;
  onChange?: (value: React.SetStateAction<string>) => void;
}

const FormGroup: React.FC<FormGroupProps> = ({
  name,
  label,
  value,
  required,
  type,
  className,
  textOnly,
  isInvalid,
  feebackType,
  feebackValue,
  onKeyPress,
  onChange
}): JSX.Element => {
  const plainText = textOnly ? true : false;

  return (
    <Form.Group as={Col} controlId={name}>
      <Form.Label sm="2">{label}</Form.Label>
      {console.log(isInvalid)}
      <Form.Control
        required={required}
        className={`${className}`}
        readOnly={plainText}
        size="sm"
        type={type ? type : 'text'}
        placeholder={label}
        name={name}
        isInvalid={isInvalid}
        value={value}
        onKeyPress={onKeyPress ? onKeyPress : undefined}
        onChange={(e: Event) => {
          if (onChange) {
            onChange(e.target.value);
          }
        }}
      />
      <Form.Control.Feedback type={feebackType}>{feebackValue}</Form.Control.Feedback>
    </Form.Group>
  );
};

export default FormGroup;
