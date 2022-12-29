package utility

import (
	"encoding/json"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/noornee/reddit-dl/model"
)

// returns true if a valid flag was passed
func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Value.String() == name {
			found = true
		}
	})
	return found
}

// parses the json body and returns the parsed url(s) and an error
func ParseJSONBody(file []byte) ([]string, error) {

	var (
		dataDump     model.Reddit   // data dump
		metaDataDump map[string]any // metadata map for reddit gallery -> line 44
		urls         []string       // slice of urls
	)

	err := json.Unmarshal(file, &dataDump)
	if err != nil {
		return urls, err
	}

	for _, v := range dataDump {

		data := v.Data.Children[0].Data

		// this is for multiple pictures posted --> reddit gallery
		if data.MediaMetadata != nil {

			metadata := data.MediaMetadata
			err := json.Unmarshal(metadata, &metaDataDump)
			if err != nil {
				return urls, err
			}

			for i := range metaDataDump {
				media_id := metaDataDump[i].(map[string]any)
				media_s := media_id["s"].(map[string]any)
				media_url := media_s["u"]
				url := strings.ReplaceAll(media_url.(string), "amp;", "")
				urls = append(urls, url)

			}

			return urls, nil
		}

		// if securemedia is nil then it's a normal image/gif
		if data.SecureMedia.RedditVideo == nil && data.SecureMedia.Oembed == nil {

			url := data.URLOverriddenByDest

			fmt.Println(url)
			urls = append(urls, url)
			return urls, nil

		}

		// for crossposts videos
		if data.CrossPost != nil {

			url := data.CrossPost[0].SecureMedia.RedditVideo.FallbackURL
			urls = append(urls, url)
			return urls, nil

		}

		// for external providers --> specifically gfycat.com
		if data.SecureMedia.Oembed != nil {

			gfycat := "https://gfycat.com"
			provider_url := data.SecureMedia.Oembed.ProviderURL

			if provider_url == gfycat {
				url := strings.ReplaceAll(data.SecureMedia.Oembed.ThumbnailURL, "size_restricted.gif", "mobile.mp4")
				urls = append(urls, url)
				fmt.Println(urls)
				return urls, nil
			} else {
				return urls, fmt.Errorf("unsupported provider \"%s\"", provider_url)
			}

		}

		// for normal reddit video
		if data.SecureMedia.RedditVideo != nil {

			url := data.SecureMedia.RedditVideo.FallbackURL
			urls = append(urls, url)
			return urls, nil

		}

	}

	return urls, nil
}

func GetMediaUrl(url string) (media, audio string) {

	// checks if its a gif
	if strings.HasSuffix(url, ".gif") {
		media = url
		audio = ""
		return media, audio
	}

	// normal video
	media = strings.Split(url, "?")[0]
	re, _ := regexp.Compile("_[0-9]+")
	audio = re.ReplaceAllString(media, "_audio")

	if media == audio {
		return media, ""
	}

	return media, audio
}
