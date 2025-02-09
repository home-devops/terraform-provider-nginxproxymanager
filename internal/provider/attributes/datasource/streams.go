// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"terraform-provider-nginxproxymanager/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var nestedStream = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Description: "The ID of the stream.",
		Computed:    true,
	},
}

var Streams = map[string]schema.Attribute{
	"streams": schema.ListNestedAttribute{
		Description: "The streams.",
		Computed:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: utils.MergeMaps(Stream, nestedStream),
		},
	},
}
