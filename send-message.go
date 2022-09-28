package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	// SLACK_BOT is in env file contains token of Slack app which lets us connect our bot
	api := slack.New(os.Getenv("SLACK_BOT_AUTH_TOKEN"))

	chID, timestamp, err := api.PostMessage("C044Z5GB63S",
		slack.MsgOptionText("Hello, Slack bot", false))
	handleError(err)
	fmt.Printf("Message successfully sent to channel %s at: %s", chID, timestamp)

}
