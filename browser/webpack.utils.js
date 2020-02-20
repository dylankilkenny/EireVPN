const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');
const ZipPlugin = require('zip-webpack-plugin');
const path = require('path');
const webpack = require('webpack');

const getHTMLPlugins = (browserDir, outputDir = 'dev', sourceDir = 'src') => [
  new HtmlWebpackPlugin({
    title: 'Popup',
    filename: path.resolve(__dirname, `${outputDir}/${browserDir}/popup/index.html`),
    template: `${sourceDir}/popup/index.html`,
    chunks: ['popup']
  })
];

const getOutput = (browserDir, outputDir = 'dev') => ({
  path: path.resolve(__dirname, `${outputDir}/${browserDir}`),
  filename: '[name]/[name].js'
});

const getEntry = (sourceDir = 'src') => ({
  popup: path.resolve(__dirname, `${sourceDir}/popup/popup.jsx`),
  background: path.resolve(__dirname, `${sourceDir}/background/background.js`),
  hotreload: path.resolve(__dirname, `${sourceDir}/utils/hot-reload.js`)
});

const getChromeEntry = (sourceDir = 'src') => ({
  popup: path.resolve(__dirname, `${sourceDir}/popup/popup.jsx`),
  background: path.resolve(__dirname, `${sourceDir}/background/background-chrome.js`),
  hotreload: path.resolve(__dirname, `${sourceDir}/utils/hot-reload.js`)
});

const getCopyPlugins = (browserDir, outputDir = 'dev', sourceDir = 'src') => [
  new CopyWebpackPlugin([
    {
      from: `${sourceDir}/assets`,
      to: path.resolve(__dirname, `${outputDir}/${browserDir}/assets`)
    },
    {
      from: `${sourceDir}/_locales`,
      to: path.resolve(__dirname, `${outputDir}/${browserDir}/_locales`)
    },
    {
      from: `${sourceDir}/manifest.json`,
      to: path.resolve(__dirname, `${outputDir}/${browserDir}/manifest.json`)
    }
  ])
];

const getDefinePlugin = (domain, api) => [
  new webpack.DefinePlugin({
    'process.env.API_URL': api,
    'process.env.DOMAIN': domain
  })
];

const getFirefoxCopyPlugins = (browserDir, outputDir = 'dev', sourceDir = 'src') => [
  new CopyWebpackPlugin([
    {
      from: `${sourceDir}/assets`,
      to: path.resolve(__dirname, `${outputDir}/${browserDir}/assets`)
    },
    {
      from: `${sourceDir}/_locales`,
      to: path.resolve(__dirname, `${outputDir}/${browserDir}/_locales`)
    },
    {
      from: `${sourceDir}/manifest-ff.json`,
      to: path.resolve(__dirname, `${outputDir}/${browserDir}/manifest.json`)
    }
  ])
];

const getZipPlugin = (browserDir, outputDir = 'dist') =>
  new ZipPlugin({
    path: path.resolve(__dirname, `${outputDir}/${browserDir}`),
    filename: browserDir,
    extension: 'zip',
    fileOptions: {
      mtime: new Date(),
      mode: 0o100664,
      compress: true,
      forceZip64Format: false
    },
    zipOptions: {
      forceZip64Format: false
    }
  });

module.exports = {
  getHTMLPlugins,
  getOutput,
  getCopyPlugins,
  getDefinePlugin,
  getFirefoxCopyPlugins,
  getZipPlugin,
  getEntry,
  getChromeEntry
};
