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
    entry: "./cnotes.scss",
    output: {
      filename: "style-cnotes.js",
    },
    module: {
      rules: [{
        test: /\.scss$/,
        use: [
          {
            loader: 'file-loader',
            options: {
              name: 'cnotes.css',
            },
          },
          { loader: 'extract-loader' },
          { loader: 'css-loader' },
          {
            loader: 'sass-loader',
            options: {
              includePaths: ['./node_modules']
            },
          },
        ]
      }],
    },
  },
{
  mode: 'development',
  entry: {
    app: [
      './src/dictionaryentry.js',
      './src/events.js',
      './src/resultparser.js',
      './src/wordsense.js'
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