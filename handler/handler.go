package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/noornee/reddit-dl/utility"
)

// Parses the url and appends .json to it
func ParseUrl(raw string) (url, title string, err error) {

	utility.InfoLog.Println("Parsing The URL")

	/*
		this supports:
		https://www.reddit.com/r/...
		https://reddit.com/r/...
		http://www.reddit.com/r/...
		http://reddit.com/r/...
		www.reddit.com/r/...
		reddit.com/r/...
	*/
	if !(strings.Contains(raw, "http://") || strings.Contains(raw, "https://")) {
		// No protocol was provided, append a https:// to the start of the url.
		raw = "https://" + raw
	}

	// drop the 1st occourance of www.
	raw = strings.Replace(raw, "www.", "", 1)
	// and replace the 1st occourance of reddit.com with old.reddit.com
	raw = strings.Replace(raw, "reddit.com", "old.reddit.com", 1)

	// if it fails after those, the url that was sent is likely nonsense.

	split_url := strings.Split(raw, "?")[0]

	// remove any trailing slashes
	trim_url := strings.TrimSuffix(split_url, "/")

	url = fmt.Sprintf("%s.json", trim_url)

	// get the title from the url path
	// i.e. https://old.reddit.com/r/<sub_reddit>/comments/blahblah/`<title>`
	url_path := strings.Split(trim_url, "/")
	title = url_path[len(url_path)-1]

	return url, title, nil
}

// Get the url response body
func GetBody(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	utility.InfoLog.Println("Getting the JSON body")

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}

func GetHead(url string) (status_code int, content_type string) {

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, ""
	}

	defer resp.Body.Close()

	status_code = resp.StatusCode
	content_type = resp.Header.Get("Content-Type")

	return status_code, content_type

}
