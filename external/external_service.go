package external

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/noornee/reddit-dl/handler"
	"github.com/noornee/reddit-dl/utility"
)

var temp_dir string = utility.CreateDir()

func Setup(video, audio, title string) {

	if audio != "" {
		status_code, mime := handler.GetHead(audio)

		if status_code == 200 && !strings.Contains(mime, "image") {
			aria2c(video, audio)
			ffmpeg(title)
			return
		}

	}

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

// download files[video/gif/image] with aria2c
func aria2c_nos(file, title string) {

	var cmd *exec.Cmd

	_, mime_type := handler.GetHead(file)

	switch mime_type {
	case "image/jpeg":
		cmd = exec.Command("aria2c", file, "-o", title+".jpg")
	case "image/png":
		cmd = exec.Command("aria2c", file, "-o", title+".png")
	case "image/gif":
		cmd = exec.Command("aria2c", file, "-o", title+".gif")
	default:
		cmd = exec.Command("aria2c", file, "-o", title+".mp4")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		utility.ErrorLog.Println(err)
	}

}

// merge downoladed files[video,audio] with ffmpeg
func ffmpeg(filename string) {

	filename = filename + ".mp4"

	files, err := ioutil.ReadDir(temp_dir)
	if err != nil {
		utility.ErrorLog.Println(err)
	}

	var aud, vid string

	for range files {
		vid = temp_dir + "/" + files[0].Name()
		aud = temp_dir + "/" + files[1].Name()
	}

	cmd := exec.Command("ffmpeg", "-y", "-v", "quiet", "-stats", "-i", vid, "-i", aud, filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	utility.InfoLog.Printf("Merging files into \t%s", filename)

	if err := cmd.Run(); err != nil {
		utility.ErrorLog.Println(err)
	}
	utility.InfoLog.Println("Done")

}
