package forge

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FORGE_TOKEN", nil),
				Description: "The token for API operations.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"forge_credential": dataSourceCredentials(),
			"forge_server":     dataSourceServer(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"forge_server": resourceServer(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Token: d.Get("token").(string),
	}

	return config.Client()
}
