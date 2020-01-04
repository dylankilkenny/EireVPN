import decode from 'jwt-decode';
import ext from '../../utils/ext';
import AuthService from './authService';

export default class ServicesUtil {
  static fetch(url, options) {
    const headers = {
      Accept: 'application/json',
      'Content-Type': 'application/json'
    };
    return AuthService.isLoggedIn()
      .then(loggedIn => {
        if (loggedIn) {
          return ServicesUtil.getCsrfToken();
        }
        return undefined;
      })
      .then(csrf => {
        if (csrf !== undefined) {
          headers['X-CSRF-Token'] = csrf;
        }

        return fetch(url, {
          headers,
          ...options
        })
          .then(resp => {
            if (resp.status === 200) {
              ServicesUtil.setCsrfToken(resp.headers.get('X-CSRF-Token'));
            }
            return resp.json();
          })
          .catch(error => {
            console.log(error);
            return error;
          });
      });
  }

  static setCsrfToken(token) {
    // Saves user token to localStorage
    ext.storage.local.set({ csrfToken: token }, () => {});
  }

  static getCsrfToken() {
    // Retrieves the user token from localStorage
    return new Promise(resolve => {
      ext.storage.local.get('csrfToken', resp => {
        const { csrfToken } = resp;
        if (csrfToken) {
          resolve(csrfToken);
        }
        resolve(false);
      });
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
