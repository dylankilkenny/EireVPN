import fetch from 'unfetch';
import Auth from './Auth';

const FORBIDDEN = 403;

interface FetcherArgs {
  method: string;
  url: string;
  body: string | FormData | undefined;
}

const fetcher = async (args: FetcherArgs) => {
  const headers: HeadersInit = new Headers();

  // if body is FormData we dont pass headers, browser will do that
  if (typeof args.body === 'string') {
    headers.set('Content-Type', 'application/json');
    headers.set('Accept', 'application/json');
  }

  let csrfToken = localStorage.getItem('X-CSRF-Token');
  if (csrfToken !== null) {
    headers.set('X-CSRF-Token', csrfToken);
  }

  const res = await fetch(args.url, {
    headers: headers,
    credentials: 'include',
    method: args.method,
    body: args.body
  });

  csrfToken = res.headers.get('X-CSRF-Token');
  if (csrfToken !== null) {
    localStorage.setItem('X-CSRF-Token', csrfToken);
  }

  if (res.status === FORBIDDEN) {
    Auth.ClearAndRedirect();
  }

  const j = await res.json();
  return j;
};

const getRequest = async (url: string) => {
  const args = {
    method: 'GET',
    url: url,
    body: undefined
  };
  const res = await fetcher(args);
  return res;
};

const postRequest = async (url: string, body: string | FormData) => {
  const args = {
    method: 'POST',
    url: url,
    body: body
  };
  const res = await fetcher(args);
  return res;
};

const putRequest = async (url: string, body: string) => {
  const args = {
    method: 'PUT',
    url: url,
    body: body
  };
  const res = await fetcher(args);
  return res;
};

const deleteRequest = async (url: string) => {
  const args = {
    method: 'DELETE',
    url: url,
    body: undefined
  };
  const res = await fetcher(args);
  return res;
};

export default {
  async Login(body: string) {
    return postRequest(`${process.env.apiDomain}/api/user/login`, body);
  },

  async Signup(body: string) {
    return postRequest(`${process.env.apiDomain}/api/user/signup`, body);
  },

  async Logout() {
    return getRequest(`${process.env.apiDomain}/api/user/logout`);
  },

  async ConfirmEmail(token: string) {
    return getRequest(`${process.env.apiDomain}/api/user/confirm_email/${token}`);
  },

  async ResendConfirmEmailLink() {
    return getRequest(`${process.env.apiDomain}/api/private/user/confirm_email_resend`);
  },

  async ForgotPasswordEmail(body: string) {
    return postRequest(`${process.env.apiDomain}/api/user/forgot_pass`, body);
  },

  async UpdatePassword(body: string, token: string) {
    return postRequest(`${process.env.apiDomain}/api/user/forgot_pass/${token}`, body);
  },

  async ContactSupport(body: string) {
    return postRequest(`${process.env.apiDomain}/api/message`, body);
  },

  async GetSettings() {
    return getRequest(`${process.env.apiDomain}/api/protected/settings`);
  },

  async GetUsersList() {
    return getRequest(`${process.env.apiDomain}/api/protected/users`);
  },

  async GetPlansList() {
    return getRequest(`${process.env.apiDomain}/api/protected/plans`);
  },

  async GetServersList() {
    return getRequest(`${process.env.apiDomain}/api/private/servers`);
  },

  async GetUserPlansList() {
    return getRequest(`${process.env.apiDomain}/api/protected/userplans`);
  },

  async GetServerByID(id: string) {
    return getRequest(`${process.env.apiDomain}/api/protected/servers/${id}`);
  },

  async GetUserByID(id: string) {
    return getRequest(`${process.env.apiDomain}/api/private/user/get/${id}`);
  },

  async GetUserPlanByUserID(id: string) {
    return getRequest(`${process.env.apiDomain}/api/private/userplans/${id}`);
  },

  async GetPlanByID(id: string) {
    return getRequest(`${process.env.apiDomain}/api/protected/plans/${id}`);
  },

  async UpdateServer(id: string, body: string) {
    return putRequest(`${process.env.apiDomain}/api/protected/servers/update/${id}`, body);
  },

  async UpdateUser(id: string, body: string) {
    return putRequest(`${process.env.apiDomain}/api/private/user/update/${id}`, body);
  },

  async UpdateUserPlan(user_id: string, body: string) {
    return putRequest(`${process.env.apiDomain}/api/protected/userplans/update/${user_id}`, body);
  },

  async UpdatePlan(id: string, body: string) {
    return putRequest(`${process.env.apiDomain}/api/protected/plans/update/${id}`, body);
  },

  async UpdateSettings(body: string) {
    return putRequest(`${process.env.apiDomain}/api/protected/settings/update`, body);
  },

  async ChangePassword(body: string) {
    return putRequest(`${process.env.apiDomain}/api/private/user/changepassword`, body);
  },

  async CreateServer(body: FormData) {
    return postRequest(`${process.env.apiDomain}/api/protected/servers/create`, body);
  },

  async CreateUser(body: string) {
    return postRequest(`${process.env.apiDomain}/api/user/signup`, body);
  },

  async CreateUserPlan(body: string) {
    return postRequest(`${process.env.apiDomain}/api/protected/userplans/create`, body);
  },

  async CreatePlan(body: string) {
    return postRequest(`${process.env.apiDomain}/api/protected/plans/create`, body);
  },

  async DeleteServer(id: string) {
    return deleteRequest(`${process.env.apiDomain}/api/protected/servers/delete/${id}`);
  },

  async DeleteUser(id: string) {
    return deleteRequest(`${process.env.apiDomain}/api/protected/users/delete/${id}`);
  },

  async DeleteUserPlan(user_id: string) {
    return deleteRequest(`${process.env.apiDomain}/api/protected/userplans/delete/${user_id}`);
  },

  async DeletePlan(id: string) {
    return deleteRequest(`${process.env.apiDomain}/api/protected/plans/delete/${id}`);
  }
};
