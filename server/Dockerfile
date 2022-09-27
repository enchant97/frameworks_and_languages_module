FROM golang:alpine as builder

    WORKDIR /usr/src/app

    COPY go.mod go.sum ./
    RUN go mod download && go mod verify

    COPY . .

    RUN go build -v -o /usr/local/bin/app

FROM alpine:latest

    COPY --from=builder /usr/local/bin/app /usr/local/bin/app

    ENV SERVER_BIND="0.0.0.0:8000"
    ENV GIN_MODE=release
    EXPOSE 8000

    CMD ["app"]
