package main

import (
	"net/http"

	"github.com/coffeloop/go-fabric5/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := gin.Default()

	// Middleware de autenticación
	authMiddleware := func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || username != "admin" || password != "password" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}

	v1 := r.Group("/api/v1", authMiddleware)
	{
		v1.POST("/create-fabric-ca", handlers.CreateFabricCA)
		v1.POST("/register-fabric-ca", handlers.RegisterFabricCA)
		v1.POST("/create-fabric-peer", handlers.CreateFabricPeer)
		v1.POST("/create-fabric-orderer", handlers.CreateFabricOrderer)

		v1.GET("/check-fabric-ca-status", handlers.CheckFabricCAStatus)
		v1.GET("/check-fabric-peer-status", handlers.CheckFabricPeerStatus)
		v1.GET("/check-fabric-orderer-status", handlers.CheckFabricOrdererStatus)
	}

	r.Run(":8080")
}
