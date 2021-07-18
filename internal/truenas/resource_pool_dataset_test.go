package truenas

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTruenasPoolDatasetBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigResourcePoolDataset(),
				Check:  resource.ComposeTestCheckFunc(
				//resource.TestCheckResourceAttr(name, "id", "9"),
				//resource.TestCheckResourceAttr(name, "start_on_boot", "true"),
				//resource.TestCheckResourceAttr(name, "state", "RUNNING"),
				//resource.TestCheckResourceAttr(name, "v4", "true"),
				),
			},
		},
	})
}

func testAccConfigResourcePoolDataset() string {
	return fmt.Sprintf(`
		resource "truenas_pool_dataset" "layer1" {
 			name = "%[1]s/layer1"
			comments = "layer1"
		}

		resource "truenas_pool_dataset" "layer2" {
 			name = "%[1]s/layer1/layer2"
			comments = "layer2"

			depends_on = [
				truenas_pool_dataset.layer1
			]
		}

		resource "truenas_pool_dataset" "zvol" {
 			name = "%[1]s/layer1/zvol"
			comments = "zvol"
			type = "VOLUME"
			volsize = "40M"

			depends_on = [
				truenas_pool_dataset.layer1
			]
		}
	`, testResourcePrefix)
}
