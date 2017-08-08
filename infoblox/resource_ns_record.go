package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/nameserver"
	"net/http"
	"strings"
)

func resourceNSRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSRecordCreate,
		Read:   resourceNSRecordRead,
		Update: resourceNSRecordUpdate,
		Delete: resourceNSRecordDelete,

		Schema: map[string]*schema.Schema{
			"zone_name": {
				Type:         schema.TypeString,
				Description:  "The name of the zone where the record should reside",
				Required:     true,
				ValidateFunc: validateNSRecordNoTrailingLeadingWhiteSpace,
			},
			"ms_delegation_name": {
				Type:        schema.TypeString,
				Description: "The MS delegation point name",
				Optional:    true,
			},
			"name_server": {
				Type:         schema.TypeString,
				Description:  "The FQDN of the name server",
				Required:     true,
				ValidateFunc: validateNSRecordNoTrailingLeadingWhiteSpace,
			},
			"view": {
				Type:         schema.TypeString,
				Description:  "The name of the DNS view in which the record resides",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateNSRecordNoTrailingLeadingWhiteSpace,
			},
			"name_server_addresses": {
				Type:        schema.TypeSet,
				Description: "The list of zone name servers",
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Description: "The IP address of the Zone Name Server",
							Required:    true,
						},
						"auto_create_PTR_record": &schema.Schema{
							Type:        schema.TypeBool,
							Description: "Flag to indicate if PTR records need to be auto created",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func validateNSRecordNoTrailingLeadingWhiteSpace(v interface{}, k string) (ws []string, errors []error) {
	stringToCheck := v.(string)
	trimmedString := strings.Trim(stringToCheck, " ")
	if trimmedString != stringToCheck {
		errors = append(errors, fmt.Errorf("%q must not contain trailing or leading white space", k))
	}
	return
}

func buildNSRecordZoneNameServerList(addresses *schema.Set) []nameserver.ZoneNameServer {

	var zoneNameServer nameserver.ZoneNameServer
	zoneNameServers := make([]nameserver.ZoneNameServer, 0)

	for _, addressItem := range addresses.List() {
		nsAddressItem := addressItem.(map[string]interface{})
		zoneNameServer.Address = nsAddressItem["ip_address"].(string)
		createPTR := nsAddressItem["auto_create_PTR_record"].(bool)
		zoneNameServer.AutoCreatePointerRecord = &createPTR
		zoneNameServers = append(zoneNameServers, zoneNameServer)
	}
	return zoneNameServers
}

func buildNSRecordNameServerAddresses(addresses []nameserver.ZoneNameServer) []map[string]interface{} {

	nameServerAddresses := make([]map[string]interface{}, len(addresses))

	for idx, addressItem := range addresses {
		nameServerAddress := make(map[string]interface{})
		nameServerAddress["ip_address"] = addressItem.Address
		nameServerAddress["auto_create_PTR_record"] = *addressItem.AutoCreatePointerRecord
		nameServerAddresses[idx] = nameServerAddress
	}
	return nameServerAddresses
}

func resourceNSRecordCreate(d *schema.ResourceData, m interface{}) error {

	var nsRecord nameserver.NSRecord
	client := m.(*skyinfoblox.InfobloxClient)

	if v, ok := d.GetOk("zone_name"); ok && v != "" {
		nsRecord.Name = v.(string)
	}
	if v, ok := d.GetOk("ms_delegation_name"); ok && v != "" {
		nsRecord.MSDelegationName = v.(string)
	}
	if v, ok := d.GetOk("name_server"); ok && v != "" {
		nsRecord.NameServer = v.(string)
	}
	if v, ok := d.GetOk("view"); ok && v != "" {
		nsRecord.View = v.(string)
	}
	if v, ok := d.GetOk("name_server_addresses"); ok && v != nil {
		nsRecord.Addresses = buildNSRecordZoneNameServerList(v.(*schema.Set))
	}

	createAPI := nameserver.NewCreate(nsRecord)
	err := client.Do(createAPI)
	httpStatus := createAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Record Create for %s failed with status code %d and error: %+v", nsRecord.NameServer, httpStatus, err)
	}
	nsRecord.Reference = *createAPI.ResponseObject().(*string)

	d.SetId(nsRecord.Reference)
	return resourceNSRecordRead(d, m)
}

func resourceNSRecordRead(d *schema.ResourceData, m interface{}) error {

	returnFields := []string{"name", "addresses", "nameserver", "view", "ms_delegation_name"}
	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	getNSRecordAPI := nameserver.NewGet(reference, returnFields)
	err := client.Do(getNSRecordAPI)
	httpStatus := getNSRecordAPI.StatusCode()

	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Record read for %s failed with status code %d and error: %+v", reference, httpStatus, err)
	}

	response := *getNSRecordAPI.ResponseObject().(*nameserver.NSRecord)

	d.SetId(response.Reference)
	d.Set("zone_name", response.Name)
	d.Set("ms_delegation_name", response.MSDelegationName)
	d.Set("name_server", response.NameServer)
	d.Set("view", response.View)
	d.Set("name_server_addresses", buildNSRecordNameServerAddresses(response.Addresses))

	return nil
}

func resourceNSRecordUpdate(d *schema.ResourceData, m interface{}) error {

	var nsRecordObject nameserver.NSRecord
	hasChanges := false

	if d.HasChange("zone_name") {
		if v, ok := d.GetOk("zone_name"); ok && v != "" {
			nsRecordObject.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("ms_delegation_name") {
		if v, ok := d.GetOk("ms_delegation_name"); ok && v != "" {
			nsRecordObject.MSDelegationName = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("name_server") {
		if v, ok := d.GetOk("name_server"); ok && v != "" {
			nsRecordObject.NameServer = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("view") {
		if v, ok := d.GetOk("view"); ok && v != "" {
			nsRecordObject.View = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("name_server_addresses") {
		if v, ok := d.GetOk("name_server_addresses"); ok && v != nil {
			nsRecordObject.Addresses = buildNSRecordZoneNameServerList(v.(*schema.Set))
		}
		hasChanges = true
	}

	if hasChanges {
		returnFields := []string{"name", "addresses", "nameserver", "view", "ms_delegation_name"}
		client := m.(*skyinfoblox.InfobloxClient)
		nsRecordObject.Reference = d.Id()

		updateNSRecordAPI := nameserver.NewUpdate(nsRecordObject, returnFields)
		err := client.Do(updateNSRecordAPI)
		httpStatus := updateNSRecordAPI.StatusCode()

		if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox NS Record update for %s failed with status code %d and error: %+v", nsRecordObject.Reference, httpStatus, err)
		}

		response := *updateNSRecordAPI.ResponseObject().(*nameserver.NSRecord)

		d.SetId(response.Reference)
		d.Set("zone_name", response.Name)
		d.Set("ms_delegation_name", response.MSDelegationName)
		d.Set("name_server", response.NameServer)
		d.Set("view", response.View)
		d.Set("name_server_addresses", buildNSRecordNameServerAddresses(response.Addresses))
	}

	return resourceNSRecordRead(d, m)
}

func resourceNSRecordDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	deleteNSRecordAPI := nameserver.NewDelete(reference)
	err := client.Do(deleteNSRecordAPI)
	httpStatus := deleteNSRecordAPI.StatusCode()

	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Record delete for %s failed with status code %d and error: %+v", reference, httpStatus, err)
	}

	return nil
}
