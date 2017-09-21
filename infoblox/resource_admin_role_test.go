package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
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
			return TestAccCheckDestroy(model.AdminroleObj, "name", adminRoleName)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxAdminRoleNoNameTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config: testAccInfobloxAdminRoleCreateTemplate(adminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxAdminRoleCheckExists("name", adminRoleName),
					resource.TestCheckResourceAttr(adminRoleResource, "name", adminRoleName),
					resource.TestCheckResourceAttr(adminRoleResource, "comment", "Infoblox Terraform Acceptance test"),
					resource.TestCheckResourceAttr(adminRoleResource, "disable", "true"),
				),
			},
			{
				Config: testAccInfobloxAdminRoleUpdateTemplate(updateAdminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxAdminRoleCheckExists("name", updateAdminRoleName),
					resource.TestCheckResourceAttr(adminRoleResource, "name", updateAdminRoleName),
					resource.TestCheckResourceAttr(adminRoleResource, "comment", "Infoblox Terraform Acceptance test - updated"),
					resource.TestCheckResourceAttr(adminRoleResource, "disable", "false"),
				),
			},
		},
	})
}

func testAccInfobloxAdminRoleCheckExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.AdminroleObj, key, value)
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
