package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"testing"
)

func TestAccResourceSRVRecord(t *testing.T) {

	randInt := acctest.RandInt()
	recordName := fmt.Sprintf("srv-recordcreated-%d.slupaas.bskyb.com", randInt)
	resourceName := "infoblox_srv_record.acctest"

	fmt.Printf("\n\nAcc Test record name is %s\n\n", recordName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.RecordSRVObj, "name", recordName)
		},
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
		return TestAccCheckExists(model.RecordSRVObj, "name", recordName)
	}
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
