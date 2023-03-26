package service

import (
	"os"

	"github.com/noornee/reddit-dl/internal/utility"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// merge downoladed files[video,audio] with ffmpeg
func video_merger(filename string) {

	const temp_dir = ".reddit_temp/"

	filename = filename + ".mp4"

	// use os.ReadDir instead of ioutil, as it's deprecated since go 1.16
	files, err := os.ReadDir(temp_dir)

	if err != nil {
		utility.ErrorLog.Println(err)
	}

	var aud, vid string

	for range files {
		vid = temp_dir + files[0].Name()
		aud = temp_dir + files[1].Name()

	}

	in1 := ffmpeg.Input(vid)
	in2 := ffmpeg.Input(aud)

	utility.InfoLog.Printf("Merging files into \t%s\r\n", filename)

	err = ffmpeg.Concat([]*ffmpeg.Stream{in1, in2}, ffmpeg.KwArgs{"v": 1, "a": 1}).
		Output(filename, ffmpeg.KwArgs{"v": "quiet"}).
		GlobalArgs("-stats").
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		utility.ErrorLog.Println(err)
	}

	utility.InfoLog.Println("Done")

	// remove temp directory that was created
	err = os.RemoveAll(temp_dir)
	if err != nil {
		utility.ErrorLog.Println(err)
	}

}
