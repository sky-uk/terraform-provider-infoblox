package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"github.com/sky-uk/skyinfoblox/api/zonestub"
	"strconv"
	"testing"
)

func TestAccInfobloxZoneStub(t *testing.T) {
	testFQDN := "acctest-infoblox-zone-stub-" + strconv.Itoa(acctest.RandInt()) + ".slupaas.bskyb.com"
	testFQDNResourceName := "infoblox_zone_stub.stub1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxZoneStubCheckDestroy(state, testFQDN)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccInfobloxZoneStubCreateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneStubExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "this is a stub zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zoneformat", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
				),
			},
			{
				Config: testAccInfobloxZoneStubUpdateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneStubExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "this is a stub comment"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zoneformat", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
				),
			},
		},
	})
}

func testAccInfobloxZoneStubExists(testFQDN, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Infoblox Zone Stub resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Infoblox Zone Stub resource ID not set in resources ")
		}
		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		getAllAPI := zonestub.NewGetAll([]string{"fqdn", "comment"})

		err := client.Do(getAllAPI)
		if err != nil {
			return fmt.Errorf("Error: %+v", err)
		}
		for _, dnsZoneReference := range *getAllAPI.ResponseObject().(*[]zonestub.ZoneStub) {
			if testFQDN == dnsZoneReference.FQDN {
				return nil
			}
		}
		return fmt.Errorf("Infoblox Zone %s wasn't found", testFQDN)
	}
}

func testAccInfobloxZoneStubCheckDestroy(state *terraform.State, fqdn string) error {

	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_zone_stub" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		zonestub.NewGetAll([]string{"fqdn", "comment"})
		api := zoneauth.NewGetAllZones()
		err := infobloxClient.Do(api)
		if err != nil {
			return nil
		}
		for _, zone := range *api.GetResponse() {
			if zone.FQDN == fqdn {
				return fmt.Errorf("Infoblox Zone %s still exists", fqdn)
			}
		}
	}

	return nil
}

func testAccInfobloxZoneStubCreateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
	resource "infoblox_zone_stub" "stub1" {
  		fqdn = "%s"
  		comment = "this is a stub zone"
  		disable = false
  		locked = false
  		stub_from = [{
          		name="dns1.example.com"
          		address="1.1.1.1"
          		}]
		}`, testFQDN)
}

func testAccInfobloxZoneStubUpdateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
	resource "infoblox_zone_stub" "stub1" {
  		fqdn = "%s"
  		comment = "this is a stub comment"
  		disable = false
  		locked = false
  		stub_from = [{
          		name="dns1.example.com"
          		address="1.1.1.1"
          		}]
		}`, testFQDN)
}
