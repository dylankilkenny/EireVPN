import Alert from 'react-bootstrap/Alert';
import APIError from '../interfaces/error';

interface SMProps {
  show: boolean;
  message: string;
  className?: string;
}

const SuccessMessage: React.FC<SMProps> = ({ show, message, className }): JSX.Element => {
  if (!className) {
    className = '';
  }
  if (!show) {
    return <div />;
  }
  return (
    <Alert className={`center ${className}`} variant="success">
      {message}
    </Alert>
  );
};

export default SuccessMessage;
