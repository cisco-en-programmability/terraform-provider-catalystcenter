---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_application_sets_count_v2 Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Application Policy.
  Get the number of all existing application sets
---

# catalystcenter_application_sets_count_v2 (Data Source)

It performs read operation on Application Policy.

- Get the number of all existing application sets

## Example Usage

```terraform
data "catalystcenter_application_sets_count_v2" "example" {
  provider            = catalystcenter
  scalable_group_type = "string"
}

output "catalystcenter_application_sets_count_v2_example" {
  value = data.catalystcenter_application_sets_count_v2.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `scalable_group_type` (String) scalableGroupType query parameter. Scalable group type to retrieve, valid value APPLICATION_GROUP

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `response` (Number)
- `version` (String)
