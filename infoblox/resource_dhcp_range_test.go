package infoblox

import (
	//	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	//	"github.com/sky-uk/skyinfoblox"
	//	"github.com/sky-uk/skyinfoblox/api/dhcp_range"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"testing"
)

func TestAccResourceDHCPRange(t *testing.T) {
	network := "10.0.0.0/24"
	//resourceName := "infoblox_network_net"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.RangeObj, "network", network)
		},
		Steps: []resource.TestStep{
		/* DO NOT UNCOMMENT THESE LINES UNLESS YOU KNOW WHAT YOU ARE DOING !!!
		{
			Config: testAccResourceDHCPRangeCreateTemplate(network),
			Check: resource.ComposeTestCheckFunc(
				testAccResourceDHCPRangeExists(network, resourceName),
				resource.TestCheckResourceAttr(resourceName, "network", network),
				resource.TestCheckResourceAttr(resourceName, "network_view", "default"),
				resource.TestCheckResourceAttr(resourceName, "start", "10.154.0.30"),
				resource.TestCheckResourceAttr(resourceName, "end", "10.154.0.40"),
				resource.TestCheckResourceAttr(resourceName, "server_association", "MEMBER"),
			),
		}, {
			Config: testAccResourceDHCPRangeUpdateTemplate(network),
			Check: resource.ComposeTestCheckFunc(
				testAccResourceDHCPRangeExists(network, resourceName),
				resource.TestCheckResourceAttr(resourceName, "network", network),
				resource.TestCheckResourceAttr(resourceName, "network_view", "new_view"),
				resource.TestCheckResourceAttr(resourceName, "start", "10.154.0.30"),
				resource.TestCheckResourceAttr(resourceName, "start", "10.154.0.50"),
				resource.TestCheckResourceAttr(resourceName, "server_association", "MEMBER"),
			),
		},*/
		},
	})

}

/*
func testAccResourceDHCPRangeDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
	for _, rs := range state.RootModule().Resources {
		if rs.Type != "infoblox_dhcp_range" {
			continue
		}
		if res, ok := rs.Primary.Attributes["ref"]; ok && res != "" {
			return nil
		}
		fields := []string{"end_addr", "start_addr", "network", "network_view", "member", "server_association_type"}

		api := dhcprange.NewGetDHCPRangeAPI(rs.Primary.ID, fields)
		err := infobloxClient.Do(api)
		if err != nil {
			return nil
		}

		if api.GetResponse().Network == "10.0.0.0/24" {
			return fmt.Errorf("DHCP Range still exists: %+v", api.GetResponse())
		}

	}
	return nil
}
*/
/*
func testAccResourceDHCPRangeExists(networkAddr, resourceName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		var fields []string

		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("\nInfoblox DHCP Range resource %s not found in resources", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("\nInfoblox DHCP Range resource %s ID not set", resourceName)
		}
		infobloxClient := testAccProvider.Meta().(*skyinfoblox.InfobloxClient)
		getRequest := dhcprange.NewGetDHCPRangeAPI(rs.Primary.ID, fields)
		err := infobloxClient.Do(getRequest)
		if err != nil {
			return fmt.Errorf("Error getting the DHCP Range: %q", err.Error())
		}
		if getRequest.GetResponse().Network == networkAddr {
			return nil
		}
		return fmt.Errorf("Could not find %s", networkAddr)
	}
}
*/

/*
func testAccResourceDHCPRangeCreateTemplate(network string) string {
	return fmt.Sprintf(`
	resource "infoblox_dhcp_range" "infoblox_network_net" {
		network = "%s"
		network_view = "default"
		start_addr = "10.154.0.30"
		end _addr= "10.154.0.40"
		member = {
			ipv4addr = "10.90.233.150"
			name  = "nonprdibxdns01.bskyb.com"
		}
		server_association_type = "MEMBER"
	}`, network)
}

func testAccResourceDHCPRangeUpdateTemplate(network string) string {
	return fmt.Sprintf(`
	resource "infoblox_dhcp_range" "infoblox_network_net" {
		network = "%s"
		network_view = "new_view"
		start_addr = "10.154.0.30"
		end_addr = "10.154.0.50"
		member = {
			ipv4addr = "10.74.233.150"
			name  = "nonprdibxdns02.bskyb.com"
		}
		server_association_type = "MEMBER"
	}`, network)
}
*/
