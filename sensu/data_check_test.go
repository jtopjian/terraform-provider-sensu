package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceCheck_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceCheck_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "name", "check_1"),
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "subscriptions.#", "3"),
				),
			},
		},
	})
}

func TestAccDataSourceCheck_subdue(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceCheck_subdue,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "subdue.1.day", "monday"),
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "subdue.1.begin", "09:00AM"),
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "subdue.1.end", "05:00PM"),
				),
			},
		},
	})
}

func TestAccDataSourceCheck_annotations(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceCheck_annotations,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "annotations.annotation1", "test1"),
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "annotations.annotation2", "test2"),
				),
			},
		},
	})
}

func TestAccDataSourceCheck_labels(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceCheck_labels,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "labels.label1", "test1"),
					resource.TestCheckResourceAttr(
						"data.sensu_check.check_1", "labels.label2", "test2"),
				),
			},
		},
	})
}

var testAccDataSourceCheck_basic = fmt.Sprintf(`
  %s

  data "sensu_check" "check_1" {
    name = "${sensu_check.check_1.name}"
  }
`, testAccResourceCheck_basic)

var testAccDataSourceCheck_subdue = fmt.Sprintf(`
  %s

  data "sensu_check" "check_1" {
    name = "${sensu_check.check_1.name}"
  }
`, testAccResourceCheck_subdue_1)

var testAccDataSourceCheck_annotations = fmt.Sprintf(`
  %s

  data "sensu_check" "check_1" {
    name = "${sensu_check.check_1.name}"
  }
`, testAccResourceCheck_annotations)

var testAccDataSourceCheck_labels = fmt.Sprintf(`
  %s

  data "sensu_check" "check_1" {
    name = "${sensu_check.check_1.name}"
  }
`, testAccResourceCheck_labels)
