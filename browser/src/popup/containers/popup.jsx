import React from 'react';
import Login from './login';
import Header from '../components/header';
import LargeAlert from '../components/largeAlert';
import PopupContainer from '../components/container';
import Settings from '../containers/settings';
import Main from '../containers/main';
import AuthService from '../services/authService';
import ext from '../../utils/ext';
import sendMessage from '../services/comunicationManager';

const views = {
  login: 0,
  main: 1,
  settings: 2,
  incognito: 3
};

class Popup extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
    this.authorise = this.authorise.bind(this);
    this.logout = this.logout.bind(this);
    this.renderSettings = this.renderSettings.bind(this);
    this.renderMain = this.renderMain.bind(this);
    this.checkAuth = this.checkAuth.bind(this);
    this.connected = this.connected.bind(this);
  }

  componentWillMount() {}

  componentDidMount() {
    const isFirefox = typeof InstallTrigger !== 'undefined';
    if (isFirefox) {
      ext.extension.isAllowedIncognitoAccess().then(allowed => {
        if (allowed) {
          this.checkAuth();
        } else {
          this.setState({ renderView: views.incognito });
        }
      });
    } else {
      this.checkAuth();
    }
  }

  checkAuth() {
    AuthService.isLoggedIn(this.state.connected).then(isLoggedIn => {
      let renderView;
      if (!isLoggedIn) {
        renderView = views.login;
      } else {
        renderView = views.main;
      }
      this.setState({ renderView });
    });
  }

  authorise(isLoggedIn) {
    if (isLoggedIn) {
      const renderView = views.main;
      this.setState({ renderView });
    }
  }

  logout() {
    AuthService.logout();
    sendMessage('disconnect', {});
    this.setState({ renderView: 0, connected: false });
  }

  connected() {
    this.setState({ connected: true });
  }

  renderMain() {
    this.setState({ renderView: 1 });
  }

  renderSettings() {
    this.setState({ renderView: 2 });
  }

  render() {
    switch (this.state.renderView) {
      case 0:
        return (
          <PopupContainer>
            <Header view={this.state.renderView} />
            <Login authorise={this.authorise} />
          </PopupContainer>
        );
      case 1:
        return (
          <PopupContainer>
            <Header
              view={this.state.renderView}
              renderSettings={this.renderSettings}
            />
            <div>
              <Main connected={this.connected} />
            </div>
          </PopupContainer>
        );
      case 2:
        return (
          <PopupContainer>
            <Header view={this.state.renderView} renderMain={this.renderMain} />
            <Settings renderMain={this.renderMain} logout={this.logout} />
          </PopupContainer>
        );
      case 3:
        return (
          <PopupContainer>
            <Header view={this.state.renderView} renderMain={this.renderMain} />
            <LargeAlert
              variant="warning"
              heading="Private Browsing Disabled"
              body="To use this extension private browsing is required. To enable private browsing right click on this extension and choose manage extension. You will see an option 'Run in Private Windows', click allow."
            />
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
    // if (!this.state.loggedIn) {
    // }
    // return (
    //   <PopupContainer>
    //     <Header settings={this.renderSettings} />
    //     <Button onClick={this.logout} variant="primary">
    //       logout
    //     </Button>
    //   </PopupContainer>
    // );
  }
}
export default Popup;
