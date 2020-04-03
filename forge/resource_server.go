package forge

import (
	"fmt"
	"log"
	"strconv"

	"github.com/diogenxs/terraform-provider-forge/goforge"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
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
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The server provider. Valid values are ocean2 for Digital Ocean, linode, vultr, aws, hetzner and custom.",
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"credential_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "This is only required when the provider is not custom.",
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the server to create",
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"region": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the region where the server will be created. This value is not required you are building a Custom VPS server.",
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"size": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The size of the serve will be created. This value is not required you are building a Custom VPS server.",
				ForceNew:    true,
			},
			"php_version": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Valid values are php74, php73,php72,php71, php70, and php56.",
			},
			"ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public IP Address of the server",
			},
			"private_ip_address": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
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
			"ready": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"root_password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"root_database_password": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*goforge.Client)

	name := d.Get("name").(string)

	server := goforge.Server{
		Name:       name,
		Platform:   d.Get("platform").(string),
		Region:     d.Get("region").(string),
		PHPVersion: d.Get("php_version").(string),
		Size:       d.Get("size").(string),
	}

	log.Printf("[DEBUG] Server create configuration: %#v", server)

	resp, err := client.CreateServer(server)

	if err != nil {
		return fmt.Errorf("Error creating Server: %s", err)
	}

	d.Set("root_password", resp.SudoPassword)
	d.Set("root_database_password", resp.DatabasePassword)
	d.SetId(strconv.Itoa(resp.ID))

	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		resp, err := client.GetServerByName(name)

		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error getting server: %s", err))
		}

		if resp.IsReady {
			return resource.NonRetryableError(resourceServerRead(d, m))
		}

		return resource.RetryableError(fmt.Errorf("Expected Server to be ready but was in state %v", resp.IsReady))
	})
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
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*goforge.Client)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Invalid Server id: %v", err)
	}

	err = client.DeleteServer(id)
	if err != nil {
		return fmt.Errorf("Error deleting server id: %v", err)
	}

	d.SetId("")
	return nil
}
