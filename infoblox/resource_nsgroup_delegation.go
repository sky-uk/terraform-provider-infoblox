package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceNSGroupDelegation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGroupDelegationCreate,
		Read:   resourceNSGroupDelegationRead,
		Update: resourceNSGroupDelegationUpdate,
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
			// Note: Only name and address in the external server struct are used by this resource.
			"delegate_to": util.ExternalServerSetSchema(false, true),
		},
	}
}

func resourceNSGroupDelegationCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.NsgroupDelegationObj, resourceNSGroupDelegation(), d, m)
}

func resourceNSGroupDelegationRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceNSGroupDelegation(), d, m)
}

func resourceNSGroupDelegationUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceNSGroupDelegation(), d, m)
}
