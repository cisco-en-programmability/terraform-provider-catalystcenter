---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_applications_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Application Policy.
  Get the number of all existing applications
---

# catalystcenter_applications_count (Data Source)

It performs read operation on Application Policy.

- Get the number of all existing applications

## Example Usage

```terraform
data "catalystcenter_applications_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_applications_count_example" {
  value = data.catalystcenter_applications_count.example.item
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
