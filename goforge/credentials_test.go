package goforge

import (
	"io"
	"net/http"
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

	result, err := client.CredentialsList()
	if err != nil {
		t.Errorf("Error getting credentials: %v", err)
	}

	expected := []Credential{
		Credential{ID: 1, Type: "test", Name: "Personal"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CredentialsList returned %+v, expected %+v", result, expected)
	}
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
