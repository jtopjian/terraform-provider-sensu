package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceUser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceUser_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_user.user_1", "name", "admin"),
					resource.TestCheckResourceAttr(
						"data.sensu_user.user_1", "roles.#", "1"),
					resource.TestCheckResourceAttr(
						"data.sensu_user.user_1", "roles.0", "admin"),
				),
			},
		},
	})
}

const testAccDataSourceUser_basic = `
  data "sensu_user" "user_1" {
    name = "admin"
  }
`
