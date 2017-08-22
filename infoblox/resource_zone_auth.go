package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceZoneAuth() *schema.Resource {
	return &schema.Resource{
		Create: resourceZoneAuthCreate,
		Read:   resourceZoneAuthRead,
		Update: resourceZoneAuthUpdate,
		Delete: resourceZoneAuthDelete,

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
				ForceNew:     true,
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
			"grid_primary":         util.MemberServerListSchema(),
			"grid_primary_shared_with_ms_parent_delegation": {
				Type:        schema.TypeBool,
				Description: "Determines if the server is duplicated with parent delegation.cannot be updated, nor written",
				Optional:    true,
				Computed:    true,
			},
			"grid_secondaries": util.MemberServerListSchema(),
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
				Computed:    true,
			},
			"soa_serial_number": {
				Type:        schema.TypeInt,
				Description: "The SOA serial number to be used in conjunction with set_soa_serial_number (read-only)",
				Computed:    true,
			},
			"soa_default_ttl": {
				Type:         schema.TypeInt,
				Description:  "The Time to Live (TTL) value of the SOA record of this zone",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soa_negative_ttl": {
				Type:         schema.TypeInt,
				Description:  "The negative Time to Live (TTL)",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soa_refresh": {
				Type:         schema.TypeInt,
				Description:  "This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soa_retry": {
				Type:         schema.TypeInt,
				Description:  "This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soa_expire": {
				Type:        schema.TypeInt,
				Description: "This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. The default is one week.",
				Optional:    true,
				Default:     2419200,
			},
			"copy_xfer_to_notify": {
				Type:        schema.TypeBool,
				Description: "If this flag is set to True then copy allowed IPs from Allow Transfer to Also Notify.",
				Default:     false,
				Optional:    true,
			},
			"use_copy_xfer_to_notify": {
				Type:        schema.TypeBool,
				Description: "Use flag for: copy_xfer_to_notify.",
				Default:     false,
				Optional:    true,
			},
			"use_check_names_policy": {
				Type:        schema.TypeBool,
				Description: "Apply policy to dynamic updates and inbound zone transfers (This value applies only if the host name restriction policy is set to “Strict Hostname Checking”.)",
				Default:     false,
				Optional:    true,
			},
			"allow_update":   util.AccessControlSchema(),
			"allow_transfer": util.AccessControlSchema(),
			"use_allow_transfer": {
				Type:        schema.TypeBool,
				Description: "allow_transfer",
				Default:     false,
				Optional:    true,
			},
		},
	}
}

