package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccResourceUser_basic(t *testing.T) {
	username := acctest.RandomWithPrefix("sensu-acctest")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceUser_basic(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_user.user_1", "name", username),
					resource.TestCheckResourceAttr(
						"sensu_user.user_1", "groups.#", "1"),
				),
			},
			resource.TestStep{
				Config: testAccResourceUser_update(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_user.user_1", "name", username),
					resource.TestCheckResourceAttr(
						"sensu_user.user_1", "groups.#", "2"),
				),
			},
		},
	})
}

func TestAccResourceUser_password(t *testing.T) {
	username := acctest.RandomWithPrefix("sensu-acctest")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceUser_password_1(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_user.user_1", "name", username),
				),
			},
			resource.TestStep{
				Config: testAccResourceUser_password_2(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_user.user_1", "name", username),
				),
			},
		},
	})
}

func TestAccResourceUser_disabled(t *testing.T) {
	username := acctest.RandomWithPrefix("sensu-acctest")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceUser_disabled(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_user.user_1", "name", username),
				),
			},
		},
	})
}

func testAccResourceUser_basic(username string) string {
	return fmt.Sprintf(`
		resource "sensu_user" "user_1" {
			name = "%s"
			password = "abcd1234"
			groups = ["admin"]
		}
	`, username)
}

func testAccResourceUser_update(username string) string {
	return fmt.Sprintf(`
		resource "sensu_user" "user_1" {
			name = "%s"
			password = "abcd1234"
			groups = ["admin", "read-only"]
		}
	`, username)
}

func testAccResourceUser_password_1(username string) string {
	return fmt.Sprintf(`
		resource "sensu_user" "user_1" {
			name = "%s"
			password = "abcd1234"
			groups = ["admin"]
		}
	`, username)
}

func testAccResourceUser_password_2(username string) string {
	return fmt.Sprintf(`
		resource "sensu_user" "user_1" {
			name = "%s"
			password = "1234abcd"
			groups = ["admin"]
		}
	`, username)
}

func testAccResourceUser_disabled(username string) string {
	return fmt.Sprintf(`
		resource "sensu_user" "user_1" {
			name = "%s"
			password = "1234abcd"
			groups = ["admin"]
			disabled = true
		}
	`, username)
}