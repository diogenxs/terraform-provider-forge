package goforge

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

// testCredentialsList
func TestCredentialsList(t *testing.T) {
	client, ts, err := createTestClient(CredentialsListSuccessfulResponse(t))

	if err != nil {
		t.Errorf("Unable to create test client: %v", err)
	}

	defer ts.Close()

	creds, err := client.CredentialsList()
	if err != nil {
		t.Errorf("Error getting credentials: %v", err)
	}

	expected := []Credential{
		Credential{ID: 1, Type: "test", Name: "Personal"},
	}

	if !reflect.DeepEqual(creds, expected) {
		t.Errorf("CredentialsList returned %+v, expected %+v", creds, expected)
	}
}

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

func CredentialsListSuccessfulResponse(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got ‘%s’", r.Method)
		}

		if r.URL.EscapedPath() != "/api/v1/credentials" {
			t.Errorf("Expected request to ‘/api/v1/credentials, got ‘%s’", r.URL.EscapedPath())
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		io.WriteString(w, `{"credentials":[{"id":1,"type":"test","name":"Personal"}]}`)
	})
}
