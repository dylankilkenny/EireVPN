import ext from '../utils/ext';

let proxyUsername;
let proxyPassword;

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
  console.log('connected');
  ext.browserAction.setBadgeText({ text: 'on' });
  ext.browserAction.setBadgeBackgroundColor({ color: 'green' });
  // create initial auth request. Issues were arising where
  // the onAuthRequired function wasnt working unless the extension was reopened
  // and a request made to our servers. Hacky workaround is to just make a request
  // to our servers when connected.
  fetch('https://api.eirevpn.ie');
}

function disconnectProxy() {
  proxyUsername = undefined;
  proxyPassword = undefined;
  ext.proxy.settings.clear({ scope: 'regular' }, () => {});
  console.log('disconnected');
  ext.browserAction.setBadgeText({ text: '' });
  ext.browserAction.setBadgeBackgroundColor({ color: '' });
}

function setAuth(details, callbackFn) {
  if (proxyUsername !== undefined && proxyPassword !== undefined) {
    callbackFn({
      authCredentials: { username: proxyUsername, password: proxyPassword }
    });
  } else {
    console.log('proxyUsername or proxyPassword not defined');
    disconnectProxy();
  }
}

function handleMessage(request, sender, sendResponse) {
  if (request.action === 'connect') {
    connectProxy(request.data);
  } else if (request.action === 'disconnect') {
    disconnectProxy();
  }
  sendResponse(true);
}

ext.runtime.onMessage.addListener(handleMessage);
ext.webRequest.onAuthRequired.addListener(setAuth, { urls: ['<all_urls>'] }, ['asyncBlocking']);
