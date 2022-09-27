package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {
	rootRoutes := engine.Group("/")
	{
		rootRoutes.GET("/", GetIndex)
	}
}

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
