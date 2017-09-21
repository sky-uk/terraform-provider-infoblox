package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"regexp"
)

func resourcePermission() *schema.Resource {
	return &schema.Resource{
		Create: resourcePermissionCreate,
		Read:   resourcePermissionRead,
		Update: resourcePermissionUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"group": {
				Type:        schema.TypeString,
				Description: "The name of the admin group this permission applies to.",
				Optional:    true,
			},
			"object": {
				Type:        schema.TypeString,
				Description: "A reference to a WAPI object, which will be the object this permission applies to.",
				Optional:    true,
				Computed:    true,
			},
			"permission": {
				Type:         schema.TypeString,
				Description:  "The type of permission.",
				Required:     true,
				ValidateFunc: validatePermissionType,
			},
			"resource_type": {
				Type:        schema.TypeString,
				Description: "The type of resource this permission applies to.",
				Optional:    true,
				ForceNew:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "The name of the role this permission applies to.",
				Optional:    true,
			},
		},
	}
}

func validatePermissionType(v interface{}, k string) (ws []string, errors []error) {
	permission := v.(string)
	permissionOptions := regexp.MustCompile(`^(DENY|READ|WRITE)$`)
	if !permissionOptions.MatchString(permission) {
		errors = append(errors, fmt.Errorf("%q must be one of DENY, READ or WRITE", k))
	}
	return
}

func resourcePermissionCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.PermissionObj, resourcePermission(), d, m)
}

func resourcePermissionRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourcePermission(), d, m)
}

func resourcePermissionUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourcePermission(), d, m)
}
