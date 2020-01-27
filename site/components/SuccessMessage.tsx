import Alert from 'react-bootstrap/Alert';
import APIError from '../interfaces/error';

interface SMProps {
  show: boolean;
  message: string;
}

const SuccessMessage: React.FC<SMProps> = ({ show, message }): JSX.Element => {
  if (!show) {
    return <div />;
  }
  return <Alert variant="success">{message}</Alert>;
};

export default SuccessMessage;
