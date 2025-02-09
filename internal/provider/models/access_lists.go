// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type AccessLists struct {
	AccessLists []AccessList `tfsdk:"access_lists"`
}

func (m *AccessLists) Load(ctx context.Context, resource *resources.AccessListCollection) diag.Diagnostics {
	diags := diag.Diagnostics{}
	m.AccessLists = make([]AccessList, len(*resource))
	for i, accessList := range *resource {
		diags.Append(m.AccessLists[i].Load(ctx, &accessList)...)
	}

	return diags
}
