package service

import (
	"strings"

	"github.com/noornee/reddit-dl/internal/handler"
)

func Setup(media_url, audio_url, title string) {

	if audio_url != "" {
		status_code, mime := handler.GetHead(audio_url)

		if status_code == 200 && !strings.Contains(mime, "image") {
			downloader([]string{media_url, audio_url})
			video_merger(title)
			return
		}

	}

	downloader_nos(media_url, title)

}
