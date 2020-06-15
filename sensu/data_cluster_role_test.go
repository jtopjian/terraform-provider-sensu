package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceClusterRole_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceClusterRole_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_cluster_role.cluster_role_1", "name", "cluster_role_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_cluster_role.cluster_role_1", "rule.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.0.verbs.0", "*"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.1.verbs.1", "list"),
				),
			},
		},
	})
}

var testAccDataSourceClusterRole_basic = fmt.Sprintf(`
  %s

  data "sensu_cluster_role" "cluster_role_1" {
    name = "${sensu_cluster_role.cluster_role_1.name}"
  }
`, testAccResourceClusterRole_basic)
