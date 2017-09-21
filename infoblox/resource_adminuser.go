package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
)

func resourceAdminUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdminUserCreate,
		Read:   resourceAdminUserRead,
		Update: resourceAdminUserUpdate,
		Delete: DeleteResource,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name for the user",
			},
			"admin_groups": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The admin_groups the user belongs to , there can be only 1 ",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email address for the user",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Should the user be disabled",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "a comment on the user",
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceAdminUserCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.AdminuserObj, resourceAdminUser(), d, m)
}

func resourceAdminUserRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceAdminUser(), d, m)
}

func resourceAdminUserUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceAdminUser(), d, m)
}
