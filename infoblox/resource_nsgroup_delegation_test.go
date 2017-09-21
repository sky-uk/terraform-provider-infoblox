package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"regexp"
	"testing"
)

func TestAccInfobloxNSGroupDelegationBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	nsGroupDelegationName := fmt.Sprintf("acctest-infoblox-ns-group-delegation-%d", randomInt)
	nsGroupNameDelegationUpdate := fmt.Sprintf("%s-updated", nsGroupDelegationName)
	nsGroupDelegationResourceInstance := "infoblox_ns_group_delegation.acctest"

	nsGroupDelegationNamePattern := regexp.MustCompile(`delegate_to\.[0-9]+\.name`)
	nsGroupDelegationAddressPattern := regexp.MustCompile(`delegate_to\.[0-9]+\.address`)

	fmt.Printf("\n\nAcceptance Test NS Group Delegation is %s\n\n", nsGroupDelegationName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.NsgroupDelegationObj, "name", nsGroupDelegationName)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxNSGroupDelegationNoNameTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testAccInfobloxNSGroupDelegationCommentLeadingTrailingSpaces(nsGroupDelegationName),
				ExpectError: regexp.MustCompile(`must not contain trailing or leading white space`),
			},
			{
				Config: testAccInfobloxNSGroupDelegationCreateTemplate(nsGroupDelegationName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSGroupDelegationCheckExists(nsGroupDelegationName, nsGroupDelegationResourceInstance),
					resource.TestCheckResourceAttr(nsGroupDelegationResourceInstance, "name", nsGroupDelegationName),
					resource.TestCheckResourceAttr(nsGroupDelegationResourceInstance, "comment", "Infoblox Terraform Acceptance test"),
					util.AccTestCheckValueInKeyPattern(nsGroupDelegationResourceInstance, nsGroupDelegationNamePattern, "ns1.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupDelegationResourceInstance, nsGroupDelegationAddressPattern, "192.168.100.1"),
					util.AccTestCheckValueInKeyPattern(nsGroupDelegationResourceInstance, nsGroupDelegationNamePattern, "ns2.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupDelegationResourceInstance, nsGroupDelegationAddressPattern, "192.168.101.1"),
				),
			},
			{
				Config: testAccInfobloxNSGroupDelegationUpdateTemplate(nsGroupNameDelegationUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSGroupDelegationCheckExists(nsGroupNameDelegationUpdate, nsGroupDelegationResourceInstance),
					resource.TestCheckResourceAttr(nsGroupDelegationResourceInstance, "name", nsGroupNameDelegationUpdate),
					resource.TestCheckResourceAttr(nsGroupDelegationResourceInstance, "comment", "Infoblox Terraform Acceptance test - updated"),
					util.AccTestCheckValueInKeyPattern(nsGroupDelegationResourceInstance, nsGroupDelegationNamePattern, "ns3.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupDelegationResourceInstance, nsGroupDelegationAddressPattern, "192.168.50.1"),
					util.AccTestCheckValueInKeyPattern(nsGroupDelegationResourceInstance, nsGroupDelegationNamePattern, "ns4.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupDelegationResourceInstance, nsGroupDelegationAddressPattern, "192.168.51.1"),
				),
			},
		},
	})
}

func testAccInfobloxNSGroupDelegationCheckExists(name, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.NsgroupDelegationObj, "name", name)
	}
}

func testAccInfobloxNSGroupDelegationNoNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_delegation" "acctest" {
  comment = "Infoblox Terraform Acceptance test"
  delegate_to = [
    {
      name = "ns1.example.com"
      address = "192.168.100.1"
    },
    {
      name = "ns2.example.com"
      address = "192.168.101.1"
    },
  ]
}
`)
}

func testAccInfobloxNSGroupDelegationCommentLeadingTrailingSpaces(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_delegation" "acctest" {
  name = "%s"
  comment = " Infoblox Terraform Acceptance test "
  delegate_to = [
    {
      name = "ns1.example.com"
      address = "192.168.100.1"
    },
    {
      name = "ns2.example.com"
      address = "192.168.101.1"
    },
  ]
}
`, name)
}

func testAccInfobloxNSGroupDelegationCreateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_delegation" "acctest" {
  name = "%s"
  comment = "Infoblox Terraform Acceptance test"
  delegate_to = [
    {
      name = "ns1.example.com"
      address = "192.168.100.1"
    },
    {
      name = "ns2.example.com"
      address = "192.168.101.1"
    },
  ]
}
`, name)
}

func testAccInfobloxNSGroupDelegationUpdateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_delegation" "acctest" {
  name = "%s"
  comment = "Infoblox Terraform Acceptance test - updated"
  delegate_to = [
    {
      name = "ns3.example.com"
      address = "192.168.50.1"
    },
    {
      name = "ns4.example.com"
      address = "192.168.51.1"
    },
    {
      name = "ns5.example.com"
      address = "192.168.52.1"
    },
  ]
}
`, name)
}
