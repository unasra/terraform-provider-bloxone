package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
)

// Ensure BloxOneProvider satisfies various provider interfaces.
var _ provider.Provider = &BloxOneProvider{}

// BloxOneProvider defines the provider implementation.
type BloxOneProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
	commit  string
}

// BloxOneProviderModel describes the provider data model.
type BloxOneProviderModel struct {
	CSPUrl string `tfsdk:"csp_url"`
	APIKey string `tfsdk:"api_key"`
}

func (p *BloxOneProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "bloxone"
	resp.Version = p.version
}

func (p *BloxOneProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"csp_url": schema.StringAttribute{
				MarkdownDescription: "URL for BloxOne Cloud Services Portal. Can also be configured using the `BLOXONE_CSP_URL` environment variable.",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API key for accessing the BloxOne API. Can also be configured by using the `BLOXONE_API_KEY` environment variable. https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys",
				Required:            true,
			},
		},
	}
}

func (p *BloxOneProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data BloxOneProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client, err := bloxoneclient.NewAPIClient(bloxoneclient.Configuration{
		ClientName: fmt.Sprintf("terraform/%s#%s", p.version, p.commit),
		APIKey:     data.APIKey,
	})
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("Unable to create new API client: %s", err))
	}

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *BloxOneProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *BloxOneProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version, commit string) func() provider.Provider {
	return func() provider.Provider {
		return &BloxOneProvider{
			version: version,
			commit:  commit,
		}
	}
}