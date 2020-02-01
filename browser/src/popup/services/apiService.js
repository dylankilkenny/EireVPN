import 'regenerator-runtime/runtime';
import storage from '../../utils/storage';

const FORBIDDEN = 403;

const getCsrfToken = async () => {
  const csrfToken = await storage.get('csrfToken');
  return csrfToken;
};

const fetcher = async args => {
  const headers = new Headers();

  headers.set('Content-Type', 'application/json');
  headers.set('Accept', 'application/json');

  let csrfToken = await getCsrfToken();
  if (csrfToken !== undefined) {
    headers.set('X-CSRF-Token', csrfToken);
  }

  const res = await fetch(args.url, {
    headers,
    credentials: 'include',
    method: args.method,
    body: args.body
  });

  csrfToken = res.headers.get('X-CSRF-Token');
  if (csrfToken !== null) {
    await storage.set({ csrfToken });
  }

  if (res.status === FORBIDDEN) {
    return {
      status: 403
    };
  }

  const j = await res.json();
  return j;
};

const getRequest = async url => {
  const args = {
    method: 'GET',
    url,
    body: undefined
  };
  const res = await fetcher(args);
  return res;
};

const postRequest = async (url, body) => {
  const args = {
    method: 'POST',
    url,
    body
  };
  const res = await fetcher(args);
  return res;
};

export default {
  async getServers() {
    return getRequest(`${process.env.API_URL}/api/private/servers`);
  },
  async Logout() {
    return getRequest(`${process.env.API_URL}/api/user/logout`);
  },
  async connectServer(serverId) {
    return getRequest(`${process.env.API_URL}/api/private/servers/connect/${serverId}`);
  },
  async login(email, password) {
    return postRequest(
      `${process.env.API_URL}/api/user/login`,
      JSON.stringify({
        email,
        password
      })
    );
  }
};
