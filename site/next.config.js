const config = {
  dev: {
    env: {
      apiDomain: 'http://localhost:3001',
      GA_KEY: 'UA-158748602-1'
    }
  },
  qa: {
    env: {
      apiDomain: 'http://api.qa.eirevpn.ie',
      GA_KEY: 'UA-158748602-1'
    }
  },
  prod: {
    env: {
      apiDomain: 'https://api.eirevpn.ie',
      GA_KEY: 'UA-158748602-1'
    }
  }
};
module.exports = config[process.env.NODE_ENV || 'dev'];
