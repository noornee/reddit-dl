package main

import (
	"flag"
	"log"

	"github.com/noornee/reddit-dl/external"
	"github.com/noornee/reddit-dl/handler"
	"github.com/noornee/reddit-dl/utility"
)

func main() {
	var raw_url string

	flag.StringVar(&raw_url, "url", "", "url to parse")
	flag.Parse()

	passed := utility.IsFlagPassed(raw_url)
	if !passed {
		flag.Usage()
		return
	}

	url := handler.ParseUrl(raw_url)

	body := handler.GetBody(url)

	fallback_url, title, err := utility.ParseJSONBody(body)

	if err != nil {
		log.Fatal(err)
		return
	}

	video, audio := utility.GetMediaUrl(fallback_url)

	external.Setup(video, audio, title)

}
