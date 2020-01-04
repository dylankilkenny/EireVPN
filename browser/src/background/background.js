import ext from '../utils/ext';

function rewriteUserAgentHeader(e) {
  console.log('e: ', e);
  if (e.documentUrl.includes('extension')) {
    for (const header of e.requestHeaders) {
      if (header.name === 'Origin') {
        console.log('back', process.env.API_URL);
        header.value = process.env.API_URL;
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
  const clearing = ext.proxy.settings.clear({});
  clearing.then(onCleared);
}
ext.runtime.onMessage.addListener(request => {
  console.log(request);
  if (request.action === 'connect') {
    connectProxy(request.data);
  } else if (request.action === 'disconnect') {
    disconnectProxy();
  }
});

ext.webRequest.onBeforeSendHeaders.addListener(
  rewriteUserAgentHeader,
  { urls: ['<all_urls>'] },
  ['blocking', 'requestHeaders']
);
