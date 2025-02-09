// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasource

import (
	"terraform-provider-nginxproxymanager/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var nestedCertificate = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		Description: "The ID of the certificate.",
		Computed:    true,
	},
}

var Certificates = map[string]schema.Attribute{
	"certificates": schema.ListNestedAttribute{
		Description: "The certificates.",
		Computed:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: utils.MergeMaps(Certificate, nestedCertificate),
		},
	},
}
