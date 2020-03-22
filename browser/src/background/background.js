import ext from '../utils/ext';
import storage from '../utils/storage';
import util from './util';

let proxyUsername;
let proxyPassword;
const pendingRequests = [];

function rewriteUserAgentHeader(e) {
  if (e.documentUrl.includes('extension')) {
    for (const header of e.requestHeaders) {
      if (header.name === 'Origin') {
        header.value = 'moz-extension://extensions@eirevpn.ie';
      }
    }
  }
  return { requestHeaders: e.requestHeaders };
}
function disconnectProxy() {
  function onCleared(result) {
    if (result) {
      console.log('proxy disconnected');
    } else {
      console.log('proxy disconnect failed');
    }
  }
  try {
    proxyUsername = undefined;
    proxyPassword = undefined;
    const clearing = ext.proxy.settings.clear({});
    clearing.then(onCleared);
    // having issues with this not working in some versions of firefox (dev)
    // to be double sure we'll set proxy back to system
    const proxySettings = {
      proxyType: 'system'
    };
    ext.proxy.settings.set({ value: proxySettings });
    ext.browserAction.setBadgeText({ text: '' });
    storage.set({ connected: false, server: undefined, ip: '' });
  } catch (error) {
    console.log('disconnectProxy: ', error);
  }
}

function connectProxy(proxy) {
  const proxySettings = {
    proxyType: 'manual',
    httpProxyAll: true,
    http: `http://${proxy.ip}:${proxy.port}`
  };
  proxyUsername = proxy.username;
  proxyPassword = proxy.password;
  ext.proxy.settings.set({ value: proxySettings });
  ext.browserAction.setBadgeText({ text: 'on' });
  ext.browserAction.setBadgeBackgroundColor({ color: 'green' });
  // create initial auth request. Issues were arising where
  // the onAuthRequired function wasnt working unless the extension was reopened
  // and a request made to our servers. Hacky workaround is to just make a request
  // to our servers when connected.
  util
    .timeout(1500, fetch('https://api.eirevpn.ie/api/plans'))
    .then(r => {
      console.log(r);
    })
    .catch(error => {
      console.log(error);
      storage.set({ proxy_error: true });
      disconnectProxy();
    });
}

function handleMessage(request) {
  if (request.action === 'connect') {
    connectProxy(request.data);
  } else if (request.action === 'disconnect') {
    disconnectProxy();
  }
}

// A request has completed.
// We can stop worrying about it.
function completed(requestDetails) {
  const index = pendingRequests.indexOf(requestDetails.requestId);
  if (index > -1) {
    pendingRequests.splice(index, 1);
  }
}

function provideAuth(requestDetails) {
  console.log('provideAuth');
  // If we have seen this request before, then
  // assume our credentials were bad, and give up.
  if (proxyUsername === undefined || proxyPassword === undefined) {
    console.log('proxyUsername or proxyPassword not defined');
    storage.set({ proxy_error: true });
    disconnectProxy();
  }
  if (pendingRequests.indexOf(requestDetails.requestId) !== -1) {
    console.log(`bad credentials: auth+${proxyUsername}+${proxyPassword}`);
    storage.set({ proxy_error: true });
    disconnectProxy();
    return { cancel: true };
  }

  pendingRequests.push(requestDetails.requestId);
  return {
    authCredentials: { username: proxyUsername, password: proxyPassword }
  };
}

ext.runtime.onMessage.addListener(handleMessage);

ext.webRequest.onBeforeSendHeaders.addListener(rewriteUserAgentHeader, { urls: ['<all_urls>'] }, [
  'blocking',
  'requestHeaders'
]);

ext.webRequest.onAuthRequired.addListener(provideAuth, { urls: ['<all_urls>'] }, ['blocking']);

ext.webRequest.onCompleted.addListener(completed, { urls: ['<all_urls>'] });

ext.webRequest.onErrorOccurred.addListener(completed, {
  urls: ['<all_urls>']
});
