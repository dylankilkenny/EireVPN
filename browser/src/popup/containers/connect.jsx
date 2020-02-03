import React from 'react';
import Image from 'react-bootstrap/Image';
import Alert from 'react-bootstrap/Alert';
import Container from 'react-bootstrap/Container';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Spinner from 'react-bootstrap/Spinner';
import { MdMyLocation } from 'react-icons/md';
import { FiShieldOff } from 'react-icons/fi';
import { IconContext } from 'react-icons';
import ApiService from '../services/apiService';
import sendMessage from '../services/comunicationManager';
import storage from '../../utils/storage';
import AuthService from '../services/authService';

class Main extends React.Component {
  state = {
    connected: false
  };

  async componentDidMount() {
    const resp = await ApiService.connectServer(this.props.server.id);
    if (resp.status === 200) {
      const { username, password, ip, port } = resp.data;
      await storage.set({ connected: true, server: this.props.server });
      sendMessage('connect', { ip, port, username, password });
      this.setState({ server: this.props.server, connected: true, ip: ip });
    } else if (resp.status === 403) {
      await AuthService.logout();
      this.cleanStorage();
      this.setState({ connected: false }, () => this.props.renderLogin());
    } else {
      this.cleanStorage();
      this.setState({
        error: true,
        alertMsg: resp.detail,
        connected: false
      });
    }
  }

  cleanStorage = async () => {
    await storage.set({ connected: false, server: undefined });
    sendMessage('disconnect', {});
  };

  disconnect = async () => {
    if (this.state.connected) {
      this.cleanStorage();
      this.props.renderMain();
    }
  };

  render() {
    if (this.state.error) {
      return (
        <Alert style={{ fontSize: 14 }} show={this.state.showAlert} variant="danger">
          {this.state.alertMsg}
        </Alert>
      );
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
                {this.state.connected ? (
                  <div className="status active" />
                ) : (
                  <div className="status disabled" />
                )}
              </Col>
              <Col className="server-country-col" xs="4">
                {this.state.connected ? 'Active' : 'Disabled'}
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
                {this.state.connected ? (
                  <IconContext.Provider value={{ size: '1.5em' }}>
                    <FiShieldOff style={{ float: 'right' }} />
                  </IconContext.Provider>
                ) : (
                  <div />
                )}
              </Col>
              <Col className="server-disc-col disc" xs="7">
                {this.state.connected ? (
                  'Disconnect'
                ) : (
                  <Spinner className="connecting-spinner" animation="border" variant="primary" />
                )}
              </Col>
            </Row>
          </Container>
        </Row>
      </Container>
    );
  }
}

export default Main;
