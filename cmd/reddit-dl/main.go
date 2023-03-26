package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/noornee/reddit-dl/internal/handler"
	"github.com/noornee/reddit-dl/internal/service"
	"github.com/noornee/reddit-dl/internal/utility"
	"github.com/urfave/cli/v2"
)

func main() {

	// default url incase the url flag isnt passed
	var raw_url string

	if len(os.Args) > 1 {
		raw_url = os.Args[1]
	}

	app := &cli.App{
		Name:    "reddit-dl",
		Usage:   "A reddit multimedia downloader",
		Version: "0.66.5",
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
				return nil
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
	url, title, err := handler.ParseUrl(raw_url)
	if err != nil {
		utility.ErrorLog.Fatal(err)
	}

	// body -> the url response body in form of []bytes
	body, err := handler.GetBody(url)
	if err != nil {
		utility.ErrorLog.Fatal(err)
	}

	reddit_data, err := utility.ParseJSONBody(body)
	if err != nil {
		utility.ErrorLog.Fatal(err)
	}

	var wg sync.WaitGroup

	if reddit_data.IsRedditGallery {
		for _, url := range reddit_data.GalleryUrls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				service.Setup(url, "", title)
			}(url)
		}
		wg.Wait()

	} else {
		media, audio := utility.GetMediaUrl(reddit_data.MediaUrl)
		service.Setup(media, audio, title)
	}
}
