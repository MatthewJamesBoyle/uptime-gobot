# UptimeGobot

Inspired by https://uptimerobot.com/, uptimeGobot is a little program that will monitor as many URLs as you like and report to Slack if they are down.

## Getting Started
Clone the repo:
`git clone https://github.com/MatthewJamesBoyle/uptime-gobot.git`

cd into the directory and then run

`go build uptime_gobot.go slack.go`

This will build the project into a binary.

Now you can run:

`./uptime_gobot files.txt`

Where `files.txt` is a list of urls you would like to monitor. An example file is included.

By default, uptimeGobot will poll all the urls in your file every second concurrently. You can increase the poll rate by passing an optional second parameter:

```
// Poll every 5 seconds instead of 1.
./uptime_gobot files.txt 5
```

One last thing! You need to add an env file containing a couple of things. An example is included in the repo.


