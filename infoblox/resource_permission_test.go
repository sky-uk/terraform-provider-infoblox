package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"testing"
)

func TestAccInfobloxPermissionBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	adminRoleName := fmt.Sprintf("acctest-infoblox-permission-role-%d", randomInt)
	permissionResource := "infoblox_permission.permission_acctest"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.PermissionObj, "role", adminRoleName)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccInfobloxPermissionCreateTemplate(adminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxPermissionCheckExists("role", adminRoleName),
					resource.TestCheckResourceAttr(permissionResource, "role", adminRoleName),
					resource.TestCheckResourceAttr(permissionResource, "permission", "READ"),
					resource.TestCheckResourceAttr(permissionResource, "resource_type", "AAAA"),
				),
			},
			{
				Config: testAccInfobloxPermissionUpdateTemplate(adminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxPermissionCheckExists("role", adminRoleName),
					resource.TestCheckResourceAttr(permissionResource, "role", adminRoleName),
					resource.TestCheckResourceAttr(permissionResource, "permission", "WRITE"),
					resource.TestCheckResourceAttr(permissionResource, "resource_type", "AAAA"),
				),
			},
		},
	})
}

func testAccInfobloxPermissionCheckExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.PermissionObj, key, value)
	}
}

func testAccInfobloxPermissionCreateTemplate(roleName string) string {
	return fmt.Sprintf(`
resource "infoblox_admin_role" "role_acctest" {
name = "%s"
comment = "Infoblox Terraform Role for Permission Acceptance test"
disable = true
}

resource "infoblox_permission" "permission_acctest" {
role = "${infoblox_admin_role.role_acctest.name}"
permission = "READ"
resource_type = "AAAA"
}
`, roleName)

}

func testAccInfobloxPermissionUpdateTemplate(roleName string) string {
	return fmt.Sprintf(`
resource "infoblox_admin_role" "role_acctest" {
name = "%s"
comment = "Infoblox Terraform Role for Permission Acceptance test"
disable = true
}

resource "infoblox_permission" "permission_acctest" {
role = "${infoblox_admin_role.role_acctest.name}"
permission = "WRITE"
resource_type = "AAAA"
}
`, roleName)
}
