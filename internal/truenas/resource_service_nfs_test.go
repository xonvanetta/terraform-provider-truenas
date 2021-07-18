package truenas

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTruenasServiceNFSBasic(t *testing.T) {
	name := "truenas_service_nfs.nfs"
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigResourceServiceNFS(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "id", "9"),
					resource.TestCheckResourceAttr(name, "start_on_boot", "true"),
					resource.TestCheckResourceAttr(name, "state", "RUNNING"),
					resource.TestCheckResourceAttr(name, "v4", "true"),
				),
			},
		},
	})
}

func testAccConfigResourceServiceNFS() string {
	return fmt.Sprintf(`
		resource "truenas_service_nfs" "nfs" {
			v4 = true
		}
	`)
}
