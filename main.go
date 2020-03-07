package main

import (
	"github.com/diogenxs/terraform-provider-forge/forge"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return forge.Provider()
		},
	})
}
