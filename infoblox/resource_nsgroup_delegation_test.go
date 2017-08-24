package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupdelegation"
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
			return testAccInfobloxNSGroupDelegationCheckDestroy(state, nsGroupDelegationName)
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

func testAccInfobloxNSGroupDelegationCheckDestroy(state *terraform.State, name string) error {

	client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_ns_group_delegation" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		api := nsgroupdelegation.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox - error occurred whilst retrieving a list of NS Group Delegation")
		}
		for _, nsGroupDelegation := range *api.ResponseObject().(*[]nsgroupdelegation.NSGroupDelegation) {
			if nsGroupDelegation.Name == name {
				return fmt.Errorf("Infoblox NS Group Delegation %s still exists", name)
			}
		}
	}
	return nil
}

func testAccInfobloxNSGroupDelegationCheckExists(name, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox NS Group Delegation %s wasn't found in resources", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox NS Group Delegation ID not set for %s in resources", name)
		}

		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		api := nsgroupdelegation.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox NS Group Delegation - error whilst retrieving a list of NS Group Delegation: %+v", err)
		}
		for _, nsGroupDelegation := range *api.ResponseObject().(*[]nsgroupdelegation.NSGroupDelegation) {
			if nsGroupDelegation.Name == name {
				return nil
			}
		}
		return fmt.Errorf("Infoblox NS Group Delegation %s wasn't found on remote Infoblox server", name)
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
