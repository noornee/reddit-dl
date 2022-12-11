package utility

import (
	"encoding/json"
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

// returns the fallback_url
func ParseJSONfile(file []byte) string {

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

	secure_media, ok := data2["secure_media"].(map[string]interface{}) // handle cross_post here

	// if it doesn have the underlying interface `ok` would be false
	// i.e its a cross_post so this handles it
	if ok != true {
		cross_post := data2["crosspost_parent_list"].([]interface{})
		data3 := cross_post[0].(map[string]interface{})
		secure_media := data3["secure_media"].(map[string]interface{})
		reddit_video := secure_media["reddit_video"].(map[string]interface{})
		fallback_url := reddit_video["fallback_url"]
		return fmt.Sprint(fallback_url)
	}

	reddit_video := secure_media["reddit_video"].(map[string]interface{})
	fallback_url := reddit_video["fallback_url"]

	return fmt.Sprint(fallback_url)

}

func GetMediaUrl(url string) (video, audio string) {
	video = strings.Split(url, "?")[0]
	re, _ := regexp.Compile("_[0-9]+")
	audio = re.ReplaceAllString(video, "_audio")

	return video, audio
}
