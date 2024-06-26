---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_td_named_list Resource - terraform-provider-bloxone"
subcategory: "Threat Defense"
description: |-
  Manages a Named List.
---

# bloxone_td_named_list (Resource)

Manages a Named List.

## Example Usage

```terraform
resource "bloxone_td_named_list" "example" {
  name = "example_named_list"
  items_described = [
    {
      item        = "tf-domain.com"
      description = "Example Domain"
    }
  ]
  type = "custom_list"

  # Other optional fields
  description = "Example Named List"
  tags = {
    location = "site1"
  }
  threat_level     = "MEDIUM"
  confidence_level = "HIGH"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `items_described` (Attributes List) The List of ItemStructs structure which contains the item and its description (see [below for nested schema](#nestedatt--items_described))
- `name` (String) The name of the named list.

### Optional

- `confidence_level` (String) The confidence level for a custom list. The possible values are ["LOW", "MEDIUM", "HIGH"]
- `description` (String) The brief description for the named list.
- `tags` (Map of String) Enables tag support for resource where tags attribute contains user-defined key value pairs
- `threat_level` (String) The threat level for a custom list. The possible values are ["INFO", "LOW", "MEDIUM", "HIGH"]
- `type` (String) The type of the named list, that can be "custom_list", "threat_insight", "fast_flux", "dga", "dnsm", "threat_insight_nde", "default_allow", "default_block" or "threat_insight_nde".

### Read-Only

- `created_time` (String) The time when this Named List object was created.
- `id` (Number) The Named List object identifier.
- `item_count` (Number) The number of items in this named list.
- `items` (List of String) The list of the FQDN or IPv4/IPv6 CIDRs to define whitelists and blacklists for additional protection.
- `policies` (List of String) The list of the security policy names with which the named list is associated.
- `updated_time` (String) The time when this Named List object was last updated.

<a id="nestedatt--items_described"></a>
### Nested Schema for `items_described`

Optional:

- `description` (String) The description of the item
- `item` (String) The data of the Item
