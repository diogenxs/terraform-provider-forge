package goforge

// Credential
type Credential struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// CredentialsListResponse
type CredentialsListResponse struct {
	Credentials []Credential `json:"credentials"`
}

// CredentialsList
func (c *Client) CredentialsList() ([]Credential, error) {
	var r CredentialsListResponse

	req, err := c.NewRequest("/api/v1/credentials")

	if err != nil {
		return nil, err
	}

	_, err = c.DoJSON(req, &r)

	if err != nil {
		return nil, err
	}

	return r.Credentials, nil
}