func resourceZoneAuthCreate(d *schema.ResourceData, m interface{}) error {

	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	// dnsZone is used for the initial request which creates the zone
	var dnsZone zoneauth.DNSZone
	// appendDNSZone is used for the second request after the zone has been created.
	var appendDNSZone zoneauth.DNSZone
	gridTimer := true

	if v, ok := d.GetOk("fqdn"); ok && v != "" {
		dnsZone.FQDN = v.(string)
	} else {
		return fmt.Errorf("Infoblox Zone Auth Create Error: name argument required")
	}
	if v, ok := d.GetOk("ns_group"); ok && v != "" {
		dnsZone.NSGroup = v.(string)
	}
	if v, ok := d.GetOk("view"); ok && v != "" {
		dnsZone.View = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		dnsZone.Comment = v.(string)
	}
	if v, ok := d.GetOk("zone_format"); ok && v != "" {
		dnsZone.ZoneFormat = v.(string)
	}
	if v, ok := d.GetOk("prefix"); ok && v != "" {
		dnsZone.Prefix = v.(string)
	}
	if v, ok := d.GetOk("disable"); ok {
		dnsZoneDisable := v.(bool)
		dnsZone.Disable = &dnsZoneDisable
	}
	if v, ok := d.GetOk("dns_integrity_enable"); ok {
		dnsIntegrityEnable := v.(bool)
		dnsZone.DNSIntegrityEnable = &dnsIntegrityEnable
	}
	if v, ok := d.GetOk("dns_integrity_member"); ok {
		dnsZone.DNSIntegrityMember = v.(string)
	}
	if v, ok := d.GetOk("external_primaries"); ok {
		servers := []map[string]interface{}{}
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		dnsZone.ExternalPrimaries = util.BuildExternalServerListFromT(servers)
	}
	if v, ok := d.GetOk("external_secondaries"); ok {
		servers := []map[string]interface{}{}
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		dnsZone.ExternalSecondaries = util.BuildExternalServerListFromT(servers)
	}
	if v, ok := d.GetOk("grid_primary"); ok {
		servers := make([]map[string]interface{}, 0)
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		dnsZone.GridPrimary = util.BuildMemberServerListFromT(servers)
	}
	if v, ok := d.GetOk("grid_secondaries"); ok {
		servers := make([]map[string]interface{}, 0)
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		dnsZone.GridSecondaries = util.BuildMemberServerListFromT(servers)
	}
	if v, ok := d.GetOk("locked"); ok {
		zoneLocked := v.(bool)
		dnsZone.Locked = &zoneLocked
	}
	if v, ok := d.GetOk("allow_update"); ok && v != nil {
		dnsZone.AllowUpdate = util.BuildAcList(v.([]interface{}))
	}
	if v, ok := d.GetOk("use_allow_transfer"); ok {
		useAllowTransfer := v.(bool)
		dnsZone.UseAllowTransfer = &useAllowTransfer
	}
	if v, ok := d.GetOk("allow_transfer"); ok && v != nil {
		dnsZone.AllowTransfer = util.BuildAcList(v.([]interface{}))
	}

	if v, ok := d.GetOk("restart_if_needed"); ok {
		restart := v.(bool)
		appendDNSZone.RestartIfNeeded = &restart
	}

	// Some attributes can't be set on creation. They need to be sent in a subsequent request after initial creation.
	if v, ok := d.GetOk("soa_default_ttl"); ok && v != nil {
		soaTTL := v.(int)
		appendDNSZone.SOADefaultTTL = uint(soaTTL)
		appendDNSZone.UseGridZoneTimer = &gridTimer
	}
	if v, ok := d.GetOk("soa_negative_ttl"); ok && v != nil {
		soaNegativeTTL := v.(int)
		appendDNSZone.SOANegativeTTL = uint(soaNegativeTTL)
		appendDNSZone.UseGridZoneTimer = &gridTimer
	}
	if v, ok := d.GetOk("soa_refresh"); ok && v != nil {
		soaRefresh := v.(int)
		appendDNSZone.SOARefresh = uint(soaRefresh)
		appendDNSZone.UseGridZoneTimer = &gridTimer
	}
	if v, ok := d.GetOk("soa_retry"); ok && v != nil {
		soaRetry := v.(int)
		appendDNSZone.SOARetry = uint(soaRetry)
		appendDNSZone.UseGridZoneTimer = &gridTimer
	}
	if v, ok := d.GetOk("soa_expire"); ok && v != nil {
		soaExpire := v.(int)
		appendDNSZone.SOAExpire = uint(soaExpire)
		appendDNSZone.UseGridZoneTimer = &gridTimer
	}
	if v, ok := d.GetOk("copy_xfer_to_notify"); ok {
		copyXferToNotify := v.(bool)
		dnsZone.CopyXferToNotify = &copyXferToNotify
	}
	if v, ok := d.GetOk("use_copy_xfer_to_notify"); ok {
		useCopyXferToNotify := v.(bool)
		dnsZone.UseCopyXferNotify = &useCopyXferToNotify
	}
	if v, ok := d.GetOk("use_check_names_policy"); ok {
		useCheckNamesPolicy := v.(bool)
		dnsZone.UseCheckNamesPolicy = &useCheckNamesPolicy
	}

	createAPI := zoneauth.NewCreate(dnsZone)
	err := infobloxClient.Do(createAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Zone Auth Create Error: %+v", err)
	}
	if createAPI.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Infoblox Zone Create Error: Invalid HTTP response code %d returned - response %s", createAPI.StatusCode(), createAPI.GetResponse())
	}

	ref := createAPI.GetResponse()
	// We can't set some attributes on create. Therefore we need to make another call.
	appendDNSZone.Reference = ref
	appendAPI := zoneauth.NewUpdate(appendDNSZone, nil)
	err = infobloxClient.Do(appendAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Zone Auth Create Append Error: %+v ", err)
	}
	if appendAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Zone Auth Create Append: Invalid HTTP response code %d returned - response %s", appendAPI.StatusCode(), appendAPI.GetResponse())
	}

	d.SetId(ref)
	return resourceZoneAuthRead(d, m)
}

func returnFields() []string {
	return []string{"fqdn", "comment", "zone_format", "view", "prefix", "soa_serial_number", "soa_default_ttl", "soa_negative_ttl", "soa_refresh", "soa_retry", "soa_expire", "copy_xfer_to_notify", "use_copy_xfer_to_notify", "disable", "dns_integrity_enable", "dns_integrity_member", "external_primaries", "external_secondaries", "grid_primary", "grid_secondaries", "grid_primary_shared_with_ms_parent_delegation", "locked", "locked_by", "network_view", "ns_group", "allow_update", "allow_transfer", "use_check_names_policy"}
}

