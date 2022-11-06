Server
======
This directory contains the prototype server code for "Freecycle-Inc" service. This server is a REST API following the OpenAPI 3 spec shown in the [../openapi.yml](../openapi.yml) file. This server does not directly have a user interface as it has been designed for another program to be built to communicate with it.

## Usage
Because this is REST API it does not have a user interface, we can instead make web requests to it. The following are examples using the curl program.

Create new item:

```sh
curl -X POST -d '{
  "user_id": "user1234",
  "keywords": [
    "hammer",
    "nails",
    "tools"
  ],
  "description": "A hammer and nails set",
  "image": "https://placekitten.com/200/300",
  "lat": 51.2798438,
  "lon": 1.0830275
}' http://127.0.0.1:8000/item/
# response - 201
'{
  "id": 0,
  "user_id": "user1234",
  "keywords": [
    "hammer",
    "nails",
    "tools"
  ],
  "description": "A hammer and nails set",
  "image": "https://placekitten.com/200/300",
  "lat": 51.2798438,
  "lon": 1.0830275,
  "date_from": "2022-11-06T11:48:12.592Z"
}'
```

Delete existing item:

```sh
curl -X DELETE http://127.0.0.1:8000/item/0
# response - 204
```

## Deployment
### Prod
This app is designed to be deployed with Docker. The following command will build the app, configured ready for production:

```sh
docker build -t server .
```

Then run like this:

```sh
docker run -p 8000:8000 server
```

### Dev
During development Docker is not required for running. The following requirements will need to be installed and accessible on the PATH:

- Golang 1.19

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

## References
- <https://go.dev/doc/>
- <https://pkg.go.dev/std>
- <https://gin-gonic.com/docs/>
- <https://pkg.go.dev/github.com/go-playground/validator/v10>
- <https://pkg.go.dev/github.com/gin-contrib/cors>
