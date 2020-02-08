import ext from '../utils/ext';

let proxyUsername;
let proxyPassword;
const pendingRequests = [];

// A request has completed.
// We can stop worrying about it.
function completed(requestDetails) {
  const index = pendingRequests.indexOf(requestDetails.requestId);
  if (index > -1) {
    pendingRequests.splice(index, 1);
  }
}

function provideAuth(requestDetails) {
  // If we have seen this request before, then
  // assume our credentials were bad, and give up.
  if (pendingRequests.indexOf(requestDetails.requestId) !== -1) {
    console.log(`bad credentials: auth+${proxyUsername}+${proxyPassword}`);
    return { cancel: true };
  }
  pendingRequests.push(requestDetails.requestId);
  return {
    authCredentials: { username: proxyUsername, password: proxyPassword }
  };
}

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

function connectProxy(proxy) {
  const proxySettings = {
    proxyType: 'manual',
    httpProxyAll: true,
    http: `http://${proxy.username}:${proxy.password}@${proxy.ip}:${proxy.port}`
  };
  proxyUsername = proxy.username;
  proxyPassword = proxy.password;
  ext.proxy.settings.set({ value: proxySettings });
}

function disconnectProxy() {
  function onCleared(result) {
    if (result) {
      console.log('proxy disconnected');
    } else {
      console.log('proxy disconnect failed');
    }
  }
  proxyUsername = undefined;
  proxyPassword = undefined;
  const clearing = ext.proxy.settings.clear({});
  clearing.then(onCleared);
}

function handleMessage(request, sender, sendResponse) {
  if (request.action === 'connect') {
    connectProxy(request.data);
  } else if (request.action === 'disconnect') {
    disconnectProxy();
  }
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
