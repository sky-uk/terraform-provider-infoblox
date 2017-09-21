package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"strconv"
	"testing"
)

func TestAccResourceNetwork(t *testing.T) {
	networkAddr := "10.0." + strconv.Itoa(acctest.RandIntRange(0, 255)) + ".0/24"
	resourceName := "infoblox_network.net3"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.NetworkObj, "network", networkAddr)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceNetworkCreateTemplate(networkAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceNetworkExists("network", networkAddr),
					resource.TestCheckResourceAttr(resourceName, "network", networkAddr),
					resource.TestCheckResourceAttr(resourceName, "comment", "a comment on a network"),
				),
			}, {
				Config: testAccResourceNetworkUpdateTemplate(networkAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceNetworkExists("network", networkAddr),
					resource.TestCheckResourceAttr(resourceName, "network", networkAddr),
					resource.TestCheckResourceAttr(resourceName, "comment", "another comment on a network"),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
		},
	})

}

func testAccResourceNetworkExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.NetworkObj, key, value)
	}
}

func testAccResourceNetworkCreateTemplate(networkAddr string) string {
	return fmt.Sprintf(`
	resource "infoblox_network" "net3"{
	network = "%s"
	comment = "a comment on a network"
    disable = true
	}`, networkAddr)
}

func testAccResourceNetworkUpdateTemplate(networkAddr string) string {
	return fmt.Sprintf(`
	resource "infoblox_network" "net3"{
	network = "%s"
	comment = "another comment on a network"
	disable = true
   	high_water_mark = 90
    high_water_mark_reset = 80
    low_water_mark = 7
    low_water_mark_reset = 11
    enable_dhcp_thresholds = false
    use_enable_dhcp_thresholds = false
    //discovery_member = "slunonprdirep01.bskyb.com"
    //enablediscovery = true
    use_enable_discovery = true
	}`, networkAddr)
}
