package truenas

//func TestAccTruenasSharingNFSBasic(t *testing.T) {
//	resource.Test(t, resource.TestCase{
//		PreCheck:          func() { testAccPreCheck(t) },
//		ProviderFactories: testAccProviders,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccConfigResourceSharingNFS(),
//				Check: resource.ComposeTestCheckFunc(
//					resource.TestCheckResourceAttr(name, "id", "9"),
//					resource.TestCheckResourceAttr(name, "start_on_boot", "true"),
//					resource.TestCheckResourceAttr(name, "state", "RUNNING"),
//					resource.TestCheckResourceAttr(name, "v4", "true"),
//				),
//			},
//		},
//	})
//}
//
//func testAccConfigResourceSharingNFS() string {
//	return fmt.Sprintf(`
//		resource "truenas_service_nfs" "nfs" {
//			v4 = true
//		}
//	`)
//}
