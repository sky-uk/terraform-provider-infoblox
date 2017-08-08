package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"net/http"
	"strconv"
	"testing"
)

func TestAccResourceTXTRecord(t *testing.T) {
	recordName := "txtrecordcreated-" + strconv.Itoa(acctest.RandInt()) + ".testzone.slupaas.bskyb.com"
	resourceName := "infoblox_txtrecord.acctest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccResourceTXTRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceTXTRecordCreateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceTXTRecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "text", "record text"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "9000"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
			{
				Config: testAccResourceTXTRecordUpdateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceTXTRecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "text", "record text updated"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "1000"),
					resource.TestCheckResourceAttr(resourceName, "use_ttl", "true"),
				),
			},
		},
	})
}

func testAccResourceTXTRecordExists(recordName, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nResource not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox TXT record resource %s ID not set", resourceName)
		}
		ref := rs.Primary.ID
		infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		fields := []string{"name", "view", "zone", "ttl", "use_ttl", "text", "comment"}
		recAPI := records.NewGetTXTRecord(ref, fields)
		err := infobloxClient.Do(recAPI)
		if err != nil {
			return fmt.Errorf("Error getting the TXT record: %q", err.Error())
		}
		if recAPI.StatusCode() == http.StatusOK {
			return nil
		}
		return fmt.Errorf("Could not find %s", resourceName)
	}
}

func testAccResourceTXTRecordDestroy(state *terraform.State) error {
	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_txtrecord" {
			continue
		}
		ref := rs.Primary.ID
		api := records.NewGetTXTRecord(ref, nil)
		err := infobloxClient.Do(api)
		if err != nil {
			return err
		}

		if api.StatusCode() == http.StatusOK {
			return fmt.Errorf("TXT record still exists, ref: %v", ref)
		}
		return nil
	}
	//return errors.New("infoblox_txtrecord resorce not found")
	return nil
}

func testAccResourceTXTRecordCreateTemplate(txtrecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_txtrecord" "acctest"{
	name = "%s"
	ttl = 9000
	use_ttl = true
	text = "record text"
	}`, txtrecordName)
}

func testAccResourceTXTRecordUpdateTemplate(txtrecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_txtrecord" "acctest"{
	name = "%s"
	ttl = 1000
	use_ttl = true
	text = "record text updated"
	}`, txtrecordName)
}
