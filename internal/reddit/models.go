package reddit

type Reddit []struct {
	Data struct {
		Children []struct {
			Data struct {
				MediaMetadata       mediaMetadata          `json:"media_metadata,omitempty"`
				SecureMedia         secureMedia            `json:"secure_media,omitempty"`
				CrossPost           []*crosspostParentList `json:"crosspost_parent_list,omitempty"`
				Preview             *preview               `json:"preview,omitempty"`
				URLOverriddenByDest string                 `json:"url_overridden_by_dest"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type secureMedia struct {
	RedditVideo *redditVideo `json:"reddit_video,omitempty"`
	Oembed      *oembed      `json:"oembed,omitempty"`
}

type preview struct {
	Images []struct {
		Variants struct {
			GIF *struct {
				Source struct {
					URL string `json:"url"`
				} `json:"source"`
			} `json:"gif"`
		} `json:"variants"`
	} `json:"images"`

	Video *struct {
		FallbackURL string `json:"fallback_url"`
	} `json:"reddit_video_preview"`
}
type redditVideo struct {
	FallbackURL string `json:"fallback_url"`
	HLS         string `json:"hls_url"`
	DASH        string `json:"dash_url"`
}

type oembed struct {
	ProviderURL  string `json:"provider_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type crosspostParentList struct {
	SecureMedia         *secureMedia `json:"secure_media"`
	URLOverriddenByDest string       `json:"url_overridden_by_dest"`
}

type mediaMetadata map[string]struct {
	S struct {
		URL string `json:"u"`
	} `json:"s"`
}

type RedditData struct {
	IsRedditGallery bool
	GalleryUrls     []string // reddit gallery (multiple photos in a post)
	MediaUrl        string   // media url (image|gif|video)
	IsDash          bool
}
