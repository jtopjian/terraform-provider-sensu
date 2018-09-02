package sensu

/*
	Disabling test temporarily.

	It works, but I need to figure out how to successfully
	run it in Travis.
*/

/*
import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceEntity_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceEntity_basicPipe,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_entity.entity_1", "name", "sensu"),
				),
			},
		},
	})
}
*/

const testAccDataSourceEntity_basicPipe = `
	data "sensu_entity" "entity_1" {
		name = "sensu"
	}
`
