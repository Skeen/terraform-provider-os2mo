package os2mo

import (
	"fmt"
	"testing"

	hc "github.com/hashicorp-demoapp/os2mo-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccOS2moOrderBasic(t *testing.T) {
	coffeeID := "1"
	quantity := "2"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOS2moOrderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckOS2moOrderConfigBasic(coffeeID, quantity),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOS2moOrderExists("os2mo_order.new"),
				),
			},
		},
	})
}

func testAccCheckOS2moOrderDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*hc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "os2mo_order" {
			continue
		}

		orderID := rs.Primary.ID

		err := c.DeleteOrder(orderID)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckOS2moOrderConfigBasic(coffeeID, quantity string) string {
	return fmt.Sprintf(`
	resource "os2mo_order" "new" {
		items {
			coffee {
				id = %s
			}
    		quantity = %s
  		}
	}
	`, coffeeID, quantity)
}

func testAccCheckOS2moOrderExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No OrderID set")
		}

		return nil
	}
}
