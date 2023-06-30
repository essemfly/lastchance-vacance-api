package config

import (
	"fmt"

	"github.com/slack-go/slack"
)

var slackClient *slack.Client

func InitSlack() {
	// Create a new Slack client with your API token
	slackClient = slack.New("xoxb-4368315466963-5185879956006-awiabYQEMGBonfUcsI6VdUMH")
}

func SendSlackMessage(message string) {

	// Set the channel ID where you want to send the message
	channelID := "C056134D541"

	// Post a message to a channel.
	_, _, err := slackClient.PostMessage(
		channelID,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAttachments(),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func SendOrderCreateMessage(mobile, userId, productID, rewardPrice, productUrl string) {
	SendSlackMessage("[Order Create] Mobile: " + mobile + "RewardPrice: " + rewardPrice + ", UserID: " + userId + ", ProductID: " + productID + ", ProductUrl: " + productUrl)
}
