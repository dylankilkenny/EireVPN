import decode from 'jwt-decode';
import ext from '../../utils/ext';
import ServicesUtil from './servicesUtil';

export default class AuthService {
  // Initializing important variables
  // constructor(domain) {
  // this.domain = domain || 'http://localhost:3001'; // API server domain
  // this.fetch = this.fetch.bind(this); // React binding stuff
  // this.login = this.login.bind(this);
  // this.getProfile = this.getProfile.bind(this);
  // }

  static login(email, password) {
    // Get a token from api server using the fetch api
    return ServicesUtil.fetch(`${process.env.API_URL}api/user/login`, {
      method: 'POST',
      body: JSON.stringify({
        email,
        password
      })
    });
  }

  static logout() {
    // Clear user token and profile data from localStorage
    ext.storage.local.set({ csrfToken: '' }, () => {});
  }

  static isLoggedIn(isConnected) {
    return ServicesUtil.getCsrfToken().then(token => {
      if (isConnected) {
        return true;
      }
      return !!token && !AuthService.isTokenExpired(token);
    });
  }

  static isTokenExpired(token) {
    try {
      const decoded = decode(token);
      if (decoded.exp < Date.now() / 1000) {
        // Checking if token is expired. N
        return true;
      }
      return false;
    } catch (err) {
      return false;
    }
  }
}
