// resource_server.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/lrwx00t/lxd_provider/lxd"
)

func resourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"container_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"desired_status": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	containerName := d.Get("container_name").(string)
	desiredStatus := d.Get("desired_status").(string)

	d.SetId(containerName)

	err := lxd.ManageContainerWithName(containerName, desiredStatus)
	if err != nil {
		return err
	}

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
