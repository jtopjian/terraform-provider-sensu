package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceEnvironment_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceEnvironment_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_environment.environment_1", "name", "environment_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_environment.environment_1", "description", "an environment"),
					resource.TestCheckResourceAttr(
						"data.sensu_environment.environment_1", "organization", "default"),
				),
			},
		},
	})
}

var testAccDataSourceEnvironment_basic = fmt.Sprintf(`
  %s

  data "sensu_environment" "environment_1" {
    name = "${sensu_environment.environment_1.name}"
  }
`, testAccResourceEnvironment_basic)
