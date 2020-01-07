import ext from '../../utils/ext';

export default function sendMessage(message, data) {
  ext.runtime.sendMessage({ action: message, data }, () => {});
}
