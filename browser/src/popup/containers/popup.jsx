import React from 'react';
import Login from './login';
import Main from './main';
import Settings from './settings';
import Connect from './connect';
import Incognito from './incognito';
import Header from '../components/header';
import PopupContainer from '../components/container';
import AuthService from '../services/authService';
import ext from '../../utils/ext';
import sendMessage from '../services/comunicationManager';
import storage from '../../utils/storage';

const views = {
  login: 0,
  main: 1,
  connect: 2,
  settings: 3,
  incognito: 4
};

class Popup extends React.Component {
  state = {
    renderView: views.login
  };

  async componentDidMount() {
    const isFirefox = typeof InstallTrigger !== 'undefined';
    console.log(isFirefox);
    if (isFirefox) {
      try {
        const allowed = await ext.extension.isAllowedIncognitoAccess();
        if (!allowed) {
          this.renderIcognito();
        } else {
          this.checkUserStatus();
        }
      } catch (error) {
        console.log(error);
      }
    } else {
      this.checkUserStatus();
    }
  }

  checkUserStatus = async () => {
    let renderView = views.login;
    const loggedIn = await AuthService.isLoggedIn();
    if (loggedIn) {
      renderView = views.main;
    }
    const connected = await storage.get('connected');
    let server;
    if (loggedIn && !!connected) {
      server = await storage.get('server');
      renderView = views.connect;
    }

    this.setState({ renderView, server });
  };

  logout = () => {
    AuthService.logout();
    sendMessage('disconnect', {});
    this.setState({ renderView: views.login });
  };

  renderMain = () => {
    this.setState({ renderView: views.main });
  };

  renderSettings = () => {
    this.setState({ renderView: views.settings });
  };

  renderConnect = server => {
    this.setState({ renderView: views.connect, server });
  };

  renderLogin = () => {
    this.setState({ renderView: views.login });
  };

  renderIcognito = () => {
    this.setState({ renderView: views.incognito });
  };

  render() {
    switch (this.state.renderView) {
      case views.login:
        return (
          <PopupContainer>
            <Header view={this.state.renderView} />
            <Login renderMain={this.renderMain} />
          </PopupContainer>
        );
      case views.main:
        return (
          <PopupContainer>
            <Header view={this.state.renderView} renderSettings={this.renderSettings} />
            <div>
              <Main renderLogin={this.renderLogin} renderConnect={this.renderConnect} />
            </div>
          </PopupContainer>
        );
      case views.connect:
        return (
          <PopupContainer>
            <Header view={this.state.renderView} renderSettings={this.renderSettings} />
            <Connect
              renderMain={this.renderMain}
              renderLogin={this.renderLogin}
              server={this.state.server}
            />
          </PopupContainer>
        );
      case views.settings:
        return (
          <PopupContainer>
            <Header view={this.state.renderView} renderMain={this.renderMain} />
            <Settings logout={this.logout} />
          </PopupContainer>
        );
      case views.incognito:
        return (
          <PopupContainer>
            <Header view={this.state.renderView} renderMain={this.renderMain} />
            <Incognito />
          </PopupContainer>
        );
      default:
        return (
          <div>
            something went wrong{' '}
            <span role="img" aria-label="cry">
              😭
            </span>
          </div>
        );
    }
  }
}
export default Popup;
