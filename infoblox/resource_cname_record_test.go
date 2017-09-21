package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"regexp"
	"testing"
)

func TestAccInfobloxCNAMEBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	cname := fmt.Sprintf("acctest-infoblox-cname-%d.slupaas.bskyb.com", randomInt)
	cnameUpdate := fmt.Sprintf("acctest-infoblox-cname-%d-renamed.slupaas.bskyb.com", randomInt)
	canonical := fmt.Sprintf("acctest-infoblox-canonical-%d.slupaas.bskyb.com", randomInt)
	canonicalUpdate := fmt.Sprintf("acctest-infoblox-canonical-update-%d.slupaas.bskyb.com", randomInt)
	cnameResourceName := "infoblox_cname_record.acctest"

	fmt.Printf("\n\nCNAME is %s\n\n", cname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.RecordCnameObj, "name", cname)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxCNAMENoNameCreateTemplate(canonical),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testAccInfobloxCNAMENegativeTTLCreateTemplate(cname, canonical),
				ExpectError: regexp.MustCompile(`can't be negative`),
			},
			{
				Config:      testAccInfobloxCNAMEEmptyTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config: testAccInfobloxCNAMECreateTemplate(cname, canonical),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxCNAMEExists("name", cname),
					resource.TestCheckResourceAttr(cnameResourceName, "name", cname),
					resource.TestCheckResourceAttr(cnameResourceName, "comment", "Terraform Acceptance Testing for CNAMEs"),
					resource.TestCheckResourceAttr(cnameResourceName, "canonical", canonical),
					resource.TestCheckResourceAttr(cnameResourceName, "view", "default"),
					resource.TestCheckResourceAttr(cnameResourceName, "ttl", "1202"),
				),
			},
			{
				Config: testAccInfobloxCNAMEUpdateTemplate(cnameUpdate, canonicalUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxCNAMEExists("name", cnameUpdate),
					resource.TestCheckResourceAttr(cnameResourceName, "name", cnameUpdate),
					resource.TestCheckResourceAttr(cnameResourceName, "comment", "Terraform Acceptance Testing for CNAMEs update test"),
					resource.TestCheckResourceAttr(cnameResourceName, "canonical", canonicalUpdate),
					resource.TestCheckResourceAttr(cnameResourceName, "view", "default"),
					resource.TestCheckResourceAttr(cnameResourceName, "ttl", "600"),
				),
			},
			{
				Config:      testAccInfobloxCNAMEBadViewUpdateTemplate(cname, canonical),
				ExpectError: regexp.MustCompile("Response status code: 404"),
			},
		},
	})
}

func testAccInfobloxCNAMEExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.RecordCnameObj, key, value)
	}
}

func testAccInfobloxCNAMECreateTemplate(cname, canonical string) string {
	return fmt.Sprintf(`
resource "infoblox_cname_record" "acctest" {
  name = "%s"
  comment = "Terraform Acceptance Testing for CNAMEs"
  canonical = "%s"
  view = "default"
  ttl = 1202
}
`, cname, canonical)
}

func testAccInfobloxCNAMEUpdateTemplate(cname, canonical string) string {
	return fmt.Sprintf(`
resource "infoblox_cname_record" "acctest" {
  name = "%s"
  comment = "Terraform Acceptance Testing for CNAMEs update test"
  canonical = "%s"
  view = "default"
  ttl = 600
}
`, cname, canonical)
}

func testAccInfobloxCNAMEBadViewUpdateTemplate(cname, canonical string) string {
	return fmt.Sprintf(`
resource "infoblox_cname_record" "acctest" {
  name = "%s"
  comment = "Terraform Acceptance Testing for CNAMEs update test"
  canonical = "%s"
  view = "A_VIEW_WHICH_DOESNT_EXIST"
  ttl = 600
}
`, cname, canonical)
}

func testAccInfobloxCNAMEEmptyTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_cname_record" "acctest" {
}
`)
}

func testAccInfobloxCNAMENegativeTTLCreateTemplate(cname, canonical string) string {
	return fmt.Sprintf(`
resource "infoblox_cname_record" "acctest" {
  name = "%s"
  comment = "Terraform Acceptance Testing for CNAMEs update test"
  canonical = "%s"
  view = "default"
  ttl = -1
}
`, cname, canonical)
}

func testAccInfobloxCNAMENoNameCreateTemplate(canonical string) string {
	return fmt.Sprintf(`
resource "infoblox_cname_record" "acctest" {
  comment = "Terraform Acceptance Testing for CNAMEs update test"
  canonical = "%s"
  view = "default"
  ttl = 5000
}
`, canonical)
}
