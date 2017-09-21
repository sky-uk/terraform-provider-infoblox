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

func TestAccInfobloxDNSViewBasic(t *testing.T) {

	randomInt := acctest.RandInt()
	dnsViewName := fmt.Sprintf("acctest-infoblox-dns-view-%d", randomInt)
	updateDNSViewName := fmt.Sprintf("%s-updated", dnsViewName)
	dnsViewResource := "infoblox_dns_view.acctest"

	fmt.Printf("\n\nAcceptance Test DNS View is %s\n\n", dnsViewName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.ViewObj, "name", dnsViewName)
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccInfobloxDNSViewNoNameTemplate(),
				ExpectError: regexp.MustCompile(`required field is not set`),
			},
			{
				Config:      testAccInfobloxDNSViewCommentTooLongTemplate(),
				ExpectError: regexp.MustCompile(`Max length is 64`),
			},
			{
				Config: testAccInfobloxDNSViewCreateTemplate(dnsViewName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxDNSViewCheckExists("name", dnsViewName),
					resource.TestCheckResourceAttr(dnsViewResource, "name", dnsViewName),
					resource.TestCheckResourceAttr(dnsViewResource, "comment", "Infoblox Terraform Acceptance test"),
				),
			},
			{
				Config: testAccInfobloxDNSViewUpdateTemplate(updateDNSViewName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxDNSViewCheckExists("name", updateDNSViewName),
					resource.TestCheckResourceAttr(dnsViewResource, "name", updateDNSViewName),
					resource.TestCheckResourceAttr(dnsViewResource, "comment", "Infoblox Terraform Acceptance test - updated"),
				),
			},
		},
	})
}

func testAccInfobloxDNSViewCheckExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.ViewObj, key, value)
	}
}

func testAccInfobloxDNSViewNoNameTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_dns_view" "acctest" {
comment = "Infoblox Terraform Acceptance test"
}
`)
}

func testAccInfobloxDNSViewCommentTooLongTemplate() string {
	return fmt.Sprintf(`
resource "infoblox_dns_view" "acctest" {
comment = "O160axAAusI8YJOOe7G2BJMHoBKoCP5fPOri3ZtgbpesYsPftjnv1gV10HRSksk4tg"
}
`)
}

func testAccInfobloxDNSViewCreateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_dns_view" "acctest" {
name = "%s"
comment = "Infoblox Terraform Acceptance test"
}
`, name)
}

func testAccInfobloxDNSViewUpdateTemplate(name string) string {
	return fmt.Sprintf(`
resource "infoblox_dns_view" "acctest" {
name = "%s"
comment = "Infoblox Terraform Acceptance test - updated"
}
`, name)
}
