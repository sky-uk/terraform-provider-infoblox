package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
)

func resourceAdminRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdminRoleCreate,
		Read:   resourceAdminRoleRead,
		Update: resourceAdminRoleUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of an admin role.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The descriptive comment of the Admin Role object.",
			},
			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceAdminRoleCreate(d *schema.ResourceData, m interface{}) error {

	return CreateResource(model.AdminroleObj, resourceAdminRole(), d, m)
}

func resourceAdminRoleRead(d *schema.ResourceData, m interface{}) error {

	return ReadResource(resourceAdminRole(), d, m)
}

func resourceAdminRoleUpdate(d *schema.ResourceData, m interface{}) error {

	return UpdateResource(resourceAdminRole(), d, m)
}
