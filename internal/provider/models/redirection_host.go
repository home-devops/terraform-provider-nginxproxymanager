// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/inputs"
	"terraform-provider-nginxproxymanager/internal/client/resources"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RedirectionHost struct {
	ID          types.Int64  `tfsdk:"id"`
	CreatedOn   types.String `tfsdk:"created_on"`
	ModifiedOn  types.String `tfsdk:"modified_on"`
	OwnerUserId types.Int64  `tfsdk:"owner_user_id"`
	Meta        types.Map    `tfsdk:"meta"`

	DomainNames       []types.String `tfsdk:"domain_names"`
	ForwardScheme     types.String   `tfsdk:"forward_scheme"`
	ForwardDomainName types.String   `tfsdk:"forward_domain_name"`
	ForwardHTTPCode   types.Int64    `tfsdk:"forward_http_code"`
	CertificateID     types.String   `tfsdk:"certificate_id"`
	SSLForced         types.Bool     `tfsdk:"ssl_forced"`
	HSTSEnabled       types.Bool     `tfsdk:"hsts_enabled"`
	HSTSSubdomains    types.Bool     `tfsdk:"hsts_subdomains"`
	HTTP2Support      types.Bool     `tfsdk:"http2_support"`
	PreservePath      types.Bool     `tfsdk:"preserve_path"`
	BlockExploits     types.Bool     `tfsdk:"block_exploits"`
	AdvancedConfig    types.String   `tfsdk:"advanced_config"`
	Enabled           types.Bool     `tfsdk:"enabled"`
}

func (m *RedirectionHost) Load(ctx context.Context, resource *resources.RedirectionHost) diag.Diagnostics {
	meta, diags := types.MapValueFrom(ctx, types.StringType, resource.Meta.Map())

	m.ID = types.Int64Value(resource.ID)
	m.CreatedOn = types.StringValue(resource.CreatedOn)
	m.ModifiedOn = types.StringValue(resource.ModifiedOn)
	m.OwnerUserId = types.Int64Value(resource.OwnerUserID)
	m.Meta = meta

	m.ForwardScheme = types.StringValue(resource.ForwardScheme)
	m.ForwardDomainName = types.StringValue(resource.ForwardDomainName)
	m.ForwardHTTPCode = types.Int64Value(resource.ForwardHTTPCode)
	m.CertificateID = types.StringValue(string(resource.CertificateID))
	m.SSLForced = types.BoolValue(resource.SSLForced)
	m.HSTSEnabled = types.BoolValue(resource.HSTSEnabled)
	m.HSTSSubdomains = types.BoolValue(resource.HSTSSubdomains)
	m.HTTP2Support = types.BoolValue(resource.HTTP2Support)
	m.PreservePath = types.BoolValue(resource.PreservePath)
	m.BlockExploits = types.BoolValue(resource.BlockExploits)
	m.AdvancedConfig = types.StringValue(resource.AdvancedConfig)
	m.Enabled = types.BoolValue(resource.Enabled)

	m.DomainNames = make([]types.String, len(resource.DomainNames))
	for i, v := range resource.DomainNames {
		m.DomainNames[i] = types.StringValue(v)
	}

	if m.ForwardScheme.Equal(types.StringValue("$scheme")) {
		m.ForwardScheme = types.StringValue("auto")
	}

	return diags
}

func (m *RedirectionHost) Save(ctx context.Context, input *inputs.RedirectionHost) diag.Diagnostics {
	diags := diag.Diagnostics{}

	input.DomainNames = make([]string, len(m.DomainNames))
	for i, v := range m.DomainNames {
		input.DomainNames[i] = v.ValueString()
	}
	input.ForwardScheme = m.ForwardScheme.ValueString()
	input.ForwardDomainName = m.ForwardDomainName.ValueString()
	input.ForwardHTTPCode = m.ForwardHTTPCode.ValueInt64()
	input.CertificateID = m.CertificateID.ValueString()
	input.SSLForced = m.SSLForced.ValueBool()
	input.HSTSEnabled = m.HSTSEnabled.ValueBool()
	input.HSTSSubdomains = m.HSTSSubdomains.ValueBool()
	input.HTTP2Support = m.HTTP2Support.ValueBool()
	input.BlockExploits = m.BlockExploits.ValueBool()
	input.AdvancedConfig = m.AdvancedConfig.ValueString()
	input.Meta = map[string]string{}

	return diags
}
