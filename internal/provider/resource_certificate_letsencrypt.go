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
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ common.IResource                    = &CertificateLetsEncryptResource{}
	_ resource.ResourceWithConfigure      = &CertificateLetsEncryptResource{}
	_ resource.ResourceWithValidateConfig = &CertificateLetsEncryptResource{}
	_ resource.ResourceWithImportState    = &CertificateLetsEncryptResource{}
)

func NewCertificateLetsEncryptResource() resource.Resource {
	b := &common.Resource{Name: "certificate_letsencrypt"}
	r := &CertificateLetsEncryptResource{b, nil}
	b.IResource = r
	return r
}

type CertificateLetsEncryptResource struct {
	*common.Resource
	client *client.Client
}

func (r *CertificateLetsEncryptResource) SchemaImpl(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "SSL Certificates --- Manage a letsencrypt certificate.",
		Attributes:  attributes.CertificateLetsEncrypt,
	}
}

func (r *CertificateLetsEncryptResource) Configure(ctx context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	var ok bool
	r.client, ok = req.ProviderData.(*client.Client)
	if !ok {
		tflog.Error(ctx, "ProviderData is not of type *client.Client")
		return
	}
}

func (r *CertificateLetsEncryptResource) CreateImpl(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.CertificateLetsEncrypt
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var item inputs.CertificateLetsEncrypt
	diags.Append(plan.Save(ctx, &item)...)

	certificate, err := r.client.CreateCertificateLetsEncrypt(ctx, &item)
	if err != nil {
		resp.Diagnostics.AddError("Error creating certificate letsencrypt", "Could not create certificate letsencrypt, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(plan.Load(ctx, certificate)...)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertificateLetsEncryptResource) ReadImpl(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *models.CertificateLetsEncrypt
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	certificate, err := r.client.GetCertificateLetsEncrypt(ctx, state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error reading certificate letsencrypt", "Could not read certificate letsencrypt, unexpected error: "+err.Error())
		return
	}
	if certificate == nil {
		state = nil
	} else if certificate.Provider != "letsencrypt" {
		resp.Diagnostics.AddError("Error reading certificate letsencrypt", "Certificate is not a letsencrypt certificate.")
	} else {
		resp.Diagnostics.Append(state.Load(ctx, certificate)...)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertificateLetsEncryptResource) UpdateImpl(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// There is no update method for certificates, so we delete and recreate
}

func (r *CertificateLetsEncryptResource) DeleteImpl(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *models.CertificateLetsEncrypt
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteCertificate(ctx, state.ID.ValueInt64Pointer())
	if err != nil {
		resp.Diagnostics.AddError("Error deleting certificate letsencrypt", "Could not delete certificate letsencrypt, unexpected error: "+err.Error())
		return
	}
}

func (r *CertificateLetsEncryptResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data models.CertificateLetsEncrypt

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CertificateLetsEncryptResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Error importing certificate letsencrypt", "Could not import certificate letsencrypt, unexpected error: "+err.Error())
		return
	}

	diags := resp.State.SetAttribute(ctx, path.Root("id"), types.Int64Value(id))
	resp.Diagnostics.Append(diags...)
}
