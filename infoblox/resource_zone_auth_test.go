package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"regexp"
	"strconv"
	"testing"
)

func TestAccInfobloxZoneAuthBasic(t *testing.T) {

	testFQDN := "acctest-infoblox-zone-auth-" + strconv.Itoa(acctest.RandInt()) + ".slupaas.bskyb.com"
	testFQDNResourceName := "infoblox_zone_auth.acctest"

	fmt.Printf("\n\nForward FQDN is %s\n\n", testFQDN)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.ZONEAUTHObj, "fqdn", testFQDN)
		},
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
				ExpectError: regexp.MustCompile("Response status code: 400"),
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
				Config:      testAccInfobloxZoneAuthInvalidAllowUpdateStruct(testFQDN),
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
					resource.TestCheckResourceAttr(testFQDNResourceName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "prefix", "128/16"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_default_ttl", "3600"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_negative_ttl", "60"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_refresh", "1200"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_retry", "300"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_expire", "444444"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "disable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_enable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_member", "nonprdibxdns01.bskyb.com"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "locked", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "copy_xfer_to_notify", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "use_copy_xfer_to_notify", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0._struct", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.address", "192.168.100.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.permission", "ALLOW"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1._struct", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.address", "192.168.101.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.permission", "ALLOW"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2._struct", "tsigac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.tsig_key", "0jnu3SdsMvzzlmTDPYRceA=="),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.tsig_key_alg", "HMAC-SHA256"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.tsig_key_name", "acc-test.key"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.use_tsig_key_name", "true"),
				),
			},
			{
				Config: testAccInfobloxZoneAuthUpdateTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneAuthExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "Updated a zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "prefix", "128-189"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_default_ttl", "7200"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_negative_ttl", "90"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_refresh", "1800"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_retry", "150"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_expire", "888888"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "disable", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_enable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_member", "nonprdibxdns01.bskyb.com"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "locked", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "copy_xfer_to_notify", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "use_copy_xfer_to_notify", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "use_external_primary", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0._struct", "tsigac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key", "0jnu3SdsMvzzlmTDPYRceA=="),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key_alg", "HMAC-MD5"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key_name", "acc-test.key"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.use_tsig_key_name", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1._struct", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.address", "192.168.120.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.permission", "DENY"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2._struct", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.address", "192.168.120.11"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.permission", "ALLOW"),
				),
			},
			{
				Config: testAccInfobloxZoneAuthUpdateTwoTemplate(testFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxZoneAuthExists(testFQDN, testFQDNResourceName),
					resource.TestCheckResourceAttr(testFQDNResourceName, "fqdn", testFQDN),
					resource.TestCheckResourceAttr(testFQDNResourceName, "comment", "Updated a zone"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "zone_format", "FORWARD"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "view", "default"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "prefix", "128-189"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_default_ttl", "7200"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_negative_ttl", "90"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_refresh", "1800"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_retry", "150"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "soa_expire", "888888"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "disable", "true"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_enable", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "dns_integrity_member", "nonprdibxdns01.bskyb.com"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "locked", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0._struct", "tsigac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key", "0jnu3SdsMvzzlmTDPYRceA=="),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key_alg", "HMAC-MD5"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.tsig_key_name", "acc-test.key"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.0.use_tsig_key_name", "false"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1._struct", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.address", "192.168.120.10"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.1.permission", "DENY"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2._struct", "addressac"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.address", "192.168.120.11"),
					resource.TestCheckResourceAttr(testFQDNResourceName, "allow_update.2.permission", "ALLOW"),
				),
			},
		},
	})
}

func testAccInfobloxZoneAuthExists(testFQDN, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.ZONEAUTHObj, "fqdn", testFQDN)
	}
}

func testAccInfobloxZoneAuthNoFQDNTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
comment = "Updated a zone"
ns_group="Sky OTT Default"
}
`)
}

func testAccInfobloxZoneAuthEmptyTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
ns_group="Sky OTT Default"
}
`)
}

func testAccInfobloxZoneAuthTooLongCommentTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
ns_group="Sky OTT Default"
fqdn = "%s"
comment = "This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string.... This is a very long string...."
}
`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidZoneFormat(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
ns_group="Sky OTT Default"
fqdn = "%s"
comment = "Created a zone"
zone_format = "SOME_INVALID_ZONE"
}`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidSOATTL(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
ns_group="Sky OTT Default"
fqdn = "%s"
soa_default_ttl = -1
}`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidAllowUpdatePermission(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
ns_group="Sky OTT Default"
fqdn = "%s"
allow_update = [
{
  _struct = "addressac"
  address = "192.168.100.10"
  permission = "SOME_INVALID_PERMISSION"
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidAllowUpdateStruct(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
ns_group="Sky OTT Default"
fqdn = "%s"
allow_update = [
{
  _struct = "SOME_INVALID__struct"
  address = "192.168.100.10"
  permission = "ALLOW"
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthInvalidAllowUpdateTSIGAlgorithm(testFQDN string) string {
	// NOTE: there is a bug in the INFOBLOX API: the use_tsig_key_name should be
	// set to true here but is not returned back...
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
ns_group="Sky OTT Default"
fqdn = "%s"
allow_update = [
{
  _struct = "tsigac"
  tsig_key = "0jnu3SdsMvzzlmTDPYRceA=="
  tsig_key_alg = "SOME_INVALID_ALGORITHM"
  tsig_key_name = "acc-test.key"
  use_tsig_key_name = true
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthLeadingTrailingWhiteSpace(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
ns_group="Sky OTT Default"
fqdn = "%s"
allow_update = [
{
  _struct = "addressac"
  address = " 192.168.100.10"
  permission = "ALLOW"
},
{
  _struct = "tsigac"
  tsig_key = " 0jnu3SdsMvzzlmTDPYRceA== "
  tsig_key_alg = "HMAC-SHA256"
  tsig_key_name = " test.key "
  use_tsig_key_name = true
},
]}`, testFQDN)
}

func testAccInfobloxZoneAuthCreateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
    fqdn = "%s"
    comment = "Created a zone"
    zone_format = "FORWARD"
    view = "default"
    prefix = "128/16"
    soa_default_ttl = 3600
    soa_negative_ttl = 60
    soa_refresh = 1200
    soa_retry = 300
    soa_expire = 444444
    disable = false
    dns_integrity_enable = false
    dns_integrity_member = "nonprdibxdns01.bskyb.com"
    locked = true
    copy_xfer_to_notify = false
    use_copy_xfer_to_notify = false
    ns_group="Sky OTT Default"
    allow_update = [
        {
            _struct = "addressac"
            address = "192.168.100.10"
            permission = "ALLOW"
        },
        {
            _struct = "addressac"
            address = "192.168.101.10"
            permission = "ALLOW"
        },
        {
            _struct = "tsigac"
            tsig_key = "0jnu3SdsMvzzlmTDPYRceA=="
            tsig_key_alg = "HMAC-SHA256"
            tsig_key_name = "acc-test.key"
            use_tsig_key_name = true
        },
    ]
}`, testFQDN)
}

func testAccInfobloxZoneAuthUpdateTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
  fqdn = "%s"
  comment = "Updated a zone"
  zone_format = "FORWARD"
  view = "default"
  prefix = "128-189"
  soa_default_ttl = 7200
  soa_negative_ttl = 90
  soa_refresh = 1800
  soa_retry = 150
  soa_expire = 888888
  disable = true
  dns_integrity_enable = false
  dns_integrity_member = "nonprdibxdns01.bskyb.com"
  locked = false
  copy_xfer_to_notify = true
  use_copy_xfer_to_notify = true
  use_external_primary = false
  ns_group="Sky OTT Default"
  allow_transfer = [
        {
          _struct = "addressac"
          address = "192.168.234.11"
          permission = "DENY"
        },
        {
          _struct = "tsigac"
          tsig_key = "0jnu3SdsMvzzlmTDPYRceA=="
          tsig_key_alg = "HMAC-SHA256"
          tsig_key_name = "acc-test.key"
          use_tsig_key_name = true
        },
        {
          _struct = "addressac"
          address = "192.168.101.11"
          permission = "ALLOW"
        },
        {
          _struct = "addressac"
          address = "192.168.111.10"
          permission = "ALLOW"
        },
  ]
  allow_update = [
    {
      _struct = "tsigac"
      tsig_key = "0jnu3SdsMvzzlmTDPYRceA=="
      tsig_key_alg = "HMAC-MD5"
      tsig_key_name = "acc-test.key"
      use_tsig_key_name = false
    },
    {
      _struct = "addressac"
      address = "192.168.120.10"
      permission = "DENY"
    },
    {
      _struct = "addressac"
      address = "192.168.120.11"
      permission = "ALLOW"
    },
  ]}`, testFQDN)
}

func testAccInfobloxZoneAuthUpdateTwoTemplate(testFQDN string) string {
	return fmt.Sprintf(`
resource "infoblox_zone_auth" "acctest" {
fqdn = "%s"
comment = "Updated a zone"
zone_format = "FORWARD"
view = "default"
prefix = "128-189"
soa_default_ttl = 7200
soa_negative_ttl = 90
soa_refresh = 1800
soa_retry = 150
soa_expire = 888888
disable = true
dns_integrity_enable = false
dns_integrity_member = "nonprdibxdns01.bskyb.com"
locked = false
ns_group="Sky OTT Default"
allow_update = [
{
  _struct = "tsigac"
  tsig_key = "0jnu3SdsMvzzlmTDPYRceA=="
  tsig_key_alg = "HMAC-MD5"
  tsig_key_name = "acc-test.key"
  use_tsig_key_name = false
},
{
  _struct = "addressac"
  address = "192.168.120.10"
  permission = "DENY"
},
{
  _struct = "addressac"
  address = "192.168.120.11"
  permission = "ALLOW"
},]}`, testFQDN)
}
