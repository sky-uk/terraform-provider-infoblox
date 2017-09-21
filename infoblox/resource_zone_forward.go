package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceZoneForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneForwardCreate,
		Read:   resourceZoneForwardRead,
		Update: resourceZoneForwardUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Description: "The IPv4 Address or IPv6 Address of the server that is serving this zone. Not writable",
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment for the zone; maximum 256 characters",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"disable": {
				Type:        schema.TypeBool,
				Description: "Determines whether a zone is disabled or not",
				Optional:    true,
				Default:     false,
			},
			"display_domain": {
				Type:        schema.TypeString,
				Description: "The displayed name of the DNS zone.Not writable",
				Optional:    true,
				Computed:    true,
			},
			"dns_fqdn": {
				Type:        schema.TypeString,
				Description: "The name of this DNS zone in punycode format.For a reverse zone, this is in “address/cidr” format.For other zones, this is in FQDN format in punycode format.Cannot be updated",
				Optional:    true,
				Computed:    true,
			},
			"forward_to": util.ExternalServerListSchema(false, true),
			"forwarders_only": {
				Type:        schema.TypeBool,
				Description: "Determines if the appliance sends queries to forwarders only and not to other internal or Internet root servers",
				Optional:    true,
				Default:     false,
			},
			"forwarding_servers": util.ForwardingMemberServerListSchema(),
			"fqdn": {
				Type:         schema.TypeString,
				Description:  "The name of this DNS zone. For a reverse zone, this is in “address/cidr” format",
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"locked": {
				Type:        schema.TypeBool,
				Description: "If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked.The default value is False.",
				Optional:    true,
				Default:     false,
			},
			"locked_by": {
				Type:         schema.TypeString,
				Description:  "The name of a superuser or the administrator who locked this zone.Not writable",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"mask_prefix": {
				Type:         schema.TypeString,
				Description:  "IPv4 Netmask or IPv6 prefix for this zone.Not Writable",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"ns_group": {
				Type:         schema.TypeString,
				Description:  "A forwarding member name server group. Values with leading or trailing white space are not valid for this field. The default value is undefined.",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"parent": {
				Type:         schema.TypeString,
				Description:  "The parent zone of this zone. Note that when searching for reverse zones, the “in-addr.arpa” notation should be used. Not writable.",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"prefix": {
				Type:         schema.TypeString,
				Description:  "The RFC2317 prefix value of this DNS zone.",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"using_srg_associations": {
				Type:        schema.TypeBool,
				Description: "This is true if the zone is associated with a shared record group. Not writable",
				Optional:    true,
				Computed:    true,
			},
			"view": {
				Type:         schema.TypeString,
				Description:  "The name of the DNS view in which the zone resides",
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"zone_format": {
				Type:         schema.TypeString,
				Description:  "Determines the format of this zone - API default FORWARD. Cannot be updated.",
				ValidateFunc: util.ValidateZoneFormat,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
			},
		},
	}
}

func resourceZoneForwardCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.ZONEForwardObj, resourceZoneForward(), d, m)
}

func resourceZoneForwardRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceZoneForward(), d, m)
}

func resourceZoneForwardUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceZoneForward(), d, m)
}
