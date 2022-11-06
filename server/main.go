/*
The prototype REST API server for the "Freecycle-Inc" service

Usage:

	server

Environment Variables:

	SERVER_BIND
		Bind the server to a different interface & port;
		other than the default 127.0.0.1:8000

	GIN_MODE
		What mode to run gin in, defaults to 'debug', accepts 'production'
*/
package main

import (
	"embed"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

// SOURCE: https://pkg.go.dev/embed
//
//go:embed templates/*
var f embed.FS

func main() {
	serverBind := "127.0.0.1:8000"

	// change server bind with environment variable
	// SOURCE: https://pkg.go.dev/os#LookupEnv
	if value, isSet := os.LookupEnv("SERVER_BIND"); isSet {
		serverBind = value
	}

	r := gin.Default()

	// init the templates
	// SOURCE: https://gin-gonic.com/docs/examples/html-rendering/
	templates := template.Must(template.New("").ParseFS(f, "templates/*"))
	r.SetHTMLTemplate(templates)

	InitRoutes(r)

	// run the actual server, now everything has been setup
	r.Run(serverBind)
}
