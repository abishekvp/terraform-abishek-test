package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &SecurdenDataSource{}

func account_data_source() datasource.DataSource {
	return &SecurdenDataSource{}
}

type SecurdenDataSource struct {
	client *http.Client
}

type SecurdenDataSourceModel struct {
	Password types.String `tfsdk:"password"`
}

func (d *SecurdenDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_keyvalue"
}

func (d *SecurdenDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Securden data source",

		Attributes: map[string]schema.Attribute{
			"password": schema.StringAttribute{
				MarkdownDescription: "Name of the account",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (d *SecurdenDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*http.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client
}

func (d *SecurdenDataSource) Create(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var account securdenProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &account)...)
	accountid := account.AccountID.ValueString()
	data, code, message := get_account(ctx, accountid)
	if code != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("%d - %s", code, message), "")
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *SecurdenDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var account securdenProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &account)...)
	accountid := account.AccountID.ValueString()
	data, code, message := get_account(ctx, accountid)
	if code != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("%d - %s", code, message), "")
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
