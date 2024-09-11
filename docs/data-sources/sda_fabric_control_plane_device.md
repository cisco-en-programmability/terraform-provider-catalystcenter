---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_fabric_control_plane_device Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Get control plane device from SDA Fabric
---

# catalystcenter_sda_fabric_control_plane_device (Data Source)

It performs read operation on SDA.

- Get control plane device from SDA Fabric

## Example Usage

```terraform
data "catalystcenter_sda_fabric_control_plane_device" "example" {
  provider                     = catalystcenter
  device_management_ip_address = "string"
}

output "catalystcenter_sda_fabric_control_plane_device_example" {
  value = data.catalystcenter_sda_fabric_control_plane_device.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_management_ip_address` (String) deviceManagementIpAddress query parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `description` (String)
- `device_management_ip_address` (String)
- `device_name` (String)
- `roles` (String)
- `route_distribution_protocol` (String)
- `site_name_hierarchy` (String)
- `status` (String)