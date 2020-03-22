import ext from '../utils/ext';
import storage from '../utils/storage';
import util from './util';

let proxyUsername;
let proxyPassword;

function disconnectProxy() {
  proxyUsername = undefined;
  proxyPassword = undefined;
  ext.proxy.settings.clear({ scope: 'regular' }, () => {});
  console.log('disconnected');
  ext.browserAction.setBadgeText({ text: '' });
  ext.browserAction.setBadgeBackgroundColor({ color: '#0000' });
  storage.set({ connected: false, server: undefined, ip: '' });
}

function connectProxy(proxy) {
  const config = {
    mode: 'fixed_servers',
    rules: {
      singleProxy: {
        host: proxy.ip,
        port: proxy.port
      }
    }
  };
  proxyUsername = proxy.username;
  proxyPassword = proxy.password;
  ext.proxy.settings.set({ value: config, scope: 'regular' }, () => {});
  ext.browserAction.setBadgeText({ text: 'on' });
  ext.browserAction.setBadgeBackgroundColor({ color: 'green' });
  util
    .timeout(1500, fetch('https://api.eirevpn.ie/api/plans'))
    .then(() => {})
    .catch(error => {
      console.log(error);
      storage.set({ proxy_error: true });
      disconnectProxy();
    });
  console.log('connected');
}

function setAuth(details, callbackFn) {
  if (proxyUsername !== undefined && proxyPassword !== undefined) {
    callbackFn({
      authCredentials: { username: proxyUsername, password: proxyPassword }
    });
  } else {
    console.log('proxyUsername or proxyPassword not defined');
    storage.set({ proxy_error: true });
    disconnectProxy();
  }
}

function handleMessage(request, sender, sendResponse) {
  if (request.action === 'connect') {
    connectProxy(request.data);
  } else if (request.action === 'disconnect') {
    storage.set({ proxy_error: true });
    disconnectProxy();
  }
  sendResponse(true);
}

ext.runtime.onMessage.addListener(handleMessage);
ext.webRequest.onAuthRequired.addListener(setAuth, { urls: ['<all_urls>'] }, ['asyncBlocking']);
