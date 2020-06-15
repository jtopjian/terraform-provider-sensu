package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceClusterRoleBinding_basic(t *testing.T) {
	username := acctest.RandomWithPrefix("sensu-acctest")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceClusterRoleBinding_basic(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_cluster_role_binding.cluster_role_binding_1", "name", "cluster_role_binding_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_cluster_role_binding.cluster_role_binding_1", "users.#", "1"),
					resource.TestCheckResourceAttr(
						"data.sensu_cluster_role_binding.cluster_role_binding_1", "groups.#", "0"),
				),
			},
		},
	})
}

func testAccDataSourceClusterRoleBinding_basic(username string) string {
	clusterRoleBindingResource := testAccResourceClusterRoleBinding_basic(username)
	return fmt.Sprintf(`
		%s

		data "sensu_cluster_role_binding" "cluster_role_binding_1" {
			name = "${sensu_cluster_role_binding.cluster_role_binding_1.name}"
		}
	`, clusterRoleBindingResource)
}
