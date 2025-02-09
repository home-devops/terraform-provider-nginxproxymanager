// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"terraform-provider-nginxproxymanager/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ provider.Provider = &NginxProxyManagerProvider{}

type NginxProxyManagerProvider struct {
	Version string
}

type NginxProxyManagerProviderModel struct {
	URL      types.String `tfsdk:"url"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *NginxProxyManagerProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "npm"
	resp.Version = p.Version
}

func (p *NginxProxyManagerProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with Nginx Proxy Manager API.",
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				Description: "URL for Nginx Proxy Manager API. May also be provided via `NGINX_PROXY_MANAGER_URL` environment variable.",
				Optional:    true,
			},
			"username": schema.StringAttribute{
				Description: "Username for Nginx Proxy Manager API. May also be provided via `NGINX_PROXY_MANAGER_USERNAME` environment variable.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password for Nginx Proxy Manager API. May also be provided via `NGINX_PROXY_MANAGER_PASSWORD` environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *NginxProxyManagerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Nginx Proxy Manager provider")

	var config NginxProxyManagerProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.URL.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Unknown Nginx Proxy Manager API URL",
			"The provider cannot create the Nginx Proxy Manager API client as there is an unknown configuration value for the Nginx Proxy Manager API url. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the NGINX_PROXY_MANAGER_URL environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Nginx Proxy Manager API Username",
			"The provider cannot create the Nginx Proxy Manager API client as there is an unknown configuration value for the Nginx Proxy Manager API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the NGINX_PROXY_MANAGER_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Nginx Proxy Manager API Password",
			"The provider cannot create the Nginx Proxy Manager API client as there is an unknown configuration value for the Nginx Proxy Manager API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the NGINX_PROXY_MANAGER_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	url := os.Getenv("NGINX_PROXY_MANAGER_URL")
	username := os.Getenv("NGINX_PROXY_MANAGER_USERNAME")
	password := os.Getenv("NGINX_PROXY_MANAGER_PASSWORD")

	if !config.URL.IsNull() {
		url = config.URL.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if url == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("url"),
			"Missing Nginx Proxy Manager API URL",
			"The provider cannot create the Nginx Proxy Manager API client as there is a missing or empty value for the Nginx Proxy Manager API url. "+
				"Set the host value in the configuration or use the NGINX_PROXY_MANAGER_URL environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Nginx Proxy Manager API Username",
			"The provider cannot create the Nginx Proxy Manager API client as there is a missing or empty value for the Nginx Proxy Manager API username. "+
				"Set the username value in the configuration or use the NGINX_PROXY_MANAGER_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Nginx Proxy Manager API Password",
			"The provider cannot create the Nginx Proxy Manager API client as there is a missing or empty value for the Nginx Proxy Manager API password. "+
				"Set the password value in the configuration or use the NGINX_PROXY_MANAGER_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating Nginx Proxy Manager client")

	npmClient, err := client.NewClient(&url, &username, &password, p.Version)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Nginx Proxy Manager API Client",
			"An unexpected error occurred when creating the Nginx Proxy Manager API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Nginx Proxy Manager Client Error: "+err.Error(),
		)
		return
	}

	ctx = tflog.SetField(ctx, "npm_url", url)
	ctx = tflog.SetField(ctx, "npm_username", username)
	ctx = tflog.SetField(ctx, "npm_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "npm_password")

	resp.DataSourceData = npmClient
	resp.ResourceData = npmClient

	tflog.Info(ctx, "Configured Nginx Proxy Manager client", map[string]any{"success": true})
}

func (p *NginxProxyManagerProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewRedirectionHostResource,
		NewProxyHostResource,
		NewCertificateCustomResource,
		NewCertificateLetsEncryptResource,
	}
}

func (p *NginxProxyManagerProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewRedirectionHostDataSource,
		NewRedirectionHostsDataSource,
		NewProxyHostDataSource,
		NewProxyHostsDataSource,
		NewCertificateDataSource,
		NewCertificatesDataSource,
		NewAccessListDataSource,
		NewAccessListsDataSource,
		NewDeadHostDataSource,
		NewDeadHostsDataSource,
		NewStreamDataSource,
		NewStreamsDataSource,
		NewUserDataSource,
		NewUserMeDataSource,
		NewUsersDataSource,
		NewVersionDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NginxProxyManagerProvider{
			Version: version,
		}
	}
}
