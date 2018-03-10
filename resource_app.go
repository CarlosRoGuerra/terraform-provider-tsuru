package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,

		Schema: map[string]*schema.Schema{
			"appname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"teamOwner": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"plan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"router": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAppCreate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceAppRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAppUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAppDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
