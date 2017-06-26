package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/zoneauth"
	"regexp"
	"testing"
)

func TestAccInfobloxZoneAuthBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	testFQDN := fmt.Sprintf("acctest-infoblox-zone-auth-%d.com", randomInt)
	testFQDNResourceName := "infoblox_zone_auth.acctest"

	fmt.Printf("\n\nForward FQDN is %s\n\n", testFQDN)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccInfobloxZoneAuthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxZoneAuthNoFQDNTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testAccInfobloxZoneAuthEmptyTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testAccInfobloxZoneAuthTooLongCommentTemplate(testFQDN),
				ExpectError: regexp.MustCompile(`Infoblox Zone Create Error: Invalid HTTP response code 400 returned`),
			},
			{
				Config:      testAccInfobloxZoneAuthInvalidZoneFormat(testFQDN),
				ExpectError: regexp.MustCompile(`must be one of FORWARD, IPV4 or IPV6`),
			},
			{
				Config:      testAccInfobloxZoneAuthInvalidSOATTL(testFQDN),
				ExpectError: regexp.MustCompile(`can't be negative`),
			},
			{
				Config:      testAccInfobloxZoneAuthInvalidAllowUpdatePermission(testFQDN),
				ExpectError: regexp.MustCompile(`must be one of ALLOW or DENY`),
			},
			{
				Config:      testAccInfobloxZoneAuthInvalidAllowUpdateType(testFQDN),
				ExpectError: regexp.MustCompile(`must be one of addressac or tsigac`),
			},
			{
				Config:      testAccInfobloxZoneAuthInvalidAllowUpdateTSIGAlgorithm(testFQDN),
				ExpectError: regexp.MustCompile(`must be one of HMAC-MD5 or HMAC-SHA256`),
			},
			{
				Config:      testAccInfobloxZoneAuthLeadingTrailingWhiteSpace(testFQDN),
				ExpectError: regexp.MustCompilePOSIX(`must not contain trailing or leading white space`),
			},
			{
				Config: testAccInfobloxZoneAuthCreateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneAuthExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "Created a zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zoneformat", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "prefix", "128/16"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soattl", "3600"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soanegativettl", "60"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soarefresh", "1200"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soaretry", "300"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "disable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dnsintegrityenable", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dnsintegritymember", "s1ins01.devops.int.ovp.bskyb.com"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "locked", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.address", "192.168.100.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.permission", "ALLOW"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.address", "192.168.101.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.permission", "ALLOW"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.type", "tsigac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.tsigkey", "0jnu3SdsMvzzlmTDPYRceA=="),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.tsigkeyalgorithm", "HMAC-SHA256"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.tsigkeyname", "abc.key"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.usetsigkeyname", "true"),
				),
			},
			{
				Config: testAccInfobloxZoneAuthUpdateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneAuthExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "Updated a zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zoneformat", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "prefix", "128-189"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soattl", "7200"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soanegativettl", "90"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soarefresh", "1800"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soaretry", "150"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "disable", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dnsintegrityenable", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dnsintegritymember", "h1ins01.devops.int.ovp.bskyb.com"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "locked", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.type", "tsigac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.tsigkey", "0jnu3SdsMvzzlmToPYRceA=="),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.tsigkeyalgorithm", "HMAC-MD5"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.tsigkeyname", "test.key"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.usetsigkeyname", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.address", "192.168.120.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.permission", "DENY"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.address", "192.168.120.11"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.permission", "ALLOW"),
				),
			},
			{
				Config: testAccInfobloxZoneAuthUpdateTwoTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneAuthExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "Updated a zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zoneformat", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "prefix", "128-189"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soattl", "7200"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soanegativettl", "90"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soarefresh", "1800"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soaretry", "150"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "disable", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dnsintegrityenable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dnsintegritymember", "h1ins01.devops.int.ovp.bskyb.com"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "locked", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.type", "tsigac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.tsigkey", "0jnu3SdsMvzzlmToPYRceA=="),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.tsigkeyalgorithm", "HMAC-MD5"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.tsigkeyname", "test.key"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.0.usetsigkeyname", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.address", "192.168.120.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.1.permission", "DENY"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.type", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.address", "192.168.120.11"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allowupdate.2.permission", "ALLOW"),
				),
			},
		},
	})
}

