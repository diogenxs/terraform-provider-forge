package forge

import (
	"fmt"
	"strconv"

	"github.com/diogenxs/terraform-provider-forge/goforge"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCredentials() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCredentialsRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the credential in the Forge",
				Optional:    true,
			},
			"platform": {
				Type:        schema.TypeString,
				Description: "The server provider what credential is assigned",
				Computed:    true,
			},
		},
	}
}

func dataSourceCredentialsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*goforge.Client)

	reqName := d.Get("name").(string)

	if reqName == "" {
		return fmt.Errorf("Crendential ID or Credential Name is required")
	}

	d.SetId("")

	credential, err := client.GetCredentialByName(reqName)
	if err != nil {
		return fmt.Errorf("Error listing Credentials: %s", err)
	}

	if credential != nil {
		d.SetId(strconv.Itoa(credential.ID))
		d.Set("name", credential.Name)
		d.Set("platform", credential.Type)
		return nil
	}

	return fmt.Errorf("Credential %s was not found", reqName)
}
