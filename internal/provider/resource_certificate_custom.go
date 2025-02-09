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
	_ common.IResource                    = &CertificateCustomResource{}
	_ resource.ResourceWithConfigure      = &CertificateCustomResource{}
	_ resource.ResourceWithValidateConfig = &CertificateCustomResource{}
	_ resource.ResourceWithImportState    = &CertificateCustomResource{}
)

func NewCertificateCustomResource() resource.Resource {
	b := &common.Resource{Name: "certificate_custom"}
	r := &CertificateCustomResource{b, nil}
	b.IResource = r
	return r
}

type CertificateCustomResource struct {
	*common.Resource
	client *client.Client
}

func (r *CertificateCustomResource) SchemaImpl(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "SSL Certificates --- Manage a custom certificate.",
		Attributes:  attributes.CertificateCustom,
	}
}

func (r *CertificateCustomResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*client.Client)
}

func (r *CertificateCustomResource) CreateImpl(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.CertificateCustom
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item inputs.CertificateCustom
	diags.Append(plan.Save(ctx, &item)...)

	certificate, err := r.client.CreateCertificateCustom(ctx, &item)
	if err != nil {
		resp.Diagnostics.AddError("Error creating certificate custom", "Could not create certificate custom, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(plan.Load(ctx, certificate)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertificateCustomResource) ReadImpl(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *models.CertificateCustom
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	certificate, err := r.client.GetCertificateCustom(ctx, state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error reading certificate custom", "Could not read certificate custom, unexpected error: "+err.Error())
		return
	}
	if certificate == nil {
		state = nil
	} else if certificate.Provider != "other" {
		resp.Diagnostics.AddError("Error reading certificate custom", "Certificate is not a custom certificate.")
	} else {
		resp.Diagnostics.Append(state.Load(ctx, certificate)...)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertificateCustomResource) UpdateImpl(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// There is no update method for certificates, so we delete and recreate
}

func (r *CertificateCustomResource) DeleteImpl(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *models.CertificateCustom
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCertificate(ctx, state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting certificate custom", "Could not delete certificate custom, unexpected error: "+err.Error())
		return
	}
}

func (r *CertificateCustomResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.CertificateCustom

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertificateCustomResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing certificate custom", "Could not import certificate custom, unexpected error: "+err.Error())
		return
	}

	diags := resp.State.SetAttribute(ctx, path.Root("id"), types.Int64Value(id))
	resp.Diagnostics.Append(diags...)
}
