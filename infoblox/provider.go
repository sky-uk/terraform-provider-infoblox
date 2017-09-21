package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"time"
)

// Provider : The infoblox terraform provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_USERNAME", nil),
				Description: "User to authenticate with Infoblox appliance",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_PASSWORD", nil),
				Description: "Password to authenticate with Infoblox appliance",
			},
			"server": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_SERVER", nil),
				Description: "Infoblox appliance to connect to eg https://192.168.0.1",
			},
			"allow_unverified_ssl": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_ALLOW_UNVERIFIED_SSL", false),
				Description: "If set, Infoblox client will permit unverifiable SSL certificates.",
			},
			"wapi_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Infoblox WAPI server version, defaults to v2.6.1",
				Default:     "v2.6.1",
			},
			"timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "http response timeout, in seconds",
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_CLIENT_TIMEOUT", 0),
			},
			"client_debug": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_CLIENT_DEBUG", false),
				Description: "infoblox client debug",
			},
		},
		ResourcesMap: map[string]*schema.Resource{

			"infoblox_cname_record":          resourceCNAMERecord(),
			"infoblox_arecord":               resourceARecord(),
			"infoblox_srv_record":            resourceSRVRecord(),
			"infoblox_txtrecord":             resourceTXTRecord(),
			"infoblox_network":               resourceNetwork(),
			"infoblox_zone_auth":             resourceZoneAuth(),
			"infoblox_dhcp_range":            resourceDHCPRange(),
			"infoblox_admin_user":            resourceAdminUser(),
			"infoblox_admin_group":           resourceAdminGroup(),
			"infoblox_admin_role":            resourceAdminRole(),
			"infoblox_ns_record":             resourceNSRecord(),
			"infoblox_zone_delegated":        resourceZoneDelegated(),
			"infoblox_permission":            resourcePermission(),
			"infoblox_zone_stub":             resourceZoneStub(),
			"infoblox_zone_forward":          resourceZoneForward(),
			"infoblox_ns_group_delegation":   resourceNSGroupDelegation(),
			"infoblox_ns_group_forward":      resourceNSGroupForward(),
			"infoblox_ns_group_stub":         resourceNSGroupStub(),
			"infoblox_ns_group_forward_stub": resourceNSGroupForwardStub(),
			"infoblox_mx_record":             resourceMxRecord(),
			"infoblox_dns_view":              resourceDNSView(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	var seconds int64
	seconds = int64(d.Get("timeout").(int))

	params := skyinfoblox.Params{
		WapiVersion: d.Get("wapi_version").(string),
		URL:         d.Get("server").(string),
		User:        d.Get("username").(string),
		Password:    d.Get("password").(string),
		IgnoreSSL:   d.Get("allow_unverified_ssl").(bool),
		Debug:       d.Get("client_debug").(bool),
		Timeout:     time.Duration(seconds),
	}

	client := skyinfoblox.Connect(params)

	outParams := make(map[string]interface{})
	outParams["ibxClient"] = client

	return outParams, nil
}
