// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type RedirectionHosts struct {
	RedirectionHosts []RedirectionHost `tfsdk:"redirection_hosts"`
}

func (m *RedirectionHosts) Load(ctx context.Context, resource *resources.RedirectionHostCollection) diag.Diagnostics {
	diags := diag.Diagnostics{}
	m.RedirectionHosts = make([]RedirectionHost, len(*resource))
	for i, redirectionHost := range *resource {
		diags.Append(m.RedirectionHosts[i].Load(ctx, &redirectionHost)...)
	}

	return diags
}
