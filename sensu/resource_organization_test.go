package sensu

/*
	It's currently not possible to test an organization
	since creating an org also implicitly creates a default
	environment.

func TestAccResourceOrganization_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceOrganization_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_organization.organization_1", "name", "organization_1"),
					resource.TestCheckResourceAttr(
						"sensu_organization.organization_1", "description", "an organization"),
				),
			},
			resource.TestStep{
				Config: testAccResourceOrganization_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_organization.organization_1", "name", "organization_1"),
					resource.TestCheckResourceAttr(
						"sensu_organization.organization_1", "description", "some organization"),
				),
			},
		},
	})
}
*/

const testAccResourceOrganization_basic = `
  resource "sensu_organization" "organization_1" {
    name = "organization_1"
    description = "an organization"
  }
`

const testAccResourceOrganization_update = `
  resource "sensu_organization" "organization_1" {
    name = "organization_1"
    description = "some organization"
  }
`
