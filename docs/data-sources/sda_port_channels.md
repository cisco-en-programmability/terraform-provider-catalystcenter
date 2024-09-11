---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_port_channels Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns a list of port channels that match the provided query parameters.
---

# catalystcenter_sda_port_channels (Data Source)

It performs read operation on SDA.

- Returns a list of port channels that match the provided query parameters.

## Example Usage

```terraform
data "catalystcenter_sda_port_channels" "example" {
  provider              = catalystcenter
  connected_device_type = "string"
  fabric_id             = "string"
  limit                 = 1
  network_device_id     = "string"
  offset                = 1
  port_channel_name     = "string"
}

output "catalystcenter_sda_port_channels_example" {
  value = data.catalystcenter_sda_port_channels.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `connected_device_type` (String) connectedDeviceType query parameter. Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
- `fabric_id` (String) fabricId query parameter. ID of the fabric the device is assigned to.
- `limit` (Number) limit query parameter. Maximum number of records to return.
- `network_device_id` (String) networkDeviceId query parameter. ID of the network device.
- `offset` (Number) offset query parameter. Starting record for pagination.
- `port_channel_name` (String) portChannelName query parameter. Name of the port channel.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `connected_device_type` (String)
- `description` (String)
- `fabric_id` (String)
- `id` (String)
- `interface_names` (List of String)
- `network_device_id` (String)
- `port_channel_name` (String)
- `protocol` (String)