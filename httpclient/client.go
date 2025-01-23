// Package wrapping Go's native http client.
package httpclient

import (
	"net/http"
	"sync"
)

var (
	client     *http.Client
	clientOnce sync.Once
)

// Example usage:
//
//	 url, _ := url.Parse("http://google.com/search?q=golang")
//
//	 request := http.Request{
//		  Method: "GET"
//		  URL:    url,
//		  Header: map[string][]string{"User-Agent": "curl"},
//		  Body: &RequestBody{Data: []byte()},
//	 }
//
//	 response, err := HTTPClient.Do(&request)
func New() *http.Client {
	clientOnce.Do(func() {
		client = &http.Client{}
	})

	return client
}
