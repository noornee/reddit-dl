package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/noornee/reddit-dl/external"
	"github.com/noornee/reddit-dl/handler"
	"github.com/noornee/reddit-dl/utility"
)

func main() {
	var raw_url string

	flag.StringVar(&raw_url, "url", "", "url to parse")
	flag.Parse()

	passed := utility.IsFlagPassed(raw_url)

	// if a flag isnt passed and a url is passed, assign the url to raw_url
	if len(os.Args) > 1 && !strings.HasSuffix(os.Args[1], "-url") {
		raw_url = os.Args[1]
	} else {
		if !passed {
			flag.Usage()
			return
		}

	}

	url, title := handler.ParseUrl(raw_url)

	// body -> the url response body in form of []bytes
	body, err := handler.GetBody(url)
	if err != nil {
		utility.ErrorLog.Fatal(err)
	}

	// media -> an array of string(s) containing the url
	media_url, err := utility.ParseJSONBody(body)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, url := range media_url {

		// if it's a reddit gallery kinda image, then it's going to contain multiple urls
		// its length would be greater than 1
		if len(media_url) < 1 {
			media, audio := utility.GetMediaUrl(url)
			external.Setup(media, audio, title)
		} else {
			external.Setup(url, "", title)
		}
	}

}
