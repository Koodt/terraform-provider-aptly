package aptly

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APTLY_TOKEN", nil),
				Description: descriptions["token"],
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APTLY_BASE_URL", "https://localhost:8989/api/"),
				Description: descriptions["base_url"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"aptly_repo":            dataSourceAptlyRepo(),
			"aptly_publish":         dataSourceRepoPublish(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)

	return p
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"token": "The OAuth token used to connect to Aptly. " +
			"`anonymous` mode is enabled if `token` is not configured.",

		"base_url": "The Aptly Base API URL",
	}
}
