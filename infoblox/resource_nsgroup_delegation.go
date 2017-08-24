package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupdelegation"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceNSGroupDelegation() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGroupDelegationCreate,
		Read:   resourceNSGroupDelegationRead,
		Update: resourceNSGroupDelegationUpdate,
		Delete: resourceNSGroupDelegationDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the name server group",
				Required:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment field",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			// Note: Only name and address in the external server struct are used by this resource.
			"delegate_to": util.ExternalServerSetSchema(false, true),
		},
	}
}

func resourceNSGroupDelegationCreate(d *schema.ResourceData, m interface{}) error {

	var nsGroupDelegationObject nsgroupdelegation.NSGroupDelegation
	client := m.(*skyinfoblox.InfobloxClient)

	if v, ok := d.GetOk("name"); ok && v != "" {
		nsGroupDelegationObject.Name = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		nsGroupDelegationObject.Comment = v.(string)
	}
	if v, ok := d.GetOk("delegate_to"); ok && v != nil {
		nsGroupDelegationObject.DelegateTo = util.BuildExternalServerSetFromT(v.(*schema.Set))
	}

	createNSGroupDelegationAPI := nsgroupdelegation.NewCreate(nsGroupDelegationObject)
	err := client.Do(createNSGroupDelegationAPI)
	httpStatus := createNSGroupDelegationAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Delegation create for %s failed with status code %d and error: %+v", nsGroupDelegationObject.Name, httpStatus, string(createNSGroupDelegationAPI.RawResponse()))
	}

	nsGroupDelegationObject.Reference = *createNSGroupDelegationAPI.ResponseObject().(*string)
	d.SetId(nsGroupDelegationObject.Reference)
	return resourceNSGroupDelegationRead(d, m)
}

func resourceNSGroupDelegationRead(d *schema.ResourceData, m interface{}) error {

	reference := d.Id()
	client := m.(*skyinfoblox.InfobloxClient)

	getNSGroupDelegationAPI := nsgroupdelegation.NewGet(reference, nsgroupdelegation.RequestReturnFields)
	err := client.Do(getNSGroupDelegationAPI)
	httpStatus := getNSGroupDelegationAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Delegation read for %s failed with status code %d and error: %+v", reference, httpStatus, string(getNSGroupDelegationAPI.RawResponse()))
	}
	response := *getNSGroupDelegationAPI.ResponseObject().(*nsgroupdelegation.NSGroupDelegation)
	d.SetId(response.Reference)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("delegate_to", util.BuildExternalServersListFromIBX(response.DelegateTo))

	return nil
}

func resourceNSGroupDelegationUpdate(d *schema.ResourceData, m interface{}) error {

	var nsGroupDelegationObject nsgroupdelegation.NSGroupDelegation
	hasChanges := false

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok && v != "" {
			nsGroupDelegationObject.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok && v != "" {
			nsGroupDelegationObject.Comment = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("delegate_to") {
		if v, ok := d.GetOk("delegate_to"); ok && v != nil {
			nsGroupDelegationObject.DelegateTo = util.BuildExternalServerSetFromT(v.(*schema.Set))
		}
		hasChanges = true
	}

	if hasChanges {
		nsGroupDelegationObject.Reference = d.Id()
		client := m.(*skyinfoblox.InfobloxClient)

		nsGroupDelegationUpdateAPI := nsgroupdelegation.NewUpdate(nsGroupDelegationObject, nsgroupdelegation.RequestReturnFields)
		err := client.Do(nsGroupDelegationUpdateAPI)
		httpStatus := nsGroupDelegationUpdateAPI.StatusCode()

		if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox NS Group Delegation update for %s failed with status code %d and error: %+v", nsGroupDelegationObject.Name, httpStatus, string(nsGroupDelegationUpdateAPI.RawResponse()))
		}
		response := *nsGroupDelegationUpdateAPI.ResponseObject().(*nsgroupdelegation.NSGroupDelegation)

		d.SetId(response.Reference)
		d.Set("name", response.Name)
		d.Set("comment", response.Comment)
		d.Set("delegate_to", util.BuildExternalServersListFromIBX(response.DelegateTo))
	}
	return resourceNSGroupDelegationRead(d, m)
}

func resourceNSGroupDelegationDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	nsGroupDelegationDeleteAPI := nsgroupdelegation.NewDelete(reference)
	err := client.Do(nsGroupDelegationDeleteAPI)
	httpStatus := nsGroupDelegationDeleteAPI.StatusCode()

	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Delegation delete for %s failed with status code %d and error: %+v", reference, httpStatus, nsGroupDelegationDeleteAPI.RawResponse())
	}
	d.SetId("")
	return nil
}
