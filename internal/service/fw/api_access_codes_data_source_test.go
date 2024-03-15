package fw_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccAccessCodesDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_access_codes.test"
	resourceName := "bloxone_td_access_code.test"
	var v fw.AtcfwAccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAccessCodesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAccessCodesDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					}, testAccCheckAccessCodesResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAccessCodesDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_td_access_codes.test"
	resourceName := "bloxone_td_access_code.test"
	var v fw.AtcfwAccessCode
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAccessCodesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAccessCodesDataSourceConfigTagFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					}, testAccCheckAccessCodesResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAccessCodesResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{}
}

func testAccAccessCodesDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test" {
	name = %[1]q
	activation = %q
	expiration = %q
	rules = [
		{
			action = "" ,
			data = "antimalware",
			description = "",
			redirect_name = "",
			type = "named_feed"
		}
	]
}

data "bloxone_td_access_codes" "test" {
  filters = {
	 name = %[1]q
  }
}
`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339))
}

func testAccAccessCodesDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_td_access_codes" "test" {
  tag_filters = {
	tag1 = bloxone_td_access_code.test.tags.tag1
  }
}
`, tagValue)
}