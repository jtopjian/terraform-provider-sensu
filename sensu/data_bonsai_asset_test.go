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
						"data.sensu_bonsai_asset.bonsai_asset_1", "build.#", "6"),
					resource.TestCheckResourceAttr(
						"data.sensu_bonsai_asset.bonsai_asset_1", "build.0.filters.0", "entity.system.os == 'linux'"),
				),
			},
		},
	})
}

func TestAccDataSourceBonsaiAsset_createAssetFromBonsai(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceBonsaiAsset_createAssetFromBonsai,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "name", "sensu-plugins-cpu-checks"),
				),
			},
		},
	})
}

func TestAccDataSourceBonsaiAsset_createAssetFromBonsaiMultipleBuilds(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceBonsaiAsset_createAssetFromBonsaiMultipleBuilds,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "name", "sensu-plugins-cpu-checks"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "build.#", "6"),
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

const testAccDataSourceBonsaiAsset_createAssetFromBonsai = `
  data "sensu_bonsai_asset" "bonsai_asset_1" {
    name = "sensu-plugins/sensu-plugins-cpu-checks"
    version = "4.1.0"
  }

  resource "sensu_asset" "asset_1" {
    name = data.sensu_bonsai_asset.bonsai_asset_1.annotations["io.sensu.bonsai.name"]
    sha512 = data.sensu_bonsai_asset.bonsai_asset_1.build.0.sha512
    url = data.sensu_bonsai_asset.bonsai_asset_1.build.0.url
    filters = data.sensu_bonsai_asset.bonsai_asset_1.build.0.filters
  }
`

const testAccDataSourceBonsaiAsset_createAssetFromBonsaiMultipleBuilds = `
  data "sensu_bonsai_asset" "bonsai_asset_1" {
    name = "sensu-plugins/sensu-plugins-cpu-checks"
    version = "4.1.0"
  }

  resource "sensu_asset" "asset_1" {
    name = data.sensu_bonsai_asset.bonsai_asset_1.annotations["io.sensu.bonsai.name"]

    dynamic "build" {
      for_each = data.sensu_bonsai_asset.bonsai_asset_1.build
      content {
        sha512 = build.value["sha512"]
        url = build.value["url"]
        filters = build.value["filters"]
        headers = build.value["headers"]
      }
    }
  }
`
