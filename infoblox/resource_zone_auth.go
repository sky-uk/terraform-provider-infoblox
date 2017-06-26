package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"net/http"
	"strings"
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
				ValidateFunc: validateZoneAuthCheckLeadingTrailingSpaces,
			},
			"view": {
				Type:         schema.TypeString,
				Description:  "The name of the DNS view in which the zone resides",
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validateZoneAuthCheckLeadingTrailingSpaces,
			},
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment for the zone; maximum 256 characters",
				Optional:     true,
				ValidateFunc: validateZoneAuthCheckLeadingTrailingSpaces,
			},
			"zoneformat": {
				Type:         schema.TypeString,
				Description:  "Determines the format of this zone - API default FORWARD",
				ValidateFunc: validateZoneFormat,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
			},
			"prefix": {
				Type:         schema.TypeString,
				Description:  "The RFC2317 prefix value of this DNS zone",
				Optional:     true,
				ValidateFunc: validateZoneAuthCheckLeadingTrailingSpaces,
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
				ValidateFunc: validateZoneAuthUnsignedInteger,
			},
			"soanegativettl": {
				Type:         schema.TypeInt,
				Description:  "The negative Time to Live (TTL)",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateZoneAuthUnsignedInteger,
			},
			"soarefresh": {
				Type:         schema.TypeInt,
				Description:  "This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateZoneAuthUnsignedInteger,
			},
			"soaretry": {
				Type:         schema.TypeInt,
				Description:  "This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateZoneAuthUnsignedInteger,
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
							ValidateFunc: validateZoneAuthCheckLeadingTrailingSpaces,
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
							ValidateFunc: validateZoneAuthCheckLeadingTrailingSpaces,
						},
						"tsigkeyalgorithm": {
							Type:         schema.TypeString,
							Description:  "The TSIG key algorithm",
							Optional:     true,
							ValidateFunc: validateZoneAuthAllowUpdateTSIGAlgorithm,
						},
						"tsigkeyname": {
							Type:         schema.TypeString,
							Description:  "The name of the TSIG key",
							Optional:     true,
							ValidateFunc: validateZoneAuthCheckLeadingTrailingSpaces,
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

func validateZoneFormat(v interface{}, k string) (ws []string, errors []error) {
	zoneFormat := v.(string)
	if zoneFormat != "FORWARD" && zoneFormat != "IPV4" && zoneFormat != "IPV6" {
		errors = append(errors, fmt.Errorf("%q must be one of FORWARD, IPV4 or IPV6", k))
	}
	return
}

func validateZoneAuthUnsignedInteger(v interface{}, k string) (ws []string, errors []error) {
	ttl := v.(int)
	if ttl < 0 {
		errors = append(errors, fmt.Errorf("%q can't be negative", k))
	}
	return
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

func validateZoneAuthAllowUpdateTSIGAlgorithm(v interface{}, k string) (ws []string, errors []error) {
	tsigAlgorithm := v.(string)
	if tsigAlgorithm != "HMAC-MD5" && tsigAlgorithm != "HMAC-SHA256" {
		errors = append(errors, fmt.Errorf("%q must be one of HMAC-MD5 or HMAC-SHA256", k))
	}
	return
}

func validateZoneAuthCheckLeadingTrailingSpaces(v interface{}, k string) (ws []string, errors []error) {
	stringToCheck := v.(string)
	trimedString := strings.Trim(stringToCheck, " ")
	if trimedString != stringToCheck {
		errors = append(errors, fmt.Errorf("%q must not contain trailing or leading white space", k))
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
	if v, ok := d.GetOk("locked"); ok {
		zoneLocked := v.(bool)
		dnsZone.Locked = &zoneLocked
	}
	if v, ok := d.GetOk("allowupdate"); ok && v != nil {
		dnsZone.AllowUpdate = buildAllowUpdateList(v.([]interface{}))
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
		return fmt.Errorf("Infoblox Zone Create Error: Invalid HTTP response code %d returned", createAPI.StatusCode())
	}

	ref := createAPI.GetResponse()
	// We can't set some attributes on create. Therefore we need to make another call.
	appendDNSZone.Reference = ref
	appendAPI := zoneauth.NewUpdate(appendDNSZone, nil)
	err = infobloxClient.Do(appendAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Zone Auth Create Append Error: %+v", err)
	}
	if appendAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Zone Auth Create Append: Invalid HTTP resposne code %d returned", appendAPI.StatusCode())
	}

	d.SetId(ref)
	return resourceZoneAuthRead(d, m)
}

func resourceZoneAuthRead(d *schema.ResourceData, m interface{}) error {

	returnFields := []string{"fqdn", "comment", "zone_format", "view", "prefix", "soa_default_ttl", "soa_negative_ttl", "soa_refresh", "soa_retry", "soa_serial_number", "disable", "dns_integrity_enable", "dns_integrity_member", "locked", "locked_by", "network_view", "allow_update"}
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
	d.Set("locked", response.DNSIntegrityMember)
	d.Set("lockedby", response.LockedBy)
	d.Set("networkview", response.NetworkView)
	d.Set("allowupdate", response.AllowUpdate)

	return nil
}

func resourceZoneAuthUpdate(d *schema.ResourceData, m interface{}) error {

	var updateZoneAuth zoneauth.DNSZone
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	hasChanges := false
	returnFields := []string{"fqdn", "comment", "prefix", "soa_default_ttl", "soa_negative_ttl", "soa_refresh", "soa_retry", "disable", "dns_integrity_enable", "dns_integrity_member", "locked", "allow_update"}
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
	if d.HasChange("locked") {
		zoneLocked := d.Get("locked").(bool)
		updateZoneAuth.Locked = &zoneLocked
		hasChanges = true
	}
	if d.HasChange("allowupdate") {
		if v, ok := d.GetOk("allowupdate"); ok && v != nil {
			updateZoneAuth.AllowUpdate = buildAllowUpdateList(v.([]interface{}))
		}
		hasChanges = true
	}

	if hasChanges {
		updateAPI := zoneauth.NewUpdate(updateZoneAuth, returnFields)
		err := infobloxClient.Do(updateAPI)
		if err != nil {
			return fmt.Errorf("Infoblox Zone Auth Update Error: %+v", err)
		}
		if updateAPI.StatusCode() != http.StatusOK {
			return fmt.Errorf("Infoblox Zone Auth Update return code != 200")
		}

		response := updateAPI.GetResponse()
		d.SetId(response.Reference)
		d.Set("comment", response.Comment)
		d.Set("prefix", response.Prefix)
		d.Set("soattl", response.SOADefaultTTL)
		d.Set("soanegativettl", response.SOANegativeTTL)
		d.Set("soarefresh", response.SOARefresh)
		d.Set("soaretry", response.SOARetry)
		d.Set("disable", *response.Disable)
		d.Set("dnsintegrityenable", *response.DNSIntegrityEnable)
		d.Set("dnsintegritymember", response.DNSIntegrityMember)
		d.Set("locked", *response.Locked)
		d.Set("allowupdate", response.AllowUpdate)
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
		return fmt.Errorf("Infoblox Delete - Error deleting resource %s - return code != 200", resourceReference)
	}

	d.SetId("")
	return nil
}
