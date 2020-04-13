package bot

import (
	"github.com/go-joe/joe"
	"github.com/nlopes/slack"
)

type ChatBot struct {
	Joe *joe.Bot 
	Slack *slack.Client 
}

type Config struct {
	DbPath string
	SlackToken string
}