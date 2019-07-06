const path = require('path');

module.exports = {
  mode: 'development',
  entry: {
    app: ['./src/index.js', './src/test.js']
  },
  devtool: 'inline-source-map',
  devServer: {
     contentBase: './dist'
  },
  output: {
    filename: '[name].bundle.js',
    jsonpScriptType: 'module',
    path: path.resolve(__dirname, 'dist')
  }
};