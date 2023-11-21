package okta

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceOktaUserAdminRoles_crud(t *testing.T) {
	mgr := newFixtureManager("resources", userAdminRoles, t.Name())
	start := mgr.GetFixtures("basic.tf", t)
	update := mgr.GetFixtures("basic_update.tf", t)
	remove := mgr.GetFixtures("basic_removal.tf", t)
	resourceName := fmt.Sprintf("%s.test", userAdminRoles)

	oktaResourceTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorChecks(t),
		ProviderFactories: testAccProvidersFactories,
		CheckDestroy:      checkUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: start,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_roles.#", "2"),
				),
			},
			{
				Config: update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_roles.#", "3"),
				),
			},
			{
				Config: remove,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "admin_roles.#", "1"),
				),
			},
		},
	})
}
