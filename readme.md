# UptimeGobot

Inspired by https://uptimerobot.com/, uptimeGobot is a little program that will monitor as many URLs as you like and report to Slack if they are down.

## Getting Started
Clone the repo:
`git clone https://github.com/MatthewJamesBoyle/uptime-gobot.git`

cd into the directory and then run

`go build gobot.go slack.go`

This will build the project into a binary.

Now you can run:

`./gobot  -files=files.txt`

Where `files.txt` is a list of urls you would like to monitor. An example file is included.

By default, uptimeGobot will poll all the urls in your file every second concurrently. You can increase the poll rate by passing an optional second parameter:

```
// Poll every 5 seconds instead of 1.
./uptime_gobot -file=files.txt -interval=5
```

One last thing! You need to add an env file containing a couple of things. An example is included in the repo.



* add a flag to see if a url is down and only send the webhook once.
* Add a timer for how long a site is down
* Add a webhook for if a site goes back up
* Add tests
