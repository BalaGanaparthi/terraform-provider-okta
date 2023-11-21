package okta

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/okta/terraform-provider-okta/sdk"
)

func TestAccDataSourceOktaPolicy_read(t *testing.T) {
	config := testAccDataSourcePolicyConfig()

	oktaResourceTest(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		ErrorCheck:        testAccErrorChecks(t),
		ProviderFactories: testAccProvidersFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.okta_policy.test", "id"),
				),
			},
		},
	})
}

func testAccDataSourcePolicyConfig() string {
	return fmt.Sprintf(`
data "okta_policy" "test" {
  type = "%s"
  name = "Default Policy"
}
`, sdk.PasswordPolicyType)
}
