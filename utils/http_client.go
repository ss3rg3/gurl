package utils

import (
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/ss3rg3/gurl/headers"
	"io"
	"log"
)

func ExecuteHttpRequest(httpUrl string) {

	profileAndHeaders := headers.GetRandomProfile()

	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(15),
		tls_client.WithClientProfile(profileAndHeaders.Profile),
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		ExitWithError(fmt.Sprintf("Error creating HTTP client: %s\n", err))
	}

	req, err := http.NewRequest(http.MethodGet, "https://tls.peet.ws/api/all", nil)
	if err != nil {
		ExitWithError(fmt.Sprintf("Error creating HTTP request: %s\n", err))
	}
	req.Header = profileAndHeaders.CreateHeaders()

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
	log.Println(profileAndHeaders.Profile.GetClientHelloId())

}
