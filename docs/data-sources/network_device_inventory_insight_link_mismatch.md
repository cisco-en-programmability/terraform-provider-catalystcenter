---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_device_inventory_insight_link_mismatch Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Find all devices with link mismatch (speed /  vlan)
---

# catalystcenter_network_device_inventory_insight_link_mismatch (Data Source)

It performs read operation on Devices.

- Find all devices with link mismatch (speed /  vlan)

## Example Usage

```terraform
data "catalystcenter_network_device_inventory_insight_link_mismatch" "example" {
  provider = catalystcenter
  category = "string"
  limit    = 1
  offset   = 1
  order    = "string"
  site_id  = "string"
  sort_by  = "string"
}

output "catalystcenter_network_device_inventory_insight_link_mismatch_example" {
  value = data.catalystcenter_network_device_inventory_insight_link_mismatch.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `category` (String) category query parameter. Links mismatch category.  Value can be speed-duplex or vlan.
- `site_id` (String) siteId path parameter.

### Optional

- `limit` (Number) limit query parameter. Default value is 500
- `offset` (Number) offset query parameter. Row Number.  Default value is 1
- `order` (String) order query parameter. Order.  Value can be asc or desc.  Default value is asc
- `sort_by` (String) sortBy query parameter. Sort By

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `avg_update_frequency` (Number)
- `end_device_host_name` (String)
- `end_device_id` (String)
- `end_device_ip_address` (String)
- `end_port_address` (String)
- `end_port_allowed_vlan_ids` (String)
- `end_port_duplex` (String)
- `end_port_id` (String)
- `end_port_mask` (String)
- `end_port_name` (String)
- `end_port_native_vlan_id` (String)
- `end_port_pep_id` (String)
- `end_port_speed` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `last_updated` (String)
- `link_status` (String)
- `num_updates` (Number)
- `start_device_host_name` (String)
- `start_device_id` (String)
- `start_device_ip_address` (String)
- `start_port_address` (String)
- `start_port_allowed_vlan_ids` (String)
- `start_port_duplex` (String)
- `start_port_id` (String)
- `start_port_mask` (String)
- `start_port_name` (String)
- `start_port_native_vlan_id` (String)
- `start_port_pep_id` (String)
- `start_port_speed` (String)
- `type` (String)