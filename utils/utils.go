package utils

import (
	"fmt"
	"net/url"
	"os"
)

// IsValidHTTPURL checks if a given string is a valid HTTP/HTTPS URL
func IsValidHTTPURL(candidate string) bool {
	u, err := url.Parse(candidate)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

func ExitWithError(message string) {
	fmt.Printf(message)
	os.Exit(1)
}
