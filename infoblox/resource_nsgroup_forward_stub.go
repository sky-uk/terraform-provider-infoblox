package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceNSGroupForwardStub() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGroupForwardStubCreate,
		Read:   resourceNSGroupForwardStubRead,
		Update: resourceNSGroupForwardStubUpdate,
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
			// Note: this resource only makes use of name and address in the external server struct. Other attributes are ignored.
			"external_servers": util.ExternalServerSetSchema(false, true),
		},
	}
}

func resourceNSGroupForwardStubCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.NsgroupForwardstubserverObj, resourceNSGroupForwardStub(), d, m)
}

func resourceNSGroupForwardStubRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceNSGroupForwardStub(), d, m)
}

func resourceNSGroupForwardStubUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceNSGroupForwardStub(), d, m)
}
