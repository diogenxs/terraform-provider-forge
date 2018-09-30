package goforge

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

func createTestClient(handler http.Handler) (*Client, *httptest.Server, error) {
	client, err := NewClient(nil)

	if err != nil {
		return nil, nil, err
	}

	ts := httptest.NewServer(handler)
	client.BaseURL, err = url.Parse(ts.URL)

	if err != nil {
		ts.Close()

		return nil, nil, err
	}

	return client, ts, nil
}
