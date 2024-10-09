package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ provider.Provider = &securdenProvider{}
var _ provider.ProviderWithFunctions = &securdenProvider{}

type securdenProvider struct {
	version string
}

var AccountID string
var ServerURL string
var Authtoken string

type securdenProviderModel struct {
	AccountID types.String `tfsdk:"account_id"`
	ServerURL types.String `tfsdk:"server_url"`
	Authtoken types.String `tfsdk:"authtoken"`
}

func (p *securdenProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "abiraj"
	resp.Version = p.version
}

func (p *securdenProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"authtoken": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
			},
			"server_url": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (p *securdenProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring abiraj client")
	var config securdenProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	AccountID = config.AccountID.ValueString()
	ServerURL = config.ServerURL.ValueString()
	Authtoken = config.Authtoken.ValueString()
}

func (p *securdenProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *securdenProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		account_data_source,
	}
}

func (p *securdenProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{}
}

func Provider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &securdenProvider{
			version: version,
		}
	}
}
