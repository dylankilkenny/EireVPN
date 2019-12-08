import ext from '../utils/ext';

let proxy = {};

function rewriteUserAgentHeader(e) {
  if (e.documentUrl.includes('extension')) {
    for (const header of e.requestHeaders) {
      if (header.name === 'Origin') {
        header.value = 'http://localhost:8080';
      }
    }
  }
  return { requestHeaders: e.requestHeaders };
}

ext.runtime.onMessage.addListener(request => {
  if (request.action === 'connect') {
    proxy.ip = request.data.ip;
    proxy.port = request.data.port;
    proxy.active = true;
  } else if (request.action === 'disconnect') {
    proxy = {};
  }
});

function isLocalHost(url) {
  let LocalHost = false;
  if (url.includes('0.0.0.0')) {
    LocalHost = true;
  } else if (url.includes('localhost')) {
    LocalHost = true;
  }
  return LocalHost;
}

function handleProxyRequest(requestInfo) {
  if (proxy.active && !isLocalHost(requestInfo.url)) {
    console.log(`Proxying: ${requestInfo.url} with ${proxy.ip}:${proxy.port} `);
    return { type: 'http', host: proxy.ip, port: proxy.port };
  }
  return { type: 'direct' };
}

ext.proxy.onRequest.addListener(handleProxyRequest, {
  urls: ['<all_urls>']
});

ext.webRequest.onBeforeSendHeaders.addListener(
  rewriteUserAgentHeader,
  { urls: ['<all_urls>'] },
  ['blocking', 'requestHeaders']
);
