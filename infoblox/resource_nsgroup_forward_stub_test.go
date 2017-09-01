package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupfwdstub"
	"github.com/sky-uk/terraform-provider-infoblox/infoblox/util"
	"regexp"
	"testing"
)

func TestAccInfobloxNSGroupForwardStubBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	nsGroupFwdStubName := fmt.Sprintf("acctest-infoblox-ns-group-fwd-stub-%d", randomInt)
	nsGroupNameFwdStubUpdate := fmt.Sprintf("%s-updated", nsGroupFwdStubName)
	nsGroupFwdStubResourceInstance := "infoblox_ns_group_forward_stub.acctest"

	nsGroupFwdStubNamePattern := regexp.MustCompile(`external_dns_servers\.[0-9]+\.name`)
	nsGroupFwdStubAddressPattern := regexp.MustCompile(`external_dns_servers\.[0-9]+\.address`)

	fmt.Printf("\n\nAcceptance Test NS Group Forward/Stub is %s\n\n", nsGroupFwdStubName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxNSGroupFwdStubCheckDestroy(state, nsGroupFwdStubName)
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

func testAccInfobloxNSGroupFwdStubCheckDestroy(state *terraform.State, name string) error {

	client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_ns_group_forward_stub" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		api := nsgroupfwdstub.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox - error occurred whilst retrieving a list of NS Group Forward/Stub")
		}
		for _, nsGroupFwdStub := range *api.ResponseObject().(*[]nsgroupfwdstub.NSGroupFwdStub) {
			if nsGroupFwdStub.Name == name {
				return fmt.Errorf("Infoblox NS Group Forward/Stub %s still exists", name)
			}
		}
	}
	return nil
}

func testAccInfobloxNSGroupFwdStubCheckExists(name, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox NS Group Forward/Stub %s wasn't found in resources", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox NS Group Forward/Stub ID not set for %s in resources", name)
		}

		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		api := nsgroupfwdstub.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox NS Group Forward/Stub - error whilst retrieving a list of NS Group Forward/Stub: %+v", err)
		}
		for _, nsGroupFwdStub := range *api.ResponseObject().(*[]nsgroupfwdstub.NSGroupFwdStub) {
			if nsGroupFwdStub.Name == name {
				return nil
			}
		}
		return fmt.Errorf("Infoblox NS Group Forward/Stub %s wasn't found on remote Infoblox server", name)
	}
}

func testAccInfobloxNSGroupFwdStubNoNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_forward_stub" "acctest" {
  comment = "Infoblox Terraform Acceptance test"
  external_dns_servers = [
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
  external_dns_servers = [
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
  external_dns_servers = [
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
  external_dns_servers = [
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
