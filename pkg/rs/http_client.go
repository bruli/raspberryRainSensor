package rs

import (
	"net/http"
)

//go:generate go tool moq -out zmock_http_client_test.go -pkg rs_test  . HTTPClient
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
