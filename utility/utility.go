package utility

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func CreateDir() string {
	dir, err := ioutil.TempDir("", "reddit")
	if err != nil {
		log.Fatal(err)
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

// returns the fallback_url and title
func ParseJSONBody(file []byte) (string, string, error) {

	var dataDump interface{}

	err := json.Unmarshal(file, &dataDump)
	if err != nil {
		log.Println(err)
	}

	// traversing through it all to get the fallback_url
	root := dataDump.([]interface{})
	edge := root[0].(map[string]interface{})
	data := edge["data"].(map[string]interface{})
	children := data["children"].([]interface{})
	data1 := children[0].(map[string]interface{})
	data2 := data1["data"].(map[string]interface{})
	title := data2["title"]
	is_reddit_media_domain := data2["is_reddit_media_domain"]

	// if the video isnt hosted on reddit
	if is_reddit_media_domain == false {
		return "", "", errors.New("Cannot download video")
	}

	secure_media, ok := data2["secure_media"].(map[string]interface{}) // handle cross_post here

	// for gifs
	if secure_media == nil {
		url_overridden_by_dest := data2["url_overridden_by_dest"]
		return fmt.Sprint(url_overridden_by_dest), fmt.Sprint(title), nil
	}

	// if it doesn have the underlying interface `ok` would be false
	// i.e its a cross_post so this handles it
	if ok != true {
		cross_post := data2["crosspost_parent_list"].([]interface{})
		data3 := cross_post[0].(map[string]interface{})
		secure_media := data3["secure_media"].(map[string]interface{})
		reddit_video := secure_media["reddit_video"].(map[string]interface{})
		fallback_url := reddit_video["fallback_url"]
		return fmt.Sprint(fallback_url), fmt.Sprint(title), nil
	}

	reddit_video := secure_media["reddit_video"].(map[string]interface{})
	fallback_url := reddit_video["fallback_url"]

	return fmt.Sprint(fallback_url), fmt.Sprint(title), nil

}


func GetMediaUrl(url string) (video, audio string) {

	// checks if its a gif
	if strings.HasSuffix(url, ".gif") {
		video = url
		audio = ""
		return video, audio
	}

	// normal video
	video = strings.Split(url, "?")[0]
	re, _ := regexp.Compile("_[0-9]+")
	audio = re.ReplaceAllString(video, "_audio")

	return video, audio
}
