resource "bloxone_dns_view" "example_view" {
  name = "example_dns_view"
}

resource "bloxone_dns_forward_zone" "example" {
  fqdn = "tf-acc-test.com."

  # Other optional fields
  name    = "example_forward_zone"
  comment = "Example of an Forward Zone"
  tags = {
    site = "Test Site"
  }
  view = bloxone_dns_view.example_view.id
}
