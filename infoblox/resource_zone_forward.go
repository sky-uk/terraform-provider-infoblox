package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneforward"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceZoneForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneForwardCreate,
		Read:   resourceZoneForwardRead,
		Update: resourceZoneForwardUpdate,
		Delete: resourceZoneForwardDelete,

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
				Required:     false,
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

	ibxClient := m.(*skyinfoblox.InfobloxClient)
	var zone zoneforward.ZoneForward

	if v, ok := d.GetOk("comment"); ok && v != "" {
		zone.Comment = v.(string)
	}

	zone.Disable = d.Get("disable").(bool)

	if v, ok := d.GetOk("forward_to"); ok {
		serverList := util.GetMapList(v.([]interface{}))
		zone.ForwardTo = util.BuildExternalServerListFromT(serverList)
	}

	zone.ForwardersOnly = d.Get("forwarders_only").(bool)

	if v, ok := d.GetOk("forwarding_servers"); ok {
		serverList := util.GetMapList(v.([]interface{}))
		zone.ForwardingServers = util.BuildForwardingMemberServerListFromT(serverList)
	}

	if v, ok := d.GetOk("fqdn"); ok && v != "" {
		zone.Fqdn = v.(string)
	}

	zone.Locked = d.Get("locked").(bool)

	if v, ok := d.GetOk("ns_group"); ok && v != "" {
		zone.NSGroup = v.(string)
	}

	if v, ok := d.GetOk("prefix"); ok && v != "" {
		zone.Prefix = v.(string)
	}

	zone.UsingSrgAssociations = d.Get("using_srg_associations").(bool)

	if v, ok := d.GetOk("view"); ok && v != "" {
		zone.View = v.(string)
	}

	if v, ok := d.GetOk("zone_format"); ok && v != "" {
		zone.ZoneFormat = v.(string)
	}

	api := zoneforward.NewCreate(zone)
	err := ibxClient.Do(api)

	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Error creating a new forward zone, error:\n%s\n", err))
	}

	if api.StatusCode() != http.StatusCreated {
		return fmt.Errorf(fmt.Sprintf("Error creating a new forward zone, status: %d, error:\n%s\n",
			api.StatusCode(), api.RawResponse()))
	}

	ref := *api.ResponseObject().(*string)
	d.SetId(ref)

	return resourceZoneForwardRead(d, m)
}

func zoneForwardReturnFields() []string {
	return []string{"address", "comment", "disable", "display_domain", "dns_fqdn", "forward_to", "forwarders_only", "forwarding_servers", "fqdn", "locked", "locked_by", "mask_prefix", "ms_ad_integrated", "ms_ddns_mode", "ms_managed", "ms_read_only", "ms_sync_master_name", "parent", "prefix", "using_srg_associations", "view", "zone_format"}
}

func resourceZoneForwardRead(d *schema.ResourceData, m interface{}) error {
	ibxClient := m.(*skyinfoblox.InfobloxClient)
	resourceReference := d.Id()
	if resourceReference == "" {
		return nil
	}
	api := zoneforward.NewGet(resourceReference, zoneForwardReturnFields())
	err := ibxClient.Do(api)
	if err != nil {
		return fmt.Errorf("Could not read the resource %s", err.Error())
	}
	if api.StatusCode() != http.StatusOK {
		return fmt.Errorf("Could not read the resource %s", string(api.RawResponse()))
	}

	zone := *api.ResponseObject().(*zoneforward.ZoneForward)

	d.SetId(zone.Ref)
	d.Set("address", zone.Address)
	d.Set("comment", zone.Comment)
	d.Set("disable", zone.Disable)
	d.Set("display_domain", zone.DisplayDomain)
	d.Set("dns_fqdn", zone.DNSFqdn)
	d.Set("forward_to", zone.ForwardTo)
	d.Set("forwarders_only", zone.ForwardersOnly)
	d.Set("forwarding_servers", zone.ForwardingServers)
	d.Set("fqdn", zone.Fqdn)
	d.Set("locked", zone.Locked)
	d.Set("locked_by", zone.LockedBy)
	d.Set("mask_prefix", zone.MaskPrefix)
	d.Set("ns_group", zone.NSGroup)
	d.Set("parent", zone.Parent)
	d.Set("prefix", zone.Prefix)
	d.Set("using_srg_associations", zone.UsingSrgAssociations)
	d.Set("view", zone.View)
	d.Set("zone_format", zone.ZoneFormat)

	return nil
}

func resourceZoneForwardUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var updatedZone zoneforward.ZoneForward
	hasChanges := false
	returnFields := zoneForwardReturnFields()
	resourceReference := d.Id()
	updatedZone.Ref = resourceReference

	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok {
			updatedZone.Comment = v.(string)
			hasChanges = true
		}
	}

	updatedZone.Disable = d.Get("disable").(bool)
	oldV, newV := d.GetChange("disable")
	updatedZone.Disable = newV.(bool)
	if oldV.(bool) != newV.(bool) {
		hasChanges = true
	}

	if d.HasChange("forward_to") {
		servers := util.GetMapList(d.Get("forward_to").([]interface{}))
		updatedZone.ForwardTo = util.BuildExternalServerListFromT(servers)
		hasChanges = true
	}

	updatedZone.ForwardersOnly = d.Get("forwarders_only").(bool)
	oldV, newV = d.GetChange("forwarders_only")
	updatedZone.ForwardersOnly = newV.(bool)
	if oldV.(bool) != newV.(bool) {
		hasChanges = true
	}

	if d.HasChange("forwarding_servers") {
		servers := util.GetMapList(d.Get("forwarding_servers").([]interface{}))
		updatedZone.ForwardingServers = util.BuildForwardingMemberServerListFromT(servers)
		hasChanges = true
	}

	updatedZone.Locked = d.Get("locked").(bool)
	oldV, newV = d.GetChange("locked")
	updatedZone.Locked = newV.(bool)
	if oldV.(bool) != newV.(bool) {
		hasChanges = true
	}

	if d.HasChange("ns_group") {
		if v, ok := d.GetOk("ns_group"); ok {
			updatedZone.NSGroup = v.(string)
			hasChanges = true
		}
	}

	if d.HasChange("prefix") {
		if v, ok := d.GetOk("prefix"); ok {
			updatedZone.Prefix = v.(string)
			hasChanges = true
		}
	}

	if d.HasChange("view") {
		if v, ok := d.GetOk("view"); ok {
			updatedZone.View = v.(string)
			hasChanges = true
		}
	}

	if hasChanges == true {
		updateAPI := zoneforward.NewUpdate(updatedZone, returnFields)
		err := infobloxClient.Do(updateAPI)
		if err != nil {
			return fmt.Errorf("Infoblox Zone Forward Update Error: %+v", err)
		}
		if updateAPI.StatusCode() != http.StatusOK {
			return fmt.Errorf("Error updating Zone Forward with objRef: %s, Error:\n%s", resourceReference, string(updateAPI.RawResponse()))
		}

		updatedZone = *updateAPI.ResponseObject().(*zoneforward.ZoneForward)

		d.SetId(updatedZone.Ref)
	}

	return resourceZoneForwardRead(d, m)
}

func resourceZoneForwardDelete(d *schema.ResourceData, m interface{}) error {
	returnFields := []string{"fqdn"}

	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	resourceReference := d.Id()
	getAPI := zoneforward.NewGet(resourceReference, returnFields)

	err := infobloxClient.Do(getAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Delete Error when fetching resource: %+v", err)
	}
	if getAPI.StatusCode() == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	deleteAPI := zoneforward.NewDelete(resourceReference)
	err = infobloxClient.Do(deleteAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Delete - Error deleting resource %+v", err)
	}

	if deleteAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Delete - Error deleting resource %s - return code != 200 error: %v", resourceReference, deleteAPI.ResponseObject())
	}

	d.SetId("")
	return nil
}
