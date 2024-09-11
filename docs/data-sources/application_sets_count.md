---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_application_sets_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Application Policy.
  Get the number of existing application-sets
---

# catalystcenter_application_sets_count (Data Source)

It performs read operation on Application Policy.

- Get the number of existing application-sets

## Example Usage

```terraform
data "catalystcenter_application_sets_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_application_sets_count_example" {
  value = data.catalystcenter_application_sets_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `response` (String)
- `version` (String)