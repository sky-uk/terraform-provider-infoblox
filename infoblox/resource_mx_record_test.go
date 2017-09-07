package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records/mxrecord"
	"net/http"
	"strconv"
	"testing"
)

func TestResourceMXRecord(t *testing.T) {
	mailExchanger := fmt.Sprintf("mx-%s.bskyb.com", strconv.Itoa(acctest.RandIntRange(0, 100)))
	resourceName := "infoblox_mx_record.mxtest1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccResourceMXRecordDestroy(state, mailExchanger)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMXRecordCreateTemplate(mailExchanger),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceMXRecordExists(mailExchanger, resourceName),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mailExchanger),
					resource.TestCheckResourceAttr(resourceName, "name", "slupaas.bskyb.com"),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "preference", "120"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			{
				Config: testAccResourceMXRecordUpdateTemplate(mailExchanger),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceMXRecordExists(mailExchanger, resourceName),
					resource.TestCheckResourceAttr(resourceName, "mail_exchanger", mailExchanger),
					resource.TestCheckResourceAttr(resourceName, "name", "slupaas.bskyb.com"),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment edited for a disabled zone"),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
					resource.TestCheckResourceAttr(resourceName, "preference", "10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "4000"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
		},
	})

}

func testAccResourceMXRecordDestroy(state *terraform.State, mailExchanger string) error {
	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_mx_record" {
			continue
		}
		if res, ok := rs.Primary.Attributes["ref"]; ok && res != "" {
			return nil
		}

		api := mxrecord.NewGetAll()
		err := infobloxClient.Do(api)
		if err != nil {
			return nil
		}

		if api.StatusCode() != http.StatusOK {
			return fmt.Errorf("MXRecord still exists: %+v", *api.ResponseObject().(*string))
		}

		for _, x := range *api.ResponseObject().(*[]mxrecord.MxRecord) {
			if x.MailExchanger == mailExchanger {
				return nil
			}
		}

	}
	return nil
}

func testAccResourceMXRecordExists(mailExchanger, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox MX Record resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox MX Record resource %s ID not set", resourceName)
		}
		infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		getAllMX := mxrecord.NewGetAll()
		err := infobloxClient.Do(getAllMX)
		if err != nil {
			return fmt.Errorf("Error getting the Infoblox MX records: %q", err.Error())
		}
		for _, x := range *getAllMX.ResponseObject().(*[]mxrecord.MxRecord) {
			if x.MailExchanger == mailExchanger {
				return nil
			}
		}
		return fmt.Errorf("Could not find %s", resourceName)

	}
}

func testAccResourceMXRecordCreateTemplate(mailExchanger string) string {
	return fmt.Sprintf(`
	resource "infoblox_mx_record" "mxtest1" {
        	name = "slupaas.bskyb.com"
        	comment = "this is a comment"
        	disable = false
        	mail_exchanger = "%s"
        	preference = 120
        	ttl = 3600
        	use_ttl = true
        	view = "default"
	}
	`, mailExchanger)

}

func testAccResourceMXRecordUpdateTemplate(mailExchanger string) string {
	return fmt.Sprintf(`
	resource "infoblox_mx_record" "mxtest1" {
        	name = "slupaas.bskyb.com"
        	comment = "this is a comment edited for a disabled zone"
        	disable = true
        	mail_exchanger = "%s"
        	preference = 10
        	ttl = 4000
        	use_ttl = true
        	view = "default"
	}
	`, mailExchanger)

}
