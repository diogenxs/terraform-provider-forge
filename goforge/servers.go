package goforge

import (
	"fmt"
	"strconv"
)

// Server
type Server struct {
	ID               int      `json:"id,omitempty"`
	CredentialID     int      `json:"credential_id"`
	Platform         string   `json:"provider"`
	PlatformID       string   `json:"provider_id,omitempty"`
	Name             string   `json:"name"`
	Size             string   `json:"size"`
	Region           string   `json:"region"` //"ams2",
	PHPVersion       string   `json:"php_version"`
	IPAddress        string   `json:"ip_address,omitempty"`         //null,
	PrivateIPAddress string   `json:"private_ip_address,omitempty"` //null,
	BlackFireStatus  string   `json:"blackfire_status,omitempty"`   //null,
	PaperTrailStatus string   `json:"papertrail_status,omitempty"`  //null,
	Revoked          bool     `json:"revoked,omitempty"`            //false,
	CreatedAt        string   `json:"created_at,omitempty"`         //"2016-12-15 15:04:05",
	IsReady          bool     `json:"is_ready,omitempty"`
	SSHPort          int      `json:"ssh_port,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	Network          []int    `json:"network,omitempty"`
	SudoPassword     string   `json:"sudo_password,omitempty"`
	DatabasePassword string   `json:"database_password,omitempty"`
}

// ServersListResponse
type ServersListResponse struct {
	Servers []Server `json:"servers"`
}

type ServerResponse struct {
	Server           Server `json:"server"`
	SudoPassword     string `json:"sudo_password,omitempty"`
	DatabasePassword string `json:"database_password,omitempty"`
}

// ListServers
func (c *Client) ListServers() ([]Server, error) {
	var r ServersListResponse

	_, err := c.DoJSONRequest("GET", "/servers", nil, &r)

	if err != nil {
		return nil, err
	}

	return r.Servers, nil
}

func (c *Client) CreateServer(server Server) (*Server, error) {
	fmt.Printf("%v", server)
	var r ServerResponse
	_, err := c.DoJSONRequest("POST", "/servers", server, &r)

	if err != nil {
		return nil, err
	}
	if r.SudoPassword != "" {
		r.Server.SudoPassword = r.SudoPassword
	}

	if r.DatabasePassword != "" {
		r.Server.DatabasePassword = r.DatabasePassword
	}
	return &r.Server, nil
}

func (c *Client) GetServerByID(id int) (*Server, error) {
	var s ServerResponse
	_, err := c.DoJSONRequest("GET", "/servers/"+strconv.Itoa(id), nil, &s)

	if err != nil {
		return nil, err
	}

	return &s.Server, nil
}

func (c *Client) GetServerByName(name string) (*Server, error) {
	servers, err := c.ListServers()
	if err != nil {
		return nil, err
	}
	for _, server := range servers {
		if name == server.Name {
			return &server, nil
		}
	}

	return nil, fmt.Errorf("server with name %v not found", name)
}

func (c *Client) DeleteServer(id int) error {
	_, err := c.DoJSONRequest("DELETE", "/servers/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}
	return nil
}
