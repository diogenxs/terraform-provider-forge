package goforge

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

// TestListServer
func TestListServer(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/servers", ServersListSuccessfulResponse(t))

	result, err := tc.Client.ListServer()
	if err != nil {
		t.Errorf("Error getting credentials: %v", err)
	}

	expected := []Server{
		Server{
			ID:               1,
			CredentialID:     1,
			Name:             "test-via-api",
			PHPVersion:       "php71",
			IsReady:          true,
			Size:             "512MB",
			Region:           "Amsterdam 2",
			Revoked:          false,
			IPAddress:        "37.139.3.148",
			PrivateIPAddress: "10.129.3.252",
			CreatedAt:        "2016-12-15 18:38:18",
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ServersList returned %+v, expected %+v", result, expected)
	}
}

func ServersListSuccessfulResponse(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testCommonHeaders(t, r)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		io.WriteString(w, `{
			"servers":[
				{
					"id": 1,
					"credential_id": 1,
					"name": "test-via-api",
					"size": "512MB",
					"region": "Amsterdam 2",
					"php_version": "php71",
					"ip_address": "37.139.3.148",
					"private_ip_address": "10.129.3.252",
					"revoked": false,
					"created_at": "2016-12-15 18:38:18",
					"is_ready": true
				}
			]
		}`)
	})
}
