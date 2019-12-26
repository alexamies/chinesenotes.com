// Karma configuration
// Generated on Thu Dec 19 2019 20:01:52 GMT-0800 (Pacific Standard Time)

module.exports = function(config) {
  config.set({

    // base path that will be used to resolve all patterns (eg. files, exclude)
    basePath: '',

    // frameworks to use
    // available frameworks: https://npmjs.org/browse/keyword/karma-adapter
    frameworks: ["jasmine", "karma-typescript"],

    // list of files / patterns to load in the browser
    files: [
      //{ pattern: "src/*.ts" },
      { pattern: "src/CNDictionaryEntry.ts" },
      { pattern: "src/CNInterfaces.ts" },
      // { pattern: "src/CNotes.ts" },  // Crashes Karma
      { pattern: "src/CNotesMenu.ts" },
      { pattern: "src/CNWordSense.ts" },
      //{ pattern: "src/CorpusDocView.ts" }, // Crashes Karma
      { pattern: "src/DocumentFinder.ts" },
      { pattern: "src/DocumentFinderView.ts" },
      { pattern: "src/HrefVariableParser.ts" },
      { pattern: "src/ResultsParser.ts" },
      { pattern: "src/ResultsView.ts" },
      { pattern: "src/SubstringApp.ts" },
      // { pattern: "src/WordFinder.ts" }, // Crashes Karma
      // { pattern: "src/WordFinderAdapter.ts" }, // Crashes Karma
      { pattern: "src/WordFinderNavigation.ts" },
      { pattern: "src/WordFinderView.ts" },
      { pattern: "test/*.spec.ts" }
    ],

    karmaTypescriptConfig: {
      compilerOptions: {
        emitDecoratorMetadata: true,
        esModuleInterop: true,
        experimentalDecorators: true,
        module: "commonjs",
        sourceMap: true,
        target: "ES6"
      },
      exclude: ["node_modules"]
    },

    // preprocess matching files before serving them to the browser
    // available preprocessors: https://npmjs.org/browse/keyword/karma-preprocessor
    preprocessors: {
      "**/*.ts": ["karma-typescript"]
    },

    // test results reporter to use
    // possible values: 'dots', 'progress'
    // available reporters: https://npmjs.org/browse/keyword/karma-reporter
    reporters: ["dots", "karma-typescript"],

    // web server port
    port: 9876,

    // enable / disable colors in the output (reporters and logs)
    colors: true,

    // level of logging
    // possible values: config.LOG_DISABLE || config.LOG_ERROR || config.LOG_WARN || config.LOG_INFO || config.LOG_DEBUG
    logLevel: config.LOG_INFO,

    // start these browsers
    // available browser launchers: https://npmjs.org/browse/keyword/karma-launcher
    browsers: ['ChromeHeadless'],

    autoWatch: false,

    // Continuous Integration mode
    // if true, Karma captures browsers, runs the tests and exits
    singleRun: true,

    // Concurrency level
    // how many browser should be started simultaneous
    concurrency: Infinity
  })
}
