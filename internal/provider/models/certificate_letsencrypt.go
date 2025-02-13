// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package models

import (
	"context"
	"encoding/json"
	"strings"
	"terraform-provider-nginxproxymanager/internal/client/inputs"
	"terraform-provider-nginxproxymanager/internal/client/resources"
	"terraform-provider-nginxproxymanager/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificateLetsEncrypt struct {
	ID         types.Int64  `tfsdk:"id"`
	CreatedOn  types.String `tfsdk:"created_on"`
	ModifiedOn types.String `tfsdk:"modified_on"`

	Name                   types.String `tfsdk:"name"`
	DomainNames            types.List   `tfsdk:"domain_names"`
	ExpiresOn              types.String `tfsdk:"expires_on"`
	DnsChallenge           types.Bool   `tfsdk:"dns_challenge"`
	DnsProvider            types.String `tfsdk:"dns_provider"`
	DnsProviderCredentials types.String `tfsdk:"dns_provider_credentials"`
	LetsEncryptAgree       types.Bool   `tfsdk:"letsencrypt_agree"`
	LetsEncryptEmail       types.String `tfsdk:"letsencrypt_email"`
}

func (m *CertificateLetsEncrypt) Load(ctx context.Context, resource *resources.Certificate) diag.Diagnostics {
	var diags diag.Diagnostics

	m.ID = types.Int64Value(resource.ID)
	m.CreatedOn = types.StringValue(resource.CreatedOn)
	m.ModifiedOn = types.StringValue(resource.ModifiedOn)
	m.Name = types.StringValue(resource.NiceName)
	m.ExpiresOn = types.StringValue(resource.ExpiresOn)
	dnsChallenge := strings.Trim(strings.ReplaceAll(resource.Meta.Map()["dns_challenge"], "\\n", "\n"), "\"")
	m.DnsChallenge = types.BoolValue(dnsChallenge == "true")
	m.DnsProvider = types.StringValue(strings.Trim(strings.ReplaceAll(resource.Meta.Map()["dns_provider"], "\\n", "\n"), "\""))

	var dnsCreds string
	err := json.Unmarshal([]byte(resource.Meta.Map()["dns_provider_credentials"]), &dnsCreds)
	if err != nil {
		diags.AddError("Failed to unmarshal dns_provider_credentials", err.Error())
	}

	m.DomainNames, diags = types.ListValueFrom(ctx, types.StringType, resource.DomainNames)

	if diags.HasError() {
		return diags
	}

	m.DnsProviderCredentials = types.StringValue(dnsCreds)
	letsEncryptAgree := strings.Trim(strings.ReplaceAll(resource.Meta.Map()["letsencrypt_agree"], "\\n", "\n"), "\"")
	m.LetsEncryptAgree = types.BoolValue(letsEncryptAgree == "true")
	m.LetsEncryptEmail = types.StringValue(strings.Trim(strings.ReplaceAll(resource.Meta.Map()["letsencrypt_email"], "\\n", "\n"), "\""))

	return diags
}

func (m *CertificateLetsEncrypt) Save(ctx context.Context, input *inputs.CertificateLetsEncrypt) diag.Diagnostics {
	var diags diag.Diagnostics

	input.Meta.DnsChallenge = m.DnsChallenge.ValueBool()
	input.Meta.DnsProvider = m.DnsProvider.ValueString()
	input.Meta.DnsProviderCredentials = m.DnsProviderCredentials.ValueString()
	input.Meta.LetsEncryptAgree = m.LetsEncryptAgree.ValueBool()
	input.Meta.LetsEncryptEmail = m.LetsEncryptEmail.ValueString()

	input.DomainNames, diags = utils.ConvertListToStringSlice(m.DomainNames)

	if diags.HasError() {
		return diags
	}

	return diags
}
