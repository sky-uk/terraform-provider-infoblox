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

func TestAccInfobloxNSGroupForwardStubBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	nsGroupFwdStubName := fmt.Sprintf("acctest-infoblox-ns-group-fwd-stub-%d", randomInt)
	nsGroupNameFwdStubUpdate := fmt.Sprintf("%s-updated", nsGroupFwdStubName)
	nsGroupFwdStubResourceInstance := "infoblox_ns_group_forward_stub.acctest"

	nsGroupFwdStubNamePattern := regexp.MustCompile(`external_servers\.[0-9]+\.name`)
	nsGroupFwdStubAddressPattern := regexp.MustCompile(`external_servers\.[0-9]+\.address`)

	fmt.Printf("\n\nAcceptance Test NS Group Forward/Stub is %s\n\n", nsGroupFwdStubName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.NsgroupForwardstubserverObj, "name", nsGroupFwdStubName)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxNSGroupFwdStubNoNameTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testAccInfobloxNSGroupFwdStubCommentLeadingTrailingSpaces(nsGroupFwdStubName),
				ExpectError: regexp.MustCompile(`must not contain trailing or leading white space`),
			},
			{
				Config: testAccInfobloxNSGroupFwdStubCreateTemplate(nsGroupFwdStubName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSGroupFwdStubCheckExists(nsGroupFwdStubName, nsGroupFwdStubResourceInstance),
					resource.TestCheckResourceAttr(nsGroupFwdStubResourceInstance, "name", nsGroupFwdStubName),
					resource.TestCheckResourceAttr(nsGroupFwdStubResourceInstance, "comment", "Infoblox Terraform Acceptance test"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubNamePattern, "ns1.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubAddressPattern, "192.168.0.3"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubNamePattern, "ns2.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubAddressPattern, "192.168.0.4"),
				),
			},
			{
				Config: testAccInfobloxNSGroupFwdStubUpdateTemplate(nsGroupNameFwdStubUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSGroupFwdStubCheckExists(nsGroupNameFwdStubUpdate, nsGroupFwdStubResourceInstance),
					resource.TestCheckResourceAttr(nsGroupFwdStubResourceInstance, "name", nsGroupNameFwdStubUpdate),
					resource.TestCheckResourceAttr(nsGroupFwdStubResourceInstance, "comment", "Infoblox Terraform Acceptance test - updated"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubNamePattern, "ns3.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubAddressPattern, "192.168.10.3"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubNamePattern, "ns2.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubAddressPattern, "192.168.0.4"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubNamePattern, "ns4.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubAddressPattern, "192.168.10.4"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubNamePattern, "ns5.example.com"),
					util.AccTestCheckValueInKeyPattern(nsGroupFwdStubResourceInstance, nsGroupFwdStubAddressPattern, "192.168.10.5"),
				),
			},
		},
	})
}

func testAccInfobloxNSGroupFwdStubCheckExists(name, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.NsgroupForwardstubserverObj, "name", name)
	}
}

func testAccInfobloxNSGroupFwdStubNoNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_forward_stub" "acctest" {
  comment = "Infoblox Terraform Acceptance test"
  external_servers = [
    {
      name = "ns1.example.com"
      address = "192.168.0.3"
    },
    {
      name = "ns2.example.com"
      address = "192.168.0.4"
    },
  ]
}
`)
}

func testAccInfobloxNSGroupFwdStubCommentLeadingTrailingSpaces(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_forward_stub" "acctest" {
  name = "%s"
  comment = " Infoblox Terraform Acceptance test "
  external_servers = [
    {
      name = "ns1.example.com"
      address = "192.168.0.3"
    },
    {
      name = "ns2.example.com"
      address = "192.168.0.4"
    },
  ]
}
`, name)
}

func testAccInfobloxNSGroupFwdStubCreateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_forward_stub" "acctest" {
  name = "%s"
  comment = "Infoblox Terraform Acceptance test"
  external_servers = [
    {
      name = "ns1.example.com"
      address = "192.168.0.3"
    },
    {
      name = "ns2.example.com"
      address = "192.168.0.4"
    },
  ]
}
`, name)
}

func testAccInfobloxNSGroupFwdStubUpdateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_forward_stub" "acctest" {
  name = "%s"
  comment = "Infoblox Terraform Acceptance test - updated"
  external_servers = [
    {
      name = "ns3.example.com"
      address = "192.168.10.3"
    },
    {
      name = "ns2.example.com"
      address = "192.168.0.4"
    },
    {
      name = "ns4.example.com"
      address = "192.168.10.4"
    },
    {
      name = "ns5.example.com"
      address = "192.168.10.5"
    },
  ]
}
`, name)
}
