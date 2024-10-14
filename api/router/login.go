package router

import (
	"fmt"
	"net/http"

	b64 "encoding/base64"

	"github.com/JadlionHD/middleware-auth-go/api/types"
	"github.com/gin-gonic/gin"
)

// harcoded value, meant for testing
var testUsername = "admin"
var testPassword = "admin"

func LoginRoutePOST(c *gin.Context) {
	var data types.LoginAuth

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid json format"})
		return
	}

	if data.Username == testUsername && data.Password == testPassword {
		// Should generating a jwt token or similar, but for now I'll using a base64 encoding
		base64Encoded := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s|%s", data.Username, data.Password)))
		c.JSON(http.StatusOK, gin.H{
			"token": base64Encoded,
		})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username & password",
		})
		return
	}

}
