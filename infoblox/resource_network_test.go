package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox"
	"github.com/sky-uk/skyinfoblox/api/network"
	"testing"
)

func TestAccResourceNetwork(t *testing.T) {
	networkAddr := "10.0.0.0/24"
	resourceName := "infoblox_network.net3"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccResourceNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceNetworkCreateTemplate(networkAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceNetworkExists(networkAddr, resourceName),
					resource.TestCheckResourceAttr(resourceName, "network", networkAddr),
					resource.TestCheckResourceAttr(resourceName, "comment", "a comment on a network"),
				),
			}, {
				Config: testAccResourceNetworkUpdateTemplate(networkAddr),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceNetworkExists(networkAddr, resourceName),
					resource.TestCheckResourceAttr(resourceName, "network", networkAddr),
					resource.TestCheckResourceAttr(resourceName, "comment", "another comment on a network"),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
		},
	})

}

func testAccResourceNetworkDestroy(state *terraform.State) error {
	infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_network" {
			continue
		}
		if res, ok := rs.Primary.Attributes["ref"]; ok && res != "" {
			return nil
		}
		fields := []string{"network", "options"}

		api := network.NewGetNetwork(rs.Primary.ID, fields)
		err := infobloxClient.Do(api)
		if err != nil {
			return nil
		}

		if api.GetResponse().Network == "arecordcreated.test-ovp.bskyb.com" {
			return fmt.Errorf("Network still exists: %+v", api.GetResponse())
		}

	}
	return nil
}

func testAccResourceNetworkExists(networkAddr, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {

		var fields []string

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox Network resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox Network resource %s ID not set", resourceName)
		}
		infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		getAllARec := network.NewGetAllNetworks(fields)
		err := infobloxClient.Do(getAllARec)
		if err != nil {
			return fmt.Errorf("Error getting the A record: %q", err.Error())
		}
		for _, x := range getAllARec.GetResponse() {
			if x.Network == networkAddr {
				return nil
			}
		}
		return fmt.Errorf("Could not find %s", networkAddr)

	}

}

func testAccResourceNetworkCreateTemplate(networkAddr string) string {
	return fmt.Sprintf(`
	resource "infoblox_network" "net3"{
	network = "%s"
	comment = "a comment on a network"
	}`, networkAddr)
}

func testAccResourceNetworkUpdateTemplate(networkAddr string) string {
	return fmt.Sprintf(`
	resource "infoblox_network" "net3"{
	network = "%s"
	comment = "another comment on a network"
	disable = true
	option {
            name = "routers",
            num = 3,
            useoption = true,
            value =  "10.0.0.1",
            vendorclass =  "DHCP"
           }
   	high_watermark = 90
        high_watermark_reset = 80
        low_watermark = 7
        low_watermark_reset = 11
        enabledhcpthresholds = false
        use_enabledhcpthresholds = false
        discovery_member = "testhost.devops.int.ovp.bskyb.com"
        enablediscovery = true
        use_enablediscovery = true
	}`, networkAddr)
}
