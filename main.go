package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/home-devops/terraform-provider-nginxproxymanager/nginxproxymanager"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nginxproxymanager.Provider,
	})
}
