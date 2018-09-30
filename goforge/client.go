package goforge

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "1.0.0"
	userAgent      = "goforge/" + libraryVersion
	mediaType      = "application/json"
	defaultBaseURL = "https://forge.laravel.com"
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

// NewRequest returns a new pre-configured HTTP Request
func (c *Client) NewRequest(path string) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

// Do performs a request and returns the response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Do performs a request and returns the response
func (c *Client) DoJSON(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, json.NewDecoder(resp.Body).Decode(&v)
}
