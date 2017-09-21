package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceDNSView() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSViewCreate,
		Read:   resourceDNSViewRead,
		Update: resourceDNSViewUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of the DNS view.",
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"comment": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Comment for the DNS view; maximum 64 characters.",
				ValidateFunc: util.ValidateMaxLength(64),
			},
			"is_default": {
				Type:        schema.TypeBool,
				Description: "The NIOS appliance provides one default DNS view. You can rename the default view and change its settings, but you cannot delete it. There must always be at least one DNS view in the appliance.",
				Computed:    true,
			},
		},
	}
}

func resourceDNSViewCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.ViewObj, resourceDNSView(), d, m)
}

func resourceDNSViewRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceDNSView(), d, m)
}

func resourceDNSViewUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceDNSView(), d, m)
}
