package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/jtopjian/terraform-provider-sensu/sensu"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sensu.Provider})
}
