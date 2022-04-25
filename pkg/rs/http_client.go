package rs

import (
	"net/http"
)

//go:generate moq -out zmock_http_client_test.go -pkg en_test . HTTPClient
// HTTPClient provides an http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
