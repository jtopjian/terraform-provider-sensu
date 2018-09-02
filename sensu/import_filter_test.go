package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccImportFilter_basic(t *testing.T) {
	resourceName := "sensu_filter.filter_1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceFilter_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
