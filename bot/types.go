package bot

import (
	"github.com/go-joe/joe"
	"github.com/nlopes/slack"
	"github.com/gin-gonic/gin"
)

type ChatBot struct {
	Joe *joe.Bot 
	Slack *slack.Client 
}

type Config struct {
	DbPath string
	SlackToken string
	SigningSecret string
	HttpConfig string
}

type BotHandlers struct {
	Motd func (c *gin.Context)
}

type WebhookResponse struct {
	Buffer string `json:"buffer"`
}