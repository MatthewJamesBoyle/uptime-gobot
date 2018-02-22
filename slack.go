package main

import (
	"fmt"
	"os"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

func Post(url string) {
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		fmt.Println("couldn't find a webhook url")
		os.Exit(1)
	}
	payload := slack.Payload{
		Text:      url + "is down",
		Username:  os.Getenv("SLACK_USERNAME"),
		Channel:   os.Getenv("SLACK_CHANNEL"),
		IconEmoji: ":cold_sweat:",
	}
	err := slack.Send(webhookURL, "", payload)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
