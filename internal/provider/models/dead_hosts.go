// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type DeadHosts struct {
	DeadHosts []DeadHost `tfsdk:"dead_hosts"`
}

func (m *DeadHosts) Load(ctx context.Context, resource *resources.DeadHostCollection) diag.Diagnostics {
	diags := diag.Diagnostics{}
	m.DeadHosts = make([]DeadHost, len(*resource))
	for i, deadHost := range *resource {
		diags.Append(m.DeadHosts[i].Load(ctx, &deadHost)...)
	}

	return diags
}
