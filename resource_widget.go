package nginxproxymanager

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWidget() *schema.Resource {
	return &schema.Resource{
		Create: resourceWidgetCreate,
		Read:   resourceWidgetRead,
		Update: resourceWidgetUpdate,
		Delete: resourceWidgetDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceWidgetCreate(d *schema.ResourceData, m interface{}) error {
	// Implement logic to create a widget.
	d.SetId("widget-id") // Set a dummy ID for now.
	return nil
}

func resourceWidgetRead(d *schema.ResourceData, m interface{}) error {
	// Implement logic to read a widget.
	return nil
}

func resourceWidgetUpdate(d *schema.ResourceData, m interface{}) error {
	// Implement logic to update a widget.
	return nil
}

func resourceWidgetDelete(d *schema.ResourceData, m interface{}) error {
	// Implement logic to delete a widget.
	return nil
}
