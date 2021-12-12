const autoprefixer = require('autoprefixer');
const webpack = require('webpack');

module.exports = [
  {
    entry: './cnotes.scss',
    output: {
      filename: 'cnotes-not-used.js',
    },
    module: {
      rules: [
        {
          test: /\.scss$/,
          use: [
            {
              loader: 'file-loader',
              options: {
                name: 'cnotes.css',
              },
            },
            {loader: 'extract-loader' },
            {loader: 'css-loader' },
            {
              loader: 'postcss-loader',
              options: {
                postcssOptions: {
                  plugins: [
                    [
                      'autoprefixer',
                    ],
                  ],
                },
              },
            },
            {
              loader: 'sass-loader',
              options: {
                // Prefer Dart Sass
                implementation: require('sass'),

                // See https://github.com/webpack-contrib/sass-loader/issues/804
                webpackImporter: false,
                sassOptions: {
                  includePaths: ['./node_modules'],
                },
              },
            }
          ]
      }],
    },
  },
  {
    mode: 'development',
    entry:  {
      app: [
        './src/index.ts',
      ]
    },
    output: {
      filename: "./cnotes-compiled.js",
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
    plugins: [
      new webpack.DefinePlugin({
        __VERSION__: JSON.stringify(require("./package.json").version)
      })
    ],
    resolve: {
      extensions: [ '.tsx', '.ts', '.js' ],
    },
  },
];
