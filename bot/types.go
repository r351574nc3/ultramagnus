package bot

import (
	"github.com/go-joe/joe"
	"github.com/nlopes/slack"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

)

type ChatBot struct {
	Joe *joe.Bot 
	Slack *slack.Client
	logger *zap.Logger
}

type Config struct {
	DbPath string
	Slack *SlackConfig
	HttpConfig string
}

type SlackConfig struct {
	AccessToken string `json:"SLACK_BOT_ACCESS_TOKEN"`
	VerifyToken string `json:"SLACK_VERIFY_TOKEN"`
	ClientId string `json:"SLACK_CLIENT_ID"`
	ClientSecret string `json:"SLACK_CLIENT_SECRET"`
	SigningSecret string `json:"SLACK_SIGNING_SECRET"`
}

type BotHandlers struct {
	Motd func (c *gin.Context)
}

type WebhookResponse struct {
	Buffer string `json:"buffer"`
}