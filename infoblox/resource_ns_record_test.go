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

func TestAccInfobloxNSRecordBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	nameServerZoneName := "paas-testing.com"
	createNameServer := fmt.Sprintf("acctest-infoblox-%d-nameserver.%s", randomInt, nameServerZoneName)
	updateNameServer := fmt.Sprintf("acctest-infoblox-%d-nameserver-update.%s", randomInt, nameServerZoneName)
	nsResourceName := "infoblox_ns_record.acctest"
	nameServerAddressIPPattern := regexp.MustCompile(`addresses\.[0-9]+\.address`)
	nameServerAddressPTRPattern := regexp.MustCompile(`addresses\.[0-9]+\.auto_create_ptr`)

	fmt.Printf("\n\nAcceptance Test Name Servers are:\n \tcreate:%s, \n\tupdate:%s\n\n", createNameServer, updateNameServer)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.RecordNSObj, "nameserver", createNameServer)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxNSRecordNoZoneNameTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testAccInfobloxNSRecordTestValidateWhiteSpaceTemplate(nameServerZoneName),
				ExpectError: regexp.MustCompile(`must not contain trailing or leading white space`),
			},
			{
				Config: testAccInfobloxNSRecordCreateTemplate(nameServerZoneName, createNameServer),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSRecordCheckExists(createNameServer, nsResourceName),
					resource.TestCheckResourceAttr(nsResourceName, "name", nameServerZoneName),
					// TODO we need a ms_delegation_name to test against resource.TestCheckResourceAttr(nsResourceName, "ms_delegation_name", "some_ms_delegation_name"),
					resource.TestCheckResourceAttr(nsResourceName, "nameserver", createNameServer),
					resource.TestCheckResourceAttr(nsResourceName, "view", "default"),
					resource.TestCheckResourceAttr(nsResourceName, "addresses.#", "2"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "true"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.0.1"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "false"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.0.2"),
				),
			},
			{
				Config: testAccInfobloxNSRecordUpdateTemplate(nameServerZoneName, updateNameServer),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSRecordCheckExists(updateNameServer, nsResourceName),
					resource.TestCheckResourceAttr(nsResourceName, "name", nameServerZoneName),
					// TODO we need a ms_delegation_name to test against resource.TestCheckResourceAttr(nsResourceName, "ms_delegation_name", "another_ms_delegation_name"),
					resource.TestCheckResourceAttr(nsResourceName, "nameserver", updateNameServer),
					resource.TestCheckResourceAttr(nsResourceName, "view", "default"),
					resource.TestCheckResourceAttr(nsResourceName, "addresses.#", "3"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "false"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.1.1"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "true"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.1.2"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "true"),
					util.AccTestCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.1.3"),
				),
			},
		},
	})
}

func testAccInfobloxNSRecordCheckExists(nameServer, nameServerResource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.RecordNSObj, "nameserver", nameServer)
	}
}

func testAccInfobloxNSRecordNoZoneNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_ns_record" "acctest" {
    // TODO we need a ms_delegation_name to test against ms_delegation_name = "some_ms_delegation_name"
    nameserver = "ns1.example.com"
    view = "default"
    addresses = [
        {
            address = "192.168.0.1"
            auto_create_ptr = true
        },
        {
            address = "192.168.0.2"
            auto_create_ptr = false
        },
    ]
}
`)
}

func testAccInfobloxNSRecordTestValidateWhiteSpaceTemplate(zoneName string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_record" "acctest" {
    name = "%s"
    // TODO we need a ms_delegation_name to test against ms_delegation_name = "some_ms_delegation_name"
    nameserver = "    ns1.example.com   "
    view = "default"
    addresses = [
        {
            address = "192.168.0.1"
            auto_create_ptr = true
        },
        {
            address = "192.168.0.2"
            auto_create_ptr = false
        },
    ]
}
`, zoneName)
}

func testAccInfobloxNSRecordCreateTemplate(zoneName, nameServer string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_record" "acctest" {
    name = "%s"
    // TODO we need a ms_delegation_name to test against ms_delegation_name = "some_ms_delegation_name"
    nameserver = "%s"
    view = "default"
    addresses = [
        {
            address = "192.168.0.1"
            auto_create_ptr = true
        },
        {
            address = "192.168.0.2"
            auto_create_ptr = false
        },
    ]
}
`, zoneName, nameServer)
}

func testAccInfobloxNSRecordUpdateTemplate(zoneName, nameServer string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_record" "acctest" {
    name = "%s"
    // TODO we need a ms_delegation_name to test against ms_delegation_name = "another_ms_delegation_name"
    nameserver = "%s"
    view = "default"
    addresses = [
        {
            address = "192.168.1.1"
            auto_create_ptr = false
        },
        {
            address = "192.168.1.2"
            auto_create_ptr = true
        },
        {
            address = "192.168.1.3"
            auto_create_ptr = true
        },
    ]
}
`, zoneName, nameServer)
}
