package forge

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diogenxs/terraform-provider-forge/goforge"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/terraform"
	"golang.org/x/oauth2"
)

type Config struct {
	Token string
}

// Client() returns a new client for accessing laravel forge
func (c *Config) Client() (*goforge.Client, error) {
	tokenSrc := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: c.Token,
	})

	oauthTransport := &oauth2.Transport{
		Source: tokenSrc,
	}
	loggingTransport := logging.NewTransport("Forge", oauthTransport)
	oauth2Client := &http.Client{
		Transport: loggingTransport,
	}

	client, err := goforge.NewClient(oauth2Client)
	if err != nil {
		return nil, err
	}

	userAgent := fmt.Sprintf("Terraform/%s", terraform.VersionString())

	err = client.SetUserAgent(userAgent)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] Laravel Forge Client configured for URL: %s", client.BaseURL.String())

	return client, nil
}
