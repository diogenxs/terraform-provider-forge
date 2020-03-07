package forge

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"platform": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server provider. Valid values are ocean2 for Digital Ocean, linode, vultr, aws, hetzner and custom.",
			},
			"credential_id": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "This is only required when the provider is not custom.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the server to create",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the region where the server will be created. This value is not required you are building a Custom VPS server.",
			},
			"size": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The size of the serve will be created. This value is not required you are building a Custom VPS server.",
			},
			"php_version": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Valid values are php74, php73,php72,php71, php70, and php56.",
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
