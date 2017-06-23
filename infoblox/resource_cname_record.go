package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"strings"
)

func resourceCNAMERecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceCNAMECreate,
		Read:   resourceCNAMERead,
		Update: resourceCNAMEUpdate,
		Delete: resourceCNAMEDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for a CNAME record in FQDN format",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique reference to Infoblox resource",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the record; maximum 256 characters",
			},
			"view": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "The name of the DNS View in which the record resides",
			},
			"ttl": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateUnsignedInteger,
				Description:  "The Time To Live assigned to CNAME",
			},
			"canonical": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Canonical name in FQDN format",
			},
		},
	}
}

func validateUnsignedInteger(v interface{}, k string) (ws []string, errors []error) {
	ttl := v.(int)
	if ttl < 0 {
		errors = append(errors, fmt.Errorf("%q can't be negative", k))
	}
	return
}

func resourceCNAMECreate(d *schema.ResourceData, m interface{}) error {

	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var cnameRecord records.GenericRecord
	recordType := "cname"

	if v, ok := d.GetOk("name"); ok {
		cnameRecord.Name = v.(string)
	} else {
		return fmt.Errorf("Infoblox Create Error: name argument required")
	}
	if v, ok := d.GetOk("comment"); ok {
		cnameRecord.Comment = v.(string)
	}
	if v, ok := d.GetOk("view"); ok {
		cnameRecord.View = v.(string)
	}
	if v, ok := d.GetOk("ttl"); ok {
		ttl := v.(int)
		cnameRecord.TTL = uint(ttl)
	}
	if v, ok := d.GetOk("canonical"); ok {
		cnameRecord.Canonical = v.(string)
	}

	createAPI := records.NewCreateRecord(recordType, cnameRecord)

	err := infobloxClient.Do(createAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Create Error: %+v", err)
	}

	if createAPI.StatusCode() != 201 {
		return fmt.Errorf("Infoblox Create Error: Invalid HTTP response code %+v returned. Response object was %+v", createAPI.StatusCode(), createAPI.ResponseObject())
	}

	id := strings.Replace(createAPI.GetResponse(), "\"", "", -1)
	d.SetId(id)
	return resourceCNAMERead(d, m)
}

func resourceCNAMERead(d *schema.ResourceData, m interface{}) error {

	returnFields := []string{"name", "comment", "view", "ttl", "canonical"}

	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	resourceReference := d.Id()
	getSingleCNAMEAPI := records.NewGetCNAMERecord(resourceReference, returnFields)

	err := infobloxClient.Do(getSingleCNAMEAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Read Error: %+v", err)
	}
	if getSingleCNAMEAPI.StatusCode() == 404 {
		d.SetId("")
		return nil
	}

	response := getSingleCNAMEAPI.GetResponse()
	d.SetId(response.Ref)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("view", response.View)
	d.Set("ttl", response.TTL)
	d.Set("canonical", response.Canonical)

	return nil
}

func resourceCNAMEUpdate(d *schema.ResourceData, m interface{}) error {

	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	hasChanges := false
	resourceReference := d.Id()
	var updateCNAME records.GenericRecord

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok {
			updateCNAME.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok {
			updateCNAME.Comment = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("view") {
		if v, ok := d.GetOk("view"); ok {
			updateCNAME.View = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("ttl") {
		if v, ok := d.GetOk("ttl"); ok {
			ttl := v.(int)
			updateCNAME.TTL = uint(ttl)
		}
		hasChanges = true
	}
	if d.HasChange("canonical") {
		if v, ok := d.GetOk("canonical"); ok {
			updateCNAME.Canonical = v.(string)
		}
		hasChanges = true
	}

	if hasChanges {
		updateAPI := records.NewUpdateRecord(resourceReference, updateCNAME)
		err := infobloxClient.Do(updateAPI)
		if err != nil {
			return fmt.Errorf("Infoblox Update Error: %+v", err)
		}
		if updateAPI.StatusCode() != 200 {
			return fmt.Errorf("Infoblox Update Error: Invalid HTTP response code %+v returned. Response was %+v", updateAPI.StatusCode(), updateAPI.GetResponse())
		}
		id := strings.Replace(updateAPI.GetResponse(), "\"", "", -1)
		d.SetId(id)
	}

	return resourceCNAMERead(d, m)
}

func resourceCNAMEDelete(d *schema.ResourceData, m interface{}) error {

	returnFields := []string{}
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	resourceReference := d.Id()
	getSingleCNAMEAPI := records.NewGetCNAMERecord(resourceReference, returnFields)

	err := infobloxClient.Do(getSingleCNAMEAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Delete Error when fetching resource: %+v", err)
	}
	if getSingleCNAMEAPI.StatusCode() == 404 {
		d.SetId("")
		return nil
	}

	deleteAPI := records.NewDelete(resourceReference)
	err = infobloxClient.Do(deleteAPI)
	if err != nil {
		return fmt.Errorf("Infobox Delete - Error deleting resource %+v", err)
	}
	if deleteAPI.StatusCode() != 200 {
		return fmt.Errorf("Infoblox Delete - Error deleting resource %s - return code != 200", resourceReference)
	}

	d.SetId("")
	return nil
}
