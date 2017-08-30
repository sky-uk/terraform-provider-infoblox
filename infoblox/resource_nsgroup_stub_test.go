package infoblox

/*

Note: test is commented out until we have a proper test environment

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/nsgroupstub"
	"regexp"
	"testing"
)

func TestAccInfobloxNSGroupStubBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	nsGroupStubName := fmt.Sprintf("acctest-infoblox-ns-group-stub-%d", randomInt)
	nsGroupStubNameUpdated := fmt.Sprintf("%s-updated", nsGroupStubName)
	nsGroupStubResourceInstance := "infoblox_ns_group_stub.acctest"

	fmt.Printf("\n\nAcceptance Test NS Group Stub is %s\n\n", nsGroupStubName)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxNSGroupStubCheckDestroy(state, nsGroupStubName)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxNSGroupStubNoNameTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testAccInfobloxNSGroupStubCommentLeadingTrailingSpaces(nsGroupStubName),
				ExpectError: regexp.MustCompile(`must not contain trailing or leading white space`),
			},
			{
				Config: testAccInfobloxNSGroupStubCreateTemplate(nsGroupStubName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSGroupStubCheckExists(nsGroupStubName, nsGroupStubResourceInstance),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "name", nsGroupStubName),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "comment", "Infoblox Terraform Acceptance test"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.#", "3"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.0.name", "grid-member01.example.com"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.1.name", "grid-member02.example.com"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.2.name", "grid-member03.example.com"),
				),
			},
			{
				Config: testAccInfobloxNSGroupStubUpdateTemplate(nsGroupStubNameUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSGroupStubCheckExists(nsGroupStubNameUpdated, nsGroupStubResourceInstance),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "name", nsGroupStubNameUpdated),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "comment", "Infoblox Terraform Acceptance test - updated"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.#", "4"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.0.name", "grid-member01.example.com"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.1.name", "grid-member02.example.com"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.2.name", "grid-member03.example.com"),
					resource.TestCheckResourceAttr(nsGroupStubResourceInstance, "stub_members.3.name", "grid-member04.example.com"),
				),
			},
		},
	})
}

func testAccInfobloxNSGroupStubCheckDestroy(state *terraform.State, name string) error {

	client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_ns_group_stub" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		api := nsgroupstub.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox - error occurred whilst retrieving a list of NS Group Stub")
		}
		for _, nsGroupStub := range *api.ResponseObject().(*[]nsgroupstub.NSGroupStub) {
			if nsGroupStub.Name == name {
				return fmt.Errorf("Infoblox NS Group Stub %s still exists", name)
			}
		}
	}
	return nil
}

func testAccInfobloxNSGroupStubCheckExists(name, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox NS Group Stub %s wasn't found in resources", name)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox NS Group Stub ID not set for %s in resources", name)
		}

		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		api := nsgroupstub.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox NS Group Stub - error whilst retrieving a list of NS Group Stub: %+v", err)
		}
		for _, nsGroupStub := range *api.ResponseObject().(*[]nsgroupstub.NSGroupStub) {
			if nsGroupStub.Name == name {
				return nil
			}
		}
		return fmt.Errorf("Infoblox NS Group Stub %s wasn't found on remote Infoblox server", name)
	}
}

func testAccInfobloxNSGroupStubNoNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_stub" "acctest" {
  comment = "Infoblox Terraform Acceptance test"
}
`)
}

func testAccInfobloxNSGroupStubCommentLeadingTrailingSpaces(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_stub" "acctest" {
  name = "%s"
  comment = " Infoblox Terraform Acceptance test "
}
`, name)
}

func testAccInfobloxNSGroupStubCreateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_stub" "acctest" {
  name = "%s"
  comment = "Infoblox Terraform Acceptance test"
  stub_members = [
    {
      name = "grid-member01.example.com"
    },
    {
      name = "grid-member02.example.com"
    },
    {
      name = "grid-member03.example.com"
    },
  ]
}
`, name)
}

func testAccInfobloxNSGroupStubUpdateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_group_stub" "acctest" {
  name = "%s"
  comment = "Infoblox Terraform Acceptance test - updated"
  stub_members = [
    {
      name = "grid-member01.example.com"
    },
    {
      name = "grid-member02.example.com"
    },
    {
      name = "grid-member03.example.com"
    },
    {
      name = "grid-member04.example.com"
    },
  ]
}
`, name)
}
*/
