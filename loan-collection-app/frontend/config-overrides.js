const webpack = require('webpack');

module.exports = function override(config) {
  // Remove DefinePlugin from existing plugins to prevent conflicts
  config.plugins = config.plugins.filter(
    plugin => !(plugin instanceof webpack.DefinePlugin)
  );

  const fallback = config.resolve.fallback || {};
  Object.assign(fallback, {
    "crypto": require.resolve('crypto-browserify'),
    "stream": require.resolve('stream-browserify'),
    "assert": require.resolve('assert'),
    "http": require.resolve('stream-http'),
    "https": require.resolve('https-browserify'),
    "os": require.resolve('os-browserify/browser'),
    "path": require.resolve('path-browserify'),
    "process": require.resolve('process/browser.js'),
    "vm": require.resolve('vm-browserify')
  });
  
  config.resolve.fallback = fallback;
  config.plugins = (config.plugins || []).concat([
    new webpack.ProvidePlugin({
      process: 'process/browser.js',
      Buffer: ['buffer', 'Buffer']
    })
  ]);

  return config;
};