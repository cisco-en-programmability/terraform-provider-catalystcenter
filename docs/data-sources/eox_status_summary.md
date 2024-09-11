---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_eox_status_summary Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on EoX.
  Retrieves EoX summary for all devices in the network
---

# catalystcenter_eox_status_summary (Data Source)

It performs read operation on EoX.

- Retrieves EoX summary for all devices in the network

## Example Usage

```terraform
data "catalystcenter_eox_status_summary" "example" {
  provider = catalystcenter
}

output "catalystcenter_eox_status_summary_example" {
  value = data.catalystcenter_eox_status_summary.example.item
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

- `hardware_count` (Number)
- `module_count` (Number)
- `software_count` (Number)
- `total_count` (Number)