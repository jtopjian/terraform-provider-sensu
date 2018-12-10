package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceRole_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceRole_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_role.role_1", "name", "role_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_role.role_1", "rule.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "rule.0.verbs.0", "*"),
					resource.TestCheckResourceAttr(
						"sensu_role.role_1", "rule.1.verbs.1", "list"),
				),
			},
		},
	})
}

var testAccDataSourceRole_basic = fmt.Sprintf(`
  %s

  data "sensu_role" "role_1" {
    name = "${sensu_role.role_1.name}"
  }
`, testAccResourceRole_basic)
