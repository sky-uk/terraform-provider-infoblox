package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwd"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceNSGroupForward() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGroupForwardCreate,
		Read:   resourceNSGroupForwardRead,
		Update: resourceNSGroupForwardUpdate,
		Delete: resourceNSGroupForwardDelete,

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
			"forwarding_servers": util.ForwardingMemberServerListSchema(),
		},
	}
}

func resourceNSGroupForwardCreate(d *schema.ResourceData, m interface{}) error {

	var nsGroupForwardObject nsgroupfwd.NSGroupFwd
	client := m.(*skyinfoblox.InfobloxClient)

	if v, ok := d.GetOk("name"); ok && v != "" {
		nsGroupForwardObject.Name = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		nsGroupForwardObject.Comment = v.(string)
	}
	if v, ok := d.GetOk("forwarding_servers"); ok && v != nil {
		forwardingServerList := util.GetMapList(v.([]interface{}))
		nsGroupForwardObject.ForwardingServers = util.BuildForwardingMemberServerListFromT(forwardingServerList)
	}

	createNSGroupForwardAPI := nsgroupfwd.NewCreate(nsGroupForwardObject)
	err := client.Do(createNSGroupForwardAPI)
	httpStatus := createNSGroupForwardAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Forward create for %s failed with status code %d and error: %+v", nsGroupForwardObject.Name, httpStatus, string(createNSGroupForwardAPI.RawResponse()))
	}

	nsGroupForwardObject.Reference = *createNSGroupForwardAPI.ResponseObject().(*string)
	d.SetId(nsGroupForwardObject.Reference)
	return resourceNSGroupForwardRead(d, m)
}

func resourceNSGroupForwardRead(d *schema.ResourceData, m interface{}) error {

	reference := d.Id()
	client := m.(*skyinfoblox.InfobloxClient)

	getNSGroupForwardAPI := nsgroupfwd.NewGet(reference, nsgroupfwd.RequestReturnFields)
	err := client.Do(getNSGroupForwardAPI)
	httpStatus := getNSGroupForwardAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Forward read for %s failed with status code %d and error: %+v", reference, httpStatus, string(getNSGroupForwardAPI.RawResponse()))
	}
	response := *getNSGroupForwardAPI.ResponseObject().(*nsgroupfwd.NSGroupFwd)
	d.SetId(response.Reference)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("forwarding_servers", util.BuildForwardingMemberServerListFromIBX(response.ForwardingServers))
	return nil
}

func resourceNSGroupForwardUpdate(d *schema.ResourceData, m interface{}) error {

	var nsGroupForwardObject nsgroupfwd.NSGroupFwd
	hasChanges := false

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok && v != "" {
			nsGroupForwardObject.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok && v != "" {
			nsGroupForwardObject.Comment = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("forwarding_servers") {
		if v, ok := d.GetOk("forwarding_servers"); ok && v != nil {
			forwardingServerList := util.GetMapList(v.([]interface{}))
			nsGroupForwardObject.ForwardingServers = util.BuildForwardingMemberServerListFromT(forwardingServerList)
		}
		hasChanges = true
	}

	if hasChanges {
		nsGroupForwardObject.Reference = d.Id()
		client := m.(*skyinfoblox.InfobloxClient)

		nsGroupForwardUpdateAPI := nsgroupfwd.NewUpdate(nsGroupForwardObject, nsgroupfwd.RequestReturnFields)
		err := client.Do(nsGroupForwardUpdateAPI)
		httpStatus := nsGroupForwardUpdateAPI.StatusCode()
		if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox NS Group Forward update for %s failed with status code %d and error: %+v", nsGroupForwardObject.Name, httpStatus, string(nsGroupForwardUpdateAPI.RawResponse()))
		}
		response := *nsGroupForwardUpdateAPI.ResponseObject().(*nsgroupfwd.NSGroupFwd)
		d.SetId(response.Reference)
		d.Set("name", response.Name)
		d.Set("comment", response.Comment)
		d.Set("forwarding_servers", util.BuildForwardingMemberServerListFromIBX(response.ForwardingServers))
	}

	return resourceNSGroupForwardRead(d, m)
}

func resourceNSGroupForwardDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	nsGroupForwardAPI := nsgroupfwd.NewDelete(reference)
	err := client.Do(nsGroupForwardAPI)
	httpStatus := nsGroupForwardAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Forward delete for %s failed with status code %d and error: %+v", reference, httpStatus, nsGroupForwardAPI.RawResponse())
	}
	d.SetId("")
	return nil
}
