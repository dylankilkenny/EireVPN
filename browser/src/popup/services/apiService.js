import ServicesUtil from './servicesUtil';

export default class ApiService {
  static getServers() {
    // Get a token from api server using the fetch api
    return ServicesUtil.fetch(`${process.env.API_URL}api/private/servers`, {
      method: 'GET'
    });
  }
  static connectServer(serverId) {
    // Get a token from api server using the fetch api
    return ServicesUtil.fetch(
      `${process.env.API_URL}api/private/servers/connect/${serverId}`,
      {
        method: 'GET'
      }
    );
  }
}
