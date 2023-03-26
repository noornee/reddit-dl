package service

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/gookit/color"
	"github.com/noornee/reddit-dl/internal/utility"
)

// create a directory temp
func createDir() string {

	const temp_dir = ".reddit_temp"

	// if temp_dir exists, delete
	if _, err := os.Stat(temp_dir); !os.IsNotExist(err) {
		os.RemoveAll(temp_dir)
	}

	err := os.Mkdir(temp_dir, os.ModePerm)
	if err != nil {
		utility.ErrorLog.Println(err)
	}

	return temp_dir
}

func downloaded_size(resp *grab.Response) string {

	size_bytes := resp.BytesComplete()
	size_kb := size_bytes / (1 << 10)
	// mb := size_bytes / (1 << 20)

	return color.Green.Sprintf("%dKB", size_kb)

}

func total_download_size(resp *grab.Response) string {

	size_bytes := resp.Size()
	size_kb := size_bytes / (1 << 10)

	return color.Blue.Sprintf("%dKB", size_kb)

}

// download progress
func download_progress(resp *grab.Response, t *time.Ticker) {
Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v\t%.2f%%\n",
				downloaded_size(resp),
				total_download_size(resp),
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

}

func downloader(urls []string) {
	var temp_dir string = createDir() + "/"

	// create client
	client := grab.NewClient()

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			req, _ := grab.NewRequest(temp_dir, url)

			resp := client.Do(req)

			// start UI loop
			t := time.NewTicker(500 * time.Millisecond)
			defer t.Stop()

			download_progress(resp, t)

			// check for errors
			if err := resp.Err(); err != nil {
				utility.ErrorLog.Fatalf("Download failed: %v\n", err)
			}
		}(url)
	}
	wg.Wait()
}

// download files[video/gif/image](files with no sound)
func downloader_nos(url, title string) {

	// create client
	client := grab.NewClient()

	req, _ := grab.NewRequest("", url)

	resp := client.Do(req)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

	download_progress(resp, t)

	// check for errors
	if err := resp.Err(); err != nil {
		utility.ErrorLog.Fatalf("Download failed: %v\n", err)
	}

	if strings.Contains(resp.Filename, ".mp4") {
		file_name := fmt.Sprintf("%s.mp4", title)
		err := os.Rename(resp.Filename, file_name)
		if err != nil {
			utility.ErrorLog.Println(err)
		}

		utility.InfoLog.Printf("Download saved to %v \n", file_name)

	} else {

		utility.InfoLog.Printf("Download saved to %v \n", resp.Filename)
	}

}
