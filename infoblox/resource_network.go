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
			"use_authority": {
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
			"use_enableddns": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"high_watermark": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"high_watermark_reset": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"low_watermark": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"low_watermark_reset": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"enabledhcpthresholds": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_enabledhcpthresholds": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enablediscovery": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_enablediscovery": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"discovery_member": {
				Type:     schema.TypeString,
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
			"use_options": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"recycleleases": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_recycleleases": {
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
	var authority, useAuthority, createReverseZone, networkDisable, enableDdns, useEnableDdns, enableDhcpThresholds, useEnableDhcpThresholds, enableDiscovery, useEnableDiscovery, recycleLeases, useRecycleLeases, useOptions bool

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
	if v, ok := d.GetOk("use_authority"); ok {
		useAuthority = v.(bool)
		networkCreate.UseAuthority = &useAuthority
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
	if v, ok := d.GetOk("use_enableddns"); ok {
		useEnableDdns = v.(bool)
		networkCreate.UseEnableDdns = &useEnableDdns
	}
	if v, ok := d.GetOk("enabledhcpthresholds"); ok {
		enableDhcpThresholds = v.(bool)
		networkCreate.EnableDhcpThresholds = &enableDhcpThresholds
	}
	if v, ok := d.GetOk("use_enabledhcpthresholds"); ok {
		useEnableDhcpThresholds = v.(bool)
		networkCreate.UseEnableDhcpThresholds = &useEnableDhcpThresholds
	}
	if v, ok := d.GetOk("high_watermark"); ok {
		networkCreate.HighWaterMark = v.(int)
	}

	if v, ok := d.GetOk("low_watermark"); ok {
		networkCreate.LowWaterMark = v.(int)
	}
	if v, ok := d.GetOk("low_watermark_reset"); ok {
		networkCreate.LowWaterMarkReset = v.(int)
	}
	if v, ok := d.GetOk("high_watermark_reset"); ok {
		networkCreate.HighWaterMarkReset = v.(int)
	}
	if v, ok := d.GetOk("enablediscovery"); ok {
		enableDiscovery = v.(bool)
		networkCreate.EnableDiscovery = &enableDiscovery
	}
	if v, ok := d.GetOk("use_enablediscovery"); ok {
		useEnableDiscovery = v.(bool)
		networkCreate.UseEnableDiscovery = &useEnableDiscovery
	}
	if v, ok := d.GetOk("discovery_member"); ok {
		networkCreate.DiscoveryMember = v.(string)
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
	if v, ok := d.GetOk("use_options"); ok {
		useOptions = v.(bool)
		networkCreate.UseOptions = &useOptions
	}
	if v, ok := d.GetOk("recycleleases"); ok {
		recycleLeases = v.(bool)
		networkCreate.RecycleLeases = &recycleLeases
	}
	if v, ok := d.GetOk("use_recycleleases"); ok {
		useRecycleLeases = v.(bool)
		networkCreate.UseRecycleLeases = &useRecycleLeases
	}

	createNetworkAPI := network.NewCreateNetwork(networkCreate)
	createNetworkError := infobloxClient.Do(createNetworkAPI)
	if createNetworkError != nil {
		return fmt.Errorf("Error Creating Network %s", createNetworkError)
	}

	if createNetworkAPI.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Infoblox Create Error: Invalid HTTP response code %+v returned. Response object was %+v", createNetworkAPI.StatusCode(), createNetworkAPI.GetResponse())
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

	if d.HasChange("use_options") {
		_, newUseOptions := d.GetChange("use_options")
		useOptions := newUseOptions.(bool)
		updateNetwork.UseOptions = &useOptions
	}

	if d.HasChange("authority") {
		_, newAuthority := d.GetChange("authority")
		authority := newAuthority.(bool)
		updateNetwork.Authority = &authority
	}

	if d.HasChange("use_authority") {
		_, newUseAuthority := d.GetChange("use_authority")
		useAuthority := newUseAuthority.(bool)
		updateNetwork.UseAuthority = &useAuthority
	}

	if d.HasChange("enableddns") {
		_, newEnableDdns := d.GetChange("enableddns")
		enableDdns := newEnableDdns.(bool)
		updateNetwork.EnableDdns = &enableDdns
	}

	if d.HasChange("use_enableddns") {
		_, newUseEnableDdns := d.GetChange("use_enableddns")
		useEnableDdns := newUseEnableDdns.(bool)
		updateNetwork.UseEnableDdns = &useEnableDdns
	}

	if d.HasChange("enabledhcpthresholds") {
		_, newEnableDhcpThreshold := d.GetChange("enabledhcpthresholds")
		enableDhcpThreshold := newEnableDhcpThreshold.(bool)
		updateNetwork.EnableDhcpThresholds = &enableDhcpThreshold
	}

	if d.HasChange("use_enabledhcpthresholds") {
		_, newUseEnableDhcpThreshold := d.GetChange("use_enabledhcpthresholds")
		useEnableDhcpThreshold := newUseEnableDhcpThreshold.(bool)
		updateNetwork.UseEnableDhcpThresholds = &useEnableDhcpThreshold
	}

	if d.HasChange("high_watermark") {
		_, newHighWatermark := d.GetChange("high_watermark")
		highWatermark := newHighWatermark.(int)
		updateNetwork.HighWaterMark = highWatermark
	}

	if d.HasChange("high_watermark_reset") {
		_, newHighWatermarkReset := d.GetChange("high_watermark_reset")
		highWatermarkReset := newHighWatermarkReset.(int)
		updateNetwork.HighWaterMarkReset = highWatermarkReset
	}

	if d.HasChange("low_watermark") {
		_, newLowWatermark := d.GetChange("low_watermark")
		lowWatermark := newLowWatermark.(int)
		updateNetwork.LowWaterMark = lowWatermark
	}
	if d.HasChange("low_watermark_reset") {
		_, newLowWatermarkReset := d.GetChange("low_watermark_reset")
		lowWatermarkReset := newLowWatermarkReset.(int)
		updateNetwork.LowWaterMarkReset = lowWatermarkReset
	}

	if d.HasChange("enablediscovery") {
		_, newEnableDiscovery := d.GetChange("enablediscovery")
		enableDiscovery := newEnableDiscovery.(bool)
		updateNetwork.EnableDiscovery = &enableDiscovery
	}

	if d.HasChange("use_enablediscovery") {
		_, newUseEnableDiscovery := d.GetChange("use_enablediscovery")
		useEnableDiscovery := newUseEnableDiscovery.(bool)
		updateNetwork.UseEnableDiscovery = &useEnableDiscovery
	}

	if d.HasChange("discovery_member") {
		if v, ok := d.GetOk("discovery_member"); ok {
			updateNetwork.DiscoveryMember = v.(string)
		}
	}

	if d.HasChange("recycleleases") {
		_, newRecycleLeases := d.GetChange("recycleleases")
		recycleLeases := newRecycleLeases.(bool)
		updateNetwork.RecycleLeases = &recycleLeases
	}

	if d.HasChange("use_recycleleases") {
		_, newUseRecycleLeases := d.GetChange("use_recycleleases")
		useRecycleLeases := newUseRecycleLeases.(bool)
		updateNetwork.UseRecycleLeases = &useRecycleLeases
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
