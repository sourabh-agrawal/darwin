package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	hc "github.com/sourabh-agrawal/darwin/src/healthcheck"
)

var status = hc.New()

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/status", getStatus)
	r.PUT("/status/:code", updateStatus)

	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getStatus(c *gin.Context) {
	interval, err := time.ParseDuration(
		fmt.Sprintf("%dms", rand.Intn(50)),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	time.Sleep(interval)
	c.JSON(status.GetStatus(), status)
}

func updateStatus(c *gin.Context) {
	code, err := strconv.Atoi(c.Param("code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	status.UpdateStatus(code)
	c.JSON(status.GetStatus(), status)
}
