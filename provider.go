package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tsuru/go-tsuruclient/pkg/client"
)

// Provider creates a new terraform provider with Tsuru resources
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tsuru_app": resourceTsuruApp(),
		},

		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			client, _ := client.ClientFromEnvironment(nil)
			return client, nil
		},
	}
}