func resourceZoneAuthRead(d *schema.ResourceData, m interface{}) error {

	returnFields := returnFields()
	resourceReference := d.Id()
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	getZone := zoneauth.NewGetSingleZone(resourceReference, returnFields)

	err := infobloxClient.Do(getZone)
	if err != nil {
		return fmt.Errorf("Error retrieving object using reference %s", resourceReference)
	}
	if getZone.StatusCode() == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	response := getZone.GetResponse()

	d.SetId(response.Reference)
	d.Set("fqdn", response.FQDN)
	d.Set("view", response.View)
	d.Set("comment", response.Comment)
	d.Set("zone_format", response.ZoneFormat)
	d.Set("prefix", response.Prefix)
	d.Set("soa_default_ttl", response.SOADefaultTTL)
	d.Set("soa_negative_ttl", response.SOANegativeTTL)
	d.Set("soa_refresh", response.SOARefresh)
	d.Set("soa_retry", response.SOARetry)
	d.Set("soa_serial_number", response.SOASerialNumber)
	d.Set("soa_expire", response.SOAExpire)
	d.Set("disable", response.Disable)
	d.Set("dns_integrity_enable", response.DNSIntegrityEnable)
	d.Set("dns_integrity_member", response.DNSIntegrityMember)
	d.Set("external_primaries", util.BuildExternalServersListFromIBX(response.ExternalPrimaries))
	d.Set("external_secondaries", util.BuildExternalServersListFromIBX(response.ExternalSecondaries))
	d.Set("grid_primary_shared_with_ms_parent_delegation", response.GridPrimarySharedWithMSParentDelegation)
	d.Set("grid_primary", util.BuildMemberServerListFromIBX(response.GridPrimary))
	d.Set("grid_secondaries", util.BuildMemberServerListFromIBX(response.GridSecondaries))
	d.Set("locked", response.Locked)
	restart, _ := d.GetOk("restart_if_needed")
	d.Set("restart_if_needed", &restart)
	d.Set("locked_by", response.LockedBy)
	d.Set("network_view", response.NetworkView)
	d.Set("ns_group", response.NSGroup)
	d.Set("copy_xfer_to_notify", response.CopyXferToNotify)
	d.Set("use_copy_xfer_to_notify", response.UseCopyXferNotify)
	d.Set("use_check_names_policy", response.UseCheckNamesPolicy)
	d.Set("allow_update", response.AllowUpdate)
	d.Set("allow_transfer", response.AllowTransfer)
	return nil
}

