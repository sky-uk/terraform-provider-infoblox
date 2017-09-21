package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
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
			return TestAccCheckDestroy(model.ZONESTUBObj, "fqdn", testFQDN)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccInfobloxZoneStubCreateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneStubExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "this is a stub zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
				),
			},
			{
				Config: testAccInfobloxZoneStubUpdateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneStubExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "this is a stub comment"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
				),
			},
		},
	})
}

func testAccInfobloxZoneStubExists(testFQDN, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.ZONESTUBObj, "fqdn", testFQDN)
	}
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
