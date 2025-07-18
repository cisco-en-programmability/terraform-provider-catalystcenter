---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_diagnostic_tasks_id Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Health and Performance.
  This data source retrieves the diagnostic task identified by the specified id.
---

# catalystcenter_diagnostic_tasks_id (Data Source)

It performs read operation on Health and Performance.

- This data source retrieves the diagnostic task identified by the specified **id**.

## Example Usage

```terraform
data "catalystcenter_diagnostic_tasks_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_diagnostic_tasks_id_example" {
  value = data.catalystcenter_diagnostic_tasks_id.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. The **id** of the diagnostic task to be retrieved

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `end_time` (Number)
- `id` (String)
- `parent_id` (String)
- `result_location` (String)
- `root_id` (String)
- `start_time` (Number)
- `status` (String)
- `updated_time` (Number)
