package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceZoneDelegated() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneDelegatedCreate,
		Read:   resourceZoneDelegatedRead,
		Update: resourceZoneDelegateUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"reference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"view": {
				Type:        schema.TypeString,
				Description: "The name of the DNS view in which the zone resides",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment for the zone; maximum 256 characters",
				Optional:    true,
			},
			"delegate_to": util.ExternalServerListSchema(false, true),
			"delegated_ttl": {
				Type:        schema.TypeInt,
				Description: "a TTL for the delegated zone",
				Optional:    true,
				ForceNew:    false,
			},
			"disable": {
				Type:        schema.TypeBool,
				Description: "Is the zone disabled",
				Optional:    true,
				ForceNew:    false,
			},
			"fqdn": {
				Type:        schema.TypeString,
				Description: "The FQDN for the zone that is being delegated",
				Required:    true,
				ForceNew:    true,
			},
			"locked": {
				Type:        schema.TypeBool,
				Description: "Is the record locked to prevent changes",
				Optional:    true,
				ForceNew:    false,
			},
			"use_delegated_ttl": {
				Type:        schema.TypeBool,
				Description: "Should we use the deletated ttl",
				Optional:    true,
			},
			"zone_format": {
				Type:         schema.TypeString,
				Description:  "Format of the zone, default is FORWARD",
				Optional:     true,
				Default:      "FORWARD",
				ValidateFunc: util.ValidateZoneFormat,
			},
			"ns_group": {
				Type:        schema.TypeString,
				Description: "NameServer group for this zone",
				Optional:    true,
			},
		},
	}
}

func resourceZoneDelegatedCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.ZONEDelegatedObj, resourceZoneDelegated(), d, m)
}

func resourceZoneDelegatedRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceZoneDelegated(), d, m)
}

func resourceZoneDelegateUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceZoneDelegated(), d, m)
}
