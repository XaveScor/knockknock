package requester

import (
	"errors"
	"net/url"
	"strings"
)

func getUrl(dirtyUrl string) (*url.URL, error) {
	parsedUrl, parsedUrlErr := url.ParseRequestURI(dirtyUrl)
	if parsedUrlErr == nil {
		return parsedUrl, nil
	}

	// For urls without protocol
	parsedUrlWithoutProto, parsedUrlWithoutProtoErr := url.ParseRequestURI("http://" + dirtyUrl)
	if parsedUrlWithoutProtoErr == nil && strings.Contains(parsedUrlWithoutProto.Hostname(), ".") {
		return parsedUrlWithoutProto, nil
	}

	// For absolute urls like //google.com
	parsedAbsUrl, parsedAbsUrlErr := url.ParseRequestURI("http:" + dirtyUrl)
	if parsedAbsUrlErr == nil && strings.Contains(parsedUrlWithoutProto.Hostname(), ".") {
		return parsedAbsUrl, nil
	}

	return nil, errors.Join(parsedUrlErr, parsedUrlWithoutProtoErr, parsedAbsUrlErr)
}

func sanitize(dirtyUrl string) (string, error) {
	parsedUrl, parsedUrlErr := getUrl(dirtyUrl)

	if parsedUrlErr != nil {
		return "", parsedUrlErr
	}

	port := parsedUrl.Port()

	if port == "" {
		return parsedUrl.Hostname(), nil
	}

	return parsedUrl.Hostname() + ":" + port, nil
}
