package sensu

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceCheck_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "name", "check_1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "command", "/bin/foo"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "interval", "60000"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subscriptions.#", "3"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_update_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "name", "check_1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "command", "/bin/foo2"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "interval", "60001"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subscriptions.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_update_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "name", "check_1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "command", "/bin/foo2"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "interval", "0"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "cron", "*/20 * * * *"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subscriptions.#", "2"),
				),
			},
		},
	})
}

/*
func TestAccResourceCheck_proxyRequests(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_proxyRequests,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "proxy_requests.0.entity_attributes.0", "entity.Class == \"proxy\""),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "proxy_requests.0.splay", "true"),
				),
			},
		},
	})
}
*/

func TestAccResourceCheck_subdue(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_subdue_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subdue.133776838.day", "monday"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subdue.133776838.begin", "09:00AM"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subdue.133776838.end", "05:00PM"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_subdue_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subdue.907108187.day", "monday"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subdue.907108187.begin", "09:00AM"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "subdue.907108187.end", "06:00PM"),
				),
			},
		},
	})
}

func TestAccResourceCheck_hook(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_hook,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "name", "hook_1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "check_hook.#", "2"),
				),
			},
		},
	})
}

func TestAccResourceCheck_labels(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_labels_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "labels.label1", "test1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "labels.label2", "test2"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_labels_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "labels.label1", "test1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "labels.label2", "test3"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_labels_3,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sensu_check.check_1", "labels"),
				),
			},
		},
	})
}

func TestAccResourceCheck_annotations(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_annotations_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "annotations.annotation1", "test1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "annotations.annotation2", "test2"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_annotations_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "annotations.annotation1", "test1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "annotations.annotation2", "test3"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_annotations_3,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sensu_check.check_1", "annotations"),
				),
			},
		},
	})
}

const testAccResourceCheck_basic = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
			"baz",
		]
	}
`

const testAccResourceCheck_update_1 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo2"
		interval = 60001
		subscriptions = [
			"foo",
			"baz",
		]
	}
`

const testAccResourceCheck_update_2 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo2"
		cron = "*/20 * * * *"
		subscriptions = [
			"foo",
			"baz",
		]
	}
`

/*
const testAccResourceCheck_proxyRequests = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]

		proxy_requests {
			entity_attributes = [
				"entity.Class == \"proxy\"",
			]
			splay = true
			splay_coverage = 90
		}
	}
`
*/

const testAccResourceCheck_subdue_1 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]

		subdue {
			day = "monday"
			begin = "09:00AM"
			end = "05:00PM"
		}

		subdue {
			day = "monday"
			begin = "07:00PM"
			end = "09:00PM"
		}

		subdue {
			day = "tuesday"
			begin = "03:00AM"
			end = "09:00AM"
		}
	}
`

const testAccResourceCheck_subdue_2 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]

		subdue {
			day = "monday"
			begin = "09:00AM"
			end = "06:00PM"
		}

		subdue {
			day = "monday"
			begin = "07:00PM"
			end = "09:00PM"
		}

		subdue {
			day = "tuesday"
			begin = "03:00AM"
			end = "09:00AM"
		}
	}
`

const testAccResourceCheck_hook = `
	resource "sensu_hook" "hook_1" {
		name = "hook_1"
		command = "/bin/foo"
	}

	resource "sensu_hook" "hook_2" {
		name = "hook_2"
		command = "/bin/bar"
	}

	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foobar"
		interval = 6000
		subscriptions = ["foo", "bar"]

		check_hook {
			hook = "${sensu_hook.hook_1.name}"
			trigger = "non-zero"
		}

		check_hook {
			hook = "${sensu_hook.hook_2.name}"
			trigger = "non-zero"
		}
	}
`

const testAccResourceCheck_labels_1 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]

		labels = {
			label1 = "test1"
			label2 = "test2"
		}
	}
`

const testAccResourceCheck_labels_2 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]

		labels = {
			label1 = "test1"
			label2 = "test3"
		}
	}
`

const testAccResourceCheck_labels_3 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]
	}
`

const testAccResourceCheck_annotations_1 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]

		annotations = {
			annotation1 = "test1"
			annotation2 = "test2"
		}
	}
`

const testAccResourceCheck_annotations_2 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]

		annotations = {
			annotation1 = "test1"
			annotation2 = "test3"
		}
	}
`

const testAccResourceCheck_annotations_3 = `
	resource "sensu_check" "check_1" {
		name = "check_1"
		command = "/bin/foo"
		interval = 60000
		subscriptions = [
			"foo",
			"bar",
		]
	}
`
