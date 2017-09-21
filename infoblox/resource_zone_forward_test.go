package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"os"
	"regexp"
	"strconv"
	"testing"
)

func TestAccInfobloxZoneForwardBasic(t *testing.T) {

	testFQDN := "acctest-infoblox-zone-forward-" + strconv.Itoa(acctest.RandInt()) + ".slupaas.bskyb.com"
	zoneForwardName := "infoblox_zone_forward.acctest"

	fmt.Printf("\n\nForward FQDN is %s\n\n", testFQDN)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.ZONEForwardObj, "fqdn", testFQDN)
		},
		Steps: []resource.TestStep{

			{
				Config:      testZoneForwardTemplateEmpty(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testZoneForwardTemplateLongComment(testFQDN),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testZoneForwardInvalidZoneFormat(testFQDN),
				ExpectError: regexp.MustCompile(`must be one of FORWARD, IPV4 or IPV6`),
			},
			{
				Config: testZoneForwardCreateComplete(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testZoneForwardExists(testFQDN, zoneForwardName),
					resource.TestCheckResourceAttr(zoneForwardName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(zoneForwardName, "comment", "Created a forward zone"),
					resource.TestCheckResourceAttr(zoneForwardName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(zoneForwardName, "view", "default"),
					resource.TestCheckResourceAttr(zoneForwardName, "prefix", "128/16"),
					resource.TestCheckResourceAttr(zoneForwardName, "disable", "false"),
					resource.TestCheckResourceAttr(zoneForwardName, "locked", "true"),
					resource.TestCheckResourceAttr(zoneForwardName, "locked_by", os.Getenv("INFOBLOX_USERNAME")),
					resource.TestCheckResourceAttr(zoneForwardName, "forwarders_only", "false"),
					resource.TestCheckResourceAttr(zoneForwardName, "forward_to.0.address", "10.90.233.150"),
					resource.TestCheckResourceAttr(zoneForwardName, "forward_to.0.name", "slupaas.bskyb.com"),
					resource.TestCheckResourceAttr(zoneForwardName, "forward_to.0.stealth", "false"),
					resource.TestCheckResourceAttr(zoneForwardName, "forward_to.0.tsig_key_alg", "HMAC-SHA256"),
				),
			},
			{
				Config: testZoneForwardUpdateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testZoneForwardExists(testFQDN, zoneForwardName),
					resource.TestCheckResourceAttr(zoneForwardName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(zoneForwardName, "comment", "Updated forward zone"),
					resource.TestCheckResourceAttr(zoneForwardName, "view", "default"),
					resource.TestCheckResourceAttr(zoneForwardName, "prefix", "128-189"),
					resource.TestCheckResourceAttr(zoneForwardName, "disable", "true"),
					resource.TestCheckResourceAttr(zoneForwardName, "locked", "false"),
					resource.TestCheckResourceAttr(zoneForwardName, "forward_to.0.address", "10.74.233.150"),
					resource.TestCheckResourceAttr(zoneForwardName, "forward_to.0.name", "hemnonprdigmc01.bskyb.com"),
					resource.TestCheckResourceAttr(zoneForwardName, "forward_to.0.stealth", "false"),
					resource.TestCheckResourceAttr(zoneForwardName, "forward_to.0.tsig_key_alg", "HMAC-MD5"),
				),
			},
		},
	})
}

func testZoneForwardTemplateEmpty() string {
	return fmt.Sprintf(`
resource "infoblox_zone_forward" "acctest-empty-template" {
}
`)
}

func testZoneForwardTemplateLongComment(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_forward" "acctest-long-comment" {
fqdn = "%s"
comment = "This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string...."
}
`, testFQDN)
}

func testZoneForwardInvalidZoneFormat(testFQDN string) string {
	return fmt.Sprintf(`
        resource "infoblox_zone_forward" "acctest-invalid-zoneformat" {
        fqdn = "%s"
        comment = "Created a zone"
        zone_format = "SOME_INVALID_ZONE"
    }`, testFQDN)
}

func testZoneForwardExists(testFQDN, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.ZONEForwardObj, "fqdn", testFQDN)
	}
}

func testZoneForwardCreateComplete(testFQDN string) string {
	return fmt.Sprintf(`
    resource "infoblox_zone_forward" "acctest" {
      fqdn = "%s"
      comment = "Created a forward zone"
      zone_format = "FORWARD"
      view = "default"
      prefix = "128/16"
      disable = false
      locked = true
      forward_to = [{
          address = "10.90.233.150"
          name = "slupaas.bskyb.com"
          stealth = false
          tsig_key_alg = "HMAC-SHA256"
      }]
      forwarders_only = false
  }`, testFQDN)
}

func testZoneForwardUpdateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
    resource "infoblox_zone_forward" "acctest" {
      fqdn = "%s"
      comment = "Updated forward zone"
      zone_format = "FORWARD"
      view = "default"
      prefix = "128-189"
      disable = true
      locked = false
      forward_to = [{
          address = "10.74.233.150"
          name = "hemnonprdigmc01.bskyb.com"
          stealth = false
          tsig_key_alg = "HMAC-MD5"
      }]
      forwarders_only = false
  }`, testFQDN)
}
