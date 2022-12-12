package main

import (
	"flag"

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

	fallback_url, title := utility.ParseJSONBody(body)

	video, audio := utility.GetMediaUrl(fallback_url)

	external.Setup(video, audio, title)

}
