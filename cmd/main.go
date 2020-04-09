package main

import (
	"github.com/gin-gonic/gin"
)

var (
	response         *HealthResponse
)

func createRouter() *gin.Engine {
	r := gin.Default()

	// Get total hours for cases value
	r.GET("/healthcheck", func(c *gin.Context) {
		response := response.New()
		c.JSON(200, response)
	})
	return r
}

func main() {
	r := createRouter()

	r.Run(":8080")
}
