package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/records"
	"testing"
)

func TestAccResourceARecord(t *testing.T) {

	randInt := acctest.RandInt()
	recordName := fmt.Sprintf("a-record-test-%d.slupaas.bskyb.com", randInt)
	resourceName := "infoblox_arecord.acctest"

	fmt.Printf("\n\nAcc Test record name is %s\n\n", recordName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return testAccResourceARecordDestroy(state, recordName)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceARecordCreateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceARecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "9000"),
				),
			},
			{
				Config: testAccResourceARecordUpdateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceARecordExists(recordName, resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "900"),
				),
			},
		},
	})
}

func testAccResourceARecordDestroy(state *terraform.State, recordName string) error {

	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_arecord" {
			continue
		}
		if res, ok := rs.Primary.Attributes["ref"]; ok && res == "" {
			return nil
		}
		fields := []string{"name", "ipv4addr", "ttl"}
		api := records.NewGetARecord(rs.Primary.Attributes["id"], fields)
		err := infobloxClient.Do(api)
		if err != nil {
			return err
		}
		response := *api.ResponseObject().(*string)

		if response == recordName {
			return fmt.Errorf("A record still exists: %+v", api.GetResponse())
		}

	}
	return nil
}

func testAccResourceARecordExists(recordName, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox A record resource %s not found in resources: ", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox A record resource %s ID not set", resourceName)
		}
		infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		fields := []string{"name", "ipv4addr", "ttl"}
		getAllARec := records.NewGetAllARecords(fields)
		err := infobloxClient.Do(getAllARec)
		if err != nil {
			return fmt.Errorf("Error getting the A record: %+v", err)
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
	address = "10.0.0.10"
	ttl = 9000
	}`, arecordName)
}

func testAccResourceARecordUpdateTemplate(arecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_arecord" "acctest"{
	name = "%s"
	address = "10.0.0.10"
	ttl = 900
	}`, arecordName)
}
