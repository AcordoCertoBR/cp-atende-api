package http

import (
	"net/http"
	"net/url"
	"time"
)

type Http struct {
	httpClient *http.Client
	clientName string
}

func New(httpClientName string, timeoutInSeconds int) *Http {
	client := http.Client{
		Timeout: time.Duration(timeoutInSeconds) * time.Second,
	}

	return &Http{
		httpClient: &client,
		clientName: httpClientName,
	}
}

func (u *Http) Do(req *http.Request) (res *http.Response, err error) {
	return u.httpClient.Do(req)
}

func (u *Http) Get(urlString string) (res *http.Response, err error) {
	parsedURL, _ := url.Parse(urlString)

	path := parsedURL.Path
	if path == "" {
		path = urlString
	}

	return u.httpClient.Get(urlString)
}
