package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/noornee/reddit-dl/utility"
)

func check(err error) {
	if err != nil {
		utility.ErrorLog.Println(err)
	}
}

// Parses the url and appends .json to it
func ParseUrl(raw string) (url, title string) {

	utility.InfoLog.Println("Parsing The URL")

	re, _ := regexp.Compile("www")

	// i.e. 'https://www.reddit.com/...' -> 'https://old.reddit.com/...'
	subdomain := re.ReplaceAllString(raw, "old")

	split_url := strings.Split(subdomain, "?")[0]

	// remove any trailing slashes
	trim_url := strings.TrimSuffix(split_url, "/")

	url = fmt.Sprintf("%s.json", trim_url)

	// get the title from the url path
	// i.e. https://old.reddit.com/r/<sub_reddit>/comments/blahblah/`<title>`
	url_path := strings.Split(trim_url, "/")
	title = url_path[len(url_path)-1]

	return url, title
}

// Get the url response body
func GetBody(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	check(err)

	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	utility.InfoLog.Println("Getting the JSON body")

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return body, nil

}

func GetHead(url string) (status_code int, content_type string) {

	req, err := http.NewRequest("HEAD", url, nil)
	check(err)

	resp, err := http.DefaultClient.Do(req)
	check(err)

	defer resp.Body.Close()

	status_code = resp.StatusCode
	content_type = resp.Header.Get("Content-Type")

	return status_code, content_type

}
