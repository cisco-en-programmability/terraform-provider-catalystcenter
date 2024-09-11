---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_applications_count_v2 Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Application Policy.
  Get the number of all existing applications
---

# catalystcenter_applications_count_v2 (Data Source)

It performs read operation on Application Policy.

- Get the number of all existing applications

## Example Usage

```terraform
data "catalystcenter_applications_count_v2" "example" {
  provider            = catalystcenter
  scalable_group_type = "string"
}

output "catalystcenter_applications_count_v2_example" {
  value = data.catalystcenter_applications_count_v2.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `scalable_group_type` (String) scalableGroupType query parameter. scalable group type to retrieve, valid value APPLICATION

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `response` (Number)
- `version` (String)