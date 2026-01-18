package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceSecret_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceSecret_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_secret.secret_1", "name", "secret_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_secret.secret_1", "id", "SENSU_TEST_SECRET"),
					resource.TestCheckResourceAttr(
						"data.sensu_secret.secret_1", "secrets_provider", "env"),
				),
			},
		},
	})
}

const testAccDataSourceSecret_basic = `
  resource "sensu_secret" "secret_1" {
    name = "secret_1"
    id = "SENSU_TEST_SECRET"
    secrets_provider = "env"
  }

  data "sensu_secret" "secret_1" {
    name = "${sensu_secret.secret_1.name}"
  }
`
