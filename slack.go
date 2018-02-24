package main

import (
	"fmt"
	"os"
	"strconv"

	slack "github.com/ashwanthkumar/slack-go-webhook"
)

// Post takes a URL and a message. It then
// posts to slack the URL and the message.
// If the webhook url can't be found, calls
// os.exit(1)
func Post(url string, status int, timer int) {
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		fmt.Println("couldn't find a webhook url")
		os.Exit(1)
	}
	payload := slack.Payload{
		Text:      getMessage(url, status, timer),
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

func getMessage(url string, status int, timer int) string {
	var message string
	switch status {
	case Up:
		message = url + " is back up after " + strconv.Itoa(timer) + " seconds."
	case Down:
		message = url + " is down."
	case BadStatus:
		message = "Recieving a bad status from " + url + ". Best check it out!"
	}
	return message
}
