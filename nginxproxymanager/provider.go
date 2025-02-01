package nginxproxymanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"nginxproxymanager_widget": resourceWidget(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
