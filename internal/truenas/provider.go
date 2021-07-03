package truenas

import (
	"context"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/pool"

	"github.com/xonvanetta/terraform-provider-truenas/internal/truenas/sharing"

	v2 "github.com/xonvanetta/terraform-provider-truenas/internal/truenas/api/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRUENAS_HOST", "http://localhost"),
			},
			"api_key": {
				Type:        schema.TypeString,
				Sensitive:   true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TRUENAS_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"truenas_sharing_nfs":  sharing.NFSResource(),
			"truenas_pool_dataset": pool.DatasetResource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"truenas_sharing_nfs":  sharing.DataNFSResource(),
			"truenas_pool_dataset": pool.DataDatasetResource(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	apiKey := d.Get("api_key").(string)

	if apiKey == "" {
		return nil, diag.Errorf("api_key must be set to use this provider")
	}

	return v2.NewClient(host, apiKey), nil
}
