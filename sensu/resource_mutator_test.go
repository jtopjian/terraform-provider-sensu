package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceMutator_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceMutator_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "name", "mutator_1"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "command", "/bin/foo"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "env_vars.FOO", "bar"),
				),
			},
			resource.TestStep{
				Config: testAccResourceMutator_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "name", "mutator_1"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "command", "/bin/foo2"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "env_vars.FOO", "baz"),
				),
			},
		},
	})
}

const testAccResourceMutator_basic = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo"

    env_vars {
      "FOO" = "bar"
    }
  }
`

const testAccResourceMutator_update = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo2"

    env_vars {
      "FOO" = "baz"
    }
 }
`
