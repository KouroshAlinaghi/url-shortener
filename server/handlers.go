package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"url-shortener/shortener"
)

const (
	domain = "localhost:8080"
)

func GetIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"action": false,
	})
}

func AddLink(c *gin.Context) {
	url := c.PostForm("url")
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	shorten := shortener.Add(url)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"action": true,
		"original": url,
		"shorten": domain + "/" + shorten,
	})
}

func Redirect(c *gin.Context) {
	token := c.Param("token")
	url, err := shortener.Get(token)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
