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
    app: [
      './src/dictionary.js',
      './src/events.js',
      './src/idictionarybuilder.js',
      './src/term.js',
      './src/test.js',
      './src/testbuilder.js',
      './src/wordfinder.js'
      ]
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