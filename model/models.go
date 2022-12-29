package model

import "encoding/json"

type Reddit []struct {
	Data struct {
		Children []struct {
			Data struct {
				MediaMetadata       json.RawMessage        `json:"media_metadata,omitempty"`
				SecureMedia         secureMedia            `json:"secure_media,omitempty"`
				CrossPost           []*crosspostParentList `json:"crosspost_parent_list,omitempty"`
				URLOverriddenByDest string                 `json:"url_overridden_by_dest"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type secureMedia struct {
	RedditVideo *redditVideo `json:"reddit_video,omitempty"`
	Oembed      *oembed      `json:"oembed,omitempty"`
}

type redditVideo struct {
	FallbackURL string `json:"fallback_url"`
	HLS         string `json:"hls_url"`
}

type oembed struct {
	ProviderURL  string `json:"provider_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type crosspostParentList struct {
	SecureMedia         *secureMedia `json:"secure_media"`
	URLOverriddenByDest string       `json:"url_overridden_by_dest"`
}
