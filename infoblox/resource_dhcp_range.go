package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
)

func resourceDHCPRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPRangeCreate,
		Read:   resourceDHCPRangeRead,
		Update: resourceDHCPRangeUpdate,
		Delete: DeleteResource,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "This field contains the name of the Microsoft scope.",
			},
			"comment": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.ValidateMaxLength(256),
			},
			"network": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The network to which this range belongs, in IPv4 Address/CIDR format.",
			},
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "The name of the network view in which this range resides.",
			},
			"start_addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IPv4 Address starting address of the range.",
			},
			"end_addr": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IPv4 Address end address of the range.",
			},
			"member": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Infoblox DHCP member that serves this range",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4addr": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The IPv4 Address of the Grid Member.",
							Optional:    true,
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Description: "DHCP Member server FQDN",
							Optional:    true,
						},
					},
				},
			},
			"restart_if_needed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Restarts any services if required by this change. Default: true.",
				Default:     false,
			},
			"server_association_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Must be set to 'MEMBER' if member is specified",
				Default:     "NONE",
			},
		},
	}
}

func resourceDHCPRangeCreate(d *schema.ResourceData, m interface{}) error {
	return CreateResource(model.RangeObj, resourceDHCPRange(), d, m)
}

func resourceDHCPRangeRead(d *schema.ResourceData, m interface{}) error {
	return ReadResource(resourceDHCPRange(), d, m)
}

func resourceDHCPRangeUpdate(d *schema.ResourceData, m interface{}) error {
	return UpdateResource(resourceDHCPRange(), d, m)
}
