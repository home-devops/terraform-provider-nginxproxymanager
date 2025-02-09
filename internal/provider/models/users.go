// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type Users struct {
	Users []User `tfsdk:"users"`
}

func (m *Users) Load(ctx context.Context, resource *resources.UserCollection) diag.Diagnostics {
	diags := diag.Diagnostics{}
	m.Users = make([]User, len(*resource))
	for i, user := range *resource {
		diags.Append(m.Users[i].Load(ctx, &user)...)
	}

	return diags
}
