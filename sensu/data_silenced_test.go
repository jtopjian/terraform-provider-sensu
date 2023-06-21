package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceSilenced_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceSilenced_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_silenced.silenced_1", "check", "check_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_silenced.silenced_1", "subscription", "subscription_1"),
				),
			},
		},
	})
}

var testAccDataSourceSilenced_basic = fmt.Sprintf(`
	%s

	data "sensu_silenced" "silenced_1" {
	check = "${sensu_silenced.silenced_1.check}"
		subscription = "${sensu_silenced.silenced_1.subscription}"
	}
`, testAccResourceSilenced_basic)
