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

func TestAccInfobloxAdminGroupBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	adminGroupName := fmt.Sprintf("acctest-infoblox-admin-group-%d", randomInt)
	updateAdminGroupName := fmt.Sprintf("%s-updated", adminGroupName)
	adminGroupResource := "infoblox_admin_group.acctest"
	emailAddressKeyPattern := regexp.MustCompile(`email_addresses\.[0-9]+`)

	fmt.Printf("\n\nAcceptance Test Admin Group is %s\n\n", adminGroupName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.AdmingroupObj, "name", adminGroupName)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxAdminGroupNoNameTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config: testAccInfobloxAdminGroupCreateTemplate(adminGroupName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxAdminGroupCheckExists("name", adminGroupName),
					resource.TestCheckResourceAttr(adminGroupResource, "name", adminGroupName),
					resource.TestCheckResourceAttr(adminGroupResource, "comment", "Infoblox Terraform Acceptance test"),
					resource.TestCheckResourceAttr(adminGroupResource, "superuser", "true"),
					resource.TestCheckResourceAttr(adminGroupResource, "disable", "true"),
					resource.TestCheckResourceAttr(adminGroupResource, "access_method.#", "3"),
					resource.TestCheckResourceAttr(adminGroupResource, "access_method.0", "GUI"),
					resource.TestCheckResourceAttr(adminGroupResource, "access_method.1", "API"),
					resource.TestCheckResourceAttr(adminGroupResource, "access_method.2", "TAXII"),
					resource.TestCheckResourceAttr(adminGroupResource, "email_addresses.#", "4"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.one@example.com"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.two@example.com"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.three@example.com"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.four@example.com"),
					resource.TestCheckResourceAttr(adminGroupResource, "roles.#", "2"),
					resource.TestCheckResourceAttr(adminGroupResource, "roles.0", "DNS Admin"),
					resource.TestCheckResourceAttr(adminGroupResource, "roles.1", "DHCP Admin"),
				),
			},
			{
				Config: testAccInfobloxAdminGroupUpdateTemplate(updateAdminGroupName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxAdminGroupCheckExists("name", updateAdminGroupName),
					resource.TestCheckResourceAttr(adminGroupResource, "name", updateAdminGroupName),
					resource.TestCheckResourceAttr(adminGroupResource, "comment", "Infoblox Terraform Acceptance test - updated"),
					resource.TestCheckResourceAttr(adminGroupResource, "superuser", "false"),
					resource.TestCheckResourceAttr(adminGroupResource, "disable", "false"),
					resource.TestCheckResourceAttr(adminGroupResource, "access_method.#", "3"),
					resource.TestCheckResourceAttr(adminGroupResource, "access_method.0", "GUI"),
					resource.TestCheckResourceAttr(adminGroupResource, "access_method.1", "API"),
					resource.TestCheckResourceAttr(adminGroupResource, "access_method.2", "TAXII"),
					resource.TestCheckResourceAttr(adminGroupResource, "email_addresses.#", "5"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.one@example.com"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.two@example.com"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.three@example.com"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.four@example.com"),
					testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource, emailAddressKeyPattern, "user.five@example.com"),
					resource.TestCheckResourceAttr(adminGroupResource, "roles.#", "1"),
					resource.TestCheckResourceAttr(adminGroupResource, "roles.0", "DHCP Admin"),
				),
			},
		},
	})
}

func testAccInfobloxAdminGroupCheckValueInKeyPattern(adminGroupResource string, keyPattern *regexp.Regexp, checkValue string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[adminGroupResource]
		if ok {
			for attributeKey, attributeValue := range rs.Primary.Attributes {
				if keyPattern.MatchString(attributeKey) {
					if attributeValue == checkValue {
						return nil
					}
				}
			}
		}
		return fmt.Errorf("Infoblox Admin Group attribute %s not found", checkValue)
	}
}

func testAccInfobloxAdminGroupCheckExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.AdmingroupObj, key, value)
	}
}

func testAccInfobloxAdminGroupNoNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_admin_group" "acctest" {
comment = "Infoblox Terraform Acceptance test"
superuser = true
disable = true
access_method = ["GUI", "API", "TAXII"]
email_addresses = ["user.one@example.com", "user.two@example.com", "user.three@example.com", "user.four@example.com"]
roles = ["DNS Admin", "DHCP Admin"]
}
`)
}

func testAccInfobloxAdminGroupCreateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_admin_group" "acctest" {
name = "%s"
comment = "Infoblox Terraform Acceptance test"
superuser = true
disable = true
access_method = ["GUI", "API", "TAXII"]
email_addresses = ["user.one@example.com", "user.two@example.com", "user.three@example.com", "user.four@example.com"]
roles = ["DNS Admin", "DHCP Admin"]
}
`, name)
}

func testAccInfobloxAdminGroupUpdateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_admin_group" "acctest" {
name = "%s"
comment = "Infoblox Terraform Acceptance test - updated"
superuser = false
disable = false
access_method = ["GUI", "API", "TAXII"]
email_addresses = ["user.one@example.com", "user.two@example.com", "user.three@example.com", "user.four@example.com", "user.five@example.com"]
roles = ["DHCP Admin"]
}
`, name)
}
