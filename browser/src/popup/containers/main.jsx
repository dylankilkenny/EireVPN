import React from 'react';
import Image from 'react-bootstrap/Image';
import Spinner from 'react-bootstrap/Spinner';
import Container from 'react-bootstrap/Container';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import { IoIosArrowForward } from 'react-icons/io';
import { IconContext } from 'react-icons';
import ApiService from '../services/apiService';
import AuthService from '../services/authService';

class Main extends React.Component {
  state = {};

  async componentDidMount() {
    const resp = await ApiService.getServers();
    if (resp.status === 200) {
      this.setState({ servers: resp.data.servers });
    } else if (resp.status === 403) {
      await AuthService.logout();
      this.props.renderLogin();
    }
  }

  connect = server => {
    this.props.renderConnect(server);
  };

  render() {
    if (!this.state.servers) {
      return (
        <div className="spinner-div">
          <Spinner className="loading-spinner" animation="border" variant="primary" />
        </div>
      );
    }
    return (
      <Container className="servers-container">
        {this.state.servers.map(server => (
          <Row onClick={() => this.connect(server)} className="server-row" key={server.id}>
            <Col className="server-col" xs="2">
              <Image src={`${process.env.API_URL}/${server.image_path}`} roundedCircle />
            </Col>
            <Col className="server-col text" xs="7">
              {server.country}
            </Col>
            <Col className="server-col" xs="3">
              <IconContext.Provider value={{ size: '2em', color: '#e16162' }}>
                <div className="server-connect-arrow" title="settings">
                  <IoIosArrowForward style={{ cursor: 'pointer' }} />
                </div>
              </IconContext.Provider>
            </Col>
          </Row>
        ))}
      </Container>
    );
  }
}
export default Main;
