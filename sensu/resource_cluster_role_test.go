package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceClusterRole_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceClusterRole_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "name", "cluster_role_1"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.0.verbs.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.0.verbs.0", "*"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.1.verbs.1", "list"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.1.resources.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.0.resources.0", "checks"),
					resource.TestCheckResourceAttr(
						"sensu_cluster_role.cluster_role_1", "rule.1.resources.1", "filters"),
				),
			},
		},
	})
}

const testAccResourceClusterRole_basic = `
  resource "sensu_cluster_role" "cluster_role_1" {
    name = "cluster_role_1"

    rule {
      verbs = ["*"]
      resources = ["checks"]
    }

    rule {
      verbs = ["get", "list"]
      resources = ["assets", "filters"]
    }

  }
`
