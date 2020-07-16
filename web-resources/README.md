# Web Resources
This is a development directory for storing and generating artifacts that will
be moved to the user-facing HTML directories, copied from the web-staging
directory to the production storage system by bin/push.sh.

Run the following commands from this directory.

## Material Design Web
Check whether you have nodejs installed

```shell
node -v
```

If needed install [nodejs](https://nodejs.org/en/).

```shell
cd web-resources
```

To install the MD Web components and dependencies:

```shell
npm install
```

To resolve module dependencies and package the JavaScript source run 

```shell
npm run build
```

## Compile TypeScript
You should have TypeScript installed already from the `npm install` command above.

Lint the TypeScript

```shell
npm run lint
```

Compile to JavaScript

```shell
npm run compile_ts
```

## Unit Testing

Run the unit tests

```shell
npm test
```
