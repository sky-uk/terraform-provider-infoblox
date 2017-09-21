package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceZoneStub() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneStubCreate,
		Read:   resourceZoneStubRead,
		Update: resourceZoneStubUpdate,
		Delete: DeleteResource,
		Schema: map[string]*schema.Schema{
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment for the zone; maximum 256 characters",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"disable": {
				Type:        schema.TypeBool,
				Description: "Is the zone disabled",
				Optional:    true,
				ForceNew:    false,
			},
			"locked": {
				Type:        schema.TypeBool,
				Description: "Is the record locked to prevent changes",
				Optional:    true,
				ForceNew:    false,
			},
			"disable_forwarding": {
				Type:        schema.TypeBool,
				Description: "Is forward disabled for this zone",
				Optional:    true,
				ForceNew:    false,
			},
			"external_nsgroup": {
				Type:         schema.TypeString,
				Description:  "Name of the external name server group",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"fqdn": {
				Type: schema.TypeString,
				Description: "Fqdn for the zone	",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"mask_prefix": {
				Type:         schema.TypeString,
				Description:  "IPv4 Netmask or IPv6 prefix for this zone.",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"nsgroup": {
				Type:         schema.TypeString,
				Description:  "Name of the  name server group",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"prefix": {
				Type:         schema.TypeString,
				Description:  "IPv4 Netmask or IPv6 prefix for this zone.",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"stub_from":    util.ExternalServerListSchema(true, false),
			"stub_members": util.MemberServerListSchema(true, false),
			"zone_format": {
				Type:         schema.TypeString,
				Description:  "Determines the format of this zone - API default FORWARD",
				ValidateFunc: util.ValidateZoneFormat,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
			},
			"view": {
				Type:         schema.TypeString,
				Description:  "The name of the DNS view in which the zone resides",
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
		},
	}
}

func resourceZoneStubCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.ZONESTUBObj, resourceZoneStub(), d, m)
}

func resourceZoneStubRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceZoneStub(), d, m)
}

func resourceZoneStubUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceZoneStub(), d, m)
}
