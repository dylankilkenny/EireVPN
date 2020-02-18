import ReactGA from 'react-ga';

declare global {
  interface Window {
    GA_INITIALIZED: boolean;
  }
}

export const initGA = () => {
  ReactGA.initialize(process.env.GA_KEY ? process.env.GA_KEY : '');
};

export const logPageView = () => {
  ReactGA.set({ page: window.location.pathname });
  ReactGA.pageview(window.location.pathname);
};

export const logEvent = (category = '', action = '') => {
  if (category && action) {
    ReactGA.event({ category, action });
  }
};

export const logException = (description = '', fatal = false) => {
  if (description) {
    ReactGA.exception({ description, fatal });
  }
};
