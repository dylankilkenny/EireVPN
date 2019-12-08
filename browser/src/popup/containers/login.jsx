import React from 'react';
import PropTypes from 'prop-types';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Alert from 'react-bootstrap/Alert';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import ext from '../../utils/ext';
// import Cont from '../components/container';
import AuthService from '../services/authService';

class Login extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      email: '',
      password: '',
      showAlert: false,
    };
    this.login = this.login.bind(this);
    this.updateFormValue = this.updateFormValue.bind(this);
  }

  componentDidMount() {
    ext.storage.local.get('email', (resp) => {
      const { email } = resp;
      if (email) {
        this.setState({ email });
      }
    });
  }

  updateFormValue(evt, element) {
    ext.storage.local.set({ [element]: evt.target.value }, () => {});
    this.setState({
      [element]: evt.target.value,
    });
  }

  login() {
    AuthService.login(this.state.email, this.state.password).then((resp) => {
      if (resp.status === 200) {
        this.props.authorise(true);
      } else if (resp.status === 400 || resp.status === 401) {
        this.setState({ showAlert: true, alertMsg: 'Email/Password Incorrect.' });
      } else {
        this.setState({ showAlert: true, alertMsg: 'Something Went Wrong.' });
      }
    });
  }

  render() {
    return (
      <div>
        <Row>
          <Col>
            <Form>
              <Form.Group as={Row} controlId="formPlaintextEmail">
                <Form.Label column sm="1">
                  Email
                </Form.Label>
                <Col sm="10">
                  <Form.Control
                    size="sm"
                    value={this.state.email}
                    type="email"
                    placeholder="Enter email"
                    onChange={evt => this.updateFormValue(evt, 'email')}
                  />
                </Col>
              </Form.Group>
              <Form.Group as={Row} controlId="formPlaintextPassword">
                <Form.Label column sm="1">
                  Password
                </Form.Label>
                <Col sm="10">
                  <Form.Control
                    size="sm"
                    value={this.state.password}
                    type="password"
                    placeholder="Password"
                    onChange={evt => this.updateFormValue(evt, 'password')}
                  />
                </Col>
              </Form.Group>
            </Form>
            <Alert style={{ fontSize: 14 }} show={this.state.showAlert} variant="danger">
              {this.state.alertMsg}
            </Alert>
            <Button onClick={this.login} variant="secondary">
              Submit
            </Button>
          </Col>
        </Row>
      </div>
    );
  }
}

Login.propTypes = {
  authorise: PropTypes.func.isRequired,
};

export default Login;
