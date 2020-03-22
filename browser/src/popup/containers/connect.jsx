import React from 'react';
import Image from 'react-bootstrap/Image';
import Alert from 'react-bootstrap/Alert';
import Container from 'react-bootstrap/Container';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import { MdMyLocation } from 'react-icons/md';
import { FiShieldOff } from 'react-icons/fi';
import { IconContext } from 'react-icons';
import ApiService from '../services/apiService';
import sendMessage from '../services/comunicationManager';
import storage from '../../utils/storage';
import AuthService from '../services/authService';
import Loading from './loading';

const checkProxyError = async () => storage.get('proxy_error');

class Main extends React.Component {
  state = {
    loading: true
  };

  async componentDidMount() {
    const connected = await storage.get('connected');
    if (connected) {
      const ip = await storage.get('ip');
      this.setState({ server: this.props.server, loading: false, ip });
    } else {
      await storage.set({ proxy_error: false });
      const resp = await ApiService.connectServer(this.props.server.id);
      if (resp.status === 200) {
        const { username, password, ip, port } = resp.data;
        await storage.set({ connected: true, server: this.props.server, ip });
        sendMessage('connect', { ip, port, username, password });
        setTimeout(async () => {
          const err = await checkProxyError();
          if (!err) {
            this.setState({ server: this.props.server, loading: false, ip });
          } else {
            await storage.set({ proxy_error: false });
            this.setState({
              loading: false,
              error: true,
              alertMsg: 'Something Went Wrong!'
            });
          }
        }, 3000);
      } else if (resp.status === 403) {
        await AuthService.logout();
        this.props.renderLogin();
      } else {
        this.setState({
          loading: false,
          error: true,
          alertMsg: resp.detail
        });
      }
    }
  }

  disconnect = async () => {
    sendMessage('disconnect', {});
    this.props.renderMain();
  };

  render() {
    if (this.state.error) {
      return (
        <div>
          <Alert style={{ fontSize: 14 }} show={this.state.showAlert} variant="danger">
            {this.state.alertMsg}
          </Alert>
          <Container className="server-cont-disc" onClick={this.disconnect}>
            <Row>
              <Col className="center">Back to Server List</Col>
            </Row>
          </Container>
        </div>
      );
    }
    if (this.state.loading) {
      return <Loading />;
    }
    if (!this.state.server) return <div />;
    return (
      <Container className="connect-cont">
        <Row>
          <Container className="server-cont">
            <Row className="server-country-row">
              <Col className="server-country-col" xs="4">
                <Image
                  src={`${process.env.API_URL}/${this.state.server.image_path}`}
                  roundedCircle
                />
              </Col>
              <Col className="server-country-col" xs="4">
                {this.state.server.country_code}
              </Col>
            </Row>
          </Container>
          <Container className="server-cont">
            <Row className="server-country-row">
              <Col className="server-country-col" xs="2">
                <div className="status active" />
              </Col>
              <Col className="server-country-col" xs="4">
                Active
              </Col>
            </Row>
          </Container>
        </Row>
        <Row>
          <Container className="server-cont-ip">
            <Row className="server-ip-row">
              <Col className="server-ip-col" xs="2">
                <IconContext.Provider value={{ size: '1.5em' }}>
                  <MdMyLocation />
                </IconContext.Provider>
              </Col>
              <Col className="server-ip-col ip" xs="4">
                {this.state.ip}
              </Col>
            </Row>
          </Container>
        </Row>
        <Row>
          <Container className="server-cont-disc" onClick={this.disconnect}>
            <Row className="server-disc-row">
              <Col className="server-disc-col shield" xs="4">
                <IconContext.Provider value={{ size: '1.5em' }}>
                  <FiShieldOff style={{ float: 'right' }} />
                </IconContext.Provider>
              </Col>
              <Col className="server-disc-col disc" xs="7">
                Disconnect
              </Col>
            </Row>
          </Container>
        </Row>
      </Container>
    );
  }
}

export default Main;
