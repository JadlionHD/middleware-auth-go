package main

import (
	"github.com/JadlionHD/middleware-auth-go/api/router"
	middleware "github.com/JadlionHD/middleware-auth-go/api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", router.LoginRoutePOST)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	api.GET("/user", router.ApiUserGET)

	r.Run("localhost:8080")
}
