package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
	"net/http"
)

func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkCreate,
		Read:   resourceNetworkRead,
		Update: resourceNetworkUpdate,
		Delete: resourceNetworkDelete,

		Schema: map[string]*schema.Schema{
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique reference to Infoblox Network resource",
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"networkview": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authority": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"autocreatereversezone": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Disable Network for DHCP",
			},
			"enableddns": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabledhcpthresholds": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enablediscovery": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ipv4addr": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"leasescavengetime": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"netmask": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Number of bits in the network mask example: 8,16,24 etc ",
			},
			"networkcontainer": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"option": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "DHCP Related] Options such as DNS servers, gateway, ntp, etc",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "DHCP Option Name",
							Optional:    true,
						},
						"num": {
							Type:        schema.TypeInt,
							Description: "DHCO Option number",
							Optional:    true,
						},
						"useoption": {
							Type:        schema.TypeBool,
							Description: "Use the option or not",
							Optional:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Value of the option",
							Optional:    true,
						},
						"vendorclass": {
							Type:        schema.TypeString,
							Description: "Vendor Class",
							Default:     "DHCP",
							Optional:    true,
						},
					},
				},
			},
			"recycleleases": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"restartifneeded": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"updatednsonleaserenewal": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

// resourceNetworkCreate  - Creates a new netowrk resource
func resourceNetworkCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var networkCreate network.Network
	var authority, createReverseZone, networkDisable, enableDdns, enableDhcpThresholds, enableDiscovery bool

	if v, ok := d.GetOk("network"); ok {
		networkCreate.Network = v.(string)
	}
	if v, ok := d.GetOk("networkview"); ok {
		networkCreate.NetworkView = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok {
		networkCreate.Comment = v.(string)
	}
	if v, ok := d.GetOk("authority"); ok {
		authority = v.(bool)
		networkCreate.Authority = &authority
	}
	if v, ok := d.GetOk("autocreatereversezone"); ok {
		createReverseZone = v.(bool)
		networkCreate.AutoCreateReversezone = &createReverseZone
	}
	if v, ok := d.GetOk("disable"); ok {
		networkDisable = v.(bool)
		networkCreate.Disable = &networkDisable
	}
	if v, ok := d.GetOk("enableddns"); ok {
		enableDdns = v.(bool)
		networkCreate.EnableDdns = &enableDdns
	}
	if v, ok := d.GetOk("enabledhcpthreshold"); ok {
		enableDhcpThresholds = v.(bool)
		networkCreate.EnableDhcpThresholds = &enableDhcpThresholds
	}
	if v, ok := d.GetOk("enablediscovery"); ok {
		enableDiscovery = v.(bool)
		networkCreate.EnableDiscovery = &enableDiscovery
	}
	if v, ok := d.GetOk("ipv4addr"); ok {
		networkCreate.Ipv4addr = v.(string)
	}

	if v, ok := d.GetOk("leasescavengetime"); ok {
		networkCreate.LeaseScavengeTime = v.(int)
	}

	if v, ok := d.GetOk("netmask"); ok {
		networkCreate.Netmask = uint(v.(int))
	}

	if v, ok := d.GetOk("networkcontainer"); ok {
		networkCreate.NetworkContainer = v.(string)
	}
	if v, ok := d.GetOk("option"); ok {
		if options, ok := v.(*schema.Set); ok {
			networkCreate.Options = buildOptionsObject(options)
		}
	}

	createNetworkAPI := network.NewCreateNetwork(networkCreate)
	createNetworkError := infobloxClient.Do(createNetworkAPI)
	if createNetworkError != nil {
		return fmt.Errorf("Error Creating Network %s", createNetworkError)
	}

	if createNetworkAPI.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Error creating the Network request, network already exists")
	}
	d.SetId(createNetworkAPI.GetResponse())
	return resourceNetworkRead(d, m)
}

// resourceNetworkDelete  - Delete a network resource
func resourceNetworkDelete(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	deleteAPI := network.NewDeleteNetwork(d.Id())
	deleteErr := infobloxClient.Do(deleteAPI)
	if deleteErr != nil {
		return fmt.Errorf("Cound not delete the network %s", deleteErr)
	}
	if deleteAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Error Deleting the Network : %s ", deleteAPI.ResponseObject())
	}
	d.SetId("")
	return nil
}

// resourceNetworkRead - Reads the resource
func resourceNetworkRead(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	fields := []string{"ipv4addr", "disable", "network", "network_view", "comment", "netmask", "authority", "enable_ddns", "options"}
	getNetworkAPI := network.NewGetNetwork(d.Id(), fields)
	networkReadErr := infobloxClient.Do(getNetworkAPI)
	if networkReadErr != nil {
		return fmt.Errorf("Could not read resource %s", networkReadErr)
	}

	if getNetworkAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Http Error Reading the resource: %s", getNetworkAPI.ResponseObject())
	}

	readNetwork := getNetworkAPI.GetResponse()
	d.Set("network", readNetwork.Network)
	d.Set("ipv4addr", readNetwork.Ipv4addr)
	d.Set("netmask", readNetwork.Netmask)
	d.Set("disable", readNetwork.Disable)
	d.Set("options", readNetwork.Options)
	d.Set("comment", readNetwork.Comment)
	d.Set("authority", readNetwork.Authority)
	d.Set("enable_ddns", readNetwork.EnableDdns)
	d.Set("network_view", readNetwork.NetworkView)
	d.Set("ref", readNetwork.Ref)

	return nil
}

// resourceNetworkUpdate - Updates the resource
func resourceNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var updateNetwork network.Network
	if v, ok := d.GetOk("ref"); ok {
		updateNetwork.Ref = v.(string)
	}
	if d.HasChange("network") {
		_, newNetwork := d.GetChange("network")
		updateNetwork.NetworkView = newNetwork.(string)
	}

	if d.HasChange("disable") {
		_, newDisable := d.GetChange("disable")
		Disable := newDisable.(bool)
		updateNetwork.Disable = &Disable
	}

	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok {
			updateNetwork.Comment = v.(string)
		}

	}

	if d.HasChange("option") {
		if v, ok := d.GetOk("option"); ok {
			if options, ok := v.(*schema.Set); ok {
				updateNetwork.Options = buildOptionsObject(options)
			}
		}

	}

	updateNetworkAPI := network.NewUpdateNetwork(updateNetwork)
	updateNetworkErr := infobloxClient.Do(updateNetworkAPI)
	if updateNetworkErr != nil {
		return updateNetworkErr
	}

	if updateNetworkAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Error updating the Network record %s ", updateNetworkAPI.GetResponse())
	}
	return resourceNetworkRead(d, m)
}

// buildOptionsObject - This is to avoid having to repeat the code every time I need to read this field
func buildOptionsObject(options *schema.Set) []network.DHCPOptions {
	optionValues := []network.DHCPOptions{}
	for _, option := range options.List() {
		optionObject := option.(map[string]interface{})
		newOption := network.DHCPOptions{}
		if optionName, ok := optionObject["name"].(string); ok {
			newOption.Name = optionName
		}

		if optionNum, ok := optionObject["num"].(int); ok {
			newOption.Num = uint(optionNum)
		}

		if optionUse, ok := optionObject["useoption"].(bool); ok {
			newOption.UseOption = &optionUse
		}

		if optionValue, ok := optionObject["value"].(string); ok {
			newOption.Value = optionValue
		}

		if optionVendorClass, ok := optionObject["vendorclass"].(string); ok {
			newOption.VendorClass = optionVendorClass
		}
		optionValues = append(optionValues, newOption)

	}
	return optionValues
}
