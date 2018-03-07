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

const (
	Up        = iota
	Down      = iota
	BadStatus = iota
)

func init() {
	gotenv.Load()

}

type trackedURL struct {
	url   string
	up    bool
	timer int
}

func main() {

	file, intervalInSeconds := readArgs()

	urlsToMonitor := parseURLs(string(file))

	c := make(chan trackedURL)

	for _, url := range urlsToMonitor {
		go checkURL(url, c)
	}

	for l := range c {
		go func(l trackedURL) {
			time.Sleep(intervalInSeconds)
			checkURL(l, c)
		}(l)
	}
}

// readArgs looks for the -file and -interval
// flag in the command line when the program
// is executed. if a file is not present, the
// program will exit. if no interval is passed,
// it defailts to 1 second.
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

// Takes the contents of the file
// and parses all urls out of it.
// If no urls are present, the program
// will exit.
func parseURLs(file string) []trackedURL {
	urls := xurls.Strict().FindAllString(file, -1)
	if len(urls) == 0 {
		fmt.Println("No urls were found in the file you provided.")
		os.Exit(1)
	}
	var urlsToMonitor []trackedURL
	for _, url := range urls {
		urlsToMonitor = append(urlsToMonitor, trackedURL{url, true, 0})
	}

	return urlsToMonitor
}

// checkURL takes a url and a channel
// to pass the data into. It does a get
// request to the url passed and if it
// does not get a 200 status code, it
// calls Post() and writes to Slack.
func checkURL(tracked trackedURL, c chan trackedURL) {

	resp, err := http.Get(tracked.url)
	if err != nil {
		if tracked.up {
			updateStatus(&tracked, false)
		} else {
			updateStatus(&tracked, false)

			Post(tracked.url, Down, tracked.timer)

		}
		c <- tracked
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if tracked.up {
			updateStatus(&tracked, false)
			Post(tracked.url, BadStatus, tracked.timer)

		} else {
			updateStatus(&tracked, false)

		}
		c <- tracked
		return
	}

	if !tracked.up {
		updateStatus(&tracked, true)
		Post(tracked.url, Up, tracked.timer)

		return
	}
	updateStatus(&tracked, true)
	c <- tracked

}

func updateStatus(tu *trackedURL, up bool) {
	if up {
		tu.up = true
		tu.timer = 0
	} else {
		tu.up = false
		tu.timer++
	}
}

func (t trackedURL) print() {
	fmt.Println(t)
}
