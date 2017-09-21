package infoblox

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/sky-uk/skyinfoblox/api/common/v261/model"
	"log"
	"strconv"
	"testing"
)

func TestAccAdminUserResource(t *testing.T) {
	randomInt := acctest.RandIntRange(1, 10000)
	recordUserName := fmt.Sprintf("testadminuser%d", randomInt)
	resourceName := "infoblox_admin_user.testadmin"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			return TestAccCheckDestroy(model.AdminuserObj, "name", recordUserName)
		},
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAdminUserNameCreateTemplate(recordUserName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceAdminUserExists("name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment"),
					resource.TestCheckResourceAttr(resourceName, "email", "exampleuser@domain.internal.com"),
					testCheckResourceStringSlice(resourceName, "admin_groups", []string{"APP-OVP-INFOBLOX-READONLY"}),
				),
			}, {
				Config: testAccResourceAdminUserNameUpdateTemplate(recordUserName),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceAdminUserExists("name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "name", recordUserName),
					resource.TestCheckResourceAttr(resourceName, "comment", "this is a comment updated"),
					resource.TestCheckResourceAttr(resourceName, "email", "user@domain.internal.com"),
					testCheckResourceStringSlice(resourceName, "admin_groups", []string{"APP-OVP-INFOBLOX-READONLY"}),
				),
			},
		},
	})

}

//testCheckResourceStringSlice - test that a slice has the same content
func testCheckResourceStringSlice(name, key string, expected []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		is, err := primaryInstanceState(s, name)
		if err != nil {
			return err
		}

		for k := range is.Attributes {
			log.Printf("attr[%s]:%+v\n", k, is.Attributes[k])
		}

		if sliceLength, found := is.Attributes[key+".#"]; found {
			i, err := strconv.Atoi(sliceLength)
			if err != nil {
				return err
			}
			if i != len(expected) {
				return fmt.Errorf("Attr %s expected lenght [ %d ] differs from actual [ %s ]", key, len(expected), sliceLength)
			}

		} else {
			return fmt.Errorf("Error getting attr %s length", key)
		}

		for index := 0; index < len(expected); index++ {
			indexKey := key + "." + strconv.Itoa(index)
			log.Println("Looking for key: ", indexKey)
			if actual, ok := is.Attributes[indexKey]; ok {
				if actual != expected[index] {
					return fmt.Errorf("%s[%d] expected this:\n%s\nActual:\n%s", key, index, expected[index], actual)
				}
			} else {
				return fmt.Errorf("Attribute %s not found in resource %s", key, name)
			}
		}
		return nil
	}
}

func testAccResourceAdminUserExists(key, value string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		return TestAccCheckExists(model.AdminuserObj, key, value)
	}
}

func testAccResourceAdminUserNameCreateTemplate(username string) string {
	return fmt.Sprintf(`
	resource "infoblox_admin_user" "testadmin" {
	name = "%s"
	comment = "this is a comment"
	email = "exampleuser@domain.internal.com"
	admin_groups = ["APP-OVP-INFOBLOX-READONLY"]
	password = "c0a6264f0f128d94cd8ef26652e7d9fd"}`, username)
}

func testAccResourceAdminUserNameUpdateTemplate(username string) string {
	return fmt.Sprintf(`
	resource "infoblox_admin_user" "testadmin" {
  		name = "%s"
		comment = "this is a comment updated"
		email = "user@domain.internal.com"
		admin_groups = ["APP-OVP-INFOBLOX-READONLY"]
		password = "c0a6264f0f128d94cd8ef26652e7d9fd"
	}
	`, username)
}

// primaryInstanceState returns the primary instance state for the given resource name.
func primaryInstanceState(s *terraform.State, name string) (*terraform.InstanceState, error) {
	ms := s.RootModule()
	rs, ok := ms.Resources[name]
	if !ok {
		return nil, fmt.Errorf("Not found: %s", name)
	}

	is := rs.Primary
	if is == nil {
		return nil, fmt.Errorf("No primary instance: %s", name)
	}

	return is, nil
}
