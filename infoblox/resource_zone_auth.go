package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceZoneAuth() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneAuthCreate,
		Read:   resourceZoneAuthRead,
		Update: resourceZoneAuthUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"fqdn": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of this DNS zone. For a reverse zone, this is in “address/cidr” format",
				ForceNew:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"view": {
				Type:         schema.TypeString,
				Description:  "The name of the DNS view in which the zone resides",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment for the zone; maximum 256 characters",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"zone_format": {
				Type:         schema.TypeString,
				Description:  "Determines the format of this zone - API default FORWARD",
				ValidateFunc: util.ValidateZoneFormat,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
			},
			"restart_if_needed": {
				Type:        schema.TypeBool,
				Description: "Restarts the member service. The default value is False. Not readable",
				Optional:    true,
				Computed:    true,
			},
			"prefix": {
				Type:         schema.TypeString,
				Description:  "The RFC2317 prefix value of this DNS zone",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"disable": {
				Type:        schema.TypeBool,
				Description: "Determines whether a zone is disabled or not",
				Optional:    true,
				Computed:    true,
			},
			"dns_integrity_enable": {
				Type:        schema.TypeBool,
				Description: "If this is set to True, DNS integrity check is enabled for this zone",
				Optional:    true,
				Computed:    true,
			},
			"dns_integrity_member": {
				Type:        schema.TypeString,
				Description: "The Grid member that performs DNS integrity checks for this zone",
				Optional:    true,
				Computed:    true,
			},

			"external_primaries":   util.ExternalServerListSchema(true, false),
			"external_secondaries": util.ExternalServerListSchema(true, false),
			"grid_primary":         util.MemberServerListSchema(true, false),
			"grid_primary_shared_with_ms_parent_delegation": {
				Type:        schema.TypeBool,
				Description: "Determines if the server is duplicated with parent delegation.cannot be updated, nor written",
				Optional:    true,
				Computed:    true,
			},
			"grid_secondaries": util.MemberServerListSchema(true, false),
			"locked": {
				Type:        schema.TypeBool,
				Description: "If you enable this flag, other administrators cannot make conflicting changes",
				Optional:    true,
				Computed:    true,
			},
			"locked_by": {
				Type:        schema.TypeString,
				Description: "The name of a superuser or the administrator who locked this zone (read-only)",
				Computed:    true,
			},
			"network_view": {
				Type:        schema.TypeString,
				Description: "The name of the network view in which this zone resides (read-only)",
				Computed:    true,
			},
			"ns_group": {
				Type:        schema.TypeString,
				Description: "The name server group that serves DNS for this zone.",
				Optional:    true,
			},
			"scavenging_settings": {
				Type:        schema.TypeMap,
				Description: "The DNS scavenging settings object provides information about scavenging configuration e.g. conditions under which records can be scavenged, periodicity of scavenging operations.",
				Optional:    true,
			},
			"soa_serial_number": {
				Type:        schema.TypeInt,
				Description: "The serial number in the SOA record incrementally changes every time the record is modified. The SOA serial number to be used in conjunction with set_soa_serial_number (read-only)",
				Optional:    true,
				Computed:    true,
			},
			"soa_default_ttl": {
				Type:         schema.TypeInt,
				Description:  "The Time to Live (TTL) value of the SOA record of this zone",
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soa_negative_ttl": {
				Type:         schema.TypeInt,
				Description:  "The negative Time to Live (TTL)",
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soa_refresh": {
				Type:         schema.TypeInt,
				Description:  "This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not",
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soa_retry": {
				Type:         schema.TypeInt,
				Description:  "This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs",
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soa_expire": {
				Type:        schema.TypeInt,
				Description: "This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. The default is one week.",
				Optional:    true,
			},
			"copy_xfer_to_notify": {
				Type:        schema.TypeBool,
				Description: "If this flag is set to True then copy allowed IPs from Allow Transfer to Also Notify.",
				Optional:    true,
			},
			"use_copy_xfer_to_notify": {
				Type:        schema.TypeBool,
				Description: "Use flag for: copy_xfer_to_notify.",
				Optional:    true,
				Computed:    true,
			},
			"use_check_names_policy": {
				Type:        schema.TypeBool,
				Description: "Apply policy to dynamic updates and inbound zone transfers (This value applies only if the host name restriction policy is set to “Strict Hostname Checking”.)",
				Optional:    true,
				Computed:    true,
			},
			"use_external_primary": {
				Type:        schema.TypeBool,
				Description: "This flag controls whether the zone is using an external primary.",
				Optional:    true,
				Computed:    true,
			},
			"allow_update":   util.AccessControlSchema(),
			"allow_transfer": util.AccessControlSchema(),
			"use_allow_transfer": {
				Type:        schema.TypeBool,
				Description: "allow_transfer",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceZoneAuthCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.ZONEAUTHObj, resourceZoneAuth(), d, m)
}

func resourceZoneAuthRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceZoneAuth(), d, m)
}

func resourceZoneAuthUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceZoneAuth(), d, m)
}
