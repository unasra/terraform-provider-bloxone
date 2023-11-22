package ipam_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccSubnetDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_subnets.test"
	resourceName := "bloxone_ipam_subnet.test"
	var v ipam.IpamsvcSubnet

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubnetDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSubnetDataSourceConfigFilters("11.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSubnetExists(context.Background(), resourceName, &v),
					}, testAccCheckSubnetResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSubnetDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_subnets.test"
	resourceName := "bloxone_ipam_subnet.test"
	var v ipam.IpamsvcSubnet
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubnetDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSubnetDataSourceConfigTagFilters("11.0.0.0", 24, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSubnetExists(context.Background(), resourceName, &v),
					}, testAccCheckSubnetResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSubnetResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "asm_config", dataSourceName, "results.0.asm_config"),
		resource.TestCheckResourceAttrPair(resourceName, "asm_scope_flag", dataSourceName, "results.0.asm_scope_flag"),
		resource.TestCheckResourceAttrPair(resourceName, "cidr", dataSourceName, "results.0.cidr"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_client_update", dataSourceName, "results.0.ddns_client_update"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_conflict_resolution_mode", dataSourceName, "results.0.ddns_conflict_resolution_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domain", dataSourceName, "results.0.ddns_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_name", dataSourceName, "results.0.ddns_generate_name"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generated_prefix", dataSourceName, "results.0.ddns_generated_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_send_updates", dataSourceName, "results.0.ddns_send_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl_percent", dataSourceName, "results.0.ddns_ttl_percent"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_update_on_renew", dataSourceName, "results.0.ddns_update_on_renew"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_use_conflict_resolution", dataSourceName, "results.0.ddns_use_conflict_resolution"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_config", dataSourceName, "results.0.dhcp_config"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_host", dataSourceName, "results.0.dhcp_host"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_options", dataSourceName, "results.0.dhcp_options"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_utilization", dataSourceName, "results.0.dhcp_utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_dhcp", dataSourceName, "results.0.disable_dhcp"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_attrs", dataSourceName, "results.0.discovery_attrs"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_metadata", dataSourceName, "results.0.discovery_metadata"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_filename", dataSourceName, "results.0.header_option_filename"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_server_address", dataSourceName, "results.0.header_option_server_address"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_server_name", dataSourceName, "results.0.header_option_server_name"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_char", dataSourceName, "results.0.hostname_rewrite_char"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_enabled", dataSourceName, "results.0.hostname_rewrite_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_regex", dataSourceName, "results.0.hostname_rewrite_regex"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_assigned_hosts", dataSourceName, "results.0.inheritance_assigned_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_parent", dataSourceName, "results.0.inheritance_parent"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "rebind_time", dataSourceName, "results.0.rebind_time"),
		resource.TestCheckResourceAttrPair(resourceName, "renew_time", dataSourceName, "results.0.renew_time"),
		resource.TestCheckResourceAttrPair(resourceName, "space", dataSourceName, "results.0.space"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "threshold", dataSourceName, "results.0.threshold"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "usage", dataSourceName, "results.0.usage"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization", dataSourceName, "results.0.utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization_v6", dataSourceName, "results.0.utilization_v6"),
	}
}

func testAccSubnetDataSourceConfigFilters(address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
}
data "bloxone_ipam_subnets" "test" {
  filters = {
	address = bloxone_ipam_subnet.test.address
  }
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")
}

func testAccSubnetDataSourceConfigTagFilters(address string, cidr int, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    tags = {
	  tag1 = %q
    }
}
data "bloxone_ipam_subnets" "test" {
  tag_filters = {
	tag1 = bloxone_ipam_subnet.test.tags.tag1
  }
}
`, address, cidr, tagValue)
	return strings.Join([]string{testAccBaseWithIPSpace(), config}, "")

}
