package goforge

import (
	"net/http"
	"reflect"
	"testing"
)

// testCredentialsList
func TestListCredentials(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/credentials", CredentialsListSuccessfulResponse(t))

	result, err := tc.Client.ListCredentials()
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
	return respondJsonWithStringBody(t, "GET", `{"credentials":[{"id":1,"type":"test","name":"Personal"}]}`)
}

func TestGetCredentialByID(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/credentials", CredentialsListSuccessfulResponse(t))

	result, err := tc.Client.GetCredentialByID(1)
	if err != nil {
		t.Errorf("Error getting credentials: %v", err)
	}

	expected := &Credential{ID: 1, Type: "test", Name: "Personal"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CredentialsList returned %+v, expected %+v", result, expected)
	}
}

func TestGetCredentialByName(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/credentials", CredentialsListSuccessfulResponse(t))

	result, err := tc.Client.GetCredentialByName("Personal")
	if err != nil {
		t.Errorf("Error getting credentials: %v", err)
	}

	expected := &Credential{ID: 1, Type: "test", Name: "Personal"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CredentialsList returned %+v, expected %+v", result, expected)
	}
}
