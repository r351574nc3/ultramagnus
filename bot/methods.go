package bot

import (
	"os"
	"github.com/go-joe/joe"
	"github.com/go-joe/slack-adapter"
	slackClient "github.com/nlopes/slack"
	"github.com/robertgzr/joe-bolt-memory"
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
		slack.Adapter(c.SlackToken),
	)
	b.Slack = slackClient.New(c.SlackToken)
}


func (c *Config) New() *Config {
	if c == nil {
		c = new(Config)
	}
	c.SlackToken = os.Getenv(SLACK_TOKEN)
	c.SigningSecret = os.Getenv(SLACK_SIGNING_SECRET)
	c.DbPath = os.Getenv(DB_PATH)
	c.HttpConfig = os.Getenv(HTTP_CONFIG)
	return c
}

func (h *BotHandlers) New() *BotHandlers {
	if h == nil {
		h = new(BotHandlers)
	}

	return h
}

func (r *WebhookResponse) New() *WebhookResponse {
	if r == nil {
		return new(WebhookResponse)
	}
	return r
}

