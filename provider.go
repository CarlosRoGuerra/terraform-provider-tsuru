package main

import (
	"github.com/go-resty/resty"
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider creates a new terraform provider with Tsuru resources
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TSURU_TOKEN", nil),
			},
			"target": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TSURU_TARGET", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"tsuru_app": resourceTsuruApp(),
		},

		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			token := d.Get("token").(string)
			target := d.Get("target").(string)

			client := resty.New()
			client.SetDebug(true)
			client.SetHostURL(target)
			client.SetAuthToken(token)
			client.SetHeader("Content-Type", "application/x-www-form-urlencoded")
			return client, nil
		},
	}
}
