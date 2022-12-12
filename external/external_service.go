package external

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/noornee/reddit-dl/handler"
	"github.com/noornee/reddit-dl/utility"
)

var temp_dir string = utility.CreateDir()

func Setup(video, audio, title string) {

	if audio != "" && handler.GetStatusCode(audio) == 200 {
		aria2c(video, audio)
		ffmpeg(title)
		return
	}

	// if the audio status code is not 200, then its most likely invalid or it doesnt exist
	aria2c_nos(video, title)

}

// download files[video,audio] with aria2c
func aria2c(video, audio string) {

	cmd := exec.Command("aria2c", "-d", temp_dir, "-Z", video, audio)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

}

// download files[video/gif] with aria2c
func aria2c_nos(video, title string) {
	var cmd *exec.Cmd

	switch {
	case strings.HasSuffix(video, ".gif"):
		cmd = exec.Command("aria2c", video, "-o", title+".gif")
	default:
		cmd = exec.Command("aria2c", video, "-o", title+".mp4")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

}

// merge downoladed files[video,audio] with ffmpeg
func ffmpeg(filename string) {

	filename = filename + ".mp4"

	files, err := ioutil.ReadDir(temp_dir)
	if err != nil {
		log.Println(err)
	}

	var aud, vid string

	for range files {
		vid = temp_dir + "/" + files[0].Name()
		aud = temp_dir + "/" + files[1].Name()
	}

	cmd := exec.Command("ffmpeg", "-y", "-v", "quiet", "-stats", "-i", vid, "-i", aud, filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("merging files together")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

}
