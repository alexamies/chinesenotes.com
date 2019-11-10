# Web Resources
This is a development directory for storing and generating artifacts that will
be moved to the user-facing HTML directories, copied from the web-staging
directory to the production storage system by bin/push.sh.

Run the following commands from this directory.

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
```

To resolve module dependencies and package the JavaScript source run 
```
npm run build
```

## Compile TypeScript
You should have TypeScript installed already from the `npm install` command above.

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
npm run start_dev
```

Navigate to the page http://localhost:8080/test/index.html

## JSUnit
Older tests use JS Unit