package goforge

// Server
type Server struct {
	ID               int    `json:"id"`
	CredentialID     int    `json:"credential_id"`
	Name             string `json:"name"`
	Size             string `json:"size"`
	Region           string `json:"region"` //"ams2",
	PHPVersion       string `json:"php_version"`
	IPAddress        string `json:"ip_address"`         //null,
	PrivateIPAddress string `json:"private_ip_address"` //null,
	// BlackFireStatus  string `json:"blackfire_status"`   //null,
	// PaperTrailStatus string `json:"papertrail_status"`  //null,
	Revoked   bool   `json:"revoked"`    //false,
	CreatedAt string `json:"created_at"` //"2016-12-15 15:04:05",
	IsReady   bool   `json:"is_ready"`
}

// ServersListResponse
type ServersListResponse struct {
	Servers []Server `json:"servers"`
}

// ListServer
func (c *Client) ListServer() ([]Server, error) {
	var r ServersListResponse

	_, err := c.DoJSONRequest("GET", "/servers", nil, &r)

	if err != nil {
		return nil, err
	}

	return r.Servers, nil
}
