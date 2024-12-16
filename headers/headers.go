package headers

import (
	http "github.com/bogdanfinn/fhttp"
	"github.com/bogdanfinn/tls-client/profiles"
	"math/rand"
)

type ProfileAndHeaders struct {
	Profile       profiles.ClientProfile
	CreateHeaders func() http.Header
}

var profileAndHeadersList = []ProfileAndHeaders{
	{Profile: profiles.Chrome_131, CreateHeaders: chrome131},
	{Profile: profiles.Chrome_124, CreateHeaders: chrome124},
	{Profile: profiles.Chrome_120, CreateHeaders: chrome120},
}

func GetRandomProfile() ProfileAndHeaders {
	randomIndex := rand.Intn(len(profileAndHeadersList))
	return profileAndHeadersList[randomIndex]
}

func chrome131() http.Header {
	secChUaList := []string{
		`"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`,
		`"Brave";v="131", "Chromium";v="131", "Not_A Brand";v="24"`,
		`"Opera";v="116", "Chromium";v="131", "Not_A Brand";v="24"`,
	}
	randomIndex := rand.Intn(len(secChUaList))

	return http.Header{
		"sec-ch-ua":                 {secChUaList[randomIndex]},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {"\"Linux\""},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"},
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"sec-fetch-site":            {"cross-site"},
		"sec-fetch-mode":            {"navigate"},
		"sec-fetch-user":            {"?1"},
		"sec-fetch-dest":            {"document"},
		"referer":                   {"https://www.google.com/"},
		"accept-encoding":           {"gzip, deflate, br, zstd"},
		"accept-language":           {"en-US,en;q=0.9"},
		"priority":                  {"u=0, i"},
		http.HeaderOrderKey: {
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"upgrade-insecure-requests",
			"user-agent",
			"accept",
			"sec-fetch-site",
			"sec-fetch-mode",
			"sec-fetch-user",
			"sec-fetch-dest",
			"referer",
			"accept-encoding",
			"accept-language",
			"priority",
		},
	}
}

func chrome124() http.Header {
	secChUaList := []string{
		`"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"`,
		`"Chromium";v="124", "Brave";v="124", "Not-A.Brand";v="99"`,
		`"Chromium";v="124", "Opera";v="110", "Not-A.Brand";v="99"`,
	}
	randomIndex := rand.Intn(len(secChUaList))

	return http.Header{
		"sec-ch-ua":                 {secChUaList[randomIndex]},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {"\"Linux\""},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"},
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"sec-fetch-site":            {"cross-site"},
		"sec-fetch-mode":            {"navigate"},
		"sec-fetch-user":            {"?1"},
		"sec-fetch-dest":            {"document"},
		"referer":                   {"https://www.google.com/"},
		"accept-encoding":           {"gzip, deflate, br, zstd"},
		"accept-language":           {"en-US,en;q=0.9"},
		"priority":                  {"u=0, i"},
		http.HeaderOrderKey: {
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"upgrade-insecure-requests",
			"user-agent",
			"accept",
			"sec-fetch-site",
			"sec-fetch-mode",
			"sec-fetch-user",
			"sec-fetch-dest",
			"referer",
			"accept-encoding",
			"accept-language",
			"priority",
		},
	}
}

func chrome120() http.Header {
	secChUaList := []string{
		`"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`,
		`"Not_A Brand";v="8", "Chromium";v="120", "Brave";v="120"`,
		`"Not_A Brand";v="8", "Chromium";v="120", "Opera";v="106"`,
	}
	randomIndex := rand.Intn(len(secChUaList))

	return http.Header{
		"sec-ch-ua":                 {secChUaList[randomIndex]},
		"sec-ch-ua-mobile":          {"?0"},
		"sec-ch-ua-platform":        {"\"Linux\""},
		"upgrade-insecure-requests": {"1"},
		"user-agent":                {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"},
		"accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"sec-fetch-site":            {"cross-site"},
		"sec-fetch-mode":            {"navigate"},
		"sec-fetch-user":            {"?1"},
		"sec-fetch-dest":            {"document"},
		"referer":                   {"https://www.google.com/"},
		"accept-encoding":           {"gzip, deflate, br, zstd"},
		"accept-language":           {"en-US,en;q=0.9"},
		"priority":                  {"u=0, i"},
		http.HeaderOrderKey: {
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"upgrade-insecure-requests",
			"user-agent",
			"accept",
			"sec-fetch-site",
			"sec-fetch-mode",
			"sec-fetch-user",
			"sec-fetch-dest",
			"referer",
			"accept-encoding",
			"accept-language",
			"priority",
		},
	}
}
