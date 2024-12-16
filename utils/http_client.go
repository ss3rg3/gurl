package utils

import (
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"io"
	"log"
)

func ExecuteHttpRequest(httpUrl string) {
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),                 // todo set all timeouts properly
		tls_client.WithClientProfile(profiles.Chrome_131), // todo use from file
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		ExitWithError(fmt.Sprintf("Error creating HTTP client: %s\n", err))
	}

	req, err := http.NewRequest(http.MethodGet, "https://tls.peet.ws/api/all", nil)
	if err != nil {
		ExitWithError(fmt.Sprintf("Error creating HTTP request: %s\n", err))
	}

	// todo use proper headers
	req.Header = http.Header{
		"accept":          {"*/*"},
		"accept-language": {"de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7"},
		"user-agent":      {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"accept",
			"accept-language",
			"user-agent",
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		ExitWithError(fmt.Sprintf("Error executing HTTP request: %s\n", err))
	}
	defer resp.Body.Close()

	readBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err) // todo return with error
		return
	}

	// todo return JSON
	log.Println(string(readBytes))
	log.Println(fmt.Sprintf("status code: %d", resp.StatusCode))

}
