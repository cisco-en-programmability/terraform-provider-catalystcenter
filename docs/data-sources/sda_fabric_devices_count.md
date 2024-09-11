---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_fabric_devices_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns the count of fabric devices that match the provided query parameters.
---

# catalystcenter_sda_fabric_devices_count (Data Source)

It performs read operation on SDA.

- Returns the count of fabric devices that match the provided query parameters.

## Example Usage

```terraform
data "catalystcenter_sda_fabric_devices_count" "example" {
  provider          = catalystcenter
  device_roles      = "string"
  fabric_id         = "string"
  network_device_id = "string"
}

output "catalystcenter_sda_fabric_devices_count_example" {
  value = data.catalystcenter_sda_fabric_devices_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `fabric_id` (String) fabricId query parameter. ID of the fabric this device belongs to.

### Optional

- `device_roles` (String) deviceRoles query parameter. Device roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
- `network_device_id` (String) networkDeviceId query parameter. Network device ID of the fabric device.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)