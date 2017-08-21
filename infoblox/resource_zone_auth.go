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
			"zoneformat": {
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
			"dnsintegrityenable": {
				Type:        schema.TypeBool,
				Description: "If this is set to True, DNS integrity check is enabled for this zone",
				Optional:    true,
				Computed:    true,
			},
			"dnsintegritymember": {
				Type:        schema.TypeString,
				Description: "The Grid member that performs DNS integrity checks for this zone",
				Optional:    true,
				Computed:    true,
			},
			"externalprimaries":   util.ExternalServerListSchema(true, false),
			"externalsecondaries": util.ExternalServerListSchema(true, false),
			"gridprimary":         util.MemberServerListSchema(),
			"gridprimarysharedwithmsparentdelegation": {
				Type:        schema.TypeBool,
				Description: "Determines if the server is duplicated with parent delegation.cannot be updated, nor written",
				Optional:    true,
				Computed:    true,
			},
			"gridsecondaries": util.MemberServerListSchema(),
			"locked": {
				Type:        schema.TypeBool,
				Description: "If you enable this flag, other administrators cannot make conflicting changes",
				Optional:    true,
				Computed:    true,
			},
			"lockedby": {
				Type:        schema.TypeString,
				Description: "The name of a superuser or the administrator who locked this zone (read-only)",
				Computed:    true,
			},
			"networkview": {
				Type:        schema.TypeString,
				Description: "The name of the network view in which this zone resides (read-only)",
				Computed:    true,
			},
			"nsgroup": {
				Type:        schema.TypeString,
				Description: "The name server group that serves DNS for this zone.",
				Optional:    true,
				Computed:    true,
			},
			"soaserialnumber": {
				Type:        schema.TypeInt,
				Description: "The SOA serial number to be used in conjunction with set_soa_serial_number (read-only)",
				Computed:    true,
			},
			"soattl": {
				Type:         schema.TypeInt,
				Description:  "The Time to Live (TTL) value of the SOA record of this zone",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soanegativettl": {
				Type:         schema.TypeInt,
				Description:  "The negative Time to Live (TTL)",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soarefresh": {
				Type:         schema.TypeInt,
				Description:  "This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"soaretry": {
				Type:         schema.TypeInt,
				Description:  "This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs",
				Optional:     true,
				Computed:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
			},
			"allowupdate": {
				Type:        schema.TypeList,
				Description: "Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Description:  "Specifies the type of struct we're passing",
							Optional:     true,
							ValidateFunc: validateZoneAuthAllowUpdateType,
						},
						"address": {
							Type:         schema.TypeString,
							Description:  "The address this rule applies to or ANY",
							Optional:     true,
							ValidateFunc: util.CheckLeadingTrailingSpaces,
						},
						"permission": {
							Type:         schema.TypeString,
							Description:  "The permission to use for this address",
							Optional:     true,
							ValidateFunc: validateZoneAuthAllowUpdatePermission,
						},
						"tsigkey": {
							Type:         schema.TypeString,
							Description:  "A generated TSIG key",
							Optional:     true,
							ValidateFunc: util.CheckLeadingTrailingSpaces,
						},
						"tsigkeyalgorithm": {
							Type:         schema.TypeString,
							Description:  "The TSIG key algorithm",
							Optional:     true,
							ValidateFunc: util.ValidateTSIGAlgorithm,
						},
						"tsigkeyname": {
							Type:         schema.TypeString,
							Description:  "The name of the TSIG key",
							Optional:     true,
							ValidateFunc: util.CheckLeadingTrailingSpaces,
						},
						"usetsigkeyname": {
							Type:        schema.TypeBool,
							Description: "Use flag for: tsigkeyname",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func validateZoneAuthAllowUpdateType(v interface{}, k string) (ws []string, errors []error) {
	allowUpdateType := v.(string)
	if allowUpdateType != "addressac" && allowUpdateType != "tsigac" {
		errors = append(errors, fmt.Errorf("%q must be one of addressac or tsigac", k))
	}
	return
}

func validateZoneAuthAllowUpdatePermission(v interface{}, k string) (ws []string, errors []error) {
	permission := v.(string)
	if permission != "ALLOW" && permission != "DENY" {
		errors = append(errors, fmt.Errorf("%q must be one of ALLOW or DENY", k))
	}
	return
}

func buildAllowUpdateList(allowUpdateList []interface{}) []interface{} {

	allowUpdatesFrom := make([]interface{}, len(allowUpdateList))
	var addressAccessControl zoneauth.AddressAC
	var tsigAccessControl zoneauth.TsigAC

	for idx, value := range allowUpdateList {
		permission, ok := value.(map[string]interface{})
		if ok {
			if permission["type"] == "addressac" {
				addressAccessControl.StructType = permission["type"].(string)
				addressAccessControl.Address = permission["address"].(string)
				addressAccessControl.Permission = permission["permission"].(string)
				allowUpdatesFrom[idx] = addressAccessControl
			}
			if permission["type"] == "tsigac" {
				tsigAccessControl.StructType = permission["type"].(string)
				tsigAccessControl.TsigKey = permission["tsigkey"].(string)
				tsigAccessControl.TsigKeyAlg = permission["tsigkeyalgorithm"].(string)
				tsigAccessControl.TsigKeyName = permission["tsigkeyname"].(string)
				useTSIGKeyName := permission["usetsigkeyname"].(bool)
				tsigAccessControl.UseTsigKeyName = &useTSIGKeyName
				allowUpdatesFrom[idx] = tsigAccessControl
			}
		}
	}
	return allowUpdatesFrom
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
	if v, ok := d.GetOk("nsgroup"); ok && v != "" {
		dnsZone.NSGroup = v.(string)
	}
	if v, ok := d.GetOk("view"); ok && v != "" {
		dnsZone.View = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		dnsZone.Comment = v.(string)
	}
	if v, ok := d.GetOk("zoneformat"); ok && v != "" {
		dnsZone.ZoneFormat = v.(string)
	}
	if v, ok := d.GetOk("prefix"); ok && v != "" {
		dnsZone.Prefix = v.(string)
	}
	if v, ok := d.GetOk("disable"); ok {
		dnsZoneDisable := v.(bool)
		dnsZone.Disable = &dnsZoneDisable
	}
	if v, ok := d.GetOk("dnsintegrityenable"); ok {
		dnsIntegrityEnable := v.(bool)
		dnsZone.DNSIntegrityEnable = &dnsIntegrityEnable
	}
	if v, ok := d.GetOk("dnsintegritymember"); ok {
		dnsZone.DNSIntegrityMember = v.(string)
	}
	if v, ok := d.GetOk("externalprimaries"); ok {
		servers := []map[string]interface{}{}
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		dnsZone.ExternalPrimaries = util.BuildExternalServerListFromT(servers)
	}
	if v, ok := d.GetOk("externalsecondaries"); ok {
		servers := []map[string]interface{}{}
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		dnsZone.ExternalSecondaries = util.BuildExternalServerListFromT(servers)
	}
	if v, ok := d.GetOk("gridprimary"); ok {
		servers := make([]map[string]interface{}, 0)
		for _, server := range v.([]interface{}) {
			servers = append(servers, server.(map[string]interface{}))
		}
		dnsZone.GridPrimary = util.BuildMemberServerListFromT(servers)
	}
	if v, ok := d.GetOk("gridsecondaries"); ok {
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
	if v, ok := d.GetOk("allowupdate"); ok && v != nil {
		dnsZone.AllowUpdate = buildAllowUpdateList(v.([]interface{}))
	}

	if v, ok := d.GetOk("restart_if_needed"); ok {
		restart := v.(bool)
		appendDNSZone.RestartIfNeeded = &restart
	}

	// Some attributes can't be set on creation. They need to be sent in a subsequent request after initial creation.
	if v, ok := d.GetOk("soattl"); ok && v != nil {
		soaTTL := v.(int)
		appendDNSZone.SOADefaultTTL = uint(soaTTL)
		appendDNSZone.UseGridZoneTimer = &gridTimer
	}
	if v, ok := d.GetOk("soanegativettl"); ok && v != nil {
		soaNegativeTTL := v.(int)
		appendDNSZone.SOANegativeTTL = uint(soaNegativeTTL)
		appendDNSZone.UseGridZoneTimer = &gridTimer
	}
	if v, ok := d.GetOk("soarefresh"); ok && v != nil {
		soaRefresh := v.(int)
		appendDNSZone.SOARefresh = uint(soaRefresh)
		appendDNSZone.UseGridZoneTimer = &gridTimer
	}
	if v, ok := d.GetOk("soaretry"); ok && v != nil {
		soaRetry := v.(int)
		appendDNSZone.SOARetry = uint(soaRetry)
		appendDNSZone.UseGridZoneTimer = &gridTimer
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
	return []string{"fqdn", "comment", "zone_format", "view", "prefix", "soa_serial_number", "soa_default_ttl", "soa_negative_ttl", "soa_refresh", "soa_retry", "disable", "dns_integrity_enable", "dns_integrity_member", "external_primaries", "external_secondaries", "grid_primary", "grid_secondaries", "grid_primary_shared_with_ms_parent_delegation", "locked", "locked_by", "network_view", "ns_group", "allow_update"}
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
	d.Set("zoneformat", response.ZoneFormat)
	d.Set("prefix", response.Prefix)
	d.Set("soattl", response.SOADefaultTTL)
	d.Set("soanegativettl", response.SOANegativeTTL)
	d.Set("soarefresh", response.SOARefresh)
	d.Set("soaretry", response.SOARetry)
	d.Set("soaserialnumber", response.SOASerialNumber)
	d.Set("disable", response.Disable)
	d.Set("dnsintegrityenable", response.DNSIntegrityEnable)
	d.Set("dnsintegritymember", response.DNSIntegrityMember)
	d.Set("externalprimaries", util.BuildExternalServersListFromIBX(response.ExternalPrimaries))
	d.Set("externalsecondaries", util.BuildExternalServersListFromIBX(response.ExternalSecondaries))
	d.Set("gridprimarysharedwithmsparentdelegation", response.GridPrimarySharedWithMSParentDelegation)
	d.Set("gridprimary", util.BuildMemberServerListFromIBX(response.GridPrimary))
	d.Set("gridsecondaries", util.BuildMemberServerListFromIBX(response.GridSecondaries))
	d.Set("locked", response.Locked)
	d.Set("lockedby", response.LockedBy)
	d.Set("networkview", response.NetworkView)
	d.Set("nsgroup", response.NSGroup)
	d.Set("allowupdate", response.AllowUpdate)
	restart, _ := d.GetOk("restart_if_needed")
	d.Set("restart_if_needed", &restart)

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
	if d.HasChange("soattl") {
		if v, ok := d.GetOk("soattl"); ok && v != nil {
			soaTTL := v.(int)
			updateZoneAuth.SOADefaultTTL = uint(soaTTL)
			updateZoneAuth.UseGridZoneTimer = &gridTimer
		}
		hasChanges = true
	}
	if d.HasChange("soanegativettl") {
		if v, ok := d.GetOk("soanegativettl"); ok && v != nil {
			soaNegativeTTL := v.(int)
			updateZoneAuth.SOANegativeTTL = uint(soaNegativeTTL)
			updateZoneAuth.UseGridZoneTimer = &gridTimer
		}
		hasChanges = true
	}
	if d.HasChange("soarefresh") {
		if v, ok := d.GetOk("soarefresh"); ok && v != nil {
			soaRefresh := v.(int)
			updateZoneAuth.SOARefresh = uint(soaRefresh)
			updateZoneAuth.UseGridZoneTimer = &gridTimer
		}
		hasChanges = true
	}
	if d.HasChange("soaretry") {
		if v, ok := d.GetOk("soaretry"); ok && v != nil {
			soaRetry := v.(int)
			updateZoneAuth.SOARetry = uint(soaRetry)
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
	if d.HasChange("dnsintegrityenable") {
		dnsIntegrityEnable := d.Get("dnsintegrityenable").(bool)
		updateZoneAuth.DNSIntegrityEnable = &dnsIntegrityEnable
		hasChanges = true
	}
	if d.HasChange("dnsintegritymember") {
		if v, ok := d.GetOk("dnsintegritymember"); ok && v != "" {
			updateZoneAuth.DNSIntegrityMember = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("externalprimaries") {
		if v, ok := d.GetOk("externalprimaries"); ok {
			servers := []map[string]interface{}{}
			for _, server := range v.([]interface{}) {
				servers = append(servers, server.(map[string]interface{}))
			}
			updateZoneAuth.ExternalPrimaries = util.BuildExternalServerListFromT(servers)
		}
		hasChanges = true
	}
	if d.HasChange("externalsecondaries") {
		if v, ok := d.GetOk("externalsecondaries"); ok {
			servers := []map[string]interface{}{}
			for _, server := range v.([]interface{}) {
				servers = append(servers, server.(map[string]interface{}))
			}
			updateZoneAuth.ExternalSecondaries = util.BuildExternalServerListFromT(servers)
		}
		hasChanges = true
	}
	if d.HasChange("gridprimary") {
		if v, ok := d.GetOk("gridprimary"); ok {
			servers := make([]map[string]interface{}, 0)
			for _, server := range v.([]interface{}) {
				servers = append(servers, server.(map[string]interface{}))
			}
			updateZoneAuth.GridPrimary = util.BuildMemberServerListFromT(servers)
		}
		hasChanges = true
	}
	if d.HasChange("gridsecondaries") {
		if v, ok := d.GetOk("gridsecondaries"); ok {
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
	if d.HasChange("nsgroup") {
		if v, ok := d.GetOk("nsgroup"); ok && v != "" {
			updateZoneAuth.NSGroup = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("allowupdate") {
		if v, ok := d.GetOk("allowupdate"); ok && v != nil {
			updateZoneAuth.AllowUpdate = buildAllowUpdateList(v.([]interface{}))
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
