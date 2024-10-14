package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var validtoken = "YWRtaW58YWRtaW4="

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") || token[7:] != validtoken {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
	}
}
