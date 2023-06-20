package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceNamespace_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceNamespace_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_namespace.namespace_1", "name", "namespace_1"),
				),
			},
			resource.TestStep{
				Config: testAccResourceNamespace_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_namespace.namespace_1", "name", "namespace_1b"),
				),
			},
		},
	})
}

func TestAccResourceNamespace_check(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceNamespace_check,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_namespace.namespace_1", "name", "namespace_1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "name", "check_1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "namespace", "namespace_1"),
				),
			},
		},
	})
}

const testAccResourceNamespace_basic = `
	resource "sensu_namespace" "namespace_1" {
		name = "namespace_1"
	}
`

const testAccResourceNamespace_update = `
	resource "sensu_namespace" "namespace_1" {
		name = "namespace_1b"
	}
`

const testAccResourceNamespace_check = `
	resource "sensu_namespace" "namespace_1" {
		name = "namespace_1"
	}

	resource "sensu_check" "check_1" {
		name = "check_1"
		namespace = "${sensu_namespace.namespace_1.name}"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
			"baz",
		]
	}
`
