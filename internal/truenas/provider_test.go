package truenas

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]func() (*schema.Provider, error)
var testAccProvider *schema.Provider

const testResourcePrefix = "tf-acc-test"

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]func() (*schema.Provider, error){
		"truenas": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TRUENAS_HOST"); v == "" {
		t.Fatal("TRUENAS_HOST must be set for acceptance tests")
	}
	if v := os.Getenv("TRUENAS_API_KEY"); v == "" {
		t.Fatal("TRUENAS_API_KEY must be set for acceptance tests")
	}
}
