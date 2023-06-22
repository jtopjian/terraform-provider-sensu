package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceSilenced_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceSilenced_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_silenced.silenced_1", "check", "check_1"),
					resource.TestCheckResourceAttr(
						"sensu_silenced.silenced_1", "subscription", "subscription_1"),
					resource.TestCheckResourceAttr(
						"sensu_silenced.silenced_1", "begin", "now"),
				),
			},
			resource.TestStep{
				Config: testAccResourceSilenced_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_silenced.silenced_1", "check", "check_1"),
					resource.TestCheckResourceAttr(
						"sensu_silenced.silenced_1", "subscription", "subscription_1"),
					resource.TestCheckResourceAttr(
						"sensu_silenced.silenced_1", "begin", "Jan 02 2020 3:04PM MST"),
				),
			},
		},
	})
}

const testAccResourceSilenced_basic = `
  resource "sensu_silenced" "silenced_1" {
    check = "check_1"
    subscription = "subscription_1"
  }
`

const testAccResourceSilenced_update = `
  resource "sensu_silenced" "silenced_1" {
    check = "check_1"
    subscription = "subscription_1"
    begin = "Jan 02 2020 3:04PM MST"
  }
`
