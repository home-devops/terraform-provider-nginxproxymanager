// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"terraform-provider-nginxproxymanager/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var meUser = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Description: "The ID of the user.",
		Computed:    true,
	},
}

var UserMe = utils.MergeMaps(User, meUser)
