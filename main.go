package main

import (
	"github.com/abhinavdahiya/terraform-provider-azurerm-srv/srv"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: srv.Provider})
}
