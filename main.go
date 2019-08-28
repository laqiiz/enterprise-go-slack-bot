package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/laqiiz/enterprise-go-slack-bot/listener"
	"github.com/nlopes/slack"
	"log"
	"os"
)

// https://api.slack.com/slack-apps
// https://api.slack.com/internal-integrations
type envConfig struct {
	// BotToken is bot user token to access to slack API.
	BotToken string `envconfig:"BOT_TOKEN" required:"true"`

	// VerificationToken is used to validate interactive messages from slack.
	VerificationToken string `envconfig:"VERIFICATION_TOKEN" required:"true"`
}

func main() {
	os.Exit(_main(os.Args[1:]))
}

func _main(args []string) int {
	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		return 1
	}

	// Listening slack event and response
	log.Printf("[INFO] Start slack event listening")
	slackListener := &listener.SlackListener{
		Client:    slack.New(env.BotToken),
	}
	slackListener.ListenAndResponse()

	return 0
}
