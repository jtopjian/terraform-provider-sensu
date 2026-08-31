package sensu

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourcePipeline_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePipeline_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "name", "pipeline_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "namespace", "default"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.#", "2"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.0.name", "workflow_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.0.handler", "pipeline_test_handler"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.1.name", "workflow_2"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.1.handler", "pipeline_test_handler"),
				),
			},
		},
	})
}

const testAccDataSourcePipeline_basic = `
resource "sensu_handler" "handler_1" {
	name    = "pipeline_test_handler"
	type    = "pipe"
	command = "/bin/foo"
}

resource "sensu_pipeline" "pipeline_1" {
	name = "pipeline_1"

	workflow {
		name    = "workflow_1"
		handler = sensu_handler.handler_1.name
	}

	workflow {
		name    = "workflow_2"
		handler = sensu_handler.handler_1.name
	}
}

data "sensu_pipeline" "pipeline_1" {
	name = sensu_pipeline.pipeline_1.name
}
`

func TestAccDataSourcePipeline_filters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePipeline_filters,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.#", "1"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.0.name", "workflow_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.0.filters.#", "2"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.0.filters.0", "is_incident"),
					resource.TestCheckResourceAttr(
						"data.sensu_pipeline.pipeline_1", "workflow.0.filters.1", "not_silenced"),
				),
			},
		},
	})
}

const testAccDataSourcePipeline_filters = `
resource "sensu_handler" "handler_1" {
	name    = "pipeline_test_handler"
	type    = "pipe"
	command = "/bin/foo"
}

resource "sensu_pipeline" "pipeline_1" {
	name = "pipeline_1"

	workflow {
		name    = "workflow_1"
		handler = sensu_handler.handler_1.name

		filters = [
			"is_incident",
			"not_silenced",
		]
	}
}

data "sensu_pipeline" "pipeline_1" {
	name = sensu_pipeline.pipeline_1.name
}
`
