package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkCreate,
		Read:   resourceNetworkRead,
		Update: resourceNetworkUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"network": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The network address in IPv4 Address/CIDR format.",
			},
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the network view in which this network resides.",
			},
			"comment": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Comment for the network, maximum 256 characters.",
				ValidateFunc: util.ValidateMaxLength(256),
			},
			"authority": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Authority for the DHCP network. Associated with the field use_authority",
			},
			"use_authority": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"auto_create_reversezone": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "This flag controls whether reverse zones are automatically created when the network is added. Cannot be updated, nor is readable",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines whether a network is disabled or not. When this is set to False, the network is enabled.",
			},
			"enable_ddns": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The dynamic DNS updates flag of a DHCP network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
			},
			"use_enable_ddns": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"high_water_mark": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "The percentage of DHCP network usage threshold above which network usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
			},
			"high_water_mark_reset": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "The percentage of DHCP network usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.",
			},
			"low_water_mark": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "The percentage of DHCP network usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
			},
			"low_water_mark_reset": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "The percentage of DHCP network usage threshold below which network usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled).  A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
			},
			"enable_dhcp_thresholds": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if DHCP thresholds are enabled for the network.",
			},
			"use_enable_dhcp_thresholds": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_discovery": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines whether a discovery is enabled or not for this network. When this is set to False, the network discovery is disabled.",
			},
			"use_enable_discovery": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"discovery_member": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The member that will run discovery for this network.",
			},
			"ipv4addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The IPv4 Address of the network.",
			},
			"lease_scavenge_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
			},
			"members": {
				Type:        schema.TypeList,
				Description: "DHCP Member which is going to serve this network.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4addr": {
							Type:        schema.TypeString,
							Description: "IPv4 address of the member pair",
							Optional:    true,
						},
						"ipv6addr": {
							Type:        schema.TypeString,
							Description: "IPv6 address of the member pair",
							Optional:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "FQDN of the member pair",
							Optional:    true,
						},
					},
				},
			},
			"netmask": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Number of bits in the network mask example: 8,16,24 etc ",
			},
			"network_container": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The network container to which this network belongs (if any). Cannot be updated nor written",
			},
			"options": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "DHCP Related] Options such as DNS servers, gateway, ntp, etc",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "DHCP Option Name",
							Optional:    true,
						},
						"num": {
							Type:         schema.TypeInt,
							Description:  "DHCO Option number",
							Optional:     true,
							ValidateFunc: util.ValidateUnsignedInteger,
						},
						"use_option": {
							Type:        schema.TypeBool,
							Description: "Use the option or not",
							Optional:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Value of the option. For an option this value is required",
							Optional:    true,
						},
						"vendor_class": {
							Type:        schema.TypeString,
							Description: "Vendor Class",
							Default:     "DHCP",
							Optional:    true,
						},
					},
				},
			},
			"use_options": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"recycle_leases": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted.",
			},
			"use_recycle_leases": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"restart_if_needed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Restarts the member service. Not readable",
			},
			"update_dns_on_lease_renewal": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
			},
		},
	}
}

func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.NetworkObj, resourceNetwork(), d, m)
}

func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceNetwork(), d, m)
}

func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceNetwork(), d, m)
}
