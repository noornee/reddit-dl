package external

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func createDir() string {
	dir, err := ioutil.TempDir("", "reddit")
	if err != nil {
		log.Fatal(err)
	}
	//defer os.RemoveAll(dir)

	return dir
}

// download files with aria2c
func CMD_aria2c(video, audio string) {

	cmd := exec.Command("aria2c", "-d", createDir(), "-Z", video, audio)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

}

// merge downoladed files with ffmpeg
func CMD_ffmpeg() {
	cwd, _ := os.Getwd()
	dir := cwd + createDir()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}

	var aud, vid string

	for range files {
		vid = dir + "/" + files[0].Name()
		aud = dir + "/" + files[1].Name()
	}

	cmd := exec.Command("ffmpeg", "-y", "-v", "quiet", "-stats", "-i", vid, "-i", aud, "new.mp4")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("merging files together")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

}
