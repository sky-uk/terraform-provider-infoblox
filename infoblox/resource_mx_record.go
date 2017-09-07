package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/mxrecord"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
	"time"
)

func resourceMxRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceMxRecordCreate,
		Read:   resourceMxRecordRead,
		Update: resourceMxRecordUpdate,
		Delete: resourceMxRecordDelete,

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(2 * time.Minute),
			Create: schema.DefaultTimeout(2 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "Name of the zone the MX record refers to",
			},
			"comment": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "A comment on the record",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Is the record disabled",
			},
			"ddns_principal": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The GSS-TSIG principal that owns this record",
			},
			"ddns_protected": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if the DDNS updates for this record are allowed or not",
			},
			"mail_exchanger": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "Mail exchanger name in FQDN format. This value can be in unicode format",
			},
			"preference": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Preference value, 0 to 65535 (inclusive) in 32-bit unsigned integer format",
			},
			"ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: util.ValidateUnsignedInteger,
				Description:  "The Time To Live value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid ",
			},
			"use_ttl": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use flag for: ttl",
			},
			"view": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
				Description:  "The name of the DNS view in which the record resides.",
			},
		},
	}

}

func resourceMxRecordCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var mxRecord mxrecord.MxRecord
	if v, ok := d.GetOk("name"); ok {
		mxRecord.Name = v.(string)
	} else {
		mxRecord.Name = d.Get("name").(string)
	}
	if v, ok := d.GetOk("comment"); ok {
		mxRecord.Comment = v.(string)
	}
	if _, ok := d.GetOk("disable"); ok {
		mxRecord.Disable = d.Get("disable").(bool)
	}
	if v, ok := d.GetOk("ddns_principal"); ok {
		mxRecord.DDNSPrincipal = v.(string)
	}
	if v, ok := d.GetOk("ddns_protected"); ok {
		mxRecord.DDNSProtected = v.(bool)
	}
	if v, ok := d.GetOk("mail_exchanger"); ok {
		mxRecord.MailExchanger = v.(string)
	}
	if v, ok := d.GetOk("preference"); ok {
		mxRecord.Preference = uint(v.(int))
	}
	if v, ok := d.GetOk("ttl"); ok {
		mxRecord.TTL = uint(v.(int))
	}
	if _, ok := d.GetOk("use_ttl"); ok {
		mxRecord.UseTTL = d.Get("use_ttl").(bool)
	}
	if v, ok := d.GetOk("view"); ok {
		mxRecord.View = v.(string)
	}

	mxRecordCreateAPI := mxrecord.NewCreate(mxRecord)
	createMXRecordErr := infobloxClient.Do(mxRecordCreateAPI)
	if createMXRecordErr != nil {
		return fmt.Errorf(createMXRecordErr.Error())
	}
	if mxRecordCreateAPI.StatusCode() != 201 {
		return fmt.Errorf("Infoblox Create Error : %d", mxRecordCreateAPI.StatusCode())
	}
	d.SetId(*mxRecordCreateAPI.ResponseObject().(*string))
	return resourceMxRecordRead(d, m)
}

func returnMXRecordFields() []string {
	return []string{"name", "comment", "disable", "ddns_principal", "ddns_protected", "mail_exchanger", "preference", "ttl", "use_ttl", "view"}
}

func resourceMxRecordRead(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	readMXRecordAPI := mxrecord.NewGet(d.Id(), returnMXRecordFields())
	readMXRecordErr := infobloxClient.Do(readMXRecordAPI)
	if readMXRecordErr != nil {
		return fmt.Errorf(readMXRecordErr.Error())
	}
	if readMXRecordAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Read Error : %d", readMXRecordAPI.StatusCode())

	}
	readMXRecord := readMXRecordAPI.ResponseObject().(*mxrecord.MxRecord)
	d.SetId(readMXRecord.Ref)
	d.Set("name", readMXRecord.Name)
	d.Set("comment", readMXRecord.Comment)
	d.Set("disable", readMXRecord.Disable)
	d.Set("ddns_principal", readMXRecord.DDNSPrincipal)
	d.Set("ddns_protected", readMXRecord.DDNSProtected)
	d.Set("mail_exchanger", readMXRecord.MailExchanger)
	d.Set("preference", readMXRecord.Preference)
	d.Set("ttl", readMXRecord.TTL)
	d.Set("use_ttl", readMXRecord.UseTTL)
	d.Set("view", readMXRecord.View)
	return nil
}

func resourceMxRecordUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var updateMXRecord mxrecord.MxRecord
	var hasChanges bool
	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok {
			updateMXRecord.Name = v.(string)
			hasChanges = true
		}
	} else {
		updateMXRecord.Name = d.Get("name").(string)
	}

	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok {
			updateMXRecord.Comment = v.(string)
			hasChanges = true
		}
	}
	updateMXRecord.Disable = d.Get("disable").(bool)
	oldDisable, newDisable := d.GetChange("disable")
	if oldDisable != newDisable {
		hasChanges = true
	}

	if d.HasChange("ddns_principal") {
		if v, ok := d.GetOk("ddns_principal"); ok {
			updateMXRecord.DDNSPrincipal = v.(string)
			hasChanges = true
		}
	}

	oldDDNSProtected, newDDNSProtected := d.GetChange("ddns_protected")
	if oldDDNSProtected != newDDNSProtected {
		updateMXRecord.DDNSProtected = d.Get("ddns_protected").(bool)
	}

	if d.HasChange("mail_exchanger") {
		if v, ok := d.GetOk("mail_exchanger"); ok {
			updateMXRecord.MailExchanger = v.(string)
			hasChanges = true
		}
	} else {
		updateMXRecord.MailExchanger = d.Get("mail_exchanger").(string)
	}

	if d.HasChange("preference") {
		if v, ok := d.GetOk("preference"); ok {
			updateMXRecord.Preference = uint(v.(int))
			hasChanges = true
		}
	} else {
		updateMXRecord.Preference = uint(d.Get("preference").(int))
	}

	if d.HasChange("ttl") {
		if v, ok := d.GetOk("ttl"); ok {
			updateMXRecord.TTL = uint(v.(int))
			hasChanges = true
		}
	}
	updateMXRecord.UseTTL = d.Get("use_ttl").(bool)
	oldUseTTL, newUseTTL := d.GetChange("use_ttl")
	if oldUseTTL != newUseTTL {
		hasChanges = true
	}

	if d.HasChange("view") {
		if v, ok := d.GetOk("view"); ok {
			updateMXRecord.View = v.(string)
			hasChanges = true
		}
	}

	if hasChanges {
		updateMXRecordAPI := mxrecord.NewUpdate(d.Id(), updateMXRecord)
		updateMXRecordErr := infobloxClient.Do(updateMXRecordAPI)
		if updateMXRecordErr != nil {
			return fmt.Errorf("Error Updating : %s", updateMXRecordErr.Error())
		}
		if updateMXRecordAPI.StatusCode() != http.StatusOK {
			return fmt.Errorf("Error Updating : %d - %s", updateMXRecordAPI.StatusCode(), *updateMXRecordAPI.ResponseObject().(*string))
		}
		d.SetId(*updateMXRecordAPI.ResponseObject().(*string))
		return resourceMxRecordRead(d, m)
	}
	return nil
}

func resourceMxRecordDelete(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	mxRecordDeleteAPI := mxrecord.NewDelete(d.Id())
	deleteMXRecordErr := infobloxClient.Do(mxRecordDeleteAPI)
	if deleteMXRecordErr != nil {
		return fmt.Errorf(deleteMXRecordErr.Error())
	}
	return nil
}
