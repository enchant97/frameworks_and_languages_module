Client
======

## Prod
This app is designed to be deployed with Docker. The following command will build the app, configured ready for production:

```sh
docker build -t client .
```

Then run like this:

```sh
docker run -p 8001:8001 client
```

## Dev
During development Docker is not required for running. The following requirements will need to be installed and accessible on the PATH:

- npm 8
- node.js 18

### Usage
First the requirements will need to be installed:

```sh
npm i
```

To run the development server the following command can be used:

```sh
npm run dev
```

This will run the client server at localhost on port 8001. You can also bind to host using this command:

```sh
npm run dev -- --host
```

If the port needs to be changed you can also use:

```sh
npm run dev -- --port 3000
```
