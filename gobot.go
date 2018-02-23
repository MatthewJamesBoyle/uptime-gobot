package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"mvdan.cc/xurls"
)

const down = " is down"
const badStatus = " is returning a non 200 status code. Worth checking? "

func init() {
	gotenv.Load()

}

func main() {

	file, intervalInSeconds := readArgs()

	urlsToMonitor := parseURLs(string(file))

	c := make(chan string)

	for _, url := range urlsToMonitor {
		go checkURL(url, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(intervalInSeconds)
			checkURL(l, c)
		}(l)
	}
}

func readArgs() ([]byte, time.Duration) {
	filePath := flag.String("file", "", "the filepath of the URLs you want to monitor")
	i := flag.Int("interval", 1, "the gap between polls at your provided urls in seconds. 1 second by default")
	flag.Parse()
	intervalInSeconds := time.Duration(time.Second) * time.Duration(*i)

	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Println("error!", err)
		os.Exit(1)
	}
	return file, intervalInSeconds
}

func parseURLs(file string) []string {
	urlsToMonitor := xurls.Strict().FindAllString(file, -1)
	if len(urlsToMonitor) == 0 {
		fmt.Println("No urls were found in the file you provided.")
		os.Exit(1)
	}
	return urlsToMonitor
}

func checkURL(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		Post(url, down)
		c <- url
		return
	}
	if resp.StatusCode != 200 {
		Post(url, badStatus)
		c <- url
		return
	}

	c <- url
	defer resp.Body.Close()

}
