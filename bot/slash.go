package bot

import (
	/*
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/slack-go/slack"
	*/
	"github.com/gin-gonic/gin"
)

func motd(c *gin.Context) {
	var (
		response *WebhookResponse
	)
	response = response.New(
		`This is the message of the day. Here we will get a briefing of what's happening concerning the project. Things it will include are:
* Who's on support.
* What's hot right now. (What's being worked on like stories, issues, or conversations).
* Reminders
* PRs that need looking at
* Stories that have spent an excessive amount of time in the bucket
* Low hanging fruit`,
	)
	c.String(200, response.Buffer)




	/*
	verifier, err := slack.NewSecretsVerifier(r.Header, signingSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &verifier))
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = verifier.Ensure(); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/echo":
		params := &slack.Msg{Text: s.Text}
		b, err := json.Marshal(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	*/
}