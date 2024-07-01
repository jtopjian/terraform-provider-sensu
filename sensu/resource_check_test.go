package sensu

import (
	"fmt"
	"os"
	"testing"

	"github.com/blang/semver/v4"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
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
				Config: testAccResourceCheck_hook_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "name", "hook_1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "check_hook.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_2", "name", "hook_2"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_hook_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_1", "name", "hook_1"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "check_hook.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_hook.hook_2", "name", "hook_2_modified"),
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

func TestAccResourceCheck_proxyRequests(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_proxyRequests_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "proxy_requests.0.entity_attributes.#", "2"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_proxyRequests_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "proxy_requests.#", "0"),
				),
			},
		},
	})
}

func TestAccResourceCheck_pipelines(t *testing.T) {
	sensuVersion := os.Getenv("SENSU_VERSION")

	envVersion, err := semver.Parse(sensuVersion)
	if err != nil {
		fmt.Printf("Error parsing version: %v", err)
		return
	}

	pipelineMinVersion, err := semver.Parse("6.5.0")
	if err != nil {
		fmt.Printf("Error parsing version: %v", err)
		return
	}

	steps := []resource.TestStep{
		resource.TestStep{
			Config: testAccResourceCheck_pipelines_1,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"sensu_check.check_1", "pipelines.0.name", "incident_alerts"),
				resource.TestCheckResourceAttr(
					"sensu_check.check_1", "pipelines.1.name", "low_priority_alerts"),
				resource.TestCheckResourceAttr(
					"sensu_check.check_1", "pipelines.1.type", "Pipeline"),
			),
		},
		resource.TestStep{
			Config: testAccResourceCheck_pipelines_2,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"sensu_check.check_1", "pipelines.#", "0"),
			),
		},
	}
	if envVersion.LT(pipelineMinVersion) {
		steps = steps[1:]
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps:     steps,
	})
}

func TestAccResourceCheck_envVars(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_envvars_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "env_vars.foo", "bar"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_envvars_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "env_vars.foo", "bar"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "env_vars.bar", "baz"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_envvars_3,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "env_vars.foo", "barr"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "env_vars.bar", "baz"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_envvars_4,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "env_vars.foo", "barr"),
					resource.TestCheckNoResourceAttr("sensu_check.check_1", "env_vars.bar"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_envvars_5,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sensu_check.check_1", "env_vars"),
				),
			},
		},
	})
}

func TestAccResourceCheck_secrets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceCheck_secrets_1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "secrets.foo", "bar"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_secrets_2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "secrets.foo", "bar"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "secrets.bar", "baz"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_secrets_3,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "secrets.foo", "barr"),
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "secrets.bar", "baz"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_secrets_4,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_check.check_1", "secrets.foo", "barr"),
					resource.TestCheckNoResourceAttr("sensu_check.check_1", "secrets.bar"),
				),
			},
			resource.TestStep{
				Config: testAccResourceCheck_secrets_5,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("sensu_check.check_1", "secrets"),
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

const testAccResourceCheck_hook_1 = `
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

const testAccResourceCheck_hook_2 = `
  resource "sensu_hook" "hook_1" {
    name = "hook_1"
    command = "/bin/foo"
  }

  resource "sensu_hook" "hook_2" {
    name = "hook_2_modified"
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

const testAccResourceCheck_proxyRequests_1 = `
  resource "sensu_entity" "entities" {
    count = 3
    name = format("entity-%02d", count.index+1)
    class = "proxy"
    labels = {
      "proxy_type" = "website"
      "url" = format("http://example-%02d.com", count.index+1)
    }
  }

  resource "sensu_check" "check_1" {
    name = "check-http"
    command = "check-http.rb -u {{ .labels url }}"
    interval = 60
    proxy_requests {
      entity_attributes = [
        "entity.entity_class == 'proxy'",
        "entity.labels.proxy_type == 'website'",
      ]
    }
    publish = true
    subscriptions = ["proxy"]
  }
`

const testAccResourceCheck_proxyRequests_2 = `
  resource "sensu_entity" "entities" {
    count = 3
    name = format("entity-%02d", count.index+1)
    class = "proxy"
    labels = {
      "proxy_type" = "website"
      "url" = format("http://example-%02d.com", count.index+1)
    }
  }

  resource "sensu_check" "check_1" {
    name = "check-http"
    command = "check-http.rb -u {{ .labels url }}"
    interval = 60
    publish = true
    subscriptions = ["proxy"]
  }
`

const testAccResourceCheck_pipelines_1 = `
  resource "sensu_entity" "entities" {
    count = 3
    name = format("entity-%02d", count.index+1)
    class = "proxy"
    labels = {
      "proxy_type" = "website"
      "url" = format("http://example-%02d.com", count.index+1)
    }
  }

  resource "sensu_check" "check_1" {
    name = "check-http"
    command = "check-http.rb -u {{ .labels url }}"
    interval = 60
    pipelines {
		api_version = "core/v2"
        type = "Pipeline"
        name = "incident_alerts"
	}
	pipelines {
		api_version = "core/v2"
        type = "Pipeline"
        name = "low_priority_alerts"
	}
    publish = true
    subscriptions = ["proxy"]
  }
`

const testAccResourceCheck_pipelines_2 = `
  resource "sensu_entity" "entities" {
    count = 3
    name = format("entity-%02d", count.index+1)
    class = "proxy"
    labels = {
      "proxy_type" = "website"
      "url" = format("http://example-%02d.com", count.index+1)
    }
  }

  resource "sensu_check" "check_1" {
    name = "check-http"
    command = "check-http.rb -u {{ .labels url }}"
    interval = 60
    publish = true
    subscriptions = ["proxy"]
  }
`

const testAccResourceCheck_envvars_1 = `
  resource "sensu_check" "check_1" {
    name = "check_1"
    command = "/bin/foo"
    interval = 60000
    subscriptions = [
      "foo",
      "bar",
      "baz",
    ]
    env_vars = {
      "foo" = "bar",
    }
  }
`

const testAccResourceCheck_envvars_2 = `
  resource "sensu_check" "check_1" {
    name = "check_1"
    command = "/bin/foo"
    interval = 60000
    subscriptions = [
      "foo",
      "bar",
      "baz",
    ]
    env_vars = {
      "foo" = "bar",
      "bar" = "baz",
    }
  }
`

const testAccResourceCheck_envvars_3 = `
  resource "sensu_check" "check_1" {
    name = "check_1"
    command = "/bin/foo"
    interval = 60000
    subscriptions = [
      "foo",
      "bar",
      "baz",
    ]
    env_vars = {
      "foo" = "barr",
      "bar" = "baz",
    }
  }
`

const testAccResourceCheck_envvars_4 = `
  resource "sensu_check" "check_1" {
    name = "check_1"
    command = "/bin/foo"
    interval = 60000
    subscriptions = [
      "foo",
      "bar",
      "baz",
    ]
    env_vars = {
      "foo" = "barr",
    }
  }
`

const testAccResourceCheck_envvars_5 = `
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

const testAccResourceCheck_secrets_1 = `
  resource "sensu_check" "check_1" {
    name = "check_1"
    command = "/bin/foo"
    interval = 60000
    subscriptions = [
      "foo",
      "bar",
      "baz",
    ]
    secrets = {
      "foo" = "bar",
    }
  }
`

const testAccResourceCheck_secrets_2 = `
  resource "sensu_check" "check_1" {
    name = "check_1"
    command = "/bin/foo"
    interval = 60000
    subscriptions = [
      "foo",
      "bar",
      "baz",
    ]
    secrets = {
      "foo" = "bar",
      "bar" = "baz",
    }
  }
`

const testAccResourceCheck_secrets_3 = `
  resource "sensu_check" "check_1" {
    name = "check_1"
    command = "/bin/foo"
    interval = 60000
    subscriptions = [
      "foo",
      "bar",
      "baz",
    ]
    secrets = {
      "foo" = "barr",
      "bar" = "baz",
    }
  }
`

const testAccResourceCheck_secrets_4 = `
  resource "sensu_check" "check_1" {
    name = "check_1"
    command = "/bin/foo"
    interval = 60000
    subscriptions = [
      "foo",
      "bar",
      "baz",
    ]
    secrets = {
      "foo" = "barr",
    }
  }
`

const testAccResourceCheck_secrets_5 = `
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
