package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminrole"
	"net/http"
)

func resourceAdminRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdminRoleCreate,
		Read:   resourceAdminRoleRead,
		Update: resourceAdminRoleUpdate,
		Delete: resourceAdminRoleDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of an admin role.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The descriptive comment of the Admin Role object.",
			},
			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceAdminRoleCreate(d *schema.ResourceData, m interface{}) error {
	var adminRoleObject adminrole.AdminRole

	client := m.(*skyinfoblox.InfobloxClient)

	if v, ok := d.GetOk("name"); ok && v != "" {
		adminRoleObject.Name = v.(string)
	}
	if v, ok := d.GetOk("comment"); ok && v != "" {
		adminRoleObject.Comment = v.(string)
	}
	if v, ok := d.GetOk("disable"); ok && v != nil {
		disable := v.(bool)
		adminRoleObject.Disable = &disable
	}

	createAPI := adminrole.NewCreate(adminRoleObject)
	err := client.Do(createAPI)
	httpStatus := createAPI.StatusCode()
	if err != nil {
		return fmt.Errorf("Infoblox Admin Role Create for %s failed with status code %d and error: %+v", adminRoleObject.Name, httpStatus, err)
	}

	if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Admin Role Create for %s failed with status code %d - %s", adminRoleObject.Name, httpStatus, *createAPI.ResponseObject().(*string))
	}

	adminRoleObject.Reference = *createAPI.ResponseObject().(*string)

	d.SetId(adminRoleObject.Reference)
	return resourceAdminRoleRead(d, m)
}

func resourceAdminRoleRead(d *schema.ResourceData, m interface{}) error {
	reference := d.Id()
	client := m.(*skyinfoblox.InfobloxClient)

	getAdminRoleAPI := adminrole.NewGet(reference)
	err := client.Do(getAdminRoleAPI)
	httpStatus := getAdminRoleAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf("Infoblox Admin Role Read for %s failed with status code %d and error: %+v", reference, httpStatus, err)
	}

	if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Admin Role Create for %s failed with status code %d - %s", d.Id(), httpStatus, *getAdminRoleAPI.ResponseObject().(*string))
	}

	response := *getAdminRoleAPI.ResponseObject().(*adminrole.AdminRole)

	d.SetId(response.Reference)
	d.Set("name", response.Name)
	d.Set("comment", response.Comment)
	d.Set("disable", *response.Disable)

	return nil
}

func resourceAdminRoleUpdate(d *schema.ResourceData, m interface{}) error {
	var adminRoleObject adminrole.AdminRole
	hasChanges := false

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok && v != "" {
			adminRoleObject.Name = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok && v != "" {
			adminRoleObject.Comment = v.(string)
		}
		hasChanges = true
	}
	if d.HasChange("disable") {
		disable := d.Get("disable").(bool)
		adminRoleObject.Disable = &disable
		hasChanges = true
	}

	if hasChanges {
		client := m.(*skyinfoblox.InfobloxClient)
		adminRoleObject.Reference = d.Id()

		updateAdminRoleAPI := adminrole.NewUpdate(adminRoleObject.Reference, adminRoleObject)
		err := client.Do(updateAdminRoleAPI)
		httpStatus := updateAdminRoleAPI.StatusCode()

		if err != nil {
			return fmt.Errorf("Infoblox Admin Role Update for %s failed with status code %d and error: %+v", adminRoleObject.Name, httpStatus, err)
		}

		if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox Admin Role Create for %s failed with status code %d - %s", d.Id(), httpStatus, *updateAdminRoleAPI.ResponseObject().(*string))
		}

		response := *updateAdminRoleAPI.ResponseObject().(*adminrole.AdminRole)

		d.SetId(response.Reference)
		d.Set("name", response.Name)
		d.Set("comment", response.Comment)
		d.Set("disable", *response.Disable)
	}
	return resourceAdminRoleRead(d, m)
}

func resourceAdminRoleDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	deleteAdminRoleAPI := adminrole.NewDelete(reference)
	err := client.Do(deleteAdminRoleAPI)
	httpStatus := deleteAdminRoleAPI.StatusCode()

	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("Infoblox Admin Role Delete for %s failed with status code %d and error: %+v", reference, httpStatus, err)
	}

	if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Admin Role Create for %s failed with status code %d - %s", d.Id(), httpStatus, *deleteAdminRoleAPI.ResponseObject().(*string))
	}

	d.SetId("")
	return nil
}
