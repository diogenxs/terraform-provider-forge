package forge

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pittfit/terraform-provider-forge/goforge"
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

	client, err := goforge.NewClient(oauth2.NewClient(oauth2.NoContext, tokenSrc))
	if err != nil {
		return nil, err
	}

	userAgent := fmt.Sprintf("Terraform/%s", terraform.VersionString())

	err = client.SetUserAgent(userAgent)
	if err != nil {
		return nil, err
	}

	if logging.IsDebugOrHigher() {
		// client.OnRequestCompleted(logRequestAndResponse)
	}

	log.Printf("[INFO] Laravel Forge Client configured for URL: %s", client.BaseURL.String())

	return client, nil
}

func logRequestAndResponse(req *http.Request, resp *http.Response) {
	reqData, err := httputil.DumpRequest(req, true)
	if err == nil {
		log.Printf("[DEBUG] "+logRequestMessage, string(reqData))
	} else {
		log.Printf("[ERROR] Laravel Forge API Request error: %#v", err)
	}

	respData, err := httputil.DumpResponse(resp, true)
	if err == nil {
		log.Printf("[DEBUG] "+logResponseMessage, string(respData))
	} else {
		log.Printf("[ERROR] Laravel Forge API Response error: %#v", err)
	}
}

const logRequestMessage = `Laravel Forge API Request Details:
---[ REQUEST ]---------------------------------------
%s
-----------------------------------------------------`

const logResponseMessage = `Laravel Forge API Response Details:
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------`
