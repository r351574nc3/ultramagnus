package main

import (
	"github.com/gin-gonic/gin"
	"github.com/r351574nc3/ultramagnus/bot"
)

var (
	config      *bot.Config
)

func createRouter() *gin.Engine {
	r := gin.Default()
	var response *HealthResponse


	// Get total hours for cases value
	r.GET("/healthcheck", func(c *gin.Context) {
		response = response.New()
		c.JSON(200, response)
	})

	r.POST("/motd", bot.SlackHandlers.Motd)
	return r
}

func StartHttp(httpConfig string) {
	r := createRouter()
	go r.Run(httpConfig)	
}

func main() {
	config = config.New()
	bot.Setup(config)
	go StartHttp(config.HttpConfig)

	b := bot.Ultramagnus.Joe
	err := b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}
}