func resourceZoneAuthUpdate(d *schema.ResourceData, m interface{}) error {

	var updateZoneAuth zoneauth.DNSZone
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	hasChanges := false
	returnFields := returnFields()
	resourceReference := d.Id()
	updateZoneAuth.Reference = resourceReference
	gridTimer := true

	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok {
			updateZoneAuth.Comment = v.(string)
		}
		hasChanges = true
	}

	if d.HasChange("prefix") {
		if v, ok := d.GetOk("prefix"); ok {
			updateZoneAuth.Prefix = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("soa_default_ttl") {
		if v, ok := d.GetOk("soa_default_ttl"); ok && v != nil {
			soaTTL := v.(int)
			updateZoneAuth.SOADefaultTTL = uint(soaTTL)
			updateZoneAuth.UseGridZoneTimer = &gridTimer
		}
		hasChanges = true
	}
	if d.HasChange("soa_negative_ttl") {
		if v, ok := d.GetOk("soa_negative_ttl"); ok && v != nil {
			soaNegativeTTL := v.(int)
			updateZoneAuth.SOANegativeTTL = uint(soaNegativeTTL)
			updateZoneAuth.UseGridZoneTimer = &gridTimer
		}
		hasChanges = true
	}
	if d.HasChange("soa_refresh") {
		if v, ok := d.GetOk("soa_refresh"); ok && v != nil {
			soaRefresh := v.(int)
			updateZoneAuth.SOARefresh = uint(soaRefresh)
			updateZoneAuth.UseGridZoneTimer = &gridTimer
		}
		hasChanges = true
	}
	if d.HasChange("soa_retry") {
		if v, ok := d.GetOk("soa_retry"); ok && v != nil {
			soaRetry := v.(int)
			updateZoneAuth.SOARetry = uint(soaRetry)
			updateZoneAuth.UseGridZoneTimer = &gridTimer
		}
		hasChanges = true
	}
	if d.HasChange("soa_expire") {
		if v, ok := d.GetOk("soa_expire"); ok && v != nil {
			soaExpire := v.(int)
			updateZoneAuth.SOAExpire = uint(soaExpire)
			updateZoneAuth.UseGridZoneTimer = &gridTimer
		}
		hasChanges = true
	}
	if d.HasChange("disable") {
		dnsZoneDisable := d.Get("disable").(bool)
		updateZoneAuth.Disable = &dnsZoneDisable
		hasChanges = true
	}
	if d.HasChange("restart_if_needed") {
		flag := d.Get("restart_if_needed").(bool)
		updateZoneAuth.RestartIfNeeded = &flag
		hasChanges = true
	}
	if d.HasChange("dns_integrity_enable") {
		dnsIntegrityEnable := d.Get("dns_integrity_enable").(bool)
		updateZoneAuth.DNSIntegrityEnable = &dnsIntegrityEnable
		hasChanges = true
	}
	if d.HasChange("dns_integrity_member") {
		if v, ok := d.GetOk("dns_integrity_member"); ok && v != "" {
			updateZoneAuth.DNSIntegrityMember = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("external_primaries") {
		if v, ok := d.GetOk("external_primaries"); ok {
			servers := []map[string]interface{}{}
			for _, server := range v.([]interface{}) {
				servers = append(servers, server.(map[string]interface{}))
			}
			updateZoneAuth.ExternalPrimaries = util.BuildExternalServerListFromT(servers)
		}
		hasChanges = true
	}
	if d.HasChange("external_secondaries") {
		if v, ok := d.GetOk("external_secondaries"); ok {
			servers := []map[string]interface{}{}
			for _, server := range v.([]interface{}) {
				servers = append(servers, server.(map[string]interface{}))
			}
			updateZoneAuth.ExternalSecondaries = util.BuildExternalServerListFromT(servers)
		}
		hasChanges = true
	}
	if d.HasChange("grid_primary") {
		if v, ok := d.GetOk("grid_primary"); ok {
			servers := make([]map[string]interface{}, 0)
			for _, server := range v.([]interface{}) {
				servers = append(servers, server.(map[string]interface{}))
			}
			updateZoneAuth.GridPrimary = util.BuildMemberServerListFromT(servers)
		}
		hasChanges = true
	}
	if d.HasChange("grid_secondaries") {
		if v, ok := d.GetOk("grid_secondaries"); ok {
			servers := make([]map[string]interface{}, 0)
			for _, server := range v.([]interface{}) {
				servers = append(servers, server.(map[string]interface{}))
			}
			updateZoneAuth.GridSecondaries = util.BuildMemberServerListFromT(servers)
		}
		hasChanges = true
	}
	if d.HasChange("locked") {
		zoneLocked := d.Get("locked").(bool)
		updateZoneAuth.Locked = &zoneLocked
		hasChanges = true
	}
	if d.HasChange("ns_group") {
		if v, ok := d.GetOk("ns_group"); ok && v != "" {
			updateZoneAuth.NSGroup = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("copy_xfer_to_notify") {
		copyXferToNotify := d.Get("copy_xfer_to_notify").(bool)
		updateZoneAuth.CopyXferToNotify = &copyXferToNotify
		hasChanges = true
	}
	if d.HasChange("use_copy_xfer_to_notify") {
		useCopyXferToNotify := d.Get("use_copy_xfer_to_notify").(bool)
		updateZoneAuth.UseCopyXferNotify = &useCopyXferToNotify
		hasChanges = true
	}
	if d.HasChange("use_check_names_policy") {
		useCheckNamesPolicy := d.Get("use_check_names_policy").(bool)
		updateZoneAuth.UseCheckNamesPolicy = &useCheckNamesPolicy
		hasChanges = true
	}
	if d.HasChange("allow_update") {
		if v, ok := d.GetOk("allow_update"); ok && v != nil {
			updateZoneAuth.AllowUpdate = util.BuildAcList(v.([]interface{}))
		}
		hasChanges = true
	}
	if d.HasChange("allow_transfer") {
		if v, ok := d.GetOk("allow_transfer"); ok && v != nil {
			updateZoneAuth.AllowTransfer = util.BuildAcList(v.([]interface{}))
		}
		hasChanges = true
	}

	if hasChanges == true {
		updateAPI := zoneauth.NewUpdate(updateZoneAuth, returnFields)
		err := infobloxClient.Do(updateAPI)
		if err != nil {
			return fmt.Errorf("Infoblox Zone Auth Update Error: %+v", err)
		}
		if updateAPI.StatusCode() != http.StatusOK {
			return fmt.Errorf("Infoblox Zone Auth Update return code != 200")
		}
	}
	return resourceZoneAuthRead(d, m)
}

func resourceZoneAuthDelete(d *schema.ResourceData, m interface{}) error {

	returnFields := []string{"fqdn"}

	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	resourceReference := d.Id()
	getZoneAuthAPI := zoneauth.NewGetSingleZone(resourceReference, returnFields)

	err := infobloxClient.Do(getZoneAuthAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Delete Error when fetching resource: %+v", err)
	}
	if getZoneAuthAPI.StatusCode() == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	deleteAPI := zoneauth.NewDelete(resourceReference)
	err = infobloxClient.Do(deleteAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Delete - Error deleting resource %+v", err)
	}

	if deleteAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Delete - Error deleting resource %s - return code != 200 error: %s", resourceReference, deleteAPI.GetResponse())
	}

	d.SetId("")
	return nil
}
