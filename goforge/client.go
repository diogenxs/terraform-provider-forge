package goforge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	libraryVersion   = "1.0.0"
	defaultUserAgent = "goforge/" + libraryVersion
	defaultBaseURL   = "https://forge.laravel.com"

	mediaType = "application/json"
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

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: defaultUserAgent}

	return c, nil
}

// NewRequest returns a new pre-configured HTTP Request
func (c *Client) NewRequest(method string, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", mediaType)
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// SetUserAgent adds an additional user agent string to all requests
func (c *Client) SetUserAgent(ua string) error {
	c.UserAgent = fmt.Sprintf("%s %s", ua, c.UserAgent)

	return nil
}

// Do performs a request and returns the response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DoJSON performs a request and returns the response
func (c *Client) DoJSON(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, json.NewDecoder(resp.Body).Decode(&v)
}
