package main

import (
	"fmt"
	"os"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

func Post(url string) {
	webhookURL := ""
	payload := slack.Payload{
		Text:      url + "is down",
		Username:  "uptime bot",
		Channel:   "#matttestingsomestuff",
		IconEmoji: ":cold_sweat:",
	}
	err := slack.Send(webhookURL, "", payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
