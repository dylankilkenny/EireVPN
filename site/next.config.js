const env = {
  dev: {
    apiDomain: 'http://localhost:3001',
    GA_KEY: 'UA-158748602-1'
  },
  qa: {
    apiDomain: 'http://api.qa.eirevpn.ie',
    GA_KEY: 'UA-158748602-1'
  },
  prod: {
    apiDomain: 'https://api.eirevpn.ie',
    GA_KEY: 'UA-158748602-1'
  }
};
module.exports = env[process.env.NODE_ENV || 'dev'];
