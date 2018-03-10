package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider creates a new terraform provider with Tsuru resources
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tsuru_app": resourceApp(),
		},
	}
}
