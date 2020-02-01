import AbortController from 'node-abort-controller';
import storage from '../../utils/storage';
import API from './apiService';

const controller = new AbortController();

export default {
  async logout() {
    controller.abort();

    await API.Logout();

    await storage.set({ csrfToken: '' });
  },

  async isLoggedIn() {
    const csrfToken = await storage.get('csrfToken');
    return !!csrfToken;
  }
};
