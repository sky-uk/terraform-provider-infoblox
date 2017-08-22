package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"log"
)

func resourceSRVRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceSRVRecordCreate,
		Read:   resourceSRVRecordRead,
		Update: resourceSRVRecordUpdate,
		Delete: resourceSRVRecordDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name for a SRV record in FQDN format",
			},
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique reference to Infoblox resource",
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"target": {
				Type:     schema.TypeString,
				Required: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Required: true,
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
				ValidateFunc: validateUnsignedInteger,
				Description:  "The Time To Live assigned to CNAME",
			},
			"use_ttl": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for the record; maximum 256 characters",
			},
		},
	}
}

func resourceSRVRecordCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var srvRecord records.GenericRecord
	recordType := "srv"

	if v, ok := d.GetOk("name"); ok {
		srvRecord.Name = v.(string)
	} else {
		return fmt.Errorf("Infoblox Create Error: name argument required")
	}
	if v, ok := d.GetOk("port"); ok {
		srvRecord.Port = v.(int)
	} else {
		return fmt.Errorf("Infoblox Create Error: port argument required")
	}
	if v, ok := d.GetOk("priority"); ok {
		srvRecord.Priority = v.(int)
	} else {
		return fmt.Errorf("Infoblox Create Error: priority argument required")
	}
	if v, ok := d.GetOk("target"); ok {
		srvRecord.Target = v.(string)
	} else {
		return fmt.Errorf("Infoblox Create Error: target argument required")
	}
	if v, ok := d.GetOk("weight"); ok {
		srvRecord.Weight = v.(int)
	} else {
		return fmt.Errorf("Infoblox Create Error: weight argument required")
	}
	if v, ok := d.GetOk("view"); ok {
		srvRecord.View = v.(string)
	}
	if v, ok := d.GetOk("ttl"); ok {
		ttl := v.(int)
		srvRecord.TTL = uint(ttl)
	}

	useTTL := false
	if v, ok := d.GetOk("use_ttl"); ok {
		useTTL = v.(bool)
	}
	srvRecord.UseTTL = &useTTL

	if v, ok := d.GetOk("comment"); ok {
		srvRecord.Comment = v.(string)
	}

	createAPI := records.NewCreateRecord(recordType, srvRecord)

	err := infobloxClient.Do(createAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Create Error: %+v", err)
	}

	if createAPI.StatusCode() != 201 {
		return fmt.Errorf("Infoblox Create Error: Invalid HTTP response code %+v returned. Response object was %+v", createAPI.StatusCode(), createAPI.GetResponse())
	}

	id := createAPI.GetResponse()
	d.SetId(id)

	return resourceSRVRecordRead(d, m)
}

func resourceSRVRecordRead(d *schema.ResourceData, m interface{}) error {
	returnFields := []string{"name", "comment", "port", "priority", "target", "weight", "zone", "use_ttl", "ttl"}

	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	resourceReference := d.Id()
	getSingleSRVAPI := records.NewGetSRVRecord(resourceReference, returnFields)

	err := infobloxClient.Do(getSingleSRVAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Read Error: %+v", err)
	}
	if getSingleSRVAPI.StatusCode() != 200 {
		d.SetId("")
		return fmt.Errorf("Infoblox Read Error: Invalid HTTP response code %+v returned. Response object was %+v", getSingleSRVAPI.StatusCode(), getSingleSRVAPI.GetResponse())
	}

	response := getSingleSRVAPI.GetResponse()
	d.SetId(response.Ref)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("port", response.Port)
	d.Set("priority", response.Priority)
	d.Set("target", response.Target)
	d.Set("weight", response.Weight)
	d.Set("zone", response.Zone)
	d.Set("ttl", response.TTL)
	d.Set("ref", response.Ref)

	return nil
}

func resourceSRVRecordUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	hasChanges := false
	var recordReference string
	var updatedSVR records.GenericRecord

	if v, ok := d.GetOk("ref"); ok {
		recordReference = v.(string)

	} else {
		return fmt.Errorf("cannot delete without reference")
	}

	fields := []string{""}
	getSVRRecordAPI := records.NewGetSRVRecord(recordReference, fields)
	getSVRErr := infobloxClient.Do(getSVRRecordAPI)
	if getSVRErr != nil {
		return fmt.Errorf("Unable to read the specefied SVR record")
	}

	if d.HasChange("port") {
		hasChanges = true
		_, newPort := d.GetChange("port")
		updatedSVR.Port = newPort.(int)
	}

	if d.HasChange("priority") {
		hasChanges = true
		_, newPriority := d.GetChange("priority")
		updatedSVR.Priority = newPriority.(int)
	}

	if d.HasChange("weight") {
		hasChanges = true
		_, newWeight := d.GetChange("weight")
		updatedSVR.Weight = newWeight.(int)
	}

	if d.HasChange("target") {
		hasChanges = true
		_, newTarget := d.GetChange("target")
		updatedSVR.Target = newTarget.(string)
	}

	if d.HasChange("ttl") {
		hasChanges = true
		var TTL int
		_, newTTL := d.GetChange("ttl")
		TTL = newTTL.(int)
		updatedSVR.TTL = uint(TTL)
	}

	useTTL := false
	if d.HasChange("use_ttl") {
		hasChanges = true
		value := d.Get("use_ttl")
		useTTL = value.(bool)
		updatedSVR.UseTTL = &useTTL
	}

	if d.HasChange("comment") {
		hasChanges = true
		_, newComment := d.GetChange("comment")
		updatedSVR.Comment = newComment.(string)
	}

	if hasChanges {
		updateAPI := records.NewUpdateRecord(recordReference, updatedSVR)
		changeErr := infobloxClient.Do(updateAPI)
		if changeErr != nil {
			log.Printf(fmt.Sprintf("[DEBUG] Error updating  SRV record: %s", changeErr))
		}
		d.SetId(updateAPI.GetResponse())
		return resourceSRVRecordRead(d, m)

	}

	return nil
}

func resourceSRVRecordDelete(d *schema.ResourceData, m interface{}) error {
	returnFields := []string{}
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	resourceReference := d.Id()
	getSingleSRVAPI := records.NewGetSRVRecord(resourceReference, returnFields)

	err := infobloxClient.Do(getSingleSRVAPI)
	if err != nil {
		return fmt.Errorf("Infoblox Delete Error when fetching resource: %+v", err)
	}
	if getSingleSRVAPI.StatusCode() != 200 {
		d.SetId("")
		return fmt.Errorf("Infoblox Read Error: Invalid HTTP response code %+v returned. Response object was %+v", getSingleSRVAPI.StatusCode(), getSingleSRVAPI.GetResponse())
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
