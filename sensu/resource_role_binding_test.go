package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceRoleBinding_basic(t *testing.T) {
	username := acctest.RandomWithPrefix("sensu-acctest")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceRoleBinding_basic(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_role_binding.role_binding_1", "name", "role_binding_1"),
					resource.TestCheckResourceAttr(
						"sensu_role_binding.role_binding_1", "binding_type", "role"),
					resource.TestCheckResourceAttr(
						"sensu_role_binding.role_binding_1", "users.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_role_binding.role_binding_1", "groups.#", "0"),
				),
			},
		},
	})
}

func testAccResourceRoleBinding_basic(username string) string {
	userResource := testAccResourceUser_basic(username)

	return fmt.Sprintf(`
		%s

		%s

		resource "sensu_role_binding" "role_binding_1" {
			name = "role_binding_1"

			binding_type = "role"
			role = "${sensu_role.role_1.name}"

			users = ["${sensu_user.user_1.name}"]
		}
	`, testAccResourceRole_basic, userResource)
}
