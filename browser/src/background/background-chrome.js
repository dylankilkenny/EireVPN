import ext from '../utils/ext';

let proxyUsername;
let proxyPassword;

function setAuth(details, callbackFn) {
  if (proxyUsername !== undefined && proxyPassword !== undefined) {
    callbackFn({
      authCredentials: { username: proxyUsername, password: proxyPassword }
    });
  }
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
  console.log('connected');
}

function disconnectProxy() {
  proxyUsername = undefined;
  proxyPassword = undefined;
  ext.proxy.settings.clear({ scope: 'regular' }, () => {});
  console.log('disconnected');
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
