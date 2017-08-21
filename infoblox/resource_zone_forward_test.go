package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneforward"
	"net/http"
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
			return zoneForwardCheckDestroy(state, testFQDN)
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

func zoneForwardCheckDestroy(state *terraform.State, fqdn string) error {

	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_zone_forward" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		api := zoneforward.NewGetAll()
		err := infobloxClient.Do(api)
		if err != nil {
			return nil
		}
		if api.StatusCode() != http.StatusOK {
			return fmt.Errorf("Error getting all zones")
		}
		zones := *api.ResponseObject().(*[]zoneforward.ZoneForward)
		for _, zone := range zones {
			if zone.Fqdn == fqdn {
				return fmt.Errorf("Infoblox Zone %s still exists", fqdn)
			}
		}
	}

	return nil
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

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Infoblox Zone Auth resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Infoblox Zone Auth resource ID not set in resources ")
		}
		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		api := zoneforward.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Error: %+v", err)
		}
		for _, zone := range *api.ResponseObject().(*[]zoneforward.ZoneForward) {
			if testFQDN == zone.Fqdn {
				return nil
			}
		}
		return fmt.Errorf("Infoblox Zone %s wasn't found", testFQDN)
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
        comment = "Updated forward zone"
        fqdn = "%s"
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
    }`, testFQDN)
}
