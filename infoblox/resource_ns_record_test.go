package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/nameserver"
	"regexp"
	"testing"
)

func TestAccInfobloxNSRecordBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	nameServerZoneName := "paas-testing.com"
	createNameServer := fmt.Sprintf("acctest-infoblox-%d-nameserver.%s", randomInt, nameServerZoneName)
	updateNameServer := fmt.Sprintf("acctest-infoblox-%d-nameserver-update.%s", randomInt, nameServerZoneName)
	nsResourceName := "infoblox_ns_record.acctest"
	nameServerAddressIPPattern := regexp.MustCompile(`name_server_addresses\.[0-9]+\.ip_address`)
	nameServerAddressPTRPattern := regexp.MustCompile(`name_server_addresses\.[0-9]+\.auto_create_PTR_record`)

	fmt.Printf("\n\nAcceptance Test Name Servers are:\n \tcreate:%s, \n\tupdate:%s\n\n", createNameServer, updateNameServer)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxNSRecordCheckDestroy(state, createNameServer, updateNameServer)
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
					resource.TestCheckResourceAttr(nsResourceName, "zone_name", nameServerZoneName),
					// TODO we need a ms_delegation_name to test against resource.TestCheckResourceAttr(nsResourceName, "ms_delegation_name", "some_ms_delegation_name"),
					resource.TestCheckResourceAttr(nsResourceName, "name_server", createNameServer),
					resource.TestCheckResourceAttr(nsResourceName, "view", "default"),
					resource.TestCheckResourceAttr(nsResourceName, "name_server_addresses.#", "2"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "true"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.0.1"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "false"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.0.2"),
				),
			},
			{
				Config: testAccInfobloxNSRecordUpdateTemplate(nameServerZoneName, updateNameServer),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxNSRecordCheckExists(updateNameServer, nsResourceName),
					resource.TestCheckResourceAttr(nsResourceName, "zone_name", nameServerZoneName),
					// TODO we need a ms_delegation_name to test against resource.TestCheckResourceAttr(nsResourceName, "ms_delegation_name", "another_ms_delegation_name"),
					resource.TestCheckResourceAttr(nsResourceName, "name_server", updateNameServer),
					resource.TestCheckResourceAttr(nsResourceName, "view", "default"),
					resource.TestCheckResourceAttr(nsResourceName, "name_server_addresses.#", "3"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "false"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.1.1"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "true"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.1.2"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressPTRPattern, "true"),
					testAccInfobloxNSRecordCheckValueInKeyPattern(nsResourceName, nameServerAddressIPPattern, "192.168.1.3"),
				),
			},
		},
	})
}

func testAccInfobloxNSRecordCheckValueInKeyPattern(nsRecordResource string, keyPattern *regexp.Regexp, checkValue string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[nsRecordResource]
		if ok {
			for attributeKey, attributeValue := range rs.Primary.Attributes {
				if keyPattern.MatchString(attributeKey) {
					if attributeValue == checkValue {
						return nil
					}
				}
			}
		}
		return fmt.Errorf("Infoblox NS Record attribute %s not found", checkValue)
	}
}

func testAccInfobloxNSRecordCheckDestroy(state *terraform.State, createNameServer, updateNameServer string) error {

	client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_ns_record" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		api := nameserver.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox - error occurred whilst retrieving a list of NS records")
		}
		for _, nsRecord := range *api.ResponseObject().(*[]nameserver.NSRecord) {
			for _, checkNameServer := range []string{createNameServer, updateNameServer} {
				if nsRecord.NameServer == checkNameServer {
					return fmt.Errorf("Infoblox NS record for %s still exists", checkNameServer)
				}
			}
		}
	}
	return nil
}

func testAccInfobloxNSRecordCheckExists(nameServer, nameServerResource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[nameServerResource]
		if !ok {
			return fmt.Errorf("\nInfoblox NS record for %s wasn't found in resources", nameServer)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox NS record ID not set for %s in resources", nameServer)
		}

		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		api := nameserver.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox NS record - error whilst retrieving a list of NS records: %+v", err)
		}
		for _, nsRecord := range *api.ResponseObject().(*[]nameserver.NSRecord) {
			if nsRecord.NameServer == nameServer {
				return nil
			}
		}
		return fmt.Errorf("Infoblox NS record for %s wasn't found on remote Infoblox server", nameServer)
	}
}

func testAccInfobloxNSRecordNoZoneNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_ns_record" "acctest" {
    // TODO we need a ms_delegation_name to test against ms_delegation_name = "some_ms_delegation_name"
    name_server = "ns1.example.com"
    view = "default"
    name_server_addresses = [
        {
            ip_address = "192.168.0.1"
            auto_create_PTR_record = true
        },
        {
            ip_address = "192.168.0.2"
            auto_create_PTR_record = false
        },
    ]
}
`)
}

func testAccInfobloxNSRecordTestValidateWhiteSpaceTemplate(zoneName string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_record" "acctest" {
    zone_name = "%s"
    // TODO we need a ms_delegation_name to test against ms_delegation_name = "some_ms_delegation_name"
    name_server = "    ns1.example.com   "
    view = "default"
    name_server_addresses = [
        {
            ip_address = "192.168.0.1"
            auto_create_PTR_record = true
        },
        {
            ip_address = "192.168.0.2"
            auto_create_PTR_record = false
        },
    ]
}
`, zoneName)
}

func testAccInfobloxNSRecordCreateTemplate(zoneName, nameServer string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_record" "acctest" {
    zone_name = "%s"
    // TODO we need a ms_delegation_name to test against ms_delegation_name = "some_ms_delegation_name"
    name_server = "%s"
    view = "default"
    name_server_addresses = [
        {
            ip_address = "192.168.0.1"
            auto_create_PTR_record = true
        },
        {
            ip_address = "192.168.0.2"
            auto_create_PTR_record = false
        },
    ]
}
`, zoneName, nameServer)
}

func testAccInfobloxNSRecordUpdateTemplate(zoneName, nameServer string) string {
	return fmt.Sprintf(`
resource "infoblox_ns_record" "acctest" {
    zone_name = "%s"
    // TODO we need a ms_delegation_name to test against ms_delegation_name = "another_ms_delegation_name"
    name_server = "%s"
    view = "default"
    name_server_addresses = [
        {
            ip_address = "192.168.1.1"
            auto_create_PTR_record = false
        },
        {
            ip_address = "192.168.1.2"
            auto_create_PTR_record = true
        },
        {
            ip_address = "192.168.1.3"
            auto_create_PTR_record = true
        },
    ]
}
`, zoneName, nameServer)
}
