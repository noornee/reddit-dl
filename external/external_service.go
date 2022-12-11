package external

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/noornee/redown/utility"
)

var temp_dir string = utility.CreateDir()

// download files with aria2c
func CMD_aria2c(video, audio string) {

	cmd := exec.Command("aria2c", "-d", temp_dir, "-Z", video, audio)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

}

// merge downoladed files with ffmpeg
func CMD_ffmpeg() {

	files, err := ioutil.ReadDir(temp_dir)
	if err != nil {
		log.Println(err)
	}

	var aud, vid string

	for range files {
		vid = temp_dir + "/" + files[0].Name()
		aud = temp_dir + "/" + files[1].Name()
	}

	cmd := exec.Command("ffmpeg", "-y", "-v", "quiet", "-stats", "-i", vid, "-i", aud, "output.mp4")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("merging files together")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

}
