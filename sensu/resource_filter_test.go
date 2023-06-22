package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceFilter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceFilter_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "name", "filter_1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "action", "allow"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "expressions.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "when.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "when.133776838.day", "monday"),
				),
			},
			resource.TestStep{
				Config: testAccResourceFilter_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "name", "filter_1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "action", "deny"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "expressions.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "when.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "when.3883494025.day", "tuesday"),
				),
			},
		},
	})
}

func TestAccResourceFilter_runtimeAssets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceFilter_runtimeAssets_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "name", "filter_1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "runtime_assets.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccResourceFilter_runtimeAssets_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "name", "filter_1"),
					resource.TestCheckResourceAttr(
						"sensu_filter.filter_1", "runtime_assets.#", "0"),
				),
			},
		},
	})
}

const testAccResourceFilter_basic = `
  resource "sensu_filter" "filter_1" {
    name = "filter_1"
    action = "allow"
    expressions = [
      "event.Check.Team == 'ops'",
    ]

    when {
      day = "monday"
      begin = "09:00AM"
      end = "05:00PM"
    }
  }
`

const testAccResourceFilter_update = `
  resource "sensu_filter" "filter_1" {
    name = "filter_1"
    action = "deny"
    expressions = [
      "event.Check.Team == 'ops'",
      "event.Check.Team == 'dev'",
    ]

    when {
      day = "tuesday"
      begin = "09:00AM"
      end = "05:00PM"
    }
  }
`

var testAccResourceFilter_runtimeAssets_1 = fmt.Sprintf(`
  %s

  resource "sensu_filter" "filter_1" {
    name = "filter_1"
    action = "allow"
    expressions = [
      "event.Check.Team == 'ops'",
    ]

    runtime_assets = [
      "${sensu_asset.asset_1.name}"
    ]
  }
`, testAccResourceAsset_basic)

var testAccResourceFilter_runtimeAssets_2 = fmt.Sprintf(`
  %s

  resource "sensu_filter" "filter_1" {
    name = "filter_1"
    action = "allow"
    expressions = [
      "event.Check.Team == 'ops'",
    ]
  }
`, testAccResourceAsset_basic)
