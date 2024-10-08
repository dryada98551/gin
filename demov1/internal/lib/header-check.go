package lib

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func HeaderCheck(c *gin.Context) (string, error) {
	JWTToken := c.GetHeader("JWTToken")
	LoginAt := c.GetHeader("LoginAt")

	if JWTToken == "" || LoginAt == "" {
		return "", errors.New("Header parameter is missing")
	}

	return "200", nil
}
