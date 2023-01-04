package utility

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/noornee/reddit-dl/model"
)

// parses the json body and returns the parsed url(s) and an error
func ParseJSONBody(file []byte) ([]string, error) {

	var (
		dataDump model.Reddit // data dump
		urls     []string     // slice of urls
	)

	err := json.Unmarshal(file, &dataDump)
	if err != nil {
		return urls, err
	}

	for _, v := range dataDump {

		data := v.Data.Children[0].Data

		// this is for multiple pictures posted --> reddit gallery
		if data.MediaMetadata != nil {
			for _, v := range data.MediaMetadata {
				image_url := v.S.URL
				url := strings.ReplaceAll(image_url, "amp;", "")
				urls = append(urls, url)
			}
			return urls, nil
		}

		// for crossposts media
		if data.CrossPost != nil {

			// for cases where the crosspost media is nil
			if data.CrossPost[0].SecureMedia == nil {
				url := data.CrossPost[0].URLOverriddenByDest
				urls = append(urls, url)
				return urls, nil

			}

			// checks if its an embeded crosspost video
			if data.CrossPost[0].SecureMedia.Oembed != nil {

				gfycat := "https://gfycat.com"
				provider_url := data.CrossPost[0].SecureMedia.Oembed.ProviderURL

				switch provider_url {
				case gfycat:
					url := strings.ReplaceAll(data.CrossPost[0].SecureMedia.Oembed.ThumbnailURL, "size_restricted.gif", "mobile.mp4")
					urls = append(urls, url)
					fmt.Println(urls)
					return urls, nil
				default:
					return urls, fmt.Errorf("unsupported provider \"%s\"", provider_url)
				}
			}

			// normal crosspost video
			url := data.CrossPost[0].SecureMedia.RedditVideo.FallbackURL
			urls = append(urls, url)
			return urls, nil

		}

		// for external providers --> specifically gfycat.com
		if data.SecureMedia.Oembed != nil {

			gfycat := "https://gfycat.com"
			provider_url := data.SecureMedia.Oembed.ProviderURL

			switch provider_url {
			case gfycat:
				url := strings.ReplaceAll(data.SecureMedia.Oembed.ThumbnailURL, "size_restricted.gif", "mobile.mp4")
				urls = append(urls, url)
				fmt.Println(urls)
				return urls, nil
			default:
				ErrorLog.Printf("%s is not a supported provider, going to fallback options", provider_url)
			}
		}

		// MP4 variant of embedded video (likely NFSW, data.Preview.Images[0].Variants.GIF cannot be used here (even though there are two structs, obfuscated and nsfw, both are "obfuscated" (blurred.)))
		if data.Preview.Video != nil {
			url := data.Preview.Video.FallbackURL
			urls = append(urls, url)
			return urls, nil
		}

		// GIF variant of embedded video
		if data.Preview.Images[0].Variants.GIF != nil {
			url := data.Preview.Images[0].Variants.GIF.Source.URL
			urls = append(urls, url)
			return urls, nil
		}

		// if securemedia is nil then it's a normal image/gif
		if data.SecureMedia.RedditVideo == nil && data.SecureMedia.Oembed == nil {
			url := data.URLOverriddenByDest
			urls = append(urls, url)
			return urls, nil
		}

		// for normal reddit video
		if data.SecureMedia.RedditVideo != nil {
			url := data.SecureMedia.RedditVideo.FallbackURL
			urls = append(urls, url)
			return urls, nil
		}
	}

	return urls, nil
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
