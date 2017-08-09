package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/permission"
	"net/http"
	"regexp"
)

func resourcePermission() *schema.Resource {
	return &schema.Resource{
		Create: resourcePermissionCreate,
		Read:   resourcePermissionRead,
		Update: resourcePermissionUpdate,
		Delete: resourcePermissionDelete,

		Schema: map[string]*schema.Schema{
			"group": {
				Type:        schema.TypeString,
				Description: "The name of the admin group this permission applies to.",
				Optional:    true,
			},
			"object": {
				Type:        schema.TypeString,
				Description: "A reference to a WAPI object, which will be the object this permission applies to.",
				Optional:    true,
			},
			"permission": {
				Type:         schema.TypeString,
				Description:  "The type of permission.",
				Required:     true,
				ValidateFunc: validatePermissionType,
			},
			"resource_type": {
				Type:        schema.TypeString,
				Description: "The type of resource this permission applies to.",
				Optional:    true,
				ForceNew:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "The name of the role this permission applies to.",
				Optional:    true,
			},
		},
	}
}

func validatePermissionType(v interface{}, k string) (ws []string, errors []error) {
	permission := v.(string)
	permissionOptions := regexp.MustCompile(`^(DENY|READ|WRITE)$`)
	if !permissionOptions.MatchString(permission) {
		errors = append(errors, fmt.Errorf("%q must be one of DENY, READ or WRITE", k))
	}
	return
}

func resourcePermissionCreate(d *schema.ResourceData, m interface{}) error {

	var permissionObject permission.Permission
	client := m.(*skyinfoblox.InfobloxClient)
	groupAndRoleCount, objectAndResourceTypeCount := 0, 0

	if v, ok := d.GetOk("group"); ok && v != "" {
		permissionObject.Group = v.(string)
		groupAndRoleCount++
	}

	if v, ok := d.GetOk("role"); ok && v != "" {
		permissionObject.Role = v.(string)
		groupAndRoleCount++
	}

	if groupAndRoleCount != 1 {
		return fmt.Errorf("One of group or role is required but both cannot be set")
	}

	if v, ok := d.GetOk("object"); ok && v != "" {
		permissionObject.Object = v.(string)
		objectAndResourceTypeCount++
	}

	if v, ok := d.GetOk("resource_type"); ok && v != "" {
		permissionObject.ResourceType = v.(string)
		objectAndResourceTypeCount++
	}

	if groupAndRoleCount < 1 {
		return fmt.Errorf("At least one of object or resource_type is required")
	}

	if v, ok := d.GetOk("permission"); ok && v != "" {
		permissionObject.Permission = v.(string)
	}

	createAPI := permission.NewCreate(permissionObject)

	err := client.Do(createAPI)
	httpStatus := createAPI.StatusCode()
	if err != nil || httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Permission Create failed with status code %d and error: %+v", httpStatus, string(createAPI.RawResponse()))
	}
	permissionReference := *createAPI.ResponseObject().(*string)
	d.SetId(permissionReference)
	return resourcePermissionRead(d, m)
}

func resourcePermissionRead(d *schema.ResourceData, m interface{}) error {
	reference := d.Id()
	client := m.(*skyinfoblox.InfobloxClient)

	getPermissionAPI := permission.NewGet(reference)
	err := client.Do(getPermissionAPI)
	httpStatus := getPermissionAPI.StatusCode()
	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf("Infoblox Permission Read failed with error: %+v", err)
	}

	if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Permission Read failed with status code %d and error: %+v", httpStatus, string(getPermissionAPI.RawResponse()))
	}

	response := *getPermissionAPI.ResponseObject().(*permission.Permission)
	d.SetId(response.Reference)
	d.Set("group", response.Group)
	d.Set("role", response.Role)
	d.Set("object", response.Object)
	d.Set("resource_type", response.ResourceType)
	d.Set("permission", response.Permission)

	return nil
}

func resourcePermissionUpdate(d *schema.ResourceData, m interface{}) error {
	var permissionObject permission.Permission
	hasChanges := false
	reference := d.Id()

	if d.HasChange("group") {
		if v, ok := d.GetOk("group"); ok && v != "" {
			permissionObject.Group = v.(string)
			if v, ok := d.GetOk("role"); ok && v != "" {
				return fmt.Errorf("Group and Role cannot both be set")
			}
		}
		hasChanges = true
	}

	if d.HasChange("role") {
		if v, ok := d.GetOk("role"); ok && v != "" {
			permissionObject.Role = v.(string)
			if v, ok := d.GetOk("group"); ok && v != "" {
				return fmt.Errorf("Group and Role cannot both be set")
			}
		}
		hasChanges = true
	}

	if d.HasChange("object") {
		if v, ok := d.GetOk("object"); ok && v != "" {
			permissionObject.Object = v.(string)
		}
		hasChanges = true
	}

	if d.HasChange("resource_type") {
		if v, ok := d.GetOk("resource_type"); ok && v != "" {
			permissionObject.ResourceType = v.(string)
		}
		hasChanges = true
	}

	if d.HasChange("permission") {
		if v, ok := d.GetOk("permission"); ok && v != "" {
			permissionObject.Permission = v.(string)
		}
		hasChanges = true
	}

	if hasChanges {
		client := m.(*skyinfoblox.InfobloxClient)
		updatePermissionAPI := permission.NewUpdate(reference, permissionObject)
		err := client.Do(updatePermissionAPI)
		httpStatus := updatePermissionAPI.StatusCode()

		if err != nil {
			return fmt.Errorf("Infoblox Permission Update failed with error: %+v", err)
		}

		if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
			return fmt.Errorf("Infoblox Permission Update failed with status code %d and error: %+v", httpStatus, string(updatePermissionAPI.RawResponse()))
		}
		response := *updatePermissionAPI.ResponseObject().(*permission.Permission)

		d.SetId(response.Reference)
		d.Set("group", response.Group)
		d.Set("role", response.Role)
		d.Set("object", response.Object)
		d.Set("resource_type", response.ResourceType)
		d.Set("permission", response.Permission)
	}
	return resourcePermissionRead(d, m)
}

func resourcePermissionDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*skyinfoblox.InfobloxClient)
	reference := d.Id()

	deletePermissionAPI := permission.NewDelete(reference)
	err := client.Do(deletePermissionAPI)
	httpStatus := deletePermissionAPI.StatusCode()

	if httpStatus == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	if err != nil {
		return fmt.Errorf("Infoblox Permission Delete failed with error: %+v", err)
	}

	if httpStatus < http.StatusOK || httpStatus >= http.StatusBadRequest {
		return fmt.Errorf("Infoblox Permission Delete failed with status code %d and error: %+v", httpStatus, string(deletePermissionAPI.RawResponse()))
	}

	d.SetId("")
	return nil
}
