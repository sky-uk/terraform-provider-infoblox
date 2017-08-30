package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupstub"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"net/http"
)

func resourceNSGroupStub() *schema.Resource {
	return &schema.Resource{
		Create: resourceNSGroupStubCreate,
		Read:   resourceNSGroupStubRead,
		Update: resourceNSGroupStubUpdate,
		Delete: resourceNSGroupStubDelete,

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
			// Note the lead, stealth, grid_replicate, preferred_primaries, override_preferred_primaries attributes are ignored by Infoblox when set for this object type.
			"stub_members": util.MemberServerListSchema(),
		},
	}
}

func resourceNSGroupStubCreate(d *schema.ResourceData, m interface{}) error {

	var nsGroupStubObject nsgroupstub.NSGroupStub
	client := m.(*skyinfoblox.InfobloxClient)

	if v, ok := d.GetOk("name"); ok && v != "" {
		nsGroupStubObject.Name = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		nsGroupStubObject.Comment = v.(string)
	}
	if v, ok := d.GetOk("stub_members"); ok && v != nil {
		stubMemberList := util.GetMapList(v.([]interface{}))
		nsGroupStubObject.StubMembers = util.BuildMemberServerListFromT(stubMemberList)
	}

	createNSGroupStubAPI := nsgroupstub.NewCreate(nsGroupStubObject)
	err := client.Do(createNSGroupStubAPI)
	httpStatus := createNSGroupStubAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Stub create for %s failed with status code %d and error: %+v", nsGroupStubObject.Name, httpStatus, string(createNSGroupStubAPI.RawResponse()))
	}

	nsGroupStubObject.Reference = *createNSGroupStubAPI.ResponseObject().(*string)
	d.SetId(nsGroupStubObject.Reference)
	return resourceNSGroupStubRead(d, m)
}

func resourceNSGroupStubRead(d *schema.ResourceData, m interface{}) error {

	reference := d.Id()
	client := m.(*skyinfoblox.InfobloxClient)

	getNSGroupStubAPI := nsgroupstub.NewGet(reference, nsgroupstub.RequestReturnFields)
	err := client.Do(getNSGroupStubAPI)
	httpStatus := getNSGroupStubAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Stub read for %s failed with status code %d and error: %+v", reference, httpStatus, string(getNSGroupStubAPI.RawResponse()))
	}
	response := *getNSGroupStubAPI.ResponseObject().(*nsgroupstub.NSGroupStub)
	d.SetId(response.Reference)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("stub_members", util.BuildMemberServerListFromIBX(response.StubMembers))

	return nil
}

func resourceNSGroupStubUpdate(d *schema.ResourceData, m interface{}) error {

	var nsGroupStubObject nsgroupstub.NSGroupStub
	hasChanges := false

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok && v != "" {
			nsGroupStubObject.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok && v != "" {
			nsGroupStubObject.Comment = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("stub_members") {
		if v, ok := d.GetOk("stub_members"); ok && v != nil {
			stubMemberList := util.GetMapList(v.([]interface{}))
			nsGroupStubObject.StubMembers = util.BuildMemberServerListFromT(stubMemberList)
		}
		hasChanges = true
	}

	if hasChanges {
		nsGroupStubObject.Reference = d.Id()
		client := m.(*skyinfoblox.InfobloxClient)

		updateNSGroupStubAPI := nsgroupstub.NewUpdate(nsGroupStubObject, nsgroupstub.RequestReturnFields)
		err := client.Do(updateNSGroupStubAPI)
		httpStatus := updateNSGroupStubAPI.StatusCode()
		if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox NS Group Stub update for %s failed with status code %d and error: %+v", nsGroupStubObject.Reference, httpStatus, string(updateNSGroupStubAPI.RawResponse()))
		}
		response := *updateNSGroupStubAPI.ResponseObject().(*nsgroupstub.NSGroupStub)
		d.SetId(response.Reference)
		d.Set("name", response.Name)
		d.Set("comment", response.Comment)
		d.Set("stub_members", util.BuildMemberServerListFromIBX(response.StubMembers))
	}

	return resourceNSGroupStubRead(d, m)
}

func resourceNSGroupStubDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	deleteNSGroupStubAPI := nsgroupstub.NewDelete(reference)
	err := client.Do(deleteNSGroupStubAPI)
	httpStatus := deleteNSGroupStubAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox NS Group Stub delete for %s failed with status code %d and error: %+v", reference, httpStatus, string(deleteNSGroupStubAPI.RawResponse()))
	}
	d.SetId("")
	return nil
}
