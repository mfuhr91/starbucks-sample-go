package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func Redirect(path string, c *gin.Context) {
	// redirect
	location := url.URL{Path: path}
	c.Redirect(http.StatusFound, location.RequestURI())
}
