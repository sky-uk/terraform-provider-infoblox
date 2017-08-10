package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zonedelegated"
	"testing"
)

func TestAccInfobloxZoneDelegated(t *testing.T) {
	zoneFqdn := fmt.Sprintf("prd%d.hempaas.bskyb.com", acctest.RandInt())
	resourceName := "infoblox_zone_delegated.delegationtest"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxZoneDelegatedCheckDestroy(state, zoneFqdn)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccInfobloxZoneDelegatedCreateTemplate(zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneDelegatedExists(zoneFqdn, resourceName),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
					resource.TestCheckResourceAttr(resourceName, "zoneformat", "FORWARD"),
				),
			}, {
				Config: testAccInfobloxZoneDelegatedUpdateTemplate(zoneFqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneDelegatedExists(zoneFqdn, resourceName),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is the comment after we changed it"),
					resource.TestCheckResourceAttr(resourceName, "zoneformat", "FORWARD"),
				),
			},
		},
	})
}

func testAccInfobloxZoneDelegatedExists(zoneFqdn, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Infoblox Zone Delegated resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Infoblox Zone Delegated resource ID not set in resources ")
		}
		fields := []string{"fqdn"}
		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		getAllAPI := zonedelegated.NewGetAll(fields)
		err := client.Do(getAllAPI)
		if err != nil {
			return fmt.Errorf("Error: %+v", err)
		}
		for _, dnsZoneDelegated := range *getAllAPI.ResponseObject().(*[]zonedelegated.ZoneDelegated) {
			if zoneFqdn == dnsZoneDelegated.Fqdn {
				return nil
			}
		}
		return fmt.Errorf("Infoblox Zone %s wasn't found", zoneFqdn)
	}
}

func testAccInfobloxZoneDelegatedCheckDestroy(state *terraform.State, zoneFqdn string) error {
	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_zone_delegated" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		fields := []string{"fqdn"}
		api := zonedelegated.NewGetAll(fields)
		err := infobloxClient.Do(api)
		if err != nil {
			return fmt.Errorf("Could not destroy the resource %s", err.Error())
		}
		for _, zone := range *api.ResponseObject().(*[]zonedelegated.ZoneDelegated) {
			if zone.Fqdn == zoneFqdn {
				return fmt.Errorf("Infoblox Zone %s still exists", zoneFqdn)
			}
		}

	}
	return nil

}

func testAccInfobloxZoneDelegatedCreateTemplate(zoneName string) string {
	return fmt.Sprintf(`
	resource "infoblox_zone_delegated" "delegationtest" {
        fqdn = "%s"
        comment = "this is a comment"
        disable = false
        zoneformat = "FORWARD"
        delegateto {
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
        zoneformat = "FORWARD"
        delegateto {
                name="prdibxdns03.bskyb.com"
                address="10.92.16.131"
        }
        }`, zoneName)
}
