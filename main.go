package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"url-shortener/server"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static/")

	router.GET("/", server.GetIndex)
	router.POST("/", server.AddLink)
	router.GET("/:token", server.Redirect)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	router.Run("localhost:8080")
}
