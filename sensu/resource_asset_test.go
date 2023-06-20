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
				Config: testAccResourceAsset_headers_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "headers.header1", "test1"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "headers.header2", "test2"),
				),
			},
			resource.TestStep{
				Config: testAccResourceAsset_headers_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "headers.header1", "test1"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "headers.header2", "test3"),
				),
			},
			resource.TestStep{
				Config: testAccResourceAsset_headers_3,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sensu_asset.asset_1", "headers"),
				),
			},
		},
	})
}

func TestAccResourceAsset_createMultipleBuild(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceAsset_createMultipleBuild_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "annotations.io.sensu.bonsai.name", "sensu-ruby-runtime"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "build.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "build.0.filters.0", "entity.system.os == 'linux'"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "build.0.url", "https://example.com/asset_0.1.0.tar.gz"),
				),
			},
			resource.TestStep{
				Config: testAccResourceAsset_createMultipleBuild_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "annotations.io.sensu.bonsai.name", "sensu-ruby-runtime"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "build.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "build.0.filters.0", "entity.system.os == 'linux'"),
					resource.TestCheckResourceAttr(
						"sensu_asset.asset_1", "build.0.url", "https://example.com/asset_0.2.0_rhel.tar.gz"),
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

const testAccResourceAsset_headers_1 = `
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

const testAccResourceAsset_headers_2 = `
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

const testAccResourceAsset_headers_3 = `
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

const testAccResourceAsset_createMultipleBuild_1 = `
	resource "sensu_asset" "asset_1" {
		name = "sensu-ruby-runtime"
		annotations = {
			"io.sensu.bonsai.url" = "https://bonsai.sensu.io/assets/sensu/sensu-ruby-runtime"
			"io.sensu.bonsai.api_url" = "https://bonsai.sensu.io/api/v1/assets/sensu/sensu-ruby-runtime"
			"io.sensu.bonsai.tier" = "Community"
			"io.sensu.bonsai.version" = "0.0.10"
			"io.sensu.bonsai.namespace" = "sensu"
			"io.sensu.bonsai.name" = "sensu-ruby-runtime"
			"io.sensu.bonsai.tags" = ""
		}

		build {
			url = "https://example.com/asset_0.1.0.tar.gz"
			sha512 = "cbee19124b7007342ce37ff9dfd4a1dde03beb1e87e61ca2aef606a7ad3c9bd0bba4e53873c07afa5ac46b0861967a9224511b4504dadb1a5e8fb687e9495304"
			filters = [
				"entity.system.os == 'linux'",
				"entity.system.arch == 'amd64'",
				"entity.system.platform_family == 'rhel'",
				"parseInt(entity.system.platform_version.split('.')[0]) == 6",
			]
									headers = {
													"Authorization" = "Bearer changeme"
									}
		}

		build {
			url = "https://assets.bonsai.sensu.io/5123017d3dadf2067fa90fc28275b92e9b586885/sensu-ruby-runtime_0.0.10_ruby-2.4.4_debian_linux_amd64.tar.gz"
			sha512 = "a28952fd93fc63db1f8988c7bc40b0ad815eb9f35ef7317d6caf5d77ecfbfd824a9db54184400aa0c81c29b34cb48c7e8c6e3f17891aaf84cafa3c134266a61a"
			filters = [
				"entity.system.os == 'linux'",
				"entity.system.arch == 'amd64'",
				"entity.system.platform_family == 'debian'",
			]
									headers = {
													"Authorization" = "Bearer changeme"
									}
		}
	}
`

const testAccResourceAsset_createMultipleBuild_2 = `
	resource "sensu_asset" "asset_1" {
		name = "sensu-ruby-runtime"
		annotations = {
			"io.sensu.bonsai.url" = "https://bonsai.sensu.io/assets/sensu/sensu-ruby-runtime"
			"io.sensu.bonsai.api_url" = "https://bonsai.sensu.io/api/v1/assets/sensu/sensu-ruby-runtime"
			"io.sensu.bonsai.tier" = "Community"
			"io.sensu.bonsai.version" = "0.0.10"
			"io.sensu.bonsai.namespace" = "sensu"
			"io.sensu.bonsai.name" = "sensu-ruby-runtime"
			"io.sensu.bonsai.tags" = ""
		}

		build {
			url = "https://example.com/asset_0.2.0_rhel.tar.gz"
			sha512 = "cbee19124b7007342ce37ff9dfd4a1dde03beb1e87e61ca2aef606a7ad3c9bd0bba4e53873c07afa5ac46b0861967a9224511b4504dadb1a5e8fb687e9495304"
			filters = [
				"entity.system.os == 'linux'",
				"entity.system.arch == 'amd64'",
				"entity.system.platform_family == 'rhel'",
				"parseInt(entity.system.platform_version.split('.')[0]) == 6",
			]
		}

		build {
			url = "https://example.com/asset_0.2.0_debian.tar.gz"
			sha512 = "a28952fd93fc63db1f8988c7bc40b0ad815eb9f35ef7317d6caf5d77ecfbfd824a9db54184400aa0c81c29b34cb48c7e8c6e3f17891aaf84cafa3c134266a61a"
			filters = [
				"entity.system.os == 'linux'",
				"entity.system.arch == 'amd64'",
				"entity.system.platform_family == 'debian'",
			]
		}

		headers = {
			"Authorization" = "Bearer changeme"
		}
	}
`
