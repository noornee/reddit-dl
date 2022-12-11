package main

import (
	"flag"
	"fmt"
	"os"

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
	fallback_url := utility.ParseJSONfile(file)

	video, audio := utility.GetMediaUrl(fallback_url)
	fmt.Printf("%s\n%s\n", video, audio)
}
