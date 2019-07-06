const path = require('path');

function getStyleUse(bundleFilename) {
  return [
    {
      loader: 'file-loader',
      options: {
        name: bundleFilename,
      },
    },
    { loader: 'extract-loader' },
    { loader: 'css-loader' },
    {
      loader: 'sass-loader',
      options: {
        includePaths: ['./node_modules'],
        implementation: require('dart-sass'),
        fiber: require('fibers'),
      }
    },
  ];
}

module.exports = [
{
  entry: "./src/wordfinder.scss",
  output: {
    // This is necessary for webpack to compile, but we never reference this js file.
    filename: 'styles-wordfinder.js',
    path: path.resolve(__dirname, 'src')
  },
  module: {
    rules: [{
      test: /wordfinder.scss/,
      use: getStyleUse('wordfinder.css')
    }]
  }
},
{
  mode: 'development',
  entry: {
    app: ['./src/wordfinder.js', './src/test.js']
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
}];