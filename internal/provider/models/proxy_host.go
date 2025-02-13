// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"terraform-provider-nginxproxymanager/internal/client/inputs"
	"terraform-provider-nginxproxymanager/internal/client/resources"
	"terraform-provider-nginxproxymanager/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ProxyHost struct {
	ID          types.Int64  `tfsdk:"id"`
	CreatedOn   types.String `tfsdk:"created_on"`
	ModifiedOn  types.String `tfsdk:"modified_on"`
	OwnerUserId types.Int64  `tfsdk:"owner_user_id"`
	Meta        types.Map    `tfsdk:"meta"`

	DomainNames           types.List           `tfsdk:"domain_names"`
	ForwardScheme         types.String         `tfsdk:"forward_scheme"`
	ForwardHost           types.String         `tfsdk:"forward_host"`
	ForwardPort           types.Int64          `tfsdk:"forward_port"`
	CertificateID         types.String         `tfsdk:"certificate_id"`
	SSLForced             types.Bool           `tfsdk:"ssl_forced"`
	HSTSEnabled           types.Bool           `tfsdk:"hsts_enabled"`
	HSTSSubdomains        types.Bool           `tfsdk:"hsts_subdomains"`
	HTTP2Support          types.Bool           `tfsdk:"http2_support"`
	BlockExploits         types.Bool           `tfsdk:"block_exploits"`
	CachingEnabled        types.Bool           `tfsdk:"caching_enabled"`
	AllowWebsocketUpgrade types.Bool           `tfsdk:"allow_websocket_upgrade"`
	AccessListID          types.Int64          `tfsdk:"access_list_id"`
	AdvancedConfig        types.String         `tfsdk:"advanced_config"`
	Enabled               types.Bool           `tfsdk:"enabled"`
	Locations             []*ProxyHostLocation `tfsdk:"locations"`
}

func (m *ProxyHost) Load(ctx context.Context, resource *resources.ProxyHost) diag.Diagnostics {
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
	m.ForwardHost = types.StringValue(resource.ForwardHost)
	m.ForwardPort = types.Int64Value(int64(resource.ForwardPort))
	m.CertificateID = types.StringValue(string(resource.CertificateID))
	m.SSLForced = types.BoolValue(resource.SSLForced)
	m.HSTSEnabled = types.BoolValue(resource.HSTSEnabled)
	m.HSTSSubdomains = types.BoolValue(resource.HSTSSubdomains)
	m.HTTP2Support = types.BoolValue(resource.HTTP2Support)
	m.BlockExploits = types.BoolValue(resource.BlockExploits)
	m.CachingEnabled = types.BoolValue(resource.CachingEnabled)
	m.AllowWebsocketUpgrade = types.BoolValue(resource.AllowWebsocketUpgrade)
	m.AccessListID = types.Int64Value(resource.AccessListID)
	m.AdvancedConfig = types.StringValue(resource.AdvancedConfig)
	m.Enabled = types.BoolValue(resource.Enabled)

	m.DomainNames, diags = types.ListValueFrom(ctx, types.StringType, resource.DomainNames)

	if diags.HasError() {
		return diags
	}

	m.Locations = make([]*ProxyHostLocation, 0, len(resource.Locations))
	for _, locationResponse := range resource.Locations {
		location := &ProxyHostLocation{}
		location.Load(ctx, &locationResponse)
		m.Locations = append(m.Locations, location)
	}

	return diags
}

func (m *ProxyHost) Save(ctx context.Context, input *inputs.ProxyHost) diag.Diagnostics {
	var diags diag.Diagnostics

	input.ForwardScheme = m.ForwardScheme.ValueString()
	input.ForwardHost = m.ForwardHost.ValueString()
	input.ForwardPort = uint16(m.ForwardPort.ValueInt64())
	input.CertificateID = m.CertificateID.ValueString()
	input.SSLForced = m.SSLForced.ValueBool()
	input.HSTSEnabled = m.HSTSEnabled.ValueBool()
	input.HSTSSubdomains = m.HSTSSubdomains.ValueBool()
	input.HTTP2Support = m.HTTP2Support.ValueBool()
	input.BlockExploits = m.BlockExploits.ValueBool()
	input.CachingEnabled = m.CachingEnabled.ValueBool()
	input.AllowWebsocketUpgrade = m.AllowWebsocketUpgrade.ValueBool()
	input.AccessListID = m.AccessListID.ValueInt64()
	input.AdvancedConfig = m.AdvancedConfig.ValueString()
	input.Meta = map[string]string{}

	input.DomainNames, diags = utils.ConvertListToStringSlice(m.DomainNames)

	if diags.HasError() {
		return diags
	}

	input.Locations = make([]inputs.ProxyHostLocation, len(m.Locations))
	for i, v := range m.Locations {
		diags.Append(v.Save(ctx, &input.Locations[i])...)
	}

	return diags
}
