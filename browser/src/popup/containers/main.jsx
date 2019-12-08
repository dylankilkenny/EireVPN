import React from 'react';
import Image from 'react-bootstrap/Image';
import Button from 'react-bootstrap/Button';
import Table from 'react-bootstrap/Table';
import styled from 'styled-components';
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
      connectedTo: ''
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

  connect(ip, port) {
    sendMessage('connect', { ip, port });
    ext.storage.local.set({ connectedTo: ip }, () => {
      this.setState({ connectedTo: ip });
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
          <Table size="sm">
            <tbody>
              {this.state.servers.map(server => (
                <tr>
                  <td style={{ verticalAlign: 'middle' }}>
                    <Image
                      src={`http://localhost:3001/${server.image_path}`}
                      roundedCircle
                    />
                  </td>
                  <td style={{ verticalAlign: 'middle' }}>{server.country}</td>
                  <td style={{ verticalAlign: 'middle' }}>
                    {this.state.connectedTo === server.ip ? (
                      <Button onClick={() => this.disconnect()} variant="link">
                        Disconnect
                      </Button>
                    ) : (
                      <Button
                        onClick={() => this.connect(server.ip, server.port)}
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
