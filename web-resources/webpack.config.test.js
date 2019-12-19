const autoprefixer = require('autoprefixer');

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
    entry:  {
      app: [
        './test/test.js',
      ]
    },
    output: {
      filename: "./test-compiled.js",
      path: __dirname + '/test'
    },
    module: {
      rules: [
        {
          test: /\.tsx?$/,
          use: 'ts-loader',
          exclude: /node_modules/,
        },
      ],
    },
    resolve: {
      extensions: [ '.tsx', '.ts', '.js' ],
    },
  },
];
