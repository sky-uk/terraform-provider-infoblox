package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"testing"
)

func TestAccInfobloxZoneDelegated(t *testing.T) {
	zoneFqdn := fmt.Sprintf("prd%d.hempaas.bskyb.com", acctest.RandInt())
	resourceName := "infoblox_zone_delegated.delegationtest"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.ZONEDelegatedObj, "fqdn", zoneFqdn)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccInfobloxZoneDelegatedCreateTemplate(zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneDelegatedExists(zoneFqdn, resourceName),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "FORWARD"),
				),
			}, {
				Config: testAccInfobloxZoneDelegatedUpdateTemplate(zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneDelegatedExists(zoneFqdn, resourceName),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is the comment after we changed it"),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "FORWARD"),
				),
			},
		},
	})
}

func testAccInfobloxZoneDelegatedExists(zoneFqdn, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.ZONEDelegatedObj, "fqdn", zoneFqdn)
	}
}

func testAccInfobloxZoneDelegatedCreateTemplate(zoneName string) string {
	return fmt.Sprintf(`
	resource "infoblox_zone_delegated" "delegationtest" {
        fqdn = "%s"
        comment = "this is a comment"
        disable = false
        zone_format = "FORWARD"
        delegate_to {
                name="prdibxdns03.bskyb.com"
                address="10.92.16.131"
        }
        }`, zoneName)
}

func testAccInfobloxZoneDelegatedUpdateTemplate(zoneName string) string {
	return fmt.Sprintf(`
	resource "infoblox_zone_delegated" "delegationtest" {
        fqdn = "%s"
        comment = "this is the comment after we changed it"
        disable = true
        zone_format = "FORWARD"
        delegate_to {
                name="prdibxdns03.bskyb.com"
                address="10.92.16.131"
        }
        }`, zoneName)
}
