---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_device_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Returns the interface count for the given deviceReturns the count of network devices based on the filter criteria by management IP address, mac address, hostname and
  location name
---

# catalystcenter_network_device_count (Data Source)

It performs read operation on Devices.

- Returns the interface count for the given device

- Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and
location name

## Example Usage

```terraform
data "catalystcenter_network_device_count" "example" {
  provider  = catalystcenter
  device_id = "string"
}

output "catalystcenter_network_device_count_example" {
  value = data.catalystcenter_network_device_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device_id` (String) deviceId path parameter. Device ID
- `hostname` (List of String) hostname query parameter.
- `location_name` (List of String) locationName query parameter.
- `mac_address` (List of String) macAddress query parameter.
- `management_ip_address` (List of String) managementIpAddress query parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `response` (Number)
- `version` (String)