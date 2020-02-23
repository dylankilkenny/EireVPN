import Router from 'next/router';
import API from './APIService';
import AbortController from 'node-abort-controller';

const controller = new AbortController();
const ADMIN_LOGIN = '/admin/login';
const USER_LOGIN = '/login';

const redirectLoginURL = (): string => {
  if (Router.pathname.includes('admin')) return ADMIN_LOGIN;
  else return USER_LOGIN;
};

const ClearAndRedirect = () => {
  controller.abort();
  localStorage.clear();
  Router.push(redirectLoginURL());
};

const IsLoggedIn = () => {
  let csrfToken = localStorage.getItem('X-CSRF-Token');
  if (csrfToken !== null) {
    return true;
  }
  return false;
};

const Logout = () => {
  const res = API.Logout();
  ClearAndRedirect();
};

export default {
  ClearAndRedirect,
  Logout,
  IsLoggedIn
};
