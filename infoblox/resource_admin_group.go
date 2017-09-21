package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
)

func resourceAdminGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdminGroupCreate,
		Read:   resourceAdminGroupRead,
		Update: resourceAdminGroupUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Admin Group",
				Required:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment field",
				Optional:    true,
			},
			"superuser": {
				Type:        schema.TypeBool,
				Description: "Whether the group is a super user group or not",
				Optional:    true,
				Computed:    true,
			},
			"disable": {
				Type:        schema.TypeBool,
				Description: "Whether the Admin Group is disabled or not",
				Optional:    true,
				Computed:    true,
			},
			"access_method": {
				Type:        schema.TypeList,
				Description: "Methods the group can use to access Infoblox",
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"email_addresses": {
				// Need to use TypeSet as the read order doesn't match the sent order.
				Type:        schema.TypeSet,
				Description: "List of email addresses to associated with the Admin Group",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"roles": {
				Type:        schema.TypeList,
				Description: "List of roles to associated with the Admin Group",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceAdminGroupCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.AdmingroupObj, resourceAdminGroup(), d, m)
}

func resourceAdminGroupRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceAdminGroup(), d, m)
}

func resourceAdminGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceAdminGroup(), d, m)
}
