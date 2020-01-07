import React from 'react';
import Image from 'react-bootstrap/Image';
import Button from 'react-bootstrap/Button';
import Table from 'react-bootstrap/Table';
import styled from 'styled-components';
import Alert from 'react-bootstrap/Alert';
import ApiService from '../services/apiService';
import sendMessage from '../services/comunicationManager';
import ext from '../../utils/ext';

const ServersContainer = styled.div`
  height: 200px;
  overflow: scroll;
  // border: 1px solid #ccc;
`;

class Main extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      connectedTo: '',
      showAlert: false
    };
  }

  componentWillMount() {}

  componentDidMount() {
    ApiService.getServers()
      .then(resp => this.setState({ servers: resp.data.servers }))
      .then(() => {
        ext.storage.local.get('connectedTo', resp => {
          const { connectedTo } = resp;
          if (connectedTo) {
            this.setState({ connectedTo });
          }
        });
      });
  }

  connect(serverid) {
    ApiService.connectServer(serverid)
      .then(resp => {
        if (resp.status === 401) {
          throw Error;
        }
        const { username } = resp.data;
        const { password } = resp.data;
        const { ip } = resp.data;
        const { port } = resp.data;
        sendMessage('connect', { ip, port, username, password });
        ext.storage.local.set({ connectedTo: serverid }, () => {
          this.setState({ connectedTo: serverid });
        });
      })
      .catch(() => {
        this.setState({
          showAlert: true,
          alertMsg:
            'You do not have an active subscription. Please purchase a subscription to continue using this service.'
        });
      });
  }

  disconnect() {
    sendMessage('disconnect', {});
    ext.storage.local.set({ connectedTo: '' }, () => {
      this.setState({ connectedTo: '' });
    });
  }

  render() {
    if (!this.state.servers) return <div />;
    return (
      <div>
        <ServersContainer>
          <Alert
            style={{ fontSize: 14 }}
            show={this.state.showAlert}
            variant="danger"
          >
            {this.state.alertMsg}
          </Alert>
          <Table size="sm">
            <tbody>
              {this.state.servers.map(server => (
                <tr key={server.ip}>
                  <td style={{ verticalAlign: 'middle' }}>
                    <Image
                      src={process.env.API_URL + server.image_path}
                      roundedCircle
                    />
                  </td>
                  <td style={{ verticalAlign: 'middle' }}>{server.country}</td>
                  <td style={{ verticalAlign: 'middle' }}>
                    {this.state.connectedTo === server.id ? (
                      <Button onClick={() => this.disconnect()} variant="link">
                        Disconnect
                      </Button>
                    ) : (
                      <Button
                        onClick={() => this.connect(server.id)}
                        variant="link"
                      >
                        Connect
                      </Button>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </Table>
        </ServersContainer>
      </div>
    );
  }
}

export default Main;
