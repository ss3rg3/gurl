package core

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"fmt"
	"github.com/andybalholm/brotli"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/klauspost/compress/zstd"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"strings"
)

func ExecuteHttpRequest(httpUrl string) {

	profileAndHeaders := GetRandom()

	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(15),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithClientProfile(profileAndHeaders.Profile),
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		ExitWithError(fmt.Sprintf("Error creating HTTP client: %s\n", err))
	}

	req, err := http.NewRequest(http.MethodGet, httpUrl, nil)
	if err != nil {
		ExitWithError(fmt.Sprintf("Error creating HTTP request: %s\n", err))
	}
	req.Header = profileAndHeaders.Headers()

	resp, err := client.Do(req)
	if err != nil {
		ExitWithError(fmt.Sprintf("Error executing HTTP request: %s\n", err))
	}
	defer resp.Body.Close()

	content, err := getDecompressedContent(resp.Body, resp.Header.Get("Content-Encoding"), resp.Header.Get("Content-Type"))
	if err != nil {
		ExitWithError(fmt.Sprintf("Error creating decompressed reader: %s\n", err))
	}

	// todo return JSON
	fmt.Println(content)
	fmt.Println(req.URL)
	fmt.Println(resp.Location())
	fmt.Println("Content-Encoding:", resp.Header.Get("Content-Encoding"))
	fmt.Println(fmt.Sprintf("status code: %d", resp.StatusCode))
	fmt.Println(profileAndHeaders.Profile.GetClientHelloId())

}

func getDecompressedContent(body io.ReadCloser, contentEncoding, contentType string) (string, error) {
	// Read all bytes first since we might need to try multiple decompression methods
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return "", fmt.Errorf("reading body: %w", err)
	}

	// First handle compression
	contentEncoding = strings.TrimSpace(contentEncoding)
	var decompressed []byte

	switch contentEncoding {
	case "gzip":
		reader, err := gzip.NewReader(bytes.NewReader(bodyBytes))
		if err != nil {
			decompressed = bodyBytes // If gzip fails, use original content
		} else {
			defer reader.Close()
			decompressed, err = io.ReadAll(reader)
			if err != nil {
				decompressed = bodyBytes
			}
		}

	case "deflate":
		reader := flate.NewReader(bytes.NewReader(bodyBytes))
		defer reader.Close()
		content, err := io.ReadAll(reader)
		if err != nil {
			decompressed = bodyBytes
		} else {
			decompressed = content
		}

	case "br":
		content, err := io.ReadAll(brotli.NewReader(bytes.NewReader(bodyBytes)))
		if err != nil {
			decompressed = bodyBytes
		} else {
			decompressed = content
		}

	case "zstd":
		reader, err := zstd.NewReader(bytes.NewReader(bodyBytes))
		if err != nil {
			decompressed = bodyBytes
		} else {
			defer reader.Close()
			content, err := io.ReadAll(reader)
			if err != nil {
				decompressed = bodyBytes
			} else {
				decompressed = content
			}
		}

	default:
		decompressed = bodyBytes
	}

	// Checks the Content-Type header first
	// If no charset is specified in the header, looks for a meta charset tag in the first 1024 bytes
	// If still not found, attempts to detect the encoding from the content
	// Falls back to UTF-8 if it can't determine the encoding
	e, name, _ := charset.DetermineEncoding(decompressed, contentType)
	if name != "utf-8" {
		decoded, err := io.ReadAll(transform.NewReader(bytes.NewReader(decompressed), e.NewDecoder()))
		if err != nil {
			// If conversion fails, return original
			return string(decompressed), nil
		}
		return string(decoded), nil
	}

	return string(decompressed), nil
}
