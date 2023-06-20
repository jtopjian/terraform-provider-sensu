package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceMutator_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceMutator_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "name", "mutator_1"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "command", "/bin/foo"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "env_vars.FOO", "bar"),
				),
			},
			resource.TestStep{
				Config: testAccResourceMutator_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "name", "mutator_1"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "command", "/bin/foo2"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "env_vars.FOO", "baz"),
				),
			},
		},
	})
}

func TestAccResourceMutator_secrets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceMutator_secrets_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "secrets.foo", "bar"),
				),
			},
			resource.TestStep{
				Config: testAccResourceMutator_secrets_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "secrets.foo", "bar"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "secrets.bar", "baz"),
				),
			},
			resource.TestStep{
				Config: testAccResourceMutator_secrets_3,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "secrets.foo", "barr"),
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "secrets.bar", "baz"),
				),
			},
			resource.TestStep{
				Config: testAccResourceMutator_secrets_4,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_mutator.mutator_1", "secrets.foo", "barr"),
					resource.TestCheckNoResourceAttr("sensu_mutator.mutator_1", "secrets.bar"),
				),
			},
			resource.TestStep{
				Config: testAccResourceMutator_secrets_5,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sensu_mutator.mutator_1", "secrets"),
				),
			},
		},
	})
}

const testAccResourceMutator_basic = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo"

    env_vars = {
      FOO = "bar"
    }

		secrets = {
			BAR = "foo"
		}
  }
`

const testAccResourceMutator_update = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo2"

    env_vars = {
      FOO = "baz"
    }
 }
`

const testAccResourceMutator_secrets_1 = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo2"
		secrets = {
			"foo" = "bar",
		}
	}
`

const testAccResourceMutator_secrets_2 = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo2"
		secrets = {
			"foo" = "bar",
			"bar" = "baz",
		}
	}
`

const testAccResourceMutator_secrets_3 = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo2"
		secrets = {
			"foo" = "barr",
			"bar" = "baz",
		}
	}
`

const testAccResourceMutator_secrets_4 = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo2"
		secrets = {
			"foo" = "barr",
		}
	}
`

const testAccResourceMutator_secrets_5 = `
  resource "sensu_mutator" "mutator_1" {
    name = "mutator_1"
    command = "/bin/foo2"
	}
`
