package goforge

import (
	"fmt"
	"strconv"
)

// Server
type Server struct {
	ID               int      `json:"id"`
	CredentialID     int      `json:"credential_id"`
	Platform         string   `json:"provider"`
	PlatformID       string   `json:"provider_id"`
	Name             string   `json:"name"`
	Size             string   `json:"size"`
	Region           string   `json:"region"` //"ams2",
	PHPVersion       string   `json:"php_version"`
	IPAddress        string   `json:"ip_address"`         //null,
	PrivateIPAddress string   `json:"private_ip_address"` //null,
	BlackFireStatus  string   `json:"blackfire_status"`   //null,
	PaperTrailStatus string   `json:"papertrail_status"`  //null,
	Revoked          bool     `json:"revoked"`            //false,
	CreatedAt        string   `json:"created_at"`         //"2016-12-15 15:04:05",
	IsReady          bool     `json:"is_ready"`
	SSHPort          int      `json:"ssh_port"`
	Tags             []string `json:"tags"`
	Network          []int    `json:"network"`
}

// ServersListResponse
type ServersListResponse struct {
	Servers []Server `json:"servers"`
}

type ServerResponse struct {
	Server Server `json:"server"`
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
