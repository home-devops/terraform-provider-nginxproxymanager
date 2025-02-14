// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"strconv"
	"terraform-provider-nginxproxymanager/internal/client/inputs"
	"terraform-provider-nginxproxymanager/internal/client/resources"
	"terraform-provider-nginxproxymanager/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RedirectionHost struct {
	ID          types.Int64  `tfsdk:"id"`
	CreatedOn   types.String `tfsdk:"created_on"`
	ModifiedOn  types.String `tfsdk:"modified_on"`
	OwnerUserId types.Int64  `tfsdk:"owner_user_id"`
	Meta        types.Map    `tfsdk:"meta"`

	DomainNames       types.List   `tfsdk:"domain_names"`
	ForwardScheme     types.String `tfsdk:"forward_scheme"`
	ForwardDomainName types.String `tfsdk:"forward_domain_name"`
	ForwardHTTPCode   types.Int64  `tfsdk:"forward_http_code"`
	CertificateID     types.Int64  `tfsdk:"certificate_id"`
	CertificateNew    types.Bool   `tfsdk:"certificate_new"`
	SSLForced         types.Bool   `tfsdk:"ssl_forced"`
	HSTSEnabled       types.Bool   `tfsdk:"hsts_enabled"`
	HSTSSubdomains    types.Bool   `tfsdk:"hsts_subdomains"`
	HTTP2Support      types.Bool   `tfsdk:"http2_support"`
	PreservePath      types.Bool   `tfsdk:"preserve_path"`
	BlockExploits     types.Bool   `tfsdk:"block_exploits"`
	AdvancedConfig    types.String `tfsdk:"advanced_config"`
	Enabled           types.Bool   `tfsdk:"enabled"`
}

func (m *RedirectionHost) Load(ctx context.Context, resource *resources.RedirectionHost) diag.Diagnostics {
	var diags diag.Diagnostics
	m.Meta, diags = types.MapValueFrom(ctx, types.StringType, resource.Meta.Map())

	if diags.HasError() {
		return diags
	}

	m.ID = types.Int64Value(resource.ID)
	m.CreatedOn = types.StringValue(resource.CreatedOn)
	m.ModifiedOn = types.StringValue(resource.ModifiedOn)
	m.OwnerUserId = types.Int64Value(resource.OwnerUserID)

	m.ForwardScheme = types.StringValue(resource.ForwardScheme)
	m.ForwardDomainName = types.StringValue(resource.ForwardDomainName)
	m.ForwardHTTPCode = types.Int64Value(resource.ForwardHTTPCode)
	m.CertificateID = types.Int64Value(resource.CertificateID)
	m.SSLForced = types.BoolValue(resource.SSLForced)
	m.HSTSEnabled = types.BoolValue(resource.HSTSEnabled)
	m.HSTSSubdomains = types.BoolValue(resource.HSTSSubdomains)
	m.HTTP2Support = types.BoolValue(resource.HTTP2Support)
	m.PreservePath = types.BoolValue(resource.PreservePath)
	m.BlockExploits = types.BoolValue(resource.BlockExploits)
	m.AdvancedConfig = types.StringValue(resource.AdvancedConfig)
	m.Enabled = types.BoolValue(resource.Enabled)

	m.DomainNames, diags = types.ListValueFrom(ctx, types.StringType, resource.DomainNames)

	if diags.HasError() {
		return diags
	}

	if m.ForwardScheme.Equal(types.StringValue("$scheme")) {
		m.ForwardScheme = types.StringValue("auto")
	}

	return diags
}

func (m *RedirectionHost) Save(ctx context.Context, input *inputs.RedirectionHost) diag.Diagnostics {
	var diags diag.Diagnostics

	input.ForwardScheme = m.ForwardScheme.ValueString()
	input.ForwardDomainName = m.ForwardDomainName.ValueString()
	input.ForwardHTTPCode = m.ForwardHTTPCode.ValueInt64()

	if m.CertificateNew.ValueBool() {
		input.CertificateID = "new"
	} else {
		input.CertificateID = strconv.FormatInt(m.CertificateID.ValueInt64(), 10)
	}

	input.SSLForced = m.SSLForced.ValueBool()
	input.HSTSEnabled = m.HSTSEnabled.ValueBool()
	input.HSTSSubdomains = m.HSTSSubdomains.ValueBool()
	input.HTTP2Support = m.HTTP2Support.ValueBool()
	input.BlockExploits = m.BlockExploits.ValueBool()
	input.AdvancedConfig = m.AdvancedConfig.ValueString()
	input.Meta = map[string]string{}

	input.DomainNames, diags = utils.ConvertListToStringSlice(m.DomainNames)

	if diags.HasError() {
		return diags
	}

	return diags
}
