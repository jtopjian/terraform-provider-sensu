package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceBonsaiAsset_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceBonsaiAsset_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_bonsai_asset.bonsai_asset_1", "annotations.io.sensu.bonsai.name", "sensu-plugins-cpu-checks"),
					resource.TestCheckResourceAttr(
						"data.sensu_bonsai_asset.bonsai_asset_1", "builds.#", "6"),
					resource.TestCheckResourceAttr(
						"data.sensu_bonsai_asset.bonsai_asset_1", "builds.0.filters.0", "entity.system.os == 'linux'"),
				),
			},
		},
	})
}

func TestAccDataSourceBonsaiAsset_createFromBonsai(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceBonsaiAsset_createFromBonsai,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "name", "sensu-plugins-cpu-checks"),
				),
			},
		},
	})
}

const testAccDataSourceBonsaiAsset_basic = `
	data "sensu_bonsai_asset" "bonsai_asset_1" {
		name = "sensu-plugins/sensu-plugins-cpu-checks"
		version = "4.1.0"
	}
`

const testAccDataSourceBonsaiAsset_createFromBonsai = `
	data "sensu_bonsai_asset" "bonsai_asset_1" {
		name = "sensu-plugins/sensu-plugins-cpu-checks"
		version = "4.1.0"
	}

	resource "sensu_asset" "asset_1" {
		name = data.sensu_bonsai_asset.bonsai_asset_1.annotations["io.sensu.bonsai.name"]
		sha512 = data.sensu_bonsai_asset.bonsai_asset_1.builds.0.sha512
		url = data.sensu_bonsai_asset.bonsai_asset_1.builds.0.url
		filters = data.sensu_bonsai_asset.bonsai_asset_1.builds.0.filters
	}
`
