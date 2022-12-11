package main

import (
	"flag"
	"os"

	"github.com/noornee/redown/external"
	"github.com/noornee/redown/utility"
)

func main() {
	var url string

	flag.StringVar(&url, "url", "", "url to parse")
	flag.Parse()

	passed := utility.IsFlagPassed(url)
	if !passed {
		flag.Usage()
		return
	}

	file, _ := os.ReadFile(url)
	fallback_url, title := utility.ParseJSONfile(file)

	video, audio := utility.GetMediaUrl(fallback_url)

	external.CMD_aria2c(video, audio)
	external.CMD_ffmpeg(title)

}
