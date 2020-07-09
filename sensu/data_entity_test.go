package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceEntity_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceEntity_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_entity.entity_1", "name", "entity_1"),
				),
			},
		},
	})
}

var testAccDataSourceEntity_basic = fmt.Sprintf(`
	%s

	data "sensu_entity" "entity_1" {
		name = "${sensu_entity.entity_1.name}"
	}
`, testAccResourceEntity_basic)
