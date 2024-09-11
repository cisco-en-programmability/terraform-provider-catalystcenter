---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_port_assignment_for_access_point Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Get Port assignment for access point in SDA Fabric
---

# catalystcenter_sda_port_assignment_for_access_point (Data Source)

It performs read operation on SDA.

- Get Port assignment for access point in SDA Fabric

## Example Usage

```terraform
data "catalystcenter_sda_port_assignment_for_access_point" "example" {
  provider                     = catalystcenter
  device_management_ip_address = "string"
  interface_name               = "string"
}

output "catalystcenter_sda_port_assignment_for_access_point_example" {
  value = data.catalystcenter_sda_port_assignment_for_access_point.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_management_ip_address` (String) deviceManagementIpAddress query parameter.
- `interface_name` (String) interfaceName query parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `authenticate_template_name` (String)
- `data_ip_address_pool_name` (String)
- `description` (String)
- `device_management_ip_address` (String)
- `interface_name` (String)
- `scalable_group_name` (String)
- `site_name_hierarchy` (String)
- `status` (String)
- `voice_ip_address_pool_name` (String)