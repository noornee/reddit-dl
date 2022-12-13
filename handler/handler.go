package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Parses the url and appends .json to it
func ParseUrl(raw string) (url string) {

	re, _ := regexp.Compile("www")

	// i.e. 'https://www.reddit.com/...' -> 'https://old.reddit.com/...'
	subdomain := re.ReplaceAllString(raw, "old")

	split_url := strings.Split(subdomain, "?")[0]

	// remove any trailing slashes
	trim_url := strings.TrimSuffix(split_url, "/")

	url = fmt.Sprintf("%s.json", trim_url)

	return url
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

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return body, nil

}

func GetStatusCode(url string) int {

	req, err := http.NewRequest("HEAD", url, nil)
	check(err)

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	return resp.StatusCode

}
