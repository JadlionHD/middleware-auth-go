package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiUserGET(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
