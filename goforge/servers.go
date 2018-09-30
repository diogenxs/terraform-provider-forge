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

	req, err := c.NewRequest("GET", "/api/v1/servers", nil)

	if err != nil {
		return nil, err
	}

	_, err = c.DoJSON(req, &r)

	if err != nil {
		return nil, err
	}

	return r.Servers, nil
}
