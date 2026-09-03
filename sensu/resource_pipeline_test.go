package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourcePipeline_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipeline_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "name", "pipeline_1"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "namespace", "default"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.0.name", "workflow_1"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.0.handler", "pipeline_test_handler"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.0.filters.#", "1"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.1.name", "workflow_2"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.1.handler", "pipeline_test_handler"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.1.filters.#", "1"),
				),
			},
			{
				Config: testAccResourcePipeline_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "name", "pipeline_1"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.#", "2"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.0.filters.#", "0"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.1.filters.#", "1"),
				),
			},
		},
	})
}

func TestAccResourcePipeline_noFilters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePipeline_noFilters,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.0.name", "workflow_1"),
					resource.TestCheckResourceAttr(
						"sensu_pipeline.pipeline_1", "workflow.0.filters.#", "0"),
				),
			},
		},
	})
}

const testAccResourcePipeline_basic = `
resource "sensu_filter" "filter_1" {
  name   = "pipeline_test_filter"
  action = "allow"

  expressions = [
    "event.Check.Team == 'ops'",
  ]
}

resource "sensu_handler" "handler_1" {
  name    = "pipeline_test_handler"
  type    = "pipe"
  command = "/bin/true"
}

resource "sensu_pipeline" "pipeline_1" {
  name = "pipeline_1"

  workflow {
    name = "workflow_1"

    filters = [
      sensu_filter.filter_1.name,
    ]

    handler = sensu_handler.handler_1.name
  }

  workflow {
    name = "workflow_2"

    filters = [
      sensu_filter.filter_1.name,
    ]

    handler = sensu_handler.handler_1.name
  }
}
`

const testAccResourcePipeline_update = `
resource "sensu_filter" "filter_1" {
  name   = "pipeline_test_filter"
  action = "allow"

  expressions = [
    "event.Check.Team == 'ops'",
  ]
}

resource "sensu_handler" "handler_1" {
  name    = "pipeline_test_handler"
  type    = "pipe"
  command = "/bin/true"
}

resource "sensu_pipeline" "pipeline_1" {
  name = "pipeline_1"

  workflow {
    name    = "workflow_1"
    handler = sensu_handler.handler_1.name
  }

  workflow {
    name = "workflow_2"

    filters = [
      sensu_filter.filter_1.name,
    ]

    handler = sensu_handler.handler_1.name
  }
}
`

const testAccResourcePipeline_noFilters = `
resource "sensu_handler" "handler_1" {
  name    = "pipeline_test_handler"
  type    = "pipe"
  command = "/bin/true"
}

resource "sensu_pipeline" "pipeline_1" {
  name = "pipeline_1"

  workflow {
    name    = "workflow_1"
    handler = sensu_handler.handler_1.name
  }
}
`
