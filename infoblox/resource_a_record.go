package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"log"
)

func resourceARecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceARecordCreate,
		Read:   resourceARecordRead,
		Update: resourceARecordUpdate,
		Delete: resourceARecordDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address for hostname",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "name for host record",
			},
			"zone": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS Zone for the record",
			},
			"ttl": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "TTL in seconds for host record",
			},
			"ref": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal Reference for the record",
			},
		},
	}
}

func resourceARecordCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var name, address string
	var ttl int
	var createARecord records.ARecord
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		createARecord.Name = name
	} else {
		return fmt.Errorf("name argument is required")
	}

	if v, ok := d.GetOk("address"); ok {
		address = v.(string)
		createARecord.IPv4 = address
	} else {
		return fmt.Errorf("address argument is required")
	}

	if v, ok := d.GetOk("ttl"); ok {
		ttl = v.(int)
		createARecord.TTL = uint(ttl)
	}

	createAPI := records.NewCreateARecord(createARecord)
	createARecordErr := infobloxClient.Do(createAPI)
	if createARecordErr != nil {
		return createARecordErr
	}
	response := createAPI.GetResponse()

	d.SetId(response)
	resourceARecordRead(d, m)
	return nil
}

func resourceARecordRead(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	fields := []string{"name", "ipv4addr", "ttl", "zone"}
	getSingleARecordAPI := records.NewGetARecord(d.Id(), fields)
	readErr := infobloxClient.Do(getSingleARecordAPI)
	if readErr != nil {
		d.SetId("")
		return fmt.Errorf("Record does not exist")
	}
	readData := getSingleARecordAPI.GetResponse()
	d.Set("name", readData.Name)
	d.Set("zone", readData.Zone)
	d.Set("address", readData.IPv4)
	d.Set("ttl", readData.TTL)
	d.Set("ref", readData.Ref)
	return nil
}

func resourceARecordUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var recordReference string
	var hasChanges bool
	if v, ok := d.GetOk("ref"); ok {
		recordReference = v.(string)

	} else {
		return fmt.Errorf("cannot delete without reference")
	}
	fields := []string{"name", "ipv4addr", "ttl"}
	updateARecordAPI := records.NewGetARecord(recordReference, fields)
	updateErr := infobloxClient.Do(updateARecordAPI)
	if updateErr != nil {
		return fmt.Errorf("Unable to read the A record")
	}
	recordToUpdate := updateARecordAPI.GetResponse()
	if d.HasChange("name") {
		hasChanges = true
		_, newName := d.GetChange("name")
		recordToUpdate.Name = newName.(string)
	}

	if d.HasChange("ttl") {
		hasChanges = true
		var TTL int
		_, newTTL := d.GetChange("ttl")
		TTL = newTTL.(int)
		recordToUpdate.TTL = uint(TTL)
	}

	if d.HasChange("address") {
		hasChanges = true
		_, newAddress := d.GetChange("address")
		recordToUpdate.IPv4 = newAddress.(string)
	}

	if hasChanges {
		updateAPI := records.NewUpdateARecord(recordReference, recordToUpdate)
		changeErr := infobloxClient.Do(updateAPI)
		if changeErr != nil {
			log.Printf(fmt.Sprintf("[DEBUG] Error updating  A record: %s", changeErr))
		}
		d.SetId(updateAPI.GetResponse())
		return resourceARecordRead(d, m)

	}

	return nil

}

func resourceARecordDelete(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	deleteAPI := records.NewDelete(d.Id())
	deleteRecordErr := infobloxClient.Do(deleteAPI)
	if deleteRecordErr != nil {
		return deleteRecordErr
	}
	d.SetId("")
	return nil
}
