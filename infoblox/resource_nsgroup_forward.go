package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceNSGroupForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGroupForwardCreate,
		Read:   resourceNSGroupForwardRead,
		Update: resourceNSGroupForwardUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the name server group",
				Required:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment field",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"forwarding_servers": util.ForwardingMemberServerListSchema(),
		},
	}
}

func resourceNSGroupForwardCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.NsgroupForwardingmemberObj, resourceNSGroupForward(), d, m)
}

func resourceNSGroupForwardRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceNSGroupForward(), d, m)
}

func resourceNSGroupForwardUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceNSGroupForward(), d, m)
}
