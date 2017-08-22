package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/admingroup"
	"net/http"
)

func resourceAdminGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdminGroupCreate,
		Read:   resourceAdminGroupRead,
		Update: resourceAdminGroupUpdate,
		Delete: resourceAdminGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Admin Group",
				Required:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment field",
				Optional:    true,
			},
			"superuser": {
				Type:        schema.TypeBool,
				Description: "Whether the group is a super user group or not",
				Optional:    true,
				Computed:    true,
			},
			"disable": {
				Type:        schema.TypeBool,
				Description: "Whether the Admin Group is disabled or not",
				Optional:    true,
				Computed:    true,
			},
			"access_method": {
				Type:        schema.TypeList,
				Description: "Methods the group can use to access Infoblox",
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"email_addresses": {
				// Need to use TypeSet as the read order doesn't match the sent order.
				Type:        schema.TypeSet,
				Description: "List of email addresses to associated with the Admin Group",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"roles": {
				Type:        schema.TypeList,
				Description: "List of roles to associated with the Admin Group",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func adminGroupBuildStringArray(stringList interface{}) []string {

	stringArray := make([]string, 0)
	for _, item := range stringList.([]interface{}) {
		stringArray = append(stringArray, item.(string))
	}
	return stringArray
}

func adminGroupBuildEmailAddressArray(emailAddresses *schema.Set) []string {

	emailAddressList := make([]string, 0)
	for _, emailAddress := range emailAddresses.List() {
		emailAddressList = append(emailAddressList, emailAddress.(string))
	}
	return emailAddressList
}

func resourceAdminGroupCreate(d *schema.ResourceData, m interface{}) error {

	var adminGroupObject admingroup.IBXAdminGroup
	client := m.(*skyinfoblox.InfobloxClient)

	if v, ok := d.GetOk("name"); ok && v != "" {
		adminGroupObject.Name = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		adminGroupObject.Comment = v.(string)
	}
	if v, ok := d.GetOk("superuser"); ok && v != nil {
		superUser := v.(bool)
		adminGroupObject.SuperUser = &superUser
	}
	if v, ok := d.GetOk("disable"); ok && v != nil {
		disable := v.(bool)
		adminGroupObject.Disable = &disable
	}
	if v, ok := d.GetOk("access_method"); ok && v != nil {
		adminGroupObject.AccessMethod = adminGroupBuildStringArray(v)
	}
	if v, ok := d.GetOk("email_addresses"); ok && v != nil {
		adminGroupObject.EmailAddresses = adminGroupBuildEmailAddressArray(v.(*schema.Set))
	}
	if v, ok := d.GetOk("roles"); ok && v != nil {
		adminGroupObject.Roles = adminGroupBuildStringArray(v)
	}

	createAPI := admingroup.NewCreate(adminGroupObject)
	err := client.Do(createAPI)
	httpStatus := createAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Admin Group Create for %s failed with status code %d and error: %+v", adminGroupObject.Name, httpStatus, err)
	}
	adminGroupObject.Reference = *createAPI.ResponseObject().(*string)

	d.SetId(adminGroupObject.Reference)
	return resourceAdminGroupRead(d, m)
}

func resourceAdminGroupRead(d *schema.ResourceData, m interface{}) error {

	returnFields := []string{"name", "comment", "disable", "roles", "email_addresses", "superuser", "access_method"}
	reference := d.Id()
	client := m.(*skyinfoblox.InfobloxClient)

	getAdminGroupAPI := admingroup.NewGet(reference, returnFields)
	err := client.Do(getAdminGroupAPI)
	httpStatus := getAdminGroupAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Admin Group Read for %s failed with status code %d and error: %+v", reference, httpStatus, err)
	}

	response := *getAdminGroupAPI.ResponseObject().(*admingroup.IBXAdminGroup)

	d.SetId(response.Reference)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("superuser", *response.SuperUser)
	d.Set("disable", *response.Disable)
	d.Set("access_method", response.AccessMethod)
	d.Set("email_addresses", response.EmailAddresses)
	d.Set("roles", response.Roles)

	return nil
}

func resourceAdminGroupUpdate(d *schema.ResourceData, m interface{}) error {

	var adminGroupObject admingroup.IBXAdminGroup
	hasChanges := false

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok && v != "" {
			adminGroupObject.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok && v != "" {
			adminGroupObject.Comment = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("superuser") {
		superUser := d.Get("superuser").(bool)
		adminGroupObject.SuperUser = &superUser
		hasChanges = true
	}
	if d.HasChange("disable") {
		disable := d.Get("disable").(bool)
		adminGroupObject.Disable = &disable
		hasChanges = true
	}
	if d.HasChange("access_method") {
		if v, ok := d.GetOk("access_method"); ok && v != nil {
			adminGroupObject.AccessMethod = adminGroupBuildStringArray(v)
		}
		hasChanges = true
	}
	if d.HasChange("email_addresses") {
		if v, ok := d.GetOk("email_addresses"); ok && v != nil {
			adminGroupObject.EmailAddresses = adminGroupBuildEmailAddressArray(v.(*schema.Set))
		}
		hasChanges = true
	}
	if d.HasChange("roles") {
		if v, ok := d.GetOk("roles"); ok && v != nil {
			adminGroupObject.Roles = adminGroupBuildStringArray(v)
		}
		hasChanges = true
	}

	if hasChanges {

		returnFields := []string{"name", "comment", "disable", "roles", "email_addresses", "superuser", "access_method"}
		client := m.(*skyinfoblox.InfobloxClient)
		adminGroupObject.Reference = d.Id()

		updateAdminGroupAPI := admingroup.NewUpdate(adminGroupObject, returnFields)
		err := client.Do(updateAdminGroupAPI)
		httpStatus := updateAdminGroupAPI.StatusCode()

		if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox Admin Group Update for %s failed with status code %d and error: %+v", adminGroupObject.Name, httpStatus, err)
		}
		response := *updateAdminGroupAPI.ResponseObject().(*admingroup.IBXAdminGroup)

		d.SetId(response.Reference)
		d.Set("name", response.Name)
		d.Set("comment", response.Comment)
		d.Set("superuser", *response.SuperUser)
		d.Set("disable", *response.Disable)
		d.Set("access_method", response.AccessMethod)
		d.Set("email_addresses", response.EmailAddresses)
		d.Set("roles", response.Roles)
	}

	return resourceAdminGroupRead(d, m)
}

func resourceAdminGroupDelete(d *schema.ResourceData, m interface{}) error {

	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	deleteAdminGroupAPI := admingroup.NewDelete(reference)
	err := client.Do(deleteAdminGroupAPI)
	httpStatus := deleteAdminGroupAPI.StatusCode()

	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Admin Group Delete for %s failed with status code %d and error: %+v", reference, httpStatus, err)
	}

	d.SetId("")
	return nil
}
