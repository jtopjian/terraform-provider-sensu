package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceOrganization_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceOrganization_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_organization.organization_1", "name", "default"),
					resource.TestCheckResourceAttr(
						"data.sensu_organization.organization_1", "description", "Default organization"),
				),
			},
		},
	})
}

const testAccDataSourceOrganization_basic = `
  data "sensu_organization" "organization_1" {
    name = "default"
  }
`
