---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_td_named_lists Data Source - terraform-provider-bloxone"
subcategory: "Threat Defense"
description: |-
  Retrieves information about existing Named Lists.
---

# bloxone_td_named_lists (Data Source)

Retrieves information about existing Named Lists.

## Example Usage

```terraform
# Get Named Lists filtered by tag
data "bloxone_td_named_lists" "example_by_tag" {
  tag_filters = {
    location = "site1"
  }
}

# Get all Named Lists
data "bloxone_td_named_lists" "example_all" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filters` (Map of String) Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.
- `tag_filters` (Map of String) Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.

### Read-Only

- `results` (Attributes List) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Required:

- `items_described` (Attributes List) The List of ItemStructs structure which contains the item and its description (see [below for nested schema](#nestedatt--results--items_described))
- `name` (String) The name of the named list.

Optional:

- `confidence_level` (String) The confidence level for a custom list. The possible values are ["LOW", "MEDIUM", "HIGH"]
- `description` (String) The brief description for the named list.
- `tags` (Map of String) Enables tag support for resource where tags attribute contains user-defined key value pairs
- `threat_level` (String) The threat level for a custom list. The possible values are ["INFO", "LOW", "MEDIUM", "HIGH"]
- `type` (String) The type of the named list, that can be "custom_list", "threat_insight", "fast_flux", "dga", "dnsm", "threat_insight_nde", "default_allow", "default_block" or "threat_insight_nde".

Read-Only:

- `created_time` (String) The time when this Named List object was created.
- `id` (Number) The Named List object identifier.
- `item_count` (Number) The number of items in this named list.
- `items` (List of String) The list of the FQDN or IPv4/IPv6 CIDRs to define whitelists and blacklists for additional protection.
- `policies` (List of String) The list of the security policy names with which the named list is associated.
- `updated_time` (String) The time when this Named List object was last updated.

<a id="nestedatt--results--items_described"></a>
### Nested Schema for `results.items_described`

Optional:

- `description` (String) The description of the item
- `item` (String) The data of the Item
