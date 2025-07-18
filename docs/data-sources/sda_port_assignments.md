---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_port_assignments Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns a list of port assignments that match the provided query parameters.
---

# catalystcenter_sda_port_assignments (Data Source)

It performs read operation on SDA.

- Returns a list of port assignments that match the provided query parameters.

## Example Usage

```terraform
data "catalystcenter_sda_port_assignments" "example" {
  provider          = catalystcenter
  data_vlan_name    = "string"
  fabric_id         = "string"
  interface_name    = "string"
  limit             = 1
  native_vlan_id    = 1.0
  network_device_id = "string"
  offset            = 1
  voice_vlan_name   = "string"
}

output "catalystcenter_sda_port_assignments_example" {
  value = data.catalystcenter_sda_port_assignments.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `data_vlan_name` (String) dataVlanName query parameter. Data VLAN name of the port assignment.
- `fabric_id` (String) fabricId query parameter. ID of the fabric the device is assigned to.
- `interface_name` (String) interfaceName query parameter. Interface name of the port assignment.
- `limit` (Number) limit query parameter. Maximum number of records to return. The maximum number of objects supported in a single request is 500.
- `native_vlan_id` (Number) nativeVlanId query parameter. Native VLAN of the port assignment, this option is only applicable to TRUNKING_DEVICE connectedDeviceType.(VLAN must be between 1 and 4094. In cases value not set when connectedDeviceType is TRUNKING_DEVICE, default value will be '1').
- `network_device_id` (String) networkDeviceId query parameter. Network device ID of the port assignment.
- `offset` (Number) offset query parameter. Starting record for pagination.
- `voice_vlan_name` (String) voiceVlanName query parameter. Voice VLAN name of the port assignment.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `allowed_vlan_ranges` (String)
- `authenticate_template_name` (String)
- `connected_device_type` (String)
- `data_vlan_name` (String)
- `fabric_id` (String)
- `id` (String)
- `interface_description` (String)
- `interface_name` (String)
- `native_vlan_id` (Number)
- `network_device_id` (String)
- `security_group_name` (String)
- `voice_vlan_name` (String)
