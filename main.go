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

	body, err := handler.GetBody(url)
	if err != nil {
		utility.ErrorLog.Fatal(err)
	}

	fallback_url, err := utility.ParseJSONBody(body)

	if err != nil {
		log.Fatal(err)
		return
	}

	video, audio := utility.GetMediaUrl(fallback_url)

	external.Setup(video, audio, title)

}
