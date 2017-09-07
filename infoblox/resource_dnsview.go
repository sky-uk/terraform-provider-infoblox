package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dnsview"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceDNSView() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSViewCreate,
		Read:   resourceDNSViewRead,
		Update: resourceDNSViewUpdate,
		Delete: resourceDNSViewDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of the DNS view.",
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"comment": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Comment for the DNS view; maximum 64 characters.",
				ValidateFunc: util.ValidateMaxLength(64),
			},
			"is_default": {
				Type:        schema.TypeBool,
				Description: "The NIOS appliance provides one default DNS view. You can rename the default view and change its settings, but you cannot delete it. There must always be at least one DNS view in the appliance.",
				Computed:    true,
			},
		},
	}
}

func resourceDNSViewCreate(d *schema.ResourceData, m interface{}) error {
	var newDNSView dnsview.DNSView

	client := m.(*skyinfoblox.InfobloxClient)

	if v, ok := d.GetOk("name"); ok && v != "" {
		newDNSView.Name = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		newDNSView.Comment = v.(string)
	}

	createAPI := dnsview.NewCreate(newDNSView)
	err := client.Do(createAPI)
	httpStatus := createAPI.StatusCode()
	if err != nil {
		return fmt.Errorf("Infoblox DNS View Create for %s failed with status code %d and error: %+v", newDNSView.Name, httpStatus, err)
	}

	if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox DNS View Create for %s failed with status code %d - %s", newDNSView.Name, httpStatus, *createAPI.ResponseObject().(*string))
	}

	newDNSView.Reference = *createAPI.ResponseObject().(*string)

	d.SetId(newDNSView.Reference)
	return resourceDNSViewRead(d, m)
}

func resourceDNSViewRead(d *schema.ResourceData, m interface{}) error {
	reference := d.Id()
	client := m.(*skyinfoblox.InfobloxClient)
	returnFields := []string{"name", "comment", "is_default"}

	getDNSViewAPI := dnsview.NewGet(reference, returnFields)
	err := client.Do(getDNSViewAPI)
	httpStatus := getDNSViewAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf("Infoblox DNS View read for %s failed with status code %d and error: %+v", reference, httpStatus, err)
	}

	if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox DNS View read for %s failed with status code %d - %s", d.Id(), httpStatus, *getDNSViewAPI.ResponseObject().(*string))
	}

	response := *getDNSViewAPI.ResponseObject().(*dnsview.DNSView)

	d.SetId(response.Reference)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("is_default", *response.IsDefault)

	return nil
}

func resourceDNSViewUpdate(d *schema.ResourceData, m interface{}) error {
	var updatedDNSView dnsview.DNSView
	hasChanges := false
	returnFields := []string{"name", "comment", "is_default"}

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok && v != "" {
			updatedDNSView.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok && v != "" {
			updatedDNSView.Comment = v.(string)
		}
		hasChanges = true
	}

	if hasChanges {
		client := m.(*skyinfoblox.InfobloxClient)
		updatedDNSView.Reference = d.Id()

		updateDNSViewAPI := dnsview.NewUpdate(updatedDNSView, returnFields)
		err := client.Do(updateDNSViewAPI)
		httpStatus := updateDNSViewAPI.StatusCode()

		if err != nil {
			return fmt.Errorf("Infoblox DNS View Update for %s failed with status code %d and error: %+v", updatedDNSView.Name, httpStatus, err)
		}

		if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox DNS View Update for %s failed with status code %d - %s", d.Id(), httpStatus, *updateDNSViewAPI.ResponseObject().(*string))
		}

		response := *updateDNSViewAPI.ResponseObject().(*dnsview.DNSView)
		d.SetId(response.Reference)
	}
	return resourceAdminRoleRead(d, m)
}

func resourceDNSViewDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	deleteDNSViewAPI := dnsview.NewDelete(reference)
	err := client.Do(deleteDNSViewAPI)
	httpStatus := deleteDNSViewAPI.StatusCode()

	if deleteDNSViewAPI.StatusCode() == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("InfobloxDNS View Delete for %s failed with status code %d and error: %+v", reference, httpStatus, err)
	}

	if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox DNS View Create for %s failed with status code %d - %s", d.Id(), httpStatus, *deleteDNSViewAPI.ResponseObject().(*string))
	}

	d.SetId("")
	return nil
}
