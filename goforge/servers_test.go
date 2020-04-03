package goforge

import (
	"fmt"
	"reflect"
	"testing"
)

const testServerString = `{
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
	"is_ready": true,
	"ssh_port": 22,
	"tags": ["test", "test2"],
	"network": [111, 222]
}`

var testServerExpected = Server{
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
	SSHPort:          22,
	Tags:             []string{"test", "test2"},
	Network:          []int{111, 222},
}

var testServerCreate = Server{
	Platform:     "ocean2",
	CredentialID: 1,
	Name:         "test-via-api",
	Size:         "512MB",
	PHPVersion:   "php71",
	Region:       "ams2",
}

// TestListServer
func TestListServer(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/servers", respondJsonWithStringBody(t, "GET", "{\"servers\":["+testServerString+"]}"))

	result, err := tc.Client.ListServers()
	if err != nil {
		t.Errorf("Error getting servers: %v", err)
	}

	expected := []Server{
		testServerExpected,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ServersList returned %+v, expected %+v", result, expected)
	}
}

func TestGetServerByID(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/servers/1", respondJsonWithStringBody(t, "GET", "{\"server\":"+testServerString+"}"))

	result, err := tc.Client.GetServerByID(1)
	if err != nil {
		t.Errorf("Error getting servers: %v", err)
	}

	expected := &testServerExpected

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ServersList returned %+v, expected %+v", result, expected)
	}
}

func TestGetServerByName(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/servers", respondJsonWithStringBody(t, "GET", "{\"servers\":["+testServerString+"]}"))

	result, err := tc.Client.GetServerByName("test-via-api")
	if err != nil {
		t.Errorf("Error getting servers: %v", err)
	}

	expected := &testServerExpected

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ServersList returned %+v, expected %+v", result, expected)
	}
}

func TestDeleteServer(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/servers/1", respondJsonWithStringBody(t, "DELETE", ""))

	err := tc.Client.DeleteServer(1)
	if err != nil {
		t.Errorf("Error deleting server: %v", err)
	}
}

func TestCreateServer(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	jsonString := "{\"server\":" + testServerString + ",\"sudo_password\": \"baracoda\",\"database_password\": \"spotted_eagle_ray\"}"
	fmt.Printf("\n%v\n\n\n", jsonString)
	tc.Server.Mux.Handle("/servers", respondJsonWithStringBody(t, "POST", jsonString))

	result, err := tc.Client.CreateServer(testServerCreate)
	if err != nil {
		t.Errorf("Error creating servers: %v", err)
	}

	expected := &testServerExpected
	expected.SudoPassword = "baracoda"
	expected.DatabasePassword = "spotted_eagle_ray"

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CreateServer\nreturned %v\nexpected %v", result, expected)
	}
}
