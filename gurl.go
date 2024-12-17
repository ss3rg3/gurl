package main

import (
	"flag"
	"fmt"
	"github.com/ss3rg3/gurl/core"
	"os"
)

type Options struct {
	url string
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Println("Go URL uses tls-client to GET an URL and prints its as contents as JSON to be read by another program.\n" +
			"Use https://tls.peet.ws/api/all for debugging.")
		flag.PrintDefaults()
	}

	opts := Options{}
	flag.StringVar(&opts.url, "url", "", "The URL you want to retrieve (required)")

	flag.Parse()
	validate(&opts)
	core.ExecuteHttpRequest(opts.url)
}

func validate(opts *Options) {
	if opts.url == "" {
		flag.Usage()
		core.ExitWithError("Error: URL is required")
	}

	if !core.IsValidHTTPURL(opts.url) {
		core.ExitWithError(fmt.Sprintf("Error: %s is not a valid HTTP/HTTPS URL\n", opts.url))
	}
}
