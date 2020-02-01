/* global browser, window, chrome */
import ext from './ext';

export default {
  async set(obj) {
    await ext.storage.local.set(obj, () => {});
  },
  async get(key) {
    try {
      if (browser['storage']) {
        const resp = await ext.storage.local.get(key);
        return resp[key];
      }
    } catch (error) {
      console.log(error);
    }
    try {
      if (chrome['storage']) {
        const val = await new Promise(resolve => {
          chrome.storage.local.get([key], resp => {
            resolve(resp[key]);
          });
        });
        return val;
      }
    } catch (error) {
      console.log(error);
    }
    return '';
  }
};
