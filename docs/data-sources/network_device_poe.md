---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_device_poe Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Returns POE details for device.
---

# catalystcenter_network_device_poe (Data Source)

It performs read operation on Devices.

- Returns POE details for device.

## Example Usage

```terraform
data "catalystcenter_network_device_poe" "example" {
  provider    = catalystcenter
  device_uuid = "string"
}

output "catalystcenter_network_device_poe_example" {
  value = data.catalystcenter_network_device_poe.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_uuid` (String) deviceUuid path parameter. UUID of the device

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `power_allocated` (String)
- `power_consumed` (String)
- `power_remaining` (String)