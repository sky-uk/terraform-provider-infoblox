package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"testing"
)

func TestAccResourceARecord(t *testing.T) {
	recordName := "arecordcreated.test-ovp.bskyb.com"
	resourceName := "infoblox_arecord.acctest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccResourceARecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceARecordCreateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceARecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "address", "1.2.3.4"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "9000"),
				),
			},
			{
				Config: testAccResourceARecordUpdateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceARecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "address", "1.2.3.4"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "900"),
				),
			},
		},
	})
}

func testAccResourceARecordDestroy(state *terraform.State) error {
	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_arecord" {
			continue
		}
		if res, ok := rs.Primary.Attributes["ref"]; ok && res != "" {
			return nil
		}
		fields := []string{"name", "ipv4addr", "ttl"}
		api := records.NewGetARecord(rs.Primary.Attributes["res"], fields)
		err := infobloxClient.Do(api)
		if err != nil {
			return nil
		}

		if api.GetResponse().Name == "arecordcreated.test-ovp.bskyb.com" {
			return fmt.Errorf("A record still exists", api.GetResponse())
		}

	}
	return nil
}

func testAccResourceARecordExists(recordName, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox A record resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox A record resource %s ID not set", resourceName)
		}
		infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		fields := []string{"name", "ipv4addr", "ttl"}
		getAllARec := records.NewGetAllARecords(fields)
		err := infobloxClient.Do(getAllARec)
		if err != nil {
			return fmt.Errorf("Error getting the A record", err)
		}
		for _, x := range getAllARec.GetResponse() {
			if x.Name == recordName {
				return nil
			}
		}
		return fmt.Errorf("Could not find %s", recordName)

	}

}

func testAccResourceARecordCreateTemplate(arecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_arecord" "acctest"{
	name = "%s"
	address = "1.2.3.4"
	ttl = 9000
	}`, arecordName)
}

func testAccResourceARecordUpdateTemplate(arecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_arecord" "acctest"{
	name = "%s"
	address = "1.2.3.4"
	ttl = 900
	}`, arecordName)
}
