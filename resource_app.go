package main

import (
	"github.com/tsuru/go-tsuruclient/pkg/tsuru"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTsuruApp() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppCreate,
		Read:   resourceAppRead,
		Update: resourceAppUpdate,
		Delete: resourceAppDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"team": &schema.Schema{
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
	client := m.(*tsuru.APIClient)
	appData := tsuru.App{
		Name: d.Get("name").(string),
		// Tag:
		Router: d.Get("router").(string),
		// Routeropts
		Plan:        d.Get("plan").(string),
		Pool:        d.Get("pool").(string),
		Platform:    d.Get("platform").(string),
		Description: d.Get("description").(string),
		TeamOwner:   d.Get("team").(string),
	}

	_, _, err := client.AppApi.AppCreate(nil, appData)
	if err != nil {
		return err
	}

	d.SetId(appData.Name)
	return nil
}

func resourceAppRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*tsuru.APIClient)
	appName := d.Get("name").(string)

	retrievedApp, response, err := client.AppApi.AppGet(nil, appName)
	if response.StatusCode == 404 {
		d.SetId("")
		return err
	}

	d.Set("name", retrievedApp.Name)
	d.Set("router", retrievedApp.Router)
	d.Set("plan", retrievedApp.Plan)
	d.Set("platform", retrievedApp.Platform)
	d.Set("description", retrievedApp.Description)
	d.Set("teamowner", retrievedApp.TeamOwner)

	return nil
}

func resourceAppUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAppDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*tsuru.APIClient)
	appName := d.Get("name").(string)

	_, err := client.AppApi.AppDelete(nil, appName)
	if err != nil {
		return err
	}
	return nil
}
