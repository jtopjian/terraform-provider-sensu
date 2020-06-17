package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceClusterRoleBinding_basic(t *testing.T) {
	username := acctest.RandomWithPrefix("sensu-acctest")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceClusterRoleBinding_basic(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_cluster_role_binding.cluster_role_binding_1", "name", "cluster_role_binding_1"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role_binding.cluster_role_binding_1", "cluster_role", "cluster_role_1"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role_binding.cluster_role_binding_1", "users.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role_binding.cluster_role_binding_1", "users.0", "${sensu_user.user_1.name}"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role_binding.cluster_role_binding_1", "groups.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role_binding.cluster_role_binding_1", "groups.1", "group_1"),
				),
			},
		},
	})
}

func testAccResourceClusterRoleBinding_basic(username string) string {
	userResource := testAccResourceUser_basic(username)

	return fmt.Sprintf(`
		%s

		%s

		resource "sensu_cluster_role_binding" "cluster_role_binding_1" {
			name = "cluster_role_binding_1"
			cluster_role = "${sensu_cluster_role.cluster_role_1.name}"
			users = ["${sensu_user.user_1.name}"]
            groups = ["group_0", "group_1"]
		}
	`, testAccResourceClusterRole_basic, userResource)
}
