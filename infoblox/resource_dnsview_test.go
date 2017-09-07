package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/dnsview"
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
			return testAccInfobloxDNSViewCheckDestroy(state, dnsViewName)
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
					testAccInfobloxDNSViewCheckExists(dnsViewName, dnsViewResource),
					resource.TestCheckResourceAttr(dnsViewResource, "name", dnsViewName),
					resource.TestCheckResourceAttr(dnsViewResource, "comment", "Infoblox Terraform Acceptance test"),
				),
			},
			{
				Config: testAccInfobloxDNSViewUpdateTemplate(updateDNSViewName),
				Check: resource.ComposeTestCheckFunc(
					testAccInfobloxDNSViewCheckExists(updateDNSViewName, dnsViewResource),
					resource.TestCheckResourceAttr(dnsViewResource, "name", updateDNSViewName),
					resource.TestCheckResourceAttr(dnsViewResource, "comment", "Infoblox Terraform Acceptance test - updated"),
				),
			},
		},
	})
}

func testAccInfobloxDNSViewCheckDestroy(state *terraform.State, dnsViewName string) error {

	client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_dns_view" {
			continue
		}
		if id, ok := rs.Primary.Attributes["id"]; ok && id == "" {
			return nil
		}
		api := dnsview.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox - error occurred whilst retrieving a list of DNS views")
		}
		for _, dnsView := range *api.ResponseObject().(*[]dnsview.DNSView) {
			if dnsView.Name == dnsViewName {
				return fmt.Errorf("Infoblox DNS View %s still exists", dnsViewName)
			}
		}
	}
	return nil
}

func testAccInfobloxDNSViewCheckExists(dnsViewName, dnsViewResource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		rs, ok := state.RootModule().Resources[dnsViewResource]
		if !ok {
			return fmt.Errorf("\nInfoblox DNS View %s wasn't found in resources", dnsViewName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox DNS View ID not set for %s in resources", dnsViewName)
		}

		client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		api := dnsview.NewGetAll()
		err := client.Do(api)
		if err != nil {
			return fmt.Errorf("Infoblox DNS View - error whilst retrieving a list of DNS Views: %+v", err)
		}
		for _, dnsView := range *api.ResponseObject().(*[]dnsview.DNSView) {
			if dnsView.Name == dnsViewName {
				return nil
			}
		}
		return fmt.Errorf("Infoblox DNS View %s wasn't found on remote Infoblox server", dnsViewName)
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
