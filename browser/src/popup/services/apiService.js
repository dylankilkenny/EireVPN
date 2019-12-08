import ServicesUtil from './servicesUtil';

export default class ApiService {
  static getServers() {
    // Get a token from api server using the fetch api
    return ServicesUtil.fetch('http://0.0.0.0:3001/api/private/servers', {
      method: 'GET',
    });
  }
}
