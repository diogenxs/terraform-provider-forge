package goforge

import "fmt"

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

// ListCredentials
func (c *Client) ListCredentials() ([]Credential, error) {
	var r CredentialsListResponse

	_, err := c.DoJSONRequest("GET", "/credentials", nil, &r)

	if err != nil {
		return nil, err
	}

	return r.Credentials, nil
}

func (c *Client) GetCredentialByID(id int) (*Credential, error) {
	credentials, err := c.ListCredentials()
	if err != nil {
		return nil, err
	}
	for _, credential := range credentials {
		if id == credential.ID {
			return &credential, nil
		}
	}

	return nil, fmt.Errorf("Credential with id %v not found", id)
}

func (c *Client) GetCredentialByName(name string) (*Credential, error) {
	credentials, err := c.ListCredentials()
	if err != nil {
		return nil, err
	}
	for _, credential := range credentials {
		if name == credential.Name {
			return &credential, nil
		}
	}

	return nil, fmt.Errorf("credential with name %v not found", name)
}
