package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceHandler_basicPipe(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceHandler_basicPipe,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "name", "handler_1"),
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "command", "/bin/foo"),
				),
			},
			resource.TestStep{
				Config: testAccResourceHandler_updatePipe,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "name", "handler_1"),
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "command", "/bin/foo2"),
				),
			},
		},
	})
}

func TestAccResourceHandler_basicTCP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceHandler_basicTCP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "name", "handler_1"),
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "socket.0.port", "80"),
				),
			},
			resource.TestStep{
				Config: testAccResourceHandler_updateTCP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "name", "handler_1"),
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "socket.0.port", "81"),
				),
			},
		},
	})
}

func TestAccResourceHandler_basicSet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceHandler_basicSet,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "name", "handler_1"),
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "handlers.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccResourceHandler_updateSet,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "name", "handler_1"),
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "handlers.#", "3"),
				),
			},
		},
	})
}

func TestAccResourceHandler_runtimeAssets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceHandler_runtimeAssets_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "name", "handler_1"),
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "runtime_assets.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccResourceHandler_runtimeAssets_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "name", "handler_1"),
					resource.TestCheckResourceAttr(
						"sensu_handler.handler_1", "runtime_assets.#", "0"),
				),
			},
		},
	})
}

const testAccResourceHandler_basicPipe = `
  resource "sensu_handler" "handler_1" {
    name = "handler_1"
    type = "pipe"
    command = "/bin/foo"

		env_vars = {
			FOO = "bar"
		}
  }
`

const testAccResourceHandler_updatePipe = `
  resource "sensu_handler" "handler_1" {
    name = "handler_1"
    type = "pipe"
    command = "/bin/foo2"
  }
`

const testAccResourceHandler_basicTCP = `
  resource "sensu_handler" "handler_1" {
    name = "handler_1"
    type = "tcp"
    socket {
      host = "localhost"
      port = 80
    }
    timeout = 30
  }
`

const testAccResourceHandler_updateTCP = `
  resource "sensu_handler" "handler_1" {
    name = "handler_1"
    type = "tcp"
    socket {
      host = "localhost"
      port = 81
    }
    timeout = 30
  }
`

const testAccResourceHandler_basicSet = `
  resource "sensu_handler" "handler_1" {
    name = "handler_1"
    type = "set"
    handlers = [
      "foo",
      "bar",
    ]
  }
`

const testAccResourceHandler_updateSet = `
  resource "sensu_handler" "handler_1" {
    name = "handler_1"
    type = "set"
    handlers = [
      "foo",
      "bar",
      "baz",
    ]
  }
`

var testAccResourceHandler_runtimeAssets_1 = fmt.Sprintf(`
  %s

  resource "sensu_handler" "handler_1" {
    name = "handler_1"
    type = "pipe"
    command = "/bin/foo"

    runtime_assets = [
      "${sensu_asset.asset_1.name}"
		]
  }
`, testAccResourceAsset_basic)

var testAccResourceHandler_runtimeAssets_2 = fmt.Sprintf(`
  %s

  resource "sensu_handler" "handler_1" {
    name = "handler_1"
    type = "pipe"
    command = "/bin/foo"
  }
`, testAccResourceAsset_basic)
