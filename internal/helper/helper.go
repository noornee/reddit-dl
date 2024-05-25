package helper

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// GetJSONBody fetches the JSON body from a secondary URL constructed from the response of the initial request.
func GetJSONBody(url string) ([]byte, error) {
	// Helper function to perform HTTP GET request and return response.
	doRequest := func(url string) (*http.Response, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("creating request: %w", err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0")
		return http.DefaultClient.Do(req)
	}

	// Perform the initial request.
	resp, err := doRequest(url)
	if err != nil {
		return nil, fmt.Errorf("initial request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code for initial request: %s", resp.Status)
	}

	// Construct the secondary URL from the response.
	location, err := resp.Request.Response.Location()
	if err != nil {
		return nil, fmt.Errorf("getting response location: %w", err)
	}
	newURL := fmt.Sprintf("https://www.reddit.com%s.json", location.Path)

	// Perform the secondary request.
	newResp, err := doRequest(newURL)
	if err != nil {
		return nil, fmt.Errorf("secondary request: %w", err)
	}
	defer newResp.Body.Close()

	if newResp.StatusCode != http.StatusOK {
		fmt.Println(newResp)
		return nil, fmt.Errorf("unexpected status code for secondary request: %s", newResp.Status)
	}

	// Read and return the JSON body.
	body, err := io.ReadAll(newResp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading JSON body: %w", err)
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

func GetMediaUrl(url string) (media, audio string) {
	url = strings.ReplaceAll(url, "amp;", "")

	// checks if its a gif
	if strings.Contains(url, ".gif") {
		media = url
		audio = ""
		return media, audio
	}

	// normal video
	media = strings.Split(url, "?")[0]
	re, _ := regexp.Compile("_[0-9]+")
	audio = re.ReplaceAllString(media, "_audio")

	// this is for external videos/gif i.e. from gfycat
	// it wouldnt match the regex pattern
	if media == audio {
		return media, ""
	}

	return media, audio
}
