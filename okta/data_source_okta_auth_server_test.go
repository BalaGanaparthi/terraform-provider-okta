package okta

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceOktaAuthServer_read(t *testing.T) {
	mgr := newFixtureManager("data-sources", authServer, t.Name())
	config := mgr.GetFixtures("datasource.tf", t)
	authServer := buildTestAuthServer(mgr.Seed)
	oktaResourceTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorChecks(t),
		ProviderFactories: testAccProvidersFactories,
		Steps: []resource.TestStep{
			{
				Config: authServer,
			},
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("okta_auth_server.test", "id"),
					resource.TestCheckResourceAttr("data.okta_auth_server.test", "name", buildResourceName(mgr.Seed)),
					resource.TestCheckResourceAttr("data.okta_auth_server.test", "status", statusActive),
					resource.TestCheckResourceAttrSet("data.okta_auth_server.test", "issuer"),
				),
			},
		},
	})
}

func buildTestAuthServer(i int) string {
	return fmt.Sprintf(`
resource "okta_auth_server" "test" {
  audiences   = ["whatever.rise.zone"]
  description = "test"
  name        = "testAcc_%d"
}`, i)
}
