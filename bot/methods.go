package bot

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/go-joe/joe"
	"github.com/go-joe/slack-adapter"
	slackClient "github.com/nlopes/slack"
	"github.com/robertgzr/joe-bolt-memory"
	"go.uber.org/zap"
)

func (b *ChatBot) New() *ChatBot {
	if b == nil {
		return new(ChatBot)
	}
	return b
}


func (b *ChatBot) Init(c *Config) {
	b.Joe = joe.New(
		"ultramagnus",
		bolt.Memory(c.DbPath),
		slack.Adapter(c.Slack.AccessToken),
	)
	b.Slack = slackClient.New(c.Slack.AccessToken)
}

func (b *ChatBot) Logger() *zap.Logger {
	if b.logger == nil {
		logger, _ := zap.NewProduction()
		b.logger = logger
		defer b.logger.Sync()
	}
	return b.logger
}


func (c *Config) New() *Config {
	if c == nil {
		c = new(Config)
		c.Slack = new(SlackConfig)
	}
	secrets := []byte(os.Getenv(SLACK_SECRETS))
	err := json.Unmarshal(secrets, c.Slack)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", c.Slack)
	c.DbPath = os.Getenv(DB_PATH)
	c.HttpConfig = os.Getenv(HTTP_CONFIG)

	return c
}

func (h *BotHandlers) New() *BotHandlers {
	if h == nil {
		h = new(BotHandlers)
	}

	h.Motd = motd

	return h
}

func (r *WebhookResponse) New(message string) *WebhookResponse {
	if r == nil {
		r = new(WebhookResponse)
	}

	r.Buffer = message
	return r
}

