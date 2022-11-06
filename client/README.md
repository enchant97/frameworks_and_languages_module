Client
======
This directory contains the code for client front-end server for the "Freecycle-Inc" service. It provides a user interface to the prototype api server. It is written to communicate with a REST API server following the OpenAPI 3 spec shown in the [../openapi.yml](../openapi.yml) file.

## Usage
Because this is a prototype; a api server base url must be set in the url (what you enter in the browser url bar), shown in the example below:

```
http://client:8001/?api=http://server:8000
```

When using the UI from a browser CORS will be used, normally the server would need to accept requests from the client server url. This is not needed if using the server prototype as by default is set to allow all origins.

## Deployment

### Prod
This app is designed to be deployed with Docker. The following command will build the app, configured ready for production:

```sh
docker build -t client .
```

Then run like this:

```sh
docker run -p 8001:8001 client
```

### Dev
During development Docker is not required for running. The following requirements will need to be installed and accessible on the PATH:

- npm 8
- node.js 18

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

## References
- <https://www.solidjs.com/guides/getting-started>
- <https://tailwindcss.com/docs/installation>
- <https://nginx.org/en/docs/beginners_guide.html>
- <https://www.typescriptlang.org/docs/handbook/typescript-from-scratch.html>
