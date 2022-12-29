package service

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/noornee/reddit-dl/utility"
)

// create a directory temp
func createDir() string {

	const temp_dir = ".reddit_temp"

	var err = os.Mkdir(temp_dir, os.ModePerm)
	if err != nil {
		utility.ErrorLog.Println(err)
	}

	return temp_dir
}

func downloader(urls []string) {

	var temp_dir string = createDir() + "/"

	// create client
	client := grab.NewClient()
	for _, url := range urls {

		req, _ := grab.NewRequest(temp_dir, url)

		//fmt.Printf("Downloading %v...\n", req.URL())

		resp := client.Do(req)

		// start UI loop
		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()

	Loop:
		for {
			select {
			case <-t.C:
				fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
					resp.BytesComplete(),
					resp.Size(),
					100*resp.Progress())

			case <-resp.Done:
				// download is complete
				break Loop
			}
		}

		// check for errors
		if err := resp.Err(); err != nil {
			utility.ErrorLog.Fatalf("Download failed: %v\n", err)
		}

	}

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

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

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
