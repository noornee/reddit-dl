package main

import (
	"flag"

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
}
