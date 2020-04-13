package main

import (
	"github.com/gin-gonic/gin"
	"github.com/r351574nc3/ultramagnus/bot"
)

var (
	response	*HealthResponse
	config      *bot.Config
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
	bot.Setup(config.New())

	b := bot.Ultramagnus.Joe
	err := b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}

	r.Run(":8080")
}
