package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminuser"
	"log"
	"net/http"
)

func resourceAdminUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAdminUserCreate,
		Read:   resourceAdminUserRead,
		Update: resourceAdminUserUpdate,
		Delete: resourceAdminUserDelete,
		Schema: map[string]*schema.Schema{
			"ref": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal Reference for the user",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name for the user",
			},
			"groups": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The groups the user bellongs to , there can be only 1 ",
				Elem:        schema.TypeString,
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email address for the user",
			},
			"disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Should the user be disabled",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "a comment on the user",
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceAdminUserCreate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var userCreate adminuser.AdminUser

	if v, ok := d.GetOk("name"); ok {
		userCreate.Name = v.(string)

	}

	if v, ok := d.GetOk("groups"); ok {
		groups := []string{v.(string)}
		userCreate.Groups = groups
	}

	if v, ok := d.GetOk("email"); ok {
		userCreate.Email = v.(string)
	}

	if v, ok := d.GetOk("disable"); ok {
		disable := v.(bool)
		userCreate.Disable = &disable
	}
	if v, ok := d.GetOk("comment"); ok {
		userCreate.Comment = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		userCreate.Password = v.(string)
	}

	userCreateAPI := adminuser.NewCreateAdminUser(userCreate)
	createErr := infobloxClient.Do(userCreateAPI)
	if createErr != nil {
		return fmt.Errorf("error creating resource %s", createErr.Error())
	}
	if userCreateAPI.StatusCode() != http.StatusCreated {
		return fmt.Errorf("Infoblox Create Error: Invalid HTTP response code %+v returned. Response object was %+v", userCreateAPI.StatusCode(), *userCreateAPI.ResponseObject().(*string))
	}
	response := *userCreateAPI.ResponseObject().(*string)
	d.SetId(response)
	return resourceAdminUserRead(d, m)
}

func resourceAdminUserRead(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var userRead adminuser.AdminUser
	fieldList := []string{"name", "email", "comment", "admin_groups", "disable"}
	readAPI := adminuser.NewGetAdminUser(d.Id(), fieldList)
	readErr := infobloxClient.Do(readAPI)
	if readErr != nil {
		return fmt.Errorf("error reading resource %s", readErr.Error())
	}
	if readAPI.StatusCode() == http.StatusNotFound {
		d.SetId("")
		return nil
	}

	if readAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Create Error: Invalid HTTP response code %+v returned. Response object was %+v", readAPI.StatusCode(), readAPI.ResponseObject().(string))
	}

	userRead = *readAPI.ResponseObject().(*adminuser.AdminUser)

	d.Set("name", userRead.Name)
	d.Set("groups", userRead.Groups)
	d.Set("email", userRead.Email)
	d.Set("disable", userRead.Disable)
	d.Set("comment", userRead.Comment)
	return nil
}

func resourceAdminUserUpdate(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var updateUser adminuser.AdminUser
	var readErr error
	var hasChanges bool
	updateUser, readErr = doReadCall(d.Id(), m)
	if readErr != nil {
		return fmt.Errorf("Error reading %s", readErr.Error())
	}
	if d.HasChange("name") {
		hasChanges = true
		_, newName := d.GetChange("name")
		updateUser.Name = newName.(string)
	}
	if d.HasChange("groups") {
		hasChanges = true
		_, newGroups := d.GetChange("groups")
		updateUser.Groups = []string{newGroups.(string)}
	}
	if d.HasChange("email") {
		hasChanges = true
		_, newEmail := d.GetChange("email")
		updateUser.Email = newEmail.(string)
	}
	if d.HasChange("disable") {
		hasChanges = true
		_, newDisable := d.GetChange("disable")
		updateUser.Disable = newDisable.(*bool)
	}
	if d.HasChange("comment") {
		hasChanges = true
		_, newComment := d.GetChange("comment")
		updateUser.Comment = newComment.(string)
	}

	if d.HasChange("password") {
		hasChanges = true
		_, newPassword := d.GetChange("password")
		updateUser.Password = newPassword.(string)
	} else {
		updateUser.Password = d.Get("password").(string)
	}

	if hasChanges {
		updateUserAPI := adminuser.NewUpdateAdminUser(updateUser)
		updateErr := infobloxClient.Do(updateUserAPI)
		log.Println(updateUserAPI.ResponseObject())
		if updateErr != nil {
			return fmt.Errorf("Error editing user : %s ", updateErr.Error())
		}

		if updateUserAPI.StatusCode() != http.StatusOK {
			return fmt.Errorf("Error updating the user %s ", *updateUserAPI.ResponseObject().(*string))
		}
		return resourceAdminUserRead(d, m)
	}
	return nil
}

func resourceAdminUserDelete(d *schema.ResourceData, m interface{}) error {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	deleteAPI := adminuser.NewDeleteAdminUser(d.Id())
	deleteErr := infobloxClient.Do(deleteAPI)
	if deleteErr != nil {
		return fmt.Errorf("Could not delete %s", deleteErr.Error())
	}
	if deleteAPI.StatusCode() == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if deleteAPI.StatusCode() != http.StatusOK {
		return fmt.Errorf("Infoblox Create Error: Invalid HTTP response code %+v returned. Response object was %+v", deleteAPI.StatusCode(), *deleteAPI.ResponseObject().(*string))
	}
	d.SetId("")
	return nil
}

func doReadCall(userRef string, m interface{}) (adminuser.AdminUser, error) {
	infobloxClient := m.(*skyinfoblox.InfobloxClient)
	var userRead adminuser.AdminUser
	fieldList := []string{"name", "email", "comment", "admin_groups", "disable"}
	readAPI := adminuser.NewGetAdminUser(userRef, fieldList)
	readErr := infobloxClient.Do(readAPI)
	if readErr != nil {
		return userRead, fmt.Errorf("error reading resource %s", readErr.Error())
	}
	if readAPI.StatusCode() != http.StatusOK {
		return userRead, fmt.Errorf("Infoblox Create Error: Invalid HTTP response code %+v returned. Response object was %+v", readAPI.StatusCode(), *readAPI.ResponseObject().(*string))
	}
	userRead = *readAPI.ResponseObject().(*adminuser.AdminUser)
	return userRead, nil

}
