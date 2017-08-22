package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminuser"
	"testing"
)

func TestAccAdminUserResource(t *testing.T) {
	randomInt := acctest.RandIntRange(1, 10000)
	recordUserName := fmt.Sprintf("testadminuser%d", randomInt)
	resourceName := "infoblox_admin_user.testadmin"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccResourceAdminUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAdminUserNameCreateTemplate(recordUserName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceAdminUserExists(recordUserName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
					resource.TestCheckResourceAttr(resourceName, "email", "exampleuser@domain.internal.com"),
					resource.TestCheckResourceAttr(resourceName, "admin_groups", "APP-OVP-INFOBLOX-READONLY"),
				),
			}, {
				Config: testAccResourceAdminUserNameUpdateTemplate(recordUserName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceAdminUserExists(recordUserName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment updated"),
					resource.TestCheckResourceAttr(resourceName, "email", "user@domain.internal.com"),
					resource.TestCheckResourceAttr(resourceName, "admin_groups", "APP-OVP-INFOBLOX-READONLY"),
				),
			},
		},
	})

}

func testAccResourceAdminUserDestroy(state *terraform.State) error {
	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_admin_user" {
			continue
		}
		if res, ok := rs.Primary.Attributes["id"]; ok && res != "" {
			return nil
		}
		fields := []string{"name"}
		fmt.Println(rs.Primary.Attributes)
		api := adminuser.NewGetAdminUser(rs.Primary.Attributes["id"], fields)
		err := infobloxClient.Do(api)
		if err != nil {
			return err
		}

		if api.ResponseObject().(*adminuser.AdminUser).Name == "testadminuser" {
			return fmt.Errorf("Record still exists %s", rs.Primary.Attributes["name"])
		}

	}
	return nil
}

func testAccResourceAdminUserExists(recordUserName, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox Admin User resource %s not found in resources: ", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox Admin Userresource %s ID not set", resourceName)
		}
		infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		fields := []string{"name", "comment", "email", "admin_groups"}
		getAdminUser := adminuser.NewGetAdminUser(rs.Primary.Attributes["id"], fields)
		err := infobloxClient.Do(getAdminUser)
		if err != nil {
			return fmt.Errorf("Could not get the resource %s", err.Error())
		}
		returnedUser := getAdminUser.ResponseObject().(*adminuser.AdminUser)
		if returnedUser.Name == recordUserName {
			return nil

		}
		return fmt.Errorf("Could not find %s", recordUserName)
	}

}

func testAccResourceAdminUserNameCreateTemplate(username string) string {
	return fmt.Sprintf(`
	resource "infoblox_admin_user" "testadmin" {
	name = "%s"
	comment = "this is a comment"
	email = "exampleuser@domain.internal.com"
	admin_groups = "APP-OVP-INFOBLOX-READONLY"
	password = "c0a6264f0f128d94cd8ef26652e7d9fd"}`, username)
}

func testAccResourceAdminUserNameUpdateTemplate(username string) string {
	return fmt.Sprintf(`
	resource "infoblox_admin_user" "testadmin" {
  		name = "%s"
		comment = "this is a comment updated"
		email = "user@domain.internal.com"
		admin_groups = "APP-OVP-INFOBLOX-READONLY"
		password = "c0a6264f0f128d94cd8ef26652e7d9fd"
	}
	`, username)
}
