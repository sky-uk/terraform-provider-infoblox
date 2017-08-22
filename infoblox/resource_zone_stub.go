package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zonestub"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceZoneStub() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneStubCreate,
		Read:   resourceZoneStubRead,
		Update: resourceZoneStubUpdate,
		Delete: resourceZoneStubDelete,
		Schema: map[string]*schema.Schema{
			"Reference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment for the zone; maximum 256 characters",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"disable": {
				Type:        schema.TypeBool,
				Description: "Is the zone disabled",
				Optional:    true,
				ForceNew:    false,
			},
			"locked": {
				Type:        schema.TypeBool,
				Description: "Is the record locked to prevent changes",
				Optional:    true,
				ForceNew:    false,
			},
			"disable_forwarding": {
				Type:        schema.TypeBool,
				Description: "Is forward disabled for this zone",
				Optional:    true,
				ForceNew:    false,
			},
			"external_nsgroup": {
				Type:         schema.TypeString,
				Description:  "Name of the external name server group",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"fqdn": {
				Type: schema.TypeString,
				Description: "Fqdn for the zone	",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"mask_prefix": {
				Type:         schema.TypeString,
				Description:  "IPv4 Netmask or IPv6 prefix for this zone.",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"nsgroup": {
				Type:         schema.TypeString,
				Description:  "Name of the  name server group",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"prefix": {
				Type:         schema.TypeString,
				Description:  "IPv4 Netmask or IPv6 prefix for this zone.",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"stub_from":    util.ExternalServerListSchema(true, false),
			"stub_members": util.MemberServerListSchema(),
			"zoneformat": {
				Type:         schema.TypeString,
				Description:  "Determines the format of this zone - API default FORWARD",
				ValidateFunc: util.ValidateZoneFormat,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
			},
			"view": {
				Type:         schema.TypeString,
				Description:  "The name of the DNS view in which the zone resides",
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
		},
	}
}

func resourceZoneStubCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var createZoneStub zonestub.ZoneStub

	if v, ok := d.GetOk("comment"); ok {
		createZoneStub.Comment = v.(string)

	}

	if v, ok := d.GetOk("disable"); ok {
		disable := v.(bool)
		createZoneStub.Disable = &disable
	}

	if v, ok := d.GetOk("locked"); ok {
		locked := v.(bool)
		createZoneStub.Locked = &locked
	}

	if v, ok := d.GetOk("disable_forwarding"); ok {
		disableForwarding := v.(bool)
		createZoneStub.DisableForwarding = &disableForwarding
	}
	if v, ok := d.GetOk("external_nsgroup"); ok {
		createZoneStub.ExternalNSGroup = v.(string)
	}

	if v, ok := d.GetOk("fqdn"); ok {
		createZoneStub.FQDN = v.(string)
	}

	if v, ok := d.GetOk("mask_prefix"); ok {
		createZoneStub.MaskPrefix = v.(string)
	}

	if v, ok := d.GetOk("nsgroup"); ok {
		createZoneStub.NsGroup = v.(string)
	}

	if v, ok := d.GetOk("prefix"); ok {
		createZoneStub.Prefix = v.(string)
	}
	if v, ok := d.GetOk("stub_from"); ok {
		servers := []map[string]interface{}{}
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		createZoneStub.StubFrom = util.BuildExternalServerListFromT(servers)
	}

	if v, ok := d.GetOk("stub_members"); ok {
		servers := []map[string]interface{}{}
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		createZoneStub.StubMembers = util.BuildMemberServerListFromT(servers)
	}

	if v, ok := d.GetOk("zoneformat"); ok {
		createZoneStub.ZoneFormat = v.(string)
	}

	if v, ok := d.GetOk("view"); ok {
		createZoneStub.View = v.(string)
	}
	createZoneStubAPI := zonestub.NewCreate(createZoneStub)
	createZoneStubErr := infobloxClient.Do(createZoneStubAPI)
	if createZoneStubErr != nil {
		return fmt.Errorf("Infoblox Error creating the Stub Zone : %s", createZoneStubErr.Error())
	}
	if createZoneStubAPI.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Infoblox Zone Create Error: Invalid HTTP response code %d returned - response %s", createZoneStubAPI.StatusCode(), *createZoneStubAPI.ResponseObject().(*string))
	}
	d.SetId(*createZoneStubAPI.ResponseObject().(*string))
	return resourceZoneStubRead(d, m)

}

func resourceZoneStubRead(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var readZoneStub zonestub.ZoneStub
	zoneReadAPI := zonestub.NewGet(d.Id(), returnZoneStubFields())
	readErr := infobloxClient.Do(zoneReadAPI)
	if readErr != nil {

	}

	if zoneReadAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Zone Create Error: Invalid HTTP response code %d returned - response %s", zoneReadAPI.StatusCode(), zoneReadAPI.ResponseObject().(*string))

	}

	readZoneStub = *zoneReadAPI.ResponseObject().(*zonestub.ZoneStub)
	d.SetId(readZoneStub.Ref)
	d.Set("comment", readZoneStub.Comment)
	d.Set("disable", readZoneStub.Disable)
	d.Set("locked", readZoneStub.Locked)
	d.Set("disable_forwarding", readZoneStub.DisableForwarding)
	d.Set("external_nsgroup", readZoneStub.ExternalNSGroup)
	d.Set("fqdn", readZoneStub.FQDN)
	d.Set("mask_prefix", readZoneStub.MaskPrefix)
	d.Set("nsgroup", readZoneStub.NsGroup)
	d.Set("prefix", readZoneStub.Prefix)
	d.Set("stub_from", util.BuildExternalServersListFromIBX(readZoneStub.StubFrom))
	d.Set("stub_members", util.BuildMemberServerListFromIBX(readZoneStub.StubMembers))
	d.Set("zoneformat", readZoneStub.ZoneFormat)
	d.Set("view", readZoneStub.View)
	return nil
}

func resourceZoneStubUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var updateStubZone zonestub.ZoneStub
	updateStubZone.Ref = d.Id()
	if d.HasChange("comment") {
		_, newComment := d.GetChange("comment")
		updateStubZone.Comment = newComment.(string)
	}

	if d.HasChange("disable") {
		_, newDisable := d.GetChange("disable")
		disable := newDisable.(bool)
		updateStubZone.Disable = &disable
	}

	if d.HasChange("locked") {
		_, newLocked := d.GetChange("locked")
		locked := newLocked.(bool)
		updateStubZone.Locked = &locked
	}
	if d.HasChange("disable_forwarding") {
		_, newDisableForwarding := d.GetChange("disable_forwarding")
		disableForwarding := newDisableForwarding.(bool)
		updateStubZone.DisableForwarding = &disableForwarding
	}
	if d.HasChange("external_nsgroup") {
		_, newExternalNSGroup := d.GetChange("external_nsgroup")
		updateStubZone.ExternalNSGroup = newExternalNSGroup.(string)
	}
	// Note , since the FQDN attribute forces the creation of a new resource I don't think I need to implement it .
	// will test though

	if d.HasChange("nsgroup") {
		_, newNSGroup := d.GetChange("nsgroup")
		updateStubZone.NsGroup = newNSGroup.(string)
	}

	if d.HasChange("prefix") {
		_, newPrefix := d.GetChange("prefix")
		updateStubZone.Prefix = newPrefix.(string)

	}

	if d.HasChange("stub_from") {
		servers := []map[string]interface{}{}
		_, ServersValue := d.GetChange("stub_from")
		for _, server := range ServersValue.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		updateStubZone.StubFrom = util.BuildExternalServerListFromT(servers)
	}

	if d.HasChange("stub_members") {
		servers := []map[string]interface{}{}
		_, ServersValue := d.GetChange("stub_members")
		for _, server := range ServersValue.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		updateStubZone.StubMembers = util.BuildMemberServerListFromT(servers)

	}

	if d.HasChange("zoneformat") {
		_, newZoneFormat := d.GetChange("zoneformat")
		updateStubZone.ZoneFormat = newZoneFormat.(string)
	}

	if d.HasChange("view") {
		_, newView := d.GetChange("view")
		updateStubZone.ZoneFormat = newView.(string)
	}
	updateStubZoneAPI := zonestub.NewUpdate(updateStubZone)
	updateStubZoneErr := infobloxClient.Do(updateStubZoneAPI)
	if updateStubZoneErr != nil {
		return fmt.Errorf("Error updating the Stub Zone %s", updateStubZoneErr.Error())
	}

	if updateStubZoneAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Error updating the Stub Zone : %s - %s ", updateStubZoneAPI.StatusCode(), updateStubZoneAPI.ResponseObject().(string))
	}

	return resourceZoneStubRead(d, m)
}

func resourceZoneStubDelete(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	deleteAPI := zonestub.NewDelete(d.Id())
	deleteErr := infobloxClient.Do(deleteAPI)
	if deleteErr != nil {
		return fmt.Errorf("Infoblox Zone Delete Error: %s", deleteErr.Error())
	}

	if deleteAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Zone Delete Error: %s - %s", deleteAPI.StatusCode(), deleteAPI.ResponseObject().(*string))
	}
	d.SetId("")
	return nil
}

func returnZoneStubFields() []string {
	return []string{"comment", "disable", "locked", "disable_forwarding", "external_ns_group", "fqdn", "mask_prefix", "ns_group", "prefix", "stub_from", "stub_members", "zone_format", "view"}

}
