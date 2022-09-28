package bot

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func sendNotification() {
	// Takes all command line arguments
	args := os.Args[1:]
	fmt.Println(args)

	api := slack.New(authToken)
	headText := "*Jenkins build has finished*"
	// Command line properties
	jenUrl := "*Build URL: *" + args[0]
	buildRes := "*" + args[1] + "*"
	buildNum := "*" + args[2] + "*"
	jenJobName := "*" + args[3] + "*"

	// Message formatting
	if buildRes == "*SUCCESS*" {
		buildRes = buildRes + " :white_check_mark:"
	} else {
		buildRes = buildRes + " :x:"
	}
	// Message part
	dividerBlock := slack.NewDividerBlock()
	buildDetails := jenJobName + " #" + buildNum + " " + buildRes + "\n" + jenUrl
	txtBlockObj := slack.NewTextBlockObject("mrkdwn", headText+"\n\n\n", false, false)
	buildDetailsArea := slack.NewTextBlockObject("mrkdwn", buildDetails, false, false)
	// Sections on message
	buildDetailsSection := slack.NewSectionBlock(buildDetailsArea, nil, nil)
	headTextSection := slack.NewSectionBlock(txtBlockObj, nil, nil)

	botMsg := slack.MsgOptionBlocks(headTextSection, dividerBlock, buildDetailsSection)

	// Send message to the Slack channel
	_, _, _, err := api.SendMessage(channelId, botMsg)
	handleError(err)
}
