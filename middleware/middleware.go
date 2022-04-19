package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	USERNAME = "topan"
	PASSWORD = "pass"
)

func Auth(c *gin.Context) bool {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Auth Required",
		})
		return false
	}

	valid := (username == USERNAME) && (password == PASSWORD)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unauth",
		})
		return false
	}
	return true
}
