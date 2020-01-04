import React from 'react';
import Login from './login';
import Header from '../components/header';
import PopupContainer from '../components/container';
import Settings from '../containers/settings';
import Main from '../containers/main';
import AuthService from '../services/authService';

const views = {
  login: 0,
  main: 1,
  settings: 2
};

class Popup extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
    this.authorise = this.authorise.bind(this);
    this.logout = this.logout.bind(this);
    this.renderSettings = this.renderSettings.bind(this);
    this.renderMain = this.renderMain.bind(this);
  }

  componentWillMount() {}

  componentDidMount() {
    console.log(process.env.API_URL);

    AuthService.isLoggedIn().then(isLoggedIn => {
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
    this.setState({ renderView: 0 });
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
              <Main />
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
