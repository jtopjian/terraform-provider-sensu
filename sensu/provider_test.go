package sensu

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var (
	SENSU_API_URL  = os.Getenv("SENSU_API_URL")
	SENSU_USERNAME = os.Getenv("SENSU_USERNAME")
	SENSU_PASSWORD = os.Getenv("SENSU_PASSWORD")
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"sensu": testAccProvider,
	}
}

func testAccPreCheckRequiredEnvVars(t *testing.T) {
	if SENSU_API_URL == "" {
		t.Fatal("SENSU_API_URL must be set for acceptance tests")
	}

	if SENSU_USERNAME == "" {
		t.Fatal("SENSU_USERNAME must be set for acceptance tests")
	}

	if SENSU_PASSWORD == "" {
		t.Fatal("SENSU_PASSWORD must be set for acceptance tests")
	}
}

func testAccPreCheck(t *testing.T) {
	testAccPreCheckRequiredEnvVars(t)
}
