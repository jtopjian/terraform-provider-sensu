package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
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
						"sensu_role.role_1", "rule.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "rule.0.type", "checks"),
				),
			},
			resource.TestStep{
				Config: testAccResourceRole_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "name", "role_1"),
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "rule.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "rule.1.type", "*"),
				),
			},
		},
	})
}

const testAccResourceRole_basic = `
  resource "sensu_role" "role_1" {
    name = "role_1"

    rule {
      type = "checks"
      environment = "*"
      organization = "*"
      permissions = ["read"]
    }
  }
`

const testAccResourceRole_update = `
  resource "sensu_role" "role_1" {
    name = "role_1"

    rule {
      type = "checks"
      environment = "*"
      organization = "*"
      permissions = ["read"]
    }

    rule {
      type = "*"
      environment = "*"
      organization = "*"
      permissions = ["create", "read"]
    }
  }
`
