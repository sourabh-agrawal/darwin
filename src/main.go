package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var status = newHealthCheck()

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/status", getStatus)

	r.Run("localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getStatus(c *gin.Context) {
	c.JSON(status.GetStatus(), status)
}
