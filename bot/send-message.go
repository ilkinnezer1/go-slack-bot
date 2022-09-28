package bot

import (
	"fmt"
	"github.com/slack-go/slack"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
func sendMessage() {
	// SLACK_BOT is in env file contains token of Slack app which lets us connect our bot
	api := slack.New(authToken)

	chID, timestamp, err := api.PostMessage(channelId,
		slack.MsgOptionText("Hello, Slack bot", false))
	handleError(err)
	fmt.Printf("Message successfully sent to channel %s at: %s", chID, timestamp)

}
