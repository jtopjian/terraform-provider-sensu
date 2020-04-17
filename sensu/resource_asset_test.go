package sensu

import (
    "testing"

    "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceAsset_basic(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:  func() { testAccPreCheck(t) },
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            resource.TestStep{
                Config: testAccResourceAsset_basic,
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "name", "asset_1"),
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "sha512", "4f926bf4328fbad2b9cac873d117f771914f4b837c9c85584c38ccf55a3ef3c2e8d154812246e5dda4a87450576b2c58ad9ab40c9e2edc31b288d066b195b21b"),
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "url", "http://example.com/asset/example.tar.gz"),
                ),
            },
            resource.TestStep{
                Config: testAccResourceAsset_update,
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "name", "asset_1"),
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "sha512", "5f926bf4328fbad2b9cac873d117f771914f4b837c9c85584c38ccf55a3ef3c2e8d154812246e5dda4a87450576b2c58ad9ab40c9e2edc31b288d066b195b21b"),
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "url", "http://example.com/asset/example2.tar.gz"),
                ),
            },
        },
    })
}

func TestAccResourceCheck_headers(t *testing.T) {
    resource.Test(t, resource.TestCase{
        PreCheck:  func() { testAccPreCheck(t) },
        Providers: testAccProviders,
        Steps: []resource.TestStep{
            resource.TestStep{
                Config: testAccResourceCheck_headers_1,
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "headers.header1", "test1"),
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "headers.header2", "test2"),
                ),
            },
            resource.TestStep{
                Config: testAccResourceCheck_headers_2,
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "headers.header1", "test1"),
                    resource.TestCheckResourceAttr(
                        "sensu_asset.asset_1", "headers.header2", "test3"),
                ),
            },
            resource.TestStep{
                Config: testAccResourceCheck_headers_3,
                Check: resource.ComposeTestCheckFunc(
                    resource.TestCheckNoResourceAttr("sensu_asset.asset_1", "headers"),
                ),
            },
        },
    })
}

const testAccResourceAsset_basic = `
  resource "sensu_asset" "asset_1" {
    name = "asset_1"
    sha512 = "4f926bf4328fbad2b9cac873d117f771914f4b837c9c85584c38ccf55a3ef3c2e8d154812246e5dda4a87450576b2c58ad9ab40c9e2edc31b288d066b195b21b"
    url = "http://example.com/asset/example.tar.gz"

    filters = [
      "System.OS=='linux'",
      "System.Arch=='amd64'",
    ]
  }
`

const testAccResourceAsset_update = `
  resource "sensu_asset" "asset_1" {
    name = "asset_1"
    sha512 = "5f926bf4328fbad2b9cac873d117f771914f4b837c9c85584c38ccf55a3ef3c2e8d154812246e5dda4a87450576b2c58ad9ab40c9e2edc31b288d066b195b21b"
    url = "http://example.com/asset/example2.tar.gz"

    filters = [
      "System.OS=='linux'",
      "System.Arch=='amd64'",
    ]
  }
`

const testAccResourceCheck_headers_1 = `
  resource "sensu_asset" "asset_1" {
    name = "asset_1"
    sha512 = "5f926bf4328fbad2b9cac873d117f771914f4b837c9c85584c38ccf55a3ef3c2e8d154812246e5dda4a87450576b2c58ad9ab40c9e2edc31b288d066b195b21b"
    url = "http://example.com/asset/example2.tar.gz"

    filters = [
      "System.OS=='linux'",
      "System.Arch=='amd64'",
    ]
    headers = {
        header1 = "test1"
        header2 = "test2"
    }
  }
`

const testAccResourceCheck_headers_2 = `
  resource "sensu_asset" "asset_1" {
    name = "asset_1"
    sha512 = "5f926bf4328fbad2b9cac873d117f771914f4b837c9c85584c38ccf55a3ef3c2e8d154812246e5dda4a87450576b2c58ad9ab40c9e2edc31b288d066b195b21b"
    url = "http://example.com/asset/example2.tar.gz"

    filters = [
      "System.OS=='linux'",
      "System.Arch=='amd64'",
    ]
    headers = {
        header1 = "test1"
        header2 = "test3"
    }
  }
`

const testAccResourceCheck_headers_3 = `
  resource "sensu_asset" "asset_1" {
    name = "asset_1"
    sha512 = "5f926bf4328fbad2b9cac873d117f771914f4b837c9c85584c38ccf55a3ef3c2e8d154812246e5dda4a87450576b2c58ad9ab40c9e2edc31b288d066b195b21b"
    url = "http://example.com/asset/example2.tar.gz"

    filters = [
      "System.OS=='linux'",
      "System.Arch=='amd64'",
    ]
  }
`