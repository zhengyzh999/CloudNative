package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var token = "1234567"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.Request.Header.Get("access_token")
		fmt.Println("access_token: ", accessToken)
		if accessToken != token {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "token校验失败",
			})
			c.Abort()
		}
		c.Next()
	}
}
