package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"strconv"
	"testing"
)

func TestAccInfobloxZoneAuthSecondary(t *testing.T) {

	testFQDN := "acctest-infoblox-zone-auth-secondary" + strconv.Itoa(acctest.RandInt()) + ".slupaas.bskyb.com"
	testFQDNResourceName := "infoblox_zone_auth.acctestsecondary"

	fmt.Printf("\n\nForward FQDN is %s\n\n", testFQDN)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccInfobloxZoneAuthSecondaryCheckDestroy(state, testFQDN)
		},
		Steps: []resource.TestStep{

			{
				Config: testAccInfobloxZoneAuthSecondaryCreateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneAuthSecondaryExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "Created a zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "prefix", "128/16"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "disable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_enable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_member", "nonprdibxdns01.bskyb.com"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "locked", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "copy_xfer_to_notify", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "use_copy_xfer_to_notify", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.address", "192.168.100.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.permission", "ALLOW"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.address", "192.168.101.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.permission", "ALLOW"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.type", "tsigac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.tsig_key", "0jnu3SdsMvzzlmTDPYRceA=="),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.tsig_key_alg", "HMAC-SHA256"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.tsig_key_name", "acc-test.key"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.use_tsig_key_name", "false"),
				),
			},
			{
				Config: testAccInfobloxZoneAuthSecondaryUpdateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneAuthSecondaryExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "Updated a zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "prefix", "128-189"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "disable", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_enable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_member", "nonprdibxdns01.bskyb.com"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "locked", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "copy_xfer_to_notify", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "use_copy_xfer_to_notify", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "use_external_primary", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.type", "tsigac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key", "0jnu3SdsMvzzlmToPYRceA=="),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key_alg", "HMAC-MD5"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key_name", "test.key"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.use_tsig_key_name", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.address", "192.168.120.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.permission", "DENY"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.address", "192.168.120.11"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.permission", "ALLOW"),
				),
			},
		},
	})
}

func testAccInfobloxZoneAuthSecondaryExists(testFQDN, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Infoblox Zone Auth resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Infoblox Zone Auth resource ID not set in resources ")
		}
		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		getAllAPI := zoneauth.NewGetAllZones()

		err := client.Do(getAllAPI)
		if err != nil {
			return fmt.Errorf("Error: %+v", err)
		}
		for _, dnsZoneReference := range *getAllAPI.GetResponse() {
			if testFQDN == dnsZoneReference.FQDN {
				return nil
			}
		}
		return fmt.Errorf("Infoblox Zone %s wasn't found", testFQDN)
	}
}

func testAccInfobloxZoneAuthSecondaryCheckDestroy(state *terraform.State, fqdn string) error {

	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_zone_auth" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
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

func testAccInfobloxZoneAuthSecondaryCreateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctestsecondary" {
fqdn = "%s"
comment = "Created a zone"
zone_format = "FORWARD"
view = "default"
prefix = "128/16"
disable = false
dns_integrity_enable = false
dns_integrity_member = "nonprdibxdns01.bskyb.com"
locked = true
copy_xfer_to_notify = false
use_copy_xfer_to_notify = false
use_external_primary = true
external_primaries = [
        {
            address = "10.0.0.2"
            name = "ns1.example.com"
            stealth = false
            tsig_key = "dFghJkcXb5tyUio3eWo021=="
            tsig_key_alg = "HMAC-SHA256"
            tsig_key_name = "example-key"
            use_tsig_key_name = false
        },
    ]
grid_secondaries = [
        {
            lead = false
            name = "slunonprdigm01.bskyb.com"
            enablepreferredprimaries = false
            stealth = false
        },
    ]
allow_update = [
{
  type = "addressac"
  address = "192.168.100.10"
  permission = "ALLOW"
},
{
  type = "addressac"
  address = "192.168.101.10"
  permission = "ALLOW"
},
{
  type = "tsigac"
  tsig_key = "0jnu3SdsMvzzlmTDPYRceA=="
  tsig_key_alg = "HMAC-SHA256"
  tsig_key_name = "acc-test.key"
  use_tsig_key_name = false
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthSecondaryUpdateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctestsecondary" {
fqdn = "%s"
comment = "Updated a zone"
zone_format = "FORWARD"
view = "default"
prefix = "128-189"
disable = true
dns_integrity_enable = false
dns_integrity_member = "nonprdibxdns01.bskyb.com"
locked = false
copy_xfer_to_notify = true
use_copy_xfer_to_notify = true
use_allow_transfer = false
use_external_primary = true
external_primaries = [
        {
            address = "10.0.0.3"
            name = "ns2.example.com"
            stealth = false
            tsig_key = "dFghJkcXb5tyUio3eWo021=="
            tsig_key_alg = "HMAC-SHA256"
            tsig_key_name = "example-key"
            use_tsig_key_name = false
        },
    ]
grid_secondaries = [
        {
            lead = false
            name = "slunonprdigm01.bskyb.com"
            enablepreferredprimaries = false
            stealth = false
        },
    ]
allow_transfer = [
      {
        type = "addressac"
        address = "192.168.234.11"
        permission = "DENY"
      },
      {
        type = "tsigac"
        tsig_key = "0jnu3SdsMvzzlmTDPYTceA=="
        tsig_key_alg = "HMAC-SHA256"
        tsig_key_name = "abc.key"
        use_tsig_key_name = false
      },
      {
        type = "addressac"
        address = "192.168.101.11"
        permission = "ALLOW"
      },
      {
        type = "addressac"
        address = "192.168.111.10"
        permission = "ALLOW"
      },
    ]
allow_update = [
{
  type = "tsigac"
  tsig_key = "0jnu3SdsMvzzlmToPYRceA=="
  tsig_key_alg = "HMAC-MD5"
  tsig_key_name = "test.key"
  use_tsig_key_name = false
},
{
  type = "addressac"
  address = "192.168.120.10"
  permission = "DENY"
},
{
  type = "addressac"
  address = "192.168.120.11"
  permission = "ALLOW"
},]}`, testFQDN)
}
