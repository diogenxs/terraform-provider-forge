package goforge

import (
	"net/http"
	"net/url"
)

const (
	libraryVersion = "1.0.0"
	userAgent      = "goforge/" + libraryVersion
	mediaType      = "application/json"
	defaultBaseURL = "https://forge.laravel.com/v1"
)

// Client enables communication with the Laravel Forge API
type Client struct {
	// HTTP client used to handle requests
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client
	UserAgent string
}

// NewClient returns a new API client.
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}

	return c, nil
}
