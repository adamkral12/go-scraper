package notifier

import (
"github.com/bluele/slack"
	"os"
)

const (
	channelName = "@adam.kral"
)

func SendMessage(message string) {
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	err := api.ChatPostMessage(channelName, message, nil)
	if err != nil {
		panic(err)
	}
}