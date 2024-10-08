---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_site_count_v2 Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Sites.
  Get the site count of the specified site's sub-hierarchy (inclusive of the provided site)
---

# catalystcenter_site_count_v2 (Data Source)

It performs read operation on Sites.

- Get the site count of the specified site's sub-hierarchy (inclusive of the provided site)

## Example Usage

```terraform
data "catalystcenter_site_count_v2" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_site_count_v2_example" {
  value = data.catalystcenter_site_count_v2.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id query parameter. Site instance UUID

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `response` (Number)
- `version` (String)
