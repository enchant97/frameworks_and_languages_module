package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	serverBind := "127.0.0.1:8000"

	if value, isSet := os.LookupEnv("SERVER_BIND"); isSet {
		serverBind = value
	}

	r := gin.Default()

	r.Run(serverBind)
}
