var ENV = process.argv[2];
const apiDomain = ENV == 'dev' ? 'http://localhost:3001' : 'https://api.eirevpn.ie';
module.exports = {
  env: {
    apiDomain,
    GA_KEY: 'UA-158748602-1'
  }
};
