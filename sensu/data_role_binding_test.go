package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceRoleBinding_basic(t *testing.T) {
	username := acctest.RandomWithPrefix("sensu-acctest")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceRoleBinding_basic(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_role_binding.role_binding_1", "name", "role_binding_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_role_binding.role_binding_1", "binding_type", "role"),
					resource.TestCheckResourceAttr(
						"data.sensu_role_binding.role_binding_1", "users.#", "1"),
					resource.TestCheckResourceAttr(
						"data.sensu_role_binding.role_binding_1", "groups.#", "0"),
				),
			},
		},
	})
}

func testAccDataSourceRoleBinding_basic(username string) string {
	roleBindingResource := testAccResourceRoleBinding_basic(username)
	return fmt.Sprintf(`
		%s

		data "sensu_role_binding" "role_binding_1" {
			name = "${sensu_role_binding.role_binding_1.name}"
		}
	`, roleBindingResource)
}
