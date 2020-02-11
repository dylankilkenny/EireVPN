import React from 'react';
import Button from 'react-bootstrap/Button';

interface ButtonMainProps {
  value: string;
  type?: 'submit';
  className?: string;
  onClick?: () => void;
}

type Event = React.ChangeEvent<HTMLInputElement>;

const ButtonMain: React.FC<ButtonMainProps> = ({
  value,
  className,
  type,
  onClick
}): JSX.Element => {
  return (
    <Button
      type={type ? type : 'button'}
      onClick={onClick}
      className={`btn-main ${className ? className : ''}`}
      variant="outline-secondary"
    >
      {value}
    </Button>
  );
};

export default ButtonMain;
