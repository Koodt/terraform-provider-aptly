package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/Koodt/terraform-provider-aptly/aptly"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aptly.Provider})
}
