package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceFilter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceFilter_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_filter.filter_1", "name", "filter_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_filter.filter_1", "action", "allow"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "expressions.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "when.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "when.133776838.day", "monday"),
				),
			},
		},
	})
}

func TestAccDataSourceFilter_runtimeAssets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceFilter_runtimeAssets,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_filter.filter_1", "name", "filter_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_filter.filter_1", "runtime_assets.#", "1"),
				),
			},
		},
	})
}

var testAccDataSourceFilter_basic = fmt.Sprintf(`
  %s

  data "sensu_filter" "filter_1" {
    name = "${sensu_filter.filter_1.name}"
  }
`, testAccResourceFilter_basic)

var testAccDataSourceFilter_runtimeAssets = fmt.Sprintf(`
  %s

  data "sensu_filter" "filter_1" {
    name = "${sensu_filter.filter_1.name}"
  }
`, testAccResourceFilter_runtimeAssets_1)
