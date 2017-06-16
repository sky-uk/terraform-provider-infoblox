package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
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
			"client_debug": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_CLIENT_DEBUG", false),
				Description: "infoblox client debug",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"infoblox_cname_record": resourceCNAMERecord(),
			"infoblox_arecord": resourceARecord(),
		},
		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	username := d.Get("username").(string)
	password := d.Get("password").(string)
	server := d.Get("server").(string)
	ignoreSSL := d.Get("allow_unverified_ssl").(bool)
	clientDebug := d.Get("client_debug").(bool)

	ibxClient := skyinfoblox.NewInfobloxClient(server, username, password, ignoreSSL, clientDebug)

	return ibxClient, nil
}
