package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceNamespace_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceNamespace_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_namespace.namespace_1", "name", "namespace_1"),
				),
			},
		},
	})
}

var testAccDataSourceNamespace_basic = fmt.Sprintf(`
  %s

  data "sensu_namespace" "namespace_1" {
    name = "${sensu_namespace.namespace_1.name}"
  }
`, testAccResourceNamespace_basic)
