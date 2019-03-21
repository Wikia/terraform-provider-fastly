package main

import (
	"github.com/Wikia/terraform-provider-fastly/fastly"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fastly.Provider})
}
