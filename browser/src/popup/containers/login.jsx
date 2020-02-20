import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Alert from 'react-bootstrap/Alert';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import { MdAccountCircle } from 'react-icons/md';
import { IconContext } from 'react-icons';
import API from '../services/apiService';
import storage from '../../utils/storage';
import ext from '../../utils/ext';

class Login extends React.Component {
  state = {
    email: '',
    password: '',
    showAlert: false
  };

  async componentDidMount() {
    console.log('login');
    const email = await storage.get('email');
    if (email) {
      this.setState({ email });
    }
  }

  updateFormValue = (evt, element) => {
    storage.set({ [element]: evt.target.value });
    this.setState({
      [element]: evt.target.value
    });
  };

  handleKeyPress = target => {
    if (target.charCode === 13) {
      this.login();
    }
  };

  login = async () => {
    const resp = await API.login(this.state.email, this.state.password);
    if (resp.status === 200) {
      this.props.renderMain();
    } else if (resp.status === 400 || resp.status === 401) {
      this.setState({
        showAlert: true,
        alertMsg: resp.detail
      });
    }
  };

  createAccount = () => {
    const newURL = `${process.env.DOMAIN}/signup`;
    ext.tabs.create({ url: newURL });
  };

  forgotPass = () => {
    const newURL = `${process.env.DOMAIN}/forgotpass`;
    ext.tabs.create({ url: newURL });
  };

  render() {
    return (
      <div>
        <Row className="login">
          <Col>
            <Form>
              <Form.Group as={Row} controlId="formPlaintextEmail">
                <Form.Label column sm="1" className="label">
                  Email
                </Form.Label>
                <Col sm="10">
                  <Form.Control
                    size="sm"
                    value={this.state.email}
                    type="email"
                    placeholder="Email"
                    onKeyPress={this.handleKeyPress}
                    onChange={evt => this.updateFormValue(evt, 'email')}
                  />
                </Col>
              </Form.Group>
              <Form.Group as={Row} controlId="formPlaintextPassword">
                <Form.Label column sm="1" className="label">
                  Password
                </Form.Label>
                <Col sm="10">
                  <Form.Control
                    size="sm"
                    value={this.state.password}
                    type="password"
                    placeholder="Password"
                    onKeyPress={this.handleKeyPress}
                    onChange={evt => this.updateFormValue(evt, 'password')}
                  />
                </Col>
              </Form.Group>
            </Form>
            <Alert style={{ fontSize: 14 }} show={this.state.showAlert} variant="danger">
              {this.state.alertMsg}
            </Alert>
            <Button onClick={this.login} variant="secondary" className="btn-custom">
              Submit
            </Button>
            <div onClick={() => this.forgotPass()} className="forgot-pass-link">
              Forgot Password?
            </div>
          </Col>
        </Row>
        <Row>
          <Col>
            <div className="create-acc-cont">
              <div onClick={() => this.createAccount()} className="create-account-link">
                <IconContext.Provider value={{ size: '1.5em' }}>
                  <MdAccountCircle />
                </IconContext.Provider>
                Create Free Account
              </div>
            </div>
          </Col>
        </Row>
      </div>
    );
  }
}

export default Login;
