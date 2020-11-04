package aptly

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/Koodt/terraform-provider-aptly/client"
)

func dataSourceAptlyRepos() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAptlyReposRead,

		Schema: map[string]*schema.Schema{
			"DefaultComponent": {
				Type:     schema.TypeString,
				Required: true,
			},
			"DefaultDistribution": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"Name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"Comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAptlyReposRead(d *schema.ResourceData, meta interface{}) error {
	aptlyClient := meta.(*client.Client)
	name := d.Get("name").(string)

	d.Set("name", String())

	return nil
}
