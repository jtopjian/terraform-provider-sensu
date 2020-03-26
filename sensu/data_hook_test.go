package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceHook_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceHook_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_hook.hook_1", "command", "/bin/foo"),
					resource.TestCheckResourceAttr(
						"data.sensu_hook.hook_1", "timeout", "60"),
				),
			},
		},
	})
}

var testAccDataSourceHook_basic = fmt.Sprintf(`
  %s

  data "sensu_hook" "hook_1" {
    name = "${sensu_hook.hook_1.name}"
  }
`, testAccResourceHook_basic)
