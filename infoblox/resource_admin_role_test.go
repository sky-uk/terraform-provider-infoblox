package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/adminrole"
	"regexp"
	"testing"
)

func TestAccInfobloxAdminRoleBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	adminRoleName := fmt.Sprintf("acctest-infoblox-admin-role-%d", randomInt)
	updateAdminRoleName := fmt.Sprintf("%s-updated", adminRoleName)
	adminRoleResource := "infoblox_admin_role.acctest"

	fmt.Printf("\n\nAcceptance Test Admin Role is %s\n\n", adminRoleName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxAdminRoleCheckDestroy(state, adminRoleName)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxAdminRoleNoNameTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config: testAccInfobloxAdminRoleCreateTemplate(adminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxAdminRoleCheckExists(adminRoleName, adminRoleResource),
					resource.TestCheckResourceAttr(adminRoleResource, "name", adminRoleName),
					resource.TestCheckResourceAttr(adminRoleResource, "comment", "Infoblox Terraform Acceptance test"),
					resource.TestCheckResourceAttr(adminRoleResource, "disable", "true"),
				),
			},
			{
				Config: testAccInfobloxAdminRoleUpdateTemplate(updateAdminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxAdminRoleCheckExists(updateAdminRoleName, adminRoleResource),
					resource.TestCheckResourceAttr(adminRoleResource, "name", updateAdminRoleName),
					resource.TestCheckResourceAttr(adminRoleResource, "comment", "Infoblox Terraform Acceptance test - updated"),
					resource.TestCheckResourceAttr(adminRoleResource, "disable", "false"),
				),
			},
		},
	})
}

func testAccInfobloxAdminRoleCheckDestroy(state *terraform.State, adminRoleName string) error {

	client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_admin_role" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		api := adminrole.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox - error occurred whilst retrieving a list of Admin Roles")
		}
		for _, adminRole := range *api.ResponseObject().(*[]adminrole.AdminRole) {
			if adminRole.Name == adminRoleName {
				return fmt.Errorf("Infoblox Admin Role %s still exists", adminRoleName)
			}
		}
	}
	return nil
}

func testAccInfobloxAdminRoleCheckExists(adminRoleName, adminRoleResource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[adminRoleResource]
		if !ok {
			return fmt.Errorf("\nInfoblox Admin Role %s wasn't found in resources", adminRoleName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox Admin Role ID not set for %s in resources", adminRoleName)
		}

		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		api := adminrole.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox Admin Role - error whilst retrieving a list of Admin Roles: %+v", err)
		}
		for _, adminRole := range *api.ResponseObject().(*[]adminrole.AdminRole) {
			if adminRole.Name == adminRoleName {
				return nil
			}
		}
		return fmt.Errorf("Infoblox Admin Role %s wasn't found on remote Infoblox server", adminRoleName)
	}
}

func testAccInfobloxAdminRoleNoNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_admin_role" "acctest" {
comment = "Infoblox Terraform Acceptance test"
disable = true
}
`)
}

func testAccInfobloxAdminRoleCreateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_admin_role" "acctest" {
name = "%s"
comment = "Infoblox Terraform Acceptance test"
disable = true
}
`, name)
}

func testAccInfobloxAdminRoleUpdateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_admin_role" "acctest" {
name = "%s"
comment = "Infoblox Terraform Acceptance test - updated"
disable = false
}
`, name)
}
