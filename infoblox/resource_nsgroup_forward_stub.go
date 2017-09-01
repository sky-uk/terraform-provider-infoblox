package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwdstub"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceNSGroupForwardStub() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGroupForwardStubCreate,
		Read:   resourceNSGroupForwardStubRead,
		Update: resourceNSGroupForwardStubUpdate,
		Delete: resourceNSGroupForwardStubDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Description:  "The name of the name server group",
				Required:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			"comment": {
				Type:         schema.TypeString,
				Description:  "Comment field",
				Optional:     true,
				ValidateFunc: util.CheckLeadingTrailingSpaces,
			},
			// Note: this resource only makes use of name and address in the external server struct. Other attributes are ignored.
			"external_dns_servers": util.ExternalServerSetSchema(false, true),
		},
	}
}

func resourceNSGroupForwardStubCreate(d *schema.ResourceData, m interface{}) error {

	var nsGroupForwardStubObject nsgroupfwdstub.NSGroupFwdStub
	client := m.(*skyinfoblox.InfobloxClient)

	if v, ok := d.GetOk("name"); ok && v != "" {
		nsGroupForwardStubObject.Name = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		nsGroupForwardStubObject.Comment = v.(string)
	}
	if v, ok := d.GetOk("external_dns_servers"); ok && v != nil {
		nsGroupForwardStubObject.ExternalServers = util.BuildExternalServerSetFromT(v.(*schema.Set))
	}

	createNSGroupForwardStubAPI := nsgroupfwdstub.NewCreate(nsGroupForwardStubObject)
	err := client.Do(createNSGroupForwardStubAPI)
	httpStatus := createNSGroupForwardStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Forward/Stub create for %s failed with status code %d and error: %+v", nsGroupForwardStubObject.Name, httpStatus, string(createNSGroupForwardStubAPI.RawResponse()))
	}

	nsGroupForwardStubObject.Reference = *createNSGroupForwardStubAPI.ResponseObject().(*string)
	d.SetId(nsGroupForwardStubObject.Reference)

	return resourceNSGroupForwardStubRead(d, m)
}

func resourceNSGroupForwardStubRead(d *schema.ResourceData, m interface{}) error {

	reference := d.Id()
	client := m.(*skyinfoblox.InfobloxClient)

	getNSGroupFwdStubAPI := nsgroupfwdstub.NewGet(reference, nsgroupfwdstub.RequestReturnFields)
	err := client.Do(getNSGroupFwdStubAPI)
	httpStatus := getNSGroupFwdStubAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Forward/Stub read for %s failed with status code %d and error: %+v", reference, httpStatus, string(getNSGroupFwdStubAPI.RawResponse()))
	}
	response := *getNSGroupFwdStubAPI.ResponseObject().(*nsgroupfwdstub.NSGroupFwdStub)
	d.SetId(response.Reference)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("external_dns_servers", util.BuildExternalServersListFromIBX(response.ExternalServers))

	return nil
}

func resourceNSGroupForwardStubUpdate(d *schema.ResourceData, m interface{}) error {

	var nsGroupFwdStubObject nsgroupfwdstub.NSGroupFwdStub
	hasChanges := false

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok && v != "" {
			nsGroupFwdStubObject.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok && v != "" {
			nsGroupFwdStubObject.Comment = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("external_dns_servers") {
		if v, ok := d.GetOk("external_dns_servers"); ok && v != nil {
			nsGroupFwdStubObject.ExternalServers = util.BuildExternalServerSetFromT(v.(*schema.Set))
		}
		hasChanges = true
	}

	if hasChanges {
		nsGroupFwdStubObject.Reference = d.Id()
		client := m.(*skyinfoblox.InfobloxClient)

		nsGroupFwdStubUpdateAPI := nsgroupfwdstub.NewUpdate(nsGroupFwdStubObject, nsgroupfwdstub.RequestReturnFields)
		err := client.Do(nsGroupFwdStubUpdateAPI)
		httpStatus := nsGroupFwdStubUpdateAPI.StatusCode()

		if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox NS Group Forward/Stub update for %s failed with status code %d and error: %+v", nsGroupFwdStubObject.Name, httpStatus, string(nsGroupFwdStubUpdateAPI.RawResponse()))
		}
		response := *nsGroupFwdStubUpdateAPI.ResponseObject().(*nsgroupfwdstub.NSGroupFwdStub)

		d.SetId(response.Reference)
		d.Set("name", response.Name)
		d.Set("comment", response.Comment)
		d.Set("external_dns_servers", util.BuildExternalServersListFromIBX(response.ExternalServers))
	}

	return resourceNSGroupForwardStubRead(d, m)
}

func resourceNSGroupForwardStubDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	nsGroupFwdStubDeleteAPI := nsgroupfwdstub.NewDelete(reference)
	err := client.Do(nsGroupFwdStubDeleteAPI)
	httpStatus := nsGroupFwdStubDeleteAPI.StatusCode()

	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Forward/Stub delete for %s failed with status code %d and error: %+v", reference, httpStatus, nsGroupFwdStubDeleteAPI.RawResponse())
	}
	d.SetId("")

	return nil
}
