# Web Resources
This is a development directory for storing and generating artifacts that will
be moved to the user-facing HTML directories, copied from the web-staging
directory to the production storage system by bin/push.sh.

## Material Design Web
Check whether you have nodejs installed
```
node -v
```

If needed install [nodejs](https://nodejs.org/en/).

```
cd web-resources
```

To install the MD Web components and dependencies:
```
npm install
npm install --save-dev babel-core babel-loader babel-preset-es2015 dart-sass fibers
```

To compile the JavaScript source run 
```
npm run build
```

## Compile TypeScript
Install TypeScript and related utilities
```
npm install --save-dev typescript tslint @types/source-map@0.5.2
```

Lint the TypeScript
```
npm run lint
```

Compile to JavaScript
```
npm run compile_ts
```

## Mocha Testing
Install Mocha
```
npm install --save-dev mocha
```

Install Chai for assertion testing
```
npm install --save-dev chai
```

Run the unit tests
```
npm test
```

For compatibility with ES6 modules, use the babel/register module
```
npm install --save-dev @babel/register
```

## Development Testing
Development testing can be done with the Webpack dev-server. Install it
```
npm install --save-dev webpack-dev-server
```

Generate development testing code
```
npm run compile_dev
```

Run the Webpack dev-server
```
cp src/wordfinder.css dist/.
cp src/wordfinder.js dist/.
npm run start_dev
```

This will open the page dist/index.html in your browser

## JSUnit
Older tests use JS Unit