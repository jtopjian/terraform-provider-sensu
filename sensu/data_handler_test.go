package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceHandler_basicPipe(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceHandler_basicPipe,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_handler.handler_1", "type", "pipe"),
					resource.TestCheckResourceAttr(
						"data.sensu_handler.handler_1", "command", "/bin/foo"),
				),
			},
		},
	})
}

func TestAccDataSourceHandler_basicTCP(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceHandler_basicTCP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_handler.handler_1", "type", "tcp"),
					resource.TestCheckResourceAttr(
						"data.sensu_handler.handler_1", "socket.0.port", "80"),
				),
			},
		},
	})
}

func TestAccDataSourceHandler_basicSet(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceHandler_basicSet,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_handler.handler_1", "type", "set"),
					resource.TestCheckResourceAttr(
						"data.sensu_handler.handler_1", "handlers.#", "2"),
				),
			},
		},
	})
}

var testAccDataSourceHandler_basicPipe = fmt.Sprintf(`
  %s

  data "sensu_handler" "handler_1" {
    name = "${sensu_handler.handler_1.name}"
  }
`, testAccResourceHandler_basicPipe)

var testAccDataSourceHandler_basicTCP = fmt.Sprintf(`
  %s

  data "sensu_handler" "handler_1" {
    name = "${sensu_handler.handler_1.name}"
  }
`, testAccResourceHandler_basicTCP)

var testAccDataSourceHandler_basicSet = fmt.Sprintf(`
  %s

  data "sensu_handler" "handler_1" {
    name = "${sensu_handler.handler_1.name}"
  }
`, testAccResourceHandler_basicSet)
