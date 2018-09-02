package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccImportRole_basic(t *testing.T) {
	resourceName := "sensu_role.role_1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceRole_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
