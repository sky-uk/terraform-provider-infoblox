package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
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
			return TestAccCheckDestroy(model.RecordAObj, "name", recordName)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceARecordCreateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceARecordExists("name", recordName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "9000"),
				),
			},
			{
				Config: testAccResourceARecordUpdateTemplate(recordName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceARecordExists("name", recordName),
					resource.TestCheckResourceAttr(resourceName, "name", recordName),
					resource.TestCheckResourceAttr(resourceName, "ipv4addr", "10.0.0.10"),
					resource.TestCheckResourceAttr(resourceName, "ttl", "900"),
				),
			},
		},
	})
}

func testAccResourceARecordExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.RecordAObj, key, value)
	}
}

func testAccResourceARecordCreateTemplate(arecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_arecord" "acctest"{
	name = "%s"
	ipv4addr = "10.0.0.10"
	ttl = 9000
	}`, arecordName)
}

func testAccResourceARecordUpdateTemplate(arecordName string) string {
	return fmt.Sprintf(`
	resource "infoblox_arecord" "acctest"{
	name = "%s"
	ipv4addr = "10.0.0.10"
	ttl = 900
    use_ttl = false
	}`, arecordName)
}