func testAccInfobloxZoneAuthExists(testFQDN, resourceName string) resource.TestCheckFunc {
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

func testAccInfobloxZoneAuthCheckDestroy(state *terraform.State) error {

	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_zone_auth" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id != "" {
			return nil
		}
		api := zoneauth.NewGetAllZones()
		err := infobloxClient.Do(api)
		if err != nil {
			return nil
		}
		for _, zone := range *api.GetResponse() {
			matched, _ := regexp.MatchString("acctest-infoblox-zone-auth-.*.ovp.bskyb.com", zone.FQDN)
			if matched {
				return fmt.Errorf("Infoblox Zone %s still exists", zone.FQDN)
			}
		}
	}

	return nil
}

func testAccInfobloxZoneAuthNoFQDNTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
comment = "Updated a zone"
}
`)
}

func testAccInfobloxZoneAuthEmptyTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
}
`)
}

func testAccInfobloxZoneAuthTooLongCommentTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
comment = "This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string...."
}
`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidZoneFormat(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
comment = "Created a zone"
zoneformat = "SOME_INVALID_ZONE"
}`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidSOATTL(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
soattl = -1
}`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidAllowUpdatePermission(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
allowupdate = [
{
  type = "addressac"
  address = "192.168.100.10"
  permission = "SOME_INVALID_PERMISSION"
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidAllowUpdateType(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
allowupdate = [
{
  type = "SOME_INVALID_TYPE"
  address = "192.168.100.10"
  permission = "ALLOW"
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidAllowUpdateTSIGAlgorithm(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
allowupdate = [
{
  type = "tsigac"
  tsigkey = "0jnu3SdsMvzzlmTDPYTceA=="
  tsigkeyalgorithm = "SOME_INVALID_ALGORITHM"
  tsigkeyname = "test.key"
  usetsigkeyname = true
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthLeadingTrailingWhiteSpace(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
allowupdate = [
{
  type = "addressac"
  address = " 192.168.100.10"
  permission = "ALLOW"
},
{
  type = "tsigac"
  tsigkey = " 0jnu3SdsMvzzlmTDPYTceA== "
  tsigkeyalgorithm = "HMAC-SHA256"
  tsigkeyname = " test.key "
  usetsigkeyname = true
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthCreateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
comment = "Created a zone"
zoneformat = "FORWARD"
view = "default"
prefix = "128/16"
soattl = 3600
soanegativettl = 60
soarefresh = 1200
soaretry = 300
disable = false
dnsintegrityenable = true
dnsintegritymember = "s1ins01.devops.int.ovp.bskyb.com"
locked = true
allowupdate = [
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
  tsigkey = "0jnu3SdsMvzzlmTDPYRceA=="
  tsigkeyalgorithm = "HMAC-SHA256"
  tsigkeyname = "abc.key"
  usetsigkeyname = true
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthUpdateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
comment = "Updated a zone"
zoneformat = "FORWARD"
view = "default"
prefix = "128-189"
soattl = 7200
soanegativettl = 90
soarefresh = 1800
soaretry = 150
disable = true
dnsintegrityenable = true
dnsintegritymember = "h1ins01.devops.int.ovp.bskyb.com"
locked = false
allowupdate = [
{
  type = "tsigac"
  tsigkey = "0jnu3SdsMvzzlmToPYRceA=="
  tsigkeyalgorithm = "HMAC-MD5"
  tsigkeyname = "test.key"
  usetsigkeyname = false
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

func testAccInfobloxZoneAuthUpdateTwoTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
comment = "Updated a zone"
zoneformat = "FORWARD"
view = "default"
prefix = "128-189"
soattl = 7200
soanegativettl = 90
soarefresh = 1800
soaretry = 150
disable = true
dnsintegrityenable = false
dnsintegritymember = "h1ins01.devops.int.ovp.bskyb.com"
locked = false
allowupdate = [
{
  type = "tsigac"
  tsigkey = "0jnu3SdsMvzzlmToPYRceA=="
  tsigkeyalgorithm = "HMAC-MD5"
  tsigkeyname = "test.key"
  usetsigkeyname = false
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
