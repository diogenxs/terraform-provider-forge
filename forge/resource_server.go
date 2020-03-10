package forge

import (
	"fmt"
	"strconv"

	"github.com/diogenxs/terraform-provider-forge/goforge"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"platform": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server provider. Valid values are ocean2 for Digital Ocean, linode, vultr, aws, hetzner and custom.",
			},
			"credential_id": &schema.Schema{
				Type:        schema.TypeString,
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
			"ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Public IP Address of the server",
			},
			"private_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Private IP Address of the server",
			},
			"tags": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "An array of tags applied to this object. Tags are for organizational purposes only.",
			},
			"network": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Optional:    true,
				Description: "An array of servers ID's that this server could connect, need to be in the same platform/region.",
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*goforge.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Invalid Server id: %v", err)
	}

	server, err := client.GetServerByID(id)
	if err != nil {
		return fmt.Errorf("Error searching for server id: %v", err)
	}
	fmt.Printf("%v", server)
	d.Set("name", server.Name)
	d.Set("platform", server.Platform)
	d.Set("credential_id", strconv.Itoa(server.CredentialID))
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
	d.Set("tags", server.Tags)
	d.Set("tags", server.Network)

	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
