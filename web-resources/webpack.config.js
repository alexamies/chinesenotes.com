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
    //mode: 'development',
    entry:  {
      app: [
        './cnotes.js',
        './script/find.js'
      ]
    },
    output: {
      filename: "./cnotes-compiled.js",
    },
    module: {
      rules: [{
        test: /\.js$/,
        loader: "babel-loader",
        query: {
          presets: ["@babel/preset-env"],
        },
      }],
    },
  },
];
