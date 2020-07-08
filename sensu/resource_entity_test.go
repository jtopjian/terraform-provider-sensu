package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceEntity_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceEntity_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "name", "entity_1"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "class", "proxy"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "labels.foo", "bar"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "subscriptions.#", "3"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "deregistration.0.handler", "foo"),
				),
			},
			resource.TestStep{
				Config: testAccResourceEntity_update_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "name", "entity_1"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "class", "proxy"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "labels.foo", "bar"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "subscriptions.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "deregistration.#", "0"),
				),
			},
			resource.TestStep{
				Config: testAccResourceEntity_update_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "name", "entity_1"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "class", "proxy"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "labels.foo", "baz"),
					resource.TestCheckResourceAttr(
						"sensu_entity.entity_1", "subscriptions.#", "2"),
				),
			},
		},
	})
}

const testAccResourceEntity_basic = `
	resource "sensu_entity" "entity_1" {
		name = "entity_1"
		class = "proxy"
		labels = {
			foo = "bar"
			password = "supersecret"
		}
		annotations = {
			password = "supersecret"
		}
		subscriptions = [
			"foo",
			"bar",
			"baz",
		]
		deregistration {
			handler = "foo"
		}
	}
`

const testAccResourceEntity_update_1 = `
	resource "sensu_entity" "entity_1" {
		name = "entity_1"
		class = "proxy"
		labels = {
			foo = "bar"
			password = "supersecret"
		}
		subscriptions = [
			"foo",
			"baz",
		]
	}
`

const testAccResourceEntity_update_2 = `
	resource "sensu_entity" "entity_1" {
		name = "entity_1"
		class = "proxy"
		labels = {
			foo = "baz"
			password = "supersecret"
		}
		subscriptions = [
			"foo",
			"baz",
		]
	}
`
