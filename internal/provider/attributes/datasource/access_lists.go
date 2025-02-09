// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"terraform-provider-nginxproxymanager/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var nestedAccessList = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Description: "The ID of the access list.",
		Computed:    true,
	},
}

var AccessLists = map[string]schema.Attribute{
	"access_lists": schema.ListNestedAttribute{
		Description: "The access lists.",
		Computed:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: utils.MergeMaps(AccessList, nestedAccessList),
		},
	},
}
