package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"testing"
)

func TestAccResourceTXTRecord(t *testing.T) {

	randInt := acctest.RandInt()
	recordName := fmt.Sprintf("txt-record-created-%d.slupaas.bskyb.com", randInt)
	resourceName := "infoblox_txtrecord.acctest"

	fmt.Printf("\n\nAcc Test record name is %s\n\n", recordName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.RecordTXTObj, "name", recordName)
		},
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
		return TestAccCheckExists(model.RecordTXTObj, "name", recordName)
	}
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
