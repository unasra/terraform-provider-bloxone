# Get auth zones filtered by an attribute
data "bloxone_dns_auth_zone" "example_by_attribute" {
  filters = {
    "fqdn"         = "tf-acc-test.com."
    "primary_type" = "cloud"
  }
}

# Get auth zones filtered by tag
data "bloxone_dns_auth_zone" "example_by_tag" {
  tag_filters = {
    "region" = "eu"
  }
}

# Get all auth zones
data "bloxone_dns_auth_zone" "example_all" {}
