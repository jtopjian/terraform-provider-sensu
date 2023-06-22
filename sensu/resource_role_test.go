package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceRole_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceRole_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "name", "role_1"),
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "rule.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "rule.0.verbs.0", "*"),
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "rule.1.verbs.1", "list"),
				),
			},
		},
	})
}

const testAccResourceRole_basic = `
  resource "sensu_role" "role_1" {
    name = "role_1"

    rule {
      verbs = ["*"]
      resources = ["checks"]
    }

    rule {
      verbs = ["get", "list"]
      resources = ["assets"]
    }

  }
`
