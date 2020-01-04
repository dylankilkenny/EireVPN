import ext from '../../utils/ext';

export default function sendMessage(message, data) {
  console.log('communicationManager:', message);
  ext.runtime.sendMessage({ action: message, data }, response => {
    console.log(response);
  });
}
