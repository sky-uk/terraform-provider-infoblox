package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/permission"
	"testing"
)

func TestAccInfobloxPermissionBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	adminRoleName := fmt.Sprintf("acctest-infoblox-permission-role-%d", randomInt)
	permissionResource := "infoblox_permission.permission_acctest"

	testPermission := permission.Permission{
		Role:         adminRoleName,
		Permission:   "READ",
		ResourceType: "AAAA",
	}

	updatedTestPermission := permission.Permission{
		Role:         adminRoleName,
		Permission:   "WRITE",
		ResourceType: "AAAA",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxPermissionCheckDestroy(state, testPermission)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccInfobloxPermissionCreateTemplate(adminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxPermissionCheckExists(testPermission),
					resource.TestCheckResourceAttr(permissionResource, "role", adminRoleName),
					resource.TestCheckResourceAttr(permissionResource, "permission", "READ"),
					resource.TestCheckResourceAttr(permissionResource, "resource_type", "AAAA"),
				),
			},
			{
				Config: testAccInfobloxPermissionUpdateTemplate(adminRoleName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxPermissionCheckExists(updatedTestPermission),
					resource.TestCheckResourceAttr(permissionResource, "role", adminRoleName),
					resource.TestCheckResourceAttr(permissionResource, "permission", "WRITE"),
					resource.TestCheckResourceAttr(permissionResource, "resource_type", "AAAA"),
				),
			},
		},
	})
}

func testAccInfobloxPermissionCheckExists(testPermision permission.Permission) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		api := permission.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox Permission - error whilst retrieving a list of Permissions: %+v", err)
		}
		for _, permission := range *api.ResponseObject().(*[]permission.Permission) {
			if permission.Role == testPermision.Role {
				if permission.Permission == testPermision.Permission {
					if permission.ResourceType == testPermision.ResourceType {
						return nil
					}
				}
			}
		}
		return fmt.Errorf("Permission wasn't found on remote Infoblox server")
	}
}

func testAccInfobloxPermissionCheckDestroy(state *terraform.State, testPermision permission.Permission) error {
	client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_permission" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		api := permission.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox - error occurred whilst retrieving a list of Admin Roles")
		}
		for _, permission := range *api.ResponseObject().(*[]permission.Permission) {
			if permission.Role == testPermision.Role {
				if permission.Permission == testPermision.Permission {
					if permission.ResourceType == testPermision.ResourceType {
						return fmt.Errorf("Infoblox Permission still exists")
					}
				}
			}
		}
	}
	return nil
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
