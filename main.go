package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/noornee/redown/external"
	"github.com/noornee/redown/utility"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var raw_url string

	flag.StringVar(&raw_url, "url", "", "url to parse")
	flag.Parse()

	passed := utility.IsFlagPassed(raw_url)
	if !passed {
		flag.Usage()
		return
	}

	url := parseUrl(raw_url)

	//START

	body := getData(url)

	//END

	//file, _ := os.ReadFile(url)
	fallback_url, title := utility.ParseJSONfile(body)

	video, audio := utility.GetMediaUrl(fallback_url)

	external.CMD_aria2c(video, audio)
	external.CMD_ffmpeg(title)

}

// Parses a url
func parseUrl(raw string) (url string) {

	re, _ := regexp.Compile("www")

	//replace www with old
	subdomain := re.ReplaceAllString(raw, "old")

	split_url := strings.Split(subdomain, "?")[0]

	// remove any trailing slashes
	trim_url := strings.TrimSuffix(split_url, "/")

	url = fmt.Sprintf("%s.json", trim_url)

	return url
}

func getData(url string) []byte {

	req, err := http.NewRequest("GET", url, nil)
	check(err)

	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return body

}
