package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"mvdan.cc/xurls"
)

func main() {

	// If no time is passed, lets just default it to 1 second.
	intervalInSeconds := time.Second

	argsLength := len(os.Args)

	if argsLength < 2 {
		fmt.Println("you didn't pass a file path. Please pass a comma seperated file containing the urls you want to check for uptime.")
		os.Exit(1)
	}
	if argsLength == 3 {
		interval, err := strconv.ParseInt(os.Args[2], 10, 0)

		if err != nil {
			fmt.Println("2nd arguement passed was meant to be an integer but wasn't. Please try again.")
			os.Exit(1)
		}
		intervalInSeconds *= time.Duration(interval)
	}

	filePath := os.Args[1]

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("error!", err)
		os.Exit(1)
	}

	urlsToMonitor := xurls.Strict().FindAllString(string(file), -1)
	if len(urlsToMonitor) == 0 {
		fmt.Println("No urls were found in the file you provided.")
		os.Exit(1)
	}

	c := make(chan string)

	for _, url := range urlsToMonitor {
		go checkUrl(url, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(intervalInSeconds)
			checkUrl(l, c)
		}(l)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		Post(url)
		c <- url
		return
	}
}
