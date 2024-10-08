---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_device_global_polling_interval Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Returns polling interval of all devices
---

# catalystcenter_network_device_global_polling_interval (Data Source)

It performs read operation on Devices.

- Returns polling interval of all devices

## Example Usage

```terraform
data "catalystcenter_network_device_global_polling_interval" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_device_global_polling_interval_example" {
  value = data.catalystcenter_network_device_global_polling_interval.example.item
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

- `response` (Number)
- `version` (String)
