package infoblox

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"log"
	"net/http"
)

// RespError : what POST/PUT/DELETE requests returns in case of error.
type RespError struct {
	Error string `json:"Error"`
	Code  string `json:"code"`
	Text  string `json:"text"`
}

func resourceTXTRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceTXTRecordCreate,
		Read:   resourceTXTRecordRead,
		Update: resourceTXTRecordUpdate,
		Delete: resourceTXTRecordDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"text": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"view": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
				Optional: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				Computed: true,
			},
			"ttl": &schema.Schema{
				Type:         schema.TypeInt,
				Required:     false,
				Optional:     true,
				ValidateFunc: validateUINT,
			},
			"use_ttl": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceTXTRecordUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	hasChanges := false
	var record records.GenericRecord
	fields := []string{"name", "text", "view", "ttl", "use_ttl", "comment"}

	getTXTRecordAPI := records.NewGetTXTRecord(d.Id(), fields)
	getTXTErr := infobloxClient.Do(getTXTRecordAPI)
	if getTXTErr != nil {
		return fmt.Errorf("Unable to read the specefied TXT record")
	}

	if d.HasChange("name") {
		hasChanges = true
		value := d.Get("name")
		record.Name = value.(string)
	}

	if d.HasChange("text") {
		hasChanges = true
		value := d.Get("text")
		record.Text = value.(string)
	}

	if d.HasChange("view") {
		hasChanges = true
		value := d.Get("view")
		record.View = value.(string)
	}

	if d.HasChange("ttl") {
		hasChanges = true
		value := d.Get("ttl")
		record.TTL = uint(value.(int))
	}

	useTTL := false
	if d.HasChange("use_ttl") {
		hasChanges = true
		value := d.Get("use_ttl")
		useTTL = value.(bool)
		record.UseTTL = &useTTL
	}

	if d.HasChange("comment") {
		hasChanges = true
		value := d.Get("comment")
		record.Comment = value.(string)
	}

	if hasChanges {
		updateAPI := records.NewUpdateRecord(d.Id(), record)
		updateErr := infobloxClient.Do(updateAPI)
		if updateErr != nil {
			return updateErr
		}

		log.Println("Status Code: ", updateAPI.StatusCode())
		response := updateAPI.GetResponse()
		log.Printf("Response:\n<%s>\n", response)

		if updateAPI.StatusCode() == http.StatusOK {
			log.Println("RECORD Updated!")
			// the object ref might have changed
			d.SetId(response)
			return nil
		}

		// if update fails first read back object from infoblox...
		resourceTXTRecordRead(d, m)
		return handleError(response)
	}
	return nil
}

func resourceTXTRecordDelete(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	deleteAPI := records.NewDelete(d.Id())

	DeleteRecordErr := infobloxClient.Do(deleteAPI)
	if DeleteRecordErr != nil {
		return DeleteRecordErr
	}

	responseObject := deleteAPI.GetResponse()
	if deleteAPI.StatusCode() == http.StatusOK {
		log.Println("Record DELETED")
		log.Println("Status Code: ", deleteAPI.StatusCode())
		log.Printf("Response:\n%s\n", responseObject)
		d.SetId("")
		return nil
	}

	return handleError(responseObject)
}

func handleError(respObj string) error {
	var respObjType RespError
	err := json.Unmarshal([]byte(respObj), &respObjType)
	if err != nil {
		return err
	}
	errorStr := fmt.Sprintf("Infoblox REST API Error: %s", respObjType.Text)
	return errors.New(errorStr)
}

func resourceTXTRecordCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var record records.TXTRecord

	if v, ok := d.GetOk("name"); ok {
		record.Name = v.(string)
	} else {
		return fmt.Errorf("Name argument is required")
	}

	if v, ok := d.GetOk("text"); ok {
		record.Text = v.(string)
	} else {
		return fmt.Errorf("Text argument is required")
	}

	if v, ok := d.GetOk("view"); ok {
		record.View = v.(string)
	}

	if v, ok := d.GetOk("ttl"); ok {
		record.TTL = uint(v.(int))
	}

	useTTL := false
	if v, ok := d.GetOk("use_ttl"); ok {
		useTTL = v.(bool)
	}
	record.UseTTL = &useTTL

	if v, ok := d.GetOk("comment"); ok {
		record.Comment = v.(string)
	}

	createAPI := records.NewCreateTXTRecord(record)
	createRecordErr := infobloxClient.Do(createAPI)
	if createRecordErr != nil {
		return createRecordErr
	}

	response := createAPI.GetResponse()
	if createAPI.StatusCode() == http.StatusCreated {
		d.SetId(response)
		resourceTXTRecordRead(d, m)
		return nil
	}

	return handleError(response)
}

func resourceTXTRecordRead(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	ref := d.Id()
	fields := []string{"name", "view", "zone", "ttl", "use_ttl", "text", "comment"}
	recordAPI := records.NewGetTXTRecord(ref, fields)
	readErr := infobloxClient.Do(recordAPI)
	if readErr != nil {
		log.Println("[resourceTXTRecordRead]:Error reading the object...")
		d.SetId("")
		return errors.New("Infoblox binding error")
	}

	if recordAPI.StatusCode() == http.StatusOK {
		record := recordAPI.GetResponse()
		log.Println("READ OK, going to set name: ", record.Name)
		d.Set("name", record.Name)
		d.Set("text", record.Text)
		d.Set("view", record.View)
		d.Set("zone", record.Zone)
		d.Set("ttl", record.TTL)
		d.Set("comment", record.Comment)
		return nil
	}

	d.SetId("")
	errStr := fmt.Sprintf("[resourceTXTRecordRead]:Return code: ", recordAPI.StatusCode())
	return errors.New(errStr)
}

func validateUINT(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value < 0 {
		errors = append(errors, fmt.Errorf("Integer field  %q must be unsigned integer", k))
	}
	return
}
