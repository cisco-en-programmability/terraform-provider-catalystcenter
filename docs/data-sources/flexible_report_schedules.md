---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_flexible_report_schedules Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Reports.
  Get all flexible report schedules
---

# catalystcenter_flexible_report_schedules (Data Source)

It performs read operation on Reports.

- Get all flexible report schedules

## Example Usage

```terraform
data "catalystcenter_flexible_report_schedules" "example" {
  provider = catalystcenter
}

output "catalystcenter_flexible_report_schedules_example" {
  value = data.catalystcenter_flexible_report_schedules.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `report_id` (String)
- `report_name` (String)
- `schedule` (String)