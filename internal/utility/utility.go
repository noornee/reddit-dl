package utility

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/noornee/reddit-dl/internal/model"
)

type RedditData struct {
	IsRedditGallery bool
	GalleryUrls     []string // reddit gallery (multiple photos in a post)
	MediaUrl        string   // media url (image|gif|video)
	IsDash          bool
}

// parses the json body and returns the parsed url(s) and an error
func ParseJSONBody(body []byte, useDash bool) (RedditData, error) {
	var (
		dataDump   model.Reddit // data dump
		redditData RedditData
	)

	err := json.Unmarshal(body, &dataDump)
	if err != nil {
		return redditData, err
	}

	for _, v := range dataDump {

		data := v.Data.Children[0].Data

		// download video using DASHPLAYLIST
		if useDash == true && data.SecureMedia.RedditVideo != nil {
			redditData.MediaUrl = data.SecureMedia.RedditVideo.DASH
			redditData.IsDash = true
			return redditData, nil
		}

		// this is for multiple pictures posted --> reddit gallery
		if data.MediaMetadata != nil {
			for _, v := range data.MediaMetadata {
				image_url := v.S.URL
				url := strings.ReplaceAll(image_url, "amp;", "")
				redditData.IsRedditGallery = true
				redditData.GalleryUrls = append(redditData.GalleryUrls, url)
			}
			return redditData, nil
		}

		// for crosspost videos
		if data.CrossPost != nil {
			if data.CrossPost[0].SecureMedia != nil && data.CrossPost[0].SecureMedia.RedditVideo != nil {
				redditData.MediaUrl = data.CrossPost[0].SecureMedia.RedditVideo.FallbackURL
				return redditData, nil
			}
		}

		// for external providers --> specifically gfycat.com
		if data.SecureMedia.Oembed != nil {

			gfycat := "https://gfycat.com"
			provider_url := data.SecureMedia.Oembed.ProviderURL

			switch provider_url {
			case gfycat:
				url := strings.ReplaceAll(data.SecureMedia.Oembed.ThumbnailURL, "size_restricted.gif", "mobile.mp4")
				redditData.MediaUrl = url
				return redditData, nil
			default:
				ErrorLog.Printf("%s is not a supported provider, going to fallback options", provider_url)
			}
		}

		// MP4 variant of embedded video (likely NFSW, data.Preview.Images[0].Variants.GIF cannot be used here (even though there are two structs, obfuscated and nsfw, both are "obfuscated" (blurred.)))
		if data.Preview != nil {

			if data.Preview.Video != nil {
				redditData.MediaUrl = data.Preview.Video.FallbackURL
				return redditData, nil
			}

			// GIF variant of embedded video
			if data.Preview.Images[0].Variants.GIF != nil {
				redditData.MediaUrl = data.Preview.Images[0].Variants.GIF.Source.URL
				return redditData, nil
			}

		}

		// if securemedia is nil then it's a normal image/gif
		if data.SecureMedia.RedditVideo == nil && data.SecureMedia.Oembed == nil {
			redditData.MediaUrl = data.URLOverriddenByDest
			return redditData, nil
		}

		// for normal reddit video
		if data.SecureMedia.RedditVideo != nil {
			redditData.MediaUrl = data.SecureMedia.RedditVideo.FallbackURL
			return redditData, nil
		}
	}

	return redditData, nil
}

func GetMediaUrl(url string) (media, audio string) {
	url = strings.ReplaceAll(url, "amp;", "")

	// checks if its a gif
	if strings.Contains(url, ".gif") {
		media = url
		audio = ""
		return media, audio
	}

	// normal video
	media = strings.Split(url, "?")[0]
	re, _ := regexp.Compile("_[0-9]+")
	audio = re.ReplaceAllString(media, "_audio")

	// this is for external videos/gif i.e. from gfycat
	// it wouldnt match the regex pattern
	if media == audio {
		return media, ""
	}

	return media, audio
}
