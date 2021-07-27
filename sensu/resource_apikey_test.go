package sensu

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccResourceAPIKey_basic(t *testing.T) {
	username := acctest.RandomWithPrefix("sensu-user1")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceAPIKey_basic(username),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"sensu_apikey.apikey_1", "username", username),
				),
			},
		},
	})
}

func testAccResourceAPIKey_basic(username string) string {
	return fmt.Sprintf(`
		resource "sensu_user" "user_1" {
			name = "%s"
			password = "abcd1234"
			groups = ["admin"]
			disabled = false
		}

		resource "sensu_apikey" "apikey_1" {
    	username = sensu_user.user_1.name
  	}
	`, username)
}
