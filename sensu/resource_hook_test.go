package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceHook_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceHook_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "name", "hook_1"),
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "command", "/bin/foo"),
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "stdin", "false"),
				),
			},
			resource.TestStep{
				Config: testAccResourceHook_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "name", "hook_1"),
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "command", "/bin/foo"),
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "stdin", "true"),
				),
			},
		},
	})
}

const testAccResourceHook_basic = `
  resource "sensu_hook" "hook_1" {
    name = "hook_1"
    command = "/bin/foo"
  }
`

const testAccResourceHook_update = `
  resource "sensu_hook" "hook_1" {
    name = "hook_1"
    command = "/bin/foo"
    stdin = "true"
  }
`
