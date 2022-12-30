package main

import (
	"fmt"
	"os"

	"github.com/noornee/reddit-dl/handler"
	"github.com/noornee/reddit-dl/service"
	"github.com/noornee/reddit-dl/utility"
	"github.com/urfave/cli/v2"
)

func main() {

	// default url incase the url flag isnt passed
	var raw_url string

	if len(os.Args) > 1 {
		raw_url = os.Args[1]
	}

	app := &cli.App{
		Name:  "reddit-dl",
		Usage: "A reddit multimedia downloader",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "url",
				Usage: "a reddit post url",
				Value: raw_url,
			},
		},
		Action: func(ctx *cli.Context) error {
			url := ctx.String("url")
			if url == "" {
				cli.ShowAppHelp(ctx)
			}

			controller(url)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println()
		utility.ErrorLog.Println(err)
	}

}

func controller(raw_url string) {
	url, title := handler.ParseUrl(raw_url)

	// body -> the url response body in form of []bytes
	body, err := handler.GetBody(url)
	if err != nil {
		utility.ErrorLog.Fatal(err)
	}

	// media -> an array of string(s) containing the url
	media_url, err := utility.ParseJSONBody(body)

	if err != nil {
		utility.ErrorLog.Fatal(err)
	}

	for _, url := range media_url {

		// if it's a reddit gallery(multiple images in a post), then it's going to contain multiple urls
		// its length would be greater than 1
		if len(media_url) <= 1 {
			media, audio := utility.GetMediaUrl(url)
			service.Setup(media, audio, title)
		} else {
			service.Setup(url, "", title)
		}
	}

}
