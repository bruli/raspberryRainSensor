package rs

import (
	"net/http"
)

// HTTPClient provides an http.Client
//
//go:generate moq -out zmock_http_client_test.go -pkg en_test . HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
