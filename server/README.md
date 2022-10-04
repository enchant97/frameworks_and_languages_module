Server
======

## Prod
This app is designed to be deployed with Docker. The following command will build the app, configured ready for production:

```sh
docker build -t server .
```

Then run like this:

```sh
docker run -p 8000:8000 server
```

## Dev
During development Docker is not required for running. The following requirements will need to be installed and accessible on the PATH:

- Golang 1.19

### Usage
To run the development server you can use this command:

```sh
go run .
```

This will run the server by default at localhost on port 8000. You can customize this behaviour using the `SERVER_BIND` environment variable:

```sh
export SERVER_BIND=127.0.0.1:8001
go run .
```

If you want to run the app in release mode for testing you can set this environment variable:

```sh
export GIN_MODE=release
```
