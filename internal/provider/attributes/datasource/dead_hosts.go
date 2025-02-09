// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"terraform-provider-nginxproxymanager/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var nestedDeadHost = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Description: "The ID of the dead host.",
		Computed:    true,
	},
}

var DeadHosts = map[string]schema.Attribute{
	"dead_hosts": schema.ListNestedAttribute{
		Description: "The dead hosts.",
		Computed:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: utils.MergeMaps(DeadHost, nestedDeadHost),
		},
	},
}
