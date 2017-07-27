package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"testing"
)

func TestAccResourceSRVRecord(t *testing.T) {
	recordName := "srv-recordcreated.slupaas.bskyb.com"
	resourceName := "infoblox_srv_record.acctest"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccResourceSRVRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSRVRecordCreateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceSRVRecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "port", "8080"),
					resource.TestCheckResourceAttr(resourceName, "priority", "99"),
					resource.TestCheckResourceAttr(resourceName, "target", "craig4test.testzone.slupaas.bskyb.com"),
					resource.TestCheckResourceAttr(resourceName, "weight", "10"),
					resource.TestCheckResourceAttr(resourceName, "comment", "test test"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "900"),
				),
			},
			{
				Config: testAccResourceSRVRecordUpdateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceSRVRecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "port", "65"),
					resource.TestCheckResourceAttr(resourceName, "priority", "50"),
					resource.TestCheckResourceAttr(resourceName, "target", "craig4test.testzone.slupaas.bskyb.com"),
					resource.TestCheckResourceAttr(resourceName, "weight", "20"),
					resource.TestCheckResourceAttr(resourceName, "comment", "test test test test"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "4000"),
				),
			},
		},
	})
}

func testAccResourceSRVRecordExists(recordName, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox SRV record resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox SRV record resource %s ID not set", resourceName)
		}

		infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

		returnFields := []string{"name"}

		getAllSRVRecords := records.NewGetAllSRVRecords(returnFields)

		err := infobloxClient.Do(getAllSRVRecords)

		if err != nil {
			return fmt.Errorf("Error getting the SRV record: %q", err.Error())
		}
		for _, x := range getAllSRVRecords.GetResponse() {
			if x.Name == recordName {
				return nil
			}
		}
		return fmt.Errorf("Could not find %s", recordName)

	}
}

func testAccResourceSRVRecordDestroy(state *terraform.State) error {
	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
	returnFields := []string{"name"}

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_srv_record" {
			continue
		}
		if res, ok := rs.Primary.Attributes["ref"]; ok && res != "" {
			return nil
		}
		api := records.NewGetSRVRecord(rs.Primary.Attributes["res"], returnFields)
		err := infobloxClient.Do(api)
		if err != nil {
			return nil
		}

		if api.GetResponse().Name == "srv-recordcreated.slupaas.bskyb.com" {
			return fmt.Errorf("A record still exists: %+v", api.GetResponse())
		}

	}
	return nil
}

func testAccResourceSRVRecordCreateTemplate(srvRecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_srv_record" "acctest" {
   	name = "%s"
    	port = 8080
    	priority = 99
    	target = "craig4test.testzone.slupaas.bskyb.com"
    	weight = 10
    	comment = "test test"
    	ttl = 900
	}`, srvRecordName)
}

func testAccResourceSRVRecordUpdateTemplate(srvRecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_srv_record" "acctest" {
   	name = "%s"
    	port = 65
    	priority = 50
    	target = "craig4test.testzone.slupaas.bskyb.com"
    	weight = 20
    	comment = "test test test test"
    	ttl = 4000
	}`, srvRecordName)
}
