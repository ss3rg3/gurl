package headers

import (
	"errors"
	http "github.com/bogdanfinn/fhttp"
)

type Color string

const (
	Red   Color = "red"
	Blue  Color = "blue"
	Green Color = "green"
)

func (c Color) AsString() string {
	return string(c) // or any custom formatting you want
}

var browserHeaders = map[Color]http.Header{
	Red: Chrome132(),
}

func Chrome132() http.Header {
	return http.Header{
		"accept":          {"*/*"},
		"accept-language": {"de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7"},
		"user-agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"user-agent",
		},
	}
}

func HeaderOf(name Color) (http.Header, error) {
	headers, exists := browserHeaders[name]
	if !exists {
		return nil, errors.New("Failed to find headers for browser " + name.AsString())
	}
	return headers, nil
}
