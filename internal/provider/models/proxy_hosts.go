// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type ProxyHosts struct {
	ProxyHosts []ProxyHost `tfsdk:"proxy_hosts"`
}

func (m *ProxyHosts) Load(ctx context.Context, resource *resources.ProxyHostCollection) diag.Diagnostics {
	diags := diag.Diagnostics{}
	m.ProxyHosts = make([]ProxyHost, len(*resource))
	for i, proxyHost := range *resource {
		diags.Append(m.ProxyHosts[i].Load(ctx, &proxyHost)...)
	}

	return diags
}
