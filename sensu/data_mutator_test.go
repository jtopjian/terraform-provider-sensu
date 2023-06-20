package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceMutator_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceMutator_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_mutator.mutator_1", "name", "mutator_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_mutator.mutator_1", "command", "/bin/foo"),
					resource.TestCheckResourceAttr(
						"data.sensu_mutator.mutator_1", "env_vars.FOO", "bar"),
					resource.TestCheckResourceAttr(
						"data.sensu_mutator.mutator_1", "secrets.BAR", "foo"),
				),
			},
		},
	})
}

var testAccDataSourceMutator_basic = fmt.Sprintf(`
  %s

  data "sensu_mutator" "mutator_1" {
    name = "${sensu_mutator.mutator_1.name}"
  }
`, testAccResourceMutator_basic)
