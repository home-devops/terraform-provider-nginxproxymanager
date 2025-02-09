// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strconv"

	"terraform-provider-nginxproxymanager/internal/client/inputs"
	attributes "terraform-provider-nginxproxymanager/internal/provider/attributes/resources"
	"terraform-provider-nginxproxymanager/internal/provider/models"

	"terraform-provider-nginxproxymanager/internal/client"
	"terraform-provider-nginxproxymanager/internal/provider/common"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ common.IResource                    = &RedirectionHostResource{}
	_ resource.ResourceWithConfigure      = &RedirectionHostResource{}
	_ resource.ResourceWithValidateConfig = &RedirectionHostResource{}
	_ resource.ResourceWithImportState    = &RedirectionHostResource{}
)

func NewRedirectionHostResource() resource.Resource {
	b := &common.Resource{Name: "redirection_host"}
	r := &RedirectionHostResource{b, nil}
	b.IResource = r
	return r
}

type RedirectionHostResource struct {
	*common.Resource
	client *client.Client
}

func (r *RedirectionHostResource) SchemaImpl(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Hosts --- Manage a redirection host.",
		Attributes:  attributes.RedirectionHost,
	}
}

func (r *RedirectionHostResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *RedirectionHostResource) CreateImpl(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.RedirectionHost
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item inputs.RedirectionHost
	diags.Append(plan.Save(ctx, &item)...)

	redirectionHost, err := r.client.CreateRedirectionHost(ctx, &item)
	if err != nil {
		resp.Diagnostics.AddError("Error creating redirection host", "Could not create redirection host, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(plan.Load(ctx, redirectionHost)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *RedirectionHostResource) ReadImpl(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *models.RedirectionHost
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	redirectionHost, err := r.client.GetRedirectionHost(ctx, state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error reading proxy host", "Could not read redirection host, unexpected error: "+err.Error())
		return
	}
	if redirectionHost == nil {
		state = nil
	} else {
		resp.Diagnostics.Append(state.Load(ctx, redirectionHost)...)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *RedirectionHostResource) UpdateImpl(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.RedirectionHost
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item inputs.RedirectionHost
	diags.Append(plan.Save(ctx, &item)...)

	redirectionHost, err := r.client.UpdateRedirectionHost(ctx, plan.ID.ValueInt64Pointer(), &item)
	if err != nil {
		resp.Diagnostics.AddError("Error updating redirection host", "Could not update redirection host, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(plan.Load(ctx, redirectionHost)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *RedirectionHostResource) DeleteImpl(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.RedirectionHost
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteRedirectionHost(ctx, state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting redirection host", "Could not delete redirection host, unexpected error: "+err.Error())
		return
	}
}

func (r *RedirectionHostResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.RedirectionHost

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.SSLForced.ValueBool() && data.CertificateID.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("ssl_forced"),
			"Certificate ID is required when SSL is forced",
			"Certificate ID is required when SSL is forced")
	}

	if data.HTTP2Support.ValueBool() && data.CertificateID.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("http2_support"),
			"Certificate ID is required when HTTP/2 Support is enabled",
			"Certificate ID is required when HTTP/2 Support is enabled")
	}

	if data.HSTSEnabled.ValueBool() && !data.SSLForced.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hsts_enabled"),
			"SSL must be forced when HSTS is enabled",
			"SSL must be forced when HSTS is enabled")
	}

	if data.HSTSSubdomains.ValueBool() && !data.HSTSEnabled.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			path.Root("hsts_subdomains"),
			"HSTS must be enabled when HSTS Subdomains is enabled",
			"HSTS must be enabled when HSTS Subdomains is enabled")
	}
}

func (r *RedirectionHostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing redirection host", "Could not import redirection host, unexpected error: "+err.Error())
		return
	}

	diags := resp.State.SetAttribute(ctx, path.Root("id"), types.Int64Value(id))
	resp.Diagnostics.Append(diags...)
}
