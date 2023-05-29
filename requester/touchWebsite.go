package requester

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"time"
)

var header = map[string][]string{
	"accept":                      {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
	"accept-language":             {"en-US,en;q=0.9"},
	"cache-control":               {"max-age=0"},
	"dnt":                         {"1"},
	"sec-ch-ua":                   {"\"Not:A-Brand\";v=\"99\", \"Chromium\";v=\"112\""},
	"sec-ch-ua-arch":              {"\"arm\""},
	"sec-ch-ua-bitness":           {"\"64\""},
	"sec-ch-ua-full-version":      {"\"112.0.5615.121\""},
	"sec-ch-ua-full-version-list": {"\"Not:A-Brand\";v=\"99.0.0.0\", \"Chromium\";v=\"112.0.5615.121\""},
	"sec-ch-ua-mobile":            {"?0"},
	"sec-ch-ua-model":             {""},
	"sec-ch-ua-platform":          {"\"macOS\""},
	"sec-ch-ua-platform-version":  {"\"13.3.1\""},
	"sec-ch-ua-wow64":             {"?0"},
	"sec-fetch-dest":              {"document"},
	"sec-fetch-mode":              {"navigate"},
	"sec-fetch-site":              {"same-origin"},
	"sec-fetch-user":              {"?1"},
	"upgrade-insecure-requests":   {"1"},
	"user-agent":                  {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36"},
}

func tryRequest(client *http.Client, method string, rawUrl string) error {
	parsedUrl, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return err
	}

	request := &http.Request{
		Method: method,
		URL:    parsedUrl,
		Header: header,
	}

	_, clientErr := client.Do(request)

	if clientErr != nil {
		return clientErr
	}

	return nil
}

type Url struct {
	method string
	url    string
}

func TouchWebsite(dirtyUrl string) error {
	rawUrl, sanitizeErr := sanitize(dirtyUrl)

	if sanitizeErr != nil {
		return sanitizeErr
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   20 * time.Second,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	httpUrl := "http://" + rawUrl
	httpsUrl := "https://" + rawUrl

	urls := []*Url{
		//{url: httpsUrl, method: "HEAD"},
		//{url: httpUrl, method: "HEAD"},
		{url: httpUrl, method: "GET"},
		{url: httpsUrl, method: "GET"},
	}
	finalError := errors.Join()
	for _, myUrl := range urls {
		res := tryRequest(client, myUrl.method, myUrl.url)
		if res != nil {
			finalError = errors.Join(finalError, res)
		} else {
			return nil
		}
	}
	return finalError
}
