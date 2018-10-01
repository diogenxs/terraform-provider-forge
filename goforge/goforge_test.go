package goforge

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// MockServer
type MockServer struct {
	Mux    *http.ServeMux
	Server *httptest.Server
}

// SetUp
func StartMockServer() (*MockServer, error) {
	mux := http.NewServeMux()
	hts := httptest.NewServer(mux)

	return &MockServer{Mux: mux, Server: hts}, nil
}

// TearDown
func (ts *MockServer) Close() {
	ts.Server.Close()
}

// TestClient
type TestClient struct {
	Server *MockServer
	Client *Client
}

func (ts *TestClient) TearDown() {
	ts.Server.Close()
}

func NewTestClient() (*TestClient, error) {
	client, err := NewClient(nil)
	if err != nil {
		return nil, err
	}

	ms, err := StartMockServer()
	if err != nil {
		return nil, err
	}

	client.BaseURL, err = url.Parse(ms.Server.URL)

	if err != nil {
		ms.Close()

		return nil, err
	}

	return &TestClient{Server: ms, Client: client}, nil
}

func SetUpTestClient(t *testing.T) *TestClient {
	client, err := NewTestClient()

	if err != nil {
		t.Errorf("Unable to create test client: %v", err)

		return nil
	}

	return client
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; want != got {
		t.Errorf("Request METHOD expected to be `%v`, got `%v`", want, got)
	}
}

func testHeader(t *testing.T, r *http.Request, name, want string) {
	if got := r.Header.Get(name); want != got {
		t.Errorf("Request() %v expected to be `%#v`, got `%#v`", name, want, got)
	}
}

func testCommonHeaders(t *testing.T, r *http.Request) {
	testHeader(t, r, "User-Agent", "goforge/1.0.0")

	testHeader(t, r, "Accept", "application/json")
	testHeader(t, r, "Content-Type", "application/json")
}
