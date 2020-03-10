package forge

import (
	"fmt"
	"strconv"

	"github.com/diogenxs/terraform-provider-forge/goforge"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"platform": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The server provider. Valid values are ocean2 for Digital Ocean, linode, vultr, aws, hetzner and custom.",
			},
			"credential_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "This is only required when the provider is not custom.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the server to create",
			},
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the region where the server will be created. This value is not required you are building a Custom VPS server.",
			},
			"size": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The size of the serve will be created. This value is not required you are building a Custom VPS server.",
			},
			"php_version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Valid values are php74, php73,php72,php71, php70, and php56.",
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"blackfire_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"papertrail_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When the server was created",
			},
			"revoked": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_ready": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*goforge.Client)

	reqName := d.Get("name").(string)

	if reqName == "" {
		return fmt.Errorf("Server Name is required")
	}

	d.SetId("")

	server, err := client.GetServerByName(reqName)
	if err != nil {
		return fmt.Errorf("Error listing servers: %s", err)
	}

	if server != nil {
		d.SetId(strconv.Itoa(server.ID))
		d.Set("name", server.Name)
		d.Set("platform", server.Platform)
		d.Set("credential_id", server.CredentialID)
		d.Set("region", server.Region)
		d.Set("size", server.Size)
		d.Set("php_version", server.PHPVersion)
		d.Set("ip_address", server.IPAddress)
		d.Set("private_ip_address", server.PrivateIPAddress)
		d.Set("blackfire_status", server.BlackFireStatus)
		d.Set("papertrail_status", server.PaperTrailStatus)
		d.Set("created_at", server.CreatedAt)
		d.Set("revoked", server.Revoked)
		d.Set("is_ready", server.IsReady)
		return nil
	}

	return fmt.Errorf("Server %s was not found", reqName)
}
