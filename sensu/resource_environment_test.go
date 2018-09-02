package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceEnvironment_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceEnvironment_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_environment.environment_1", "name", "environment_1"),
					resource.TestCheckResourceAttr(
						"sensu_environment.environment_1", "description", "an environment"),
				),
			},
			resource.TestStep{
				Config: testAccResourceEnvironment_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_environment.environment_1", "name", "environment_1"),
					resource.TestCheckResourceAttr(
						"sensu_environment.environment_1", "description", "some environment"),
				),
			},
		},
	})
}

const testAccResourceEnvironment_basic = `
  resource "sensu_environment" "environment_1" {
    name = "environment_1"
    description = "an environment"
  }
`

const testAccResourceEnvironment_update = `
  resource "sensu_environment" "environment_1" {
    name = "environment_1"
    description = "some environment"
  }
`
