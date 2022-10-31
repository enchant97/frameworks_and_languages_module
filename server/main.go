package main

import (
	"embed"
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

func main() {
	serverBind := "127.0.0.1:8000"

	// change server bind with environment variable
	if value, isSet := os.LookupEnv("SERVER_BIND"); isSet {
		serverBind = value
	}

	r := gin.Default()

	// init the templates
	templates := template.Must(template.New("").ParseFS(f, "templates/*"))
	r.SetHTMLTemplate(templates)

	InitRoutes(r)

	// run the actual server, now everything has been setup
	r.Run(serverBind)
}
