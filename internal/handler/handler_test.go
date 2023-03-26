package handler

import (
	"testing"
)

type parseUrlTest struct {
	name           string
	raw_url        string
	expected_title string
	expected_url   string
	expected_error error
}

var parseUrlTestCases = []parseUrlTest{
	{
		name:           "first",
		raw_url:        "https://www.reddit.com/r/commandline/comments/zw09yk/created_a_cli_tool_redditdl_to_help_download/?utm_source=share&utm_medium=web2x&context=3",
		expected_url:   "https://old.reddit.com/r/commandline/comments/zw09yk/created_a_cli_tool_redditdl_to_help_download.json",
		expected_title: "created_a_cli_tool_redditdl_to_help_download",
		expected_error: nil,
	},
	{
		name:           "second",
		raw_url:        "https://www.reddit.com/r/AnimeFunny/comments/yxrgmp/the_greatest_estate_developer",
		expected_url:   "https://old.reddit.com/r/AnimeFunny/comments/yxrgmp/the_greatest_estate_developer.json",
		expected_title: "the_greatest_estate_developer",
		expected_error: nil,
	},
}

func TestParseUrl(t *testing.T) {
	for _, test := range parseUrlTestCases {

		url, title, err := ParseUrl(test.raw_url)

		t.Run(test.name, func(t *testing.T) {
			if err != test.expected_error {
				t.Errorf("expected %v got %q", test.expected_error, err)

			} else {

				if url != test.expected_url {
					t.Errorf("expected url %q got %q", test.expected_url, url)

				}

				if title != test.expected_title {
					t.Errorf("expected title %q got %q", test.expected_title, title)

				}

			}

		})

	}
}

func TestGetBody(t *testing.T) {
	for _, test := range parseUrlTestCases {
		body, _ := GetBody(test.expected_url)
		if len(body) < 1 {
			t.Errorf("expected bytes slice instead got %v", body)

		}

	}
}

func TestGetHead(t *testing.T) {
	code, mime := GetHead("https://v.redd.it/iymk1c6ptf9a1/DASH_audio.mp4")

	expected_code := 200
	expected_mime := "video/mp4"

	t.Run("status_code", func(t *testing.T) {
		if expected_code != code {
			t.Errorf("want %d, got %d", expected_code, code)
		}

	})

	t.Run("mime_type", func(t *testing.T) {
		if expected_mime != mime {
			t.Errorf("want %q, got %q", expected_mime, mime)
		}

	})

}
