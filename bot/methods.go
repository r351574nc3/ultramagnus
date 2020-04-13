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
	c.DbPath = os.Getenv(DB_PATH)
	return c
}
