package main

import (
	"github.com/google/go-querystring/query"

	"github.com/go-resty/resty"

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

type tsuruApp struct {
	Name        string `url:"Name,omitempty"`
	Platform    string `url:"Platform,omitempty"`
	TeamOwner   string `url:"TeamOwner,omitempty"`
	Plan        string `url:"Plan,omitempty"`
	Router      string `url:"Router,omitempty"`
	Pool        string `url:"Pool,omitempty"`
	Description string `url:"Description,omitempty"`
}

func resourceAppCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*resty.Client)
	appData := &tsuruApp{
		Name:        d.Get("name").(string),
		Platform:    d.Get("platform").(string),
		TeamOwner:   d.Get("team").(string),
		Plan:        d.Get("plan").(string),
		Router:      d.Get("router").(string),
		Pool:        d.Get("pool").(string),
		Description: d.Get("description").(string),
	}

	body, _ := query.Values(appData)
	_, err := client.R().SetBody(body.Encode()).Post("/apps")
	if err != nil {
		return err
	}

	d.SetId(appData.Name)
	return nil
}

func resourceAppRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAppUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAppDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*resty.Client)
	appName := d.Get("name").(string)

	_, err := client.R().Delete("/apps/" + appName)
	if err != nil {
		return err
	}
	return nil
}
