package utility

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func CreateDir() string {
	dir, err := ioutil.TempDir("", "reddit")
	if err != nil {
		ErrorLog.Fatal(err)
	}
	defer os.RemoveAll(dir)

	return dir
}

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

	var urls []string

	var dataDump any

	//json.Unmarshal(file, &dataDump)
	err := json.NewDecoder(bytes.NewReader(file)).Decode(&dataDump)
	if err != nil {
		return urls, err
	}

	// ---------------------------------------------------------------------------------------------------- //
	// traversing through it all to get the fallback_url
	root, ok := dataDump.([]interface{})
	if ok != true {
		return urls, errors.New("cannot parse body")
	}

	edge := root[0].(map[string]any)
	data := edge["data"].(map[string]any)
	children := data["children"].([]any)
	data1 := children[0].(map[string]any)
	data2 := data1["data"].(map[string]any)

	// ------------------------------------GALLERY---------------------------------------------------------- //
	// this is for multiple pictures that are posted --> reddit gallery
	metadata, ok := data2["media_metadata"].(map[string]any)
	if ok {
		for i := range metadata {
			media_id := metadata[i].(map[string]any)
			media_s := media_id["s"].(map[string]any)
			media_url := media_s["u"]
			new_media_url := strings.ReplaceAll(fmt.Sprint(media_url), "amp;", "")
			urls = append(urls, fmt.Sprint(new_media_url))
		}
		return urls, nil
	}

	// ---------------------------------------------------------------------------------------------------- //

	secure_media, ok := data2["secure_media"].(map[string]any)

	if secure_media == nil {
		// for normal image/gif
		url_overridden_by_dest := data2["url_overridden_by_dest"]
		urls = append(urls, fmt.Sprint(url_overridden_by_dest))
		return urls, nil
	}

	// ----------------------------------------CROSSPOST--------------------------------------------------- //
	// if it doesn't have the underlying interface `ok` would be false then its a crosspost
	// for reddit cross post video
	if ok != true {
		cross_post := data2["crosspost_parent_list"].([]any)
		data3 := cross_post[0].(map[string]any)
		secure_media := data3["secure_media"].(map[string]any)
		reddit_video := secure_media["reddit_video"].(map[string]any)
		fallback_url := reddit_video["fallback_url"]
		urls = append(urls, fmt.Sprint(fallback_url))
		return urls, nil

	}

	// --------------------------------FOR GIFS/VIDEO HOSTED ON GFYCAT.COM----------------------------------- //
	oembed, ok := secure_media["oembed"].(map[string]any)
	if ok {
		provider_url := oembed["provider_url"]
		thumbnail_url := oembed["thumbnail_url"]
		if provider_url == "https://gfycat.com" {
			new_url := strings.ReplaceAll(fmt.Sprint(thumbnail_url), "size_restricted.gif", "mobile.mp4")
			urls = append(urls, fmt.Sprint(new_url))
			return urls, nil

		} else {
			return urls, errors.New(fmt.Sprintf("unsupported provider \"%s\"", provider_url))
		}
	}

	// --------------------------------NORMAL REDDIT VIDEO------------------------------------------------- //
	reddit_video := secure_media["reddit_video"].(map[string]any)
	fallback_url := reddit_video["fallback_url"]
	urls = append(urls, fmt.Sprint(fallback_url))
	return urls, nil

	// ---------------------------------------------------------------------------------------------------- //
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

	return media, audio
}
