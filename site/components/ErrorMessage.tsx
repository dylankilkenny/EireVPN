import Alert from 'react-bootstrap/Alert';
import APIError from '../interfaces/error';

interface EMProps {
  error: APIError | undefined;
  show: boolean;
}

const ErrorMessage: React.FC<EMProps> = ({ error, show }): JSX.Element => {
  if (!show) {
    return <div />;
  }
  return <Alert variant="danger">{error?.detail}</Alert>;
};

export default ErrorMessage;
