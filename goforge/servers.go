package goforge

// Server
type Server struct {
	ID           int    `json:"id"`
	CredentialID int    `json:"credential_id"`
	Name         string `json:"name"`
	PHPVersion   string `json:"php_version"`
	IsReady      bool   `json:"is_ready"`
}

// ServersListResponse
type ServersListResponse struct {
	Servers []Server `json:"servers"`
}

// ServersList
func (c *Client) ServersList() ([]Server, error) {
	var r ServersListResponse

	_, err := c.DoJSONRequest("GET", "/servers", nil, &r)

	if err != nil {
		return nil, err
	}

	return r.Servers, nil
}
