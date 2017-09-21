package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceNSGroupStub() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGroupStubCreate,
		Read:   resourceNSGroupStubRead,
		Update: resourceNSGroupStubUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Description:  "The name of the name server group",
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment field",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			// Note the lead, stealth, grid_replicate, preferred_primaries, override_preferred_primaries attributes are ignored by Infoblox when set for this object type.
			// The field is required on creation.
			"stub_members": util.MemberServerListSchema(false, true),
		},
	}
}

func resourceNSGroupStubCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.NsgroupStubmemberObj, resourceNSGroupStub(), d, m)
}

func resourceNSGroupStubRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceNSGroupStub(), d, m)
}

func resourceNSGroupStubUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceNSGroupStub(), d, m)
}
