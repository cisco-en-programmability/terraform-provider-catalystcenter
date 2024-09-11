---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_multicast_virtual_networks Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Returns a list of multicast configurations for virtual networks that match the provided query parameters.
---

# catalystcenter_sda_multicast_virtual_networks (Data Source)

It performs read operation on SDA.

- Returns a list of multicast configurations for virtual networks that match the provided query parameters.

## Example Usage

```terraform
data "catalystcenter_sda_multicast_virtual_networks" "example" {
  provider             = catalystcenter
  fabric_id            = "string"
  limit                = 1
  offset               = 1
  virtual_network_name = "string"
}

output "catalystcenter_sda_multicast_virtual_networks_example" {
  value = data.catalystcenter_sda_multicast_virtual_networks.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `fabric_id` (String) fabricId query parameter. ID of the fabric site where multicast is configured.
- `limit` (Number) limit query parameter. Maximum number of records to return.
- `offset` (Number) offset query parameter. Starting record for pagination.
- `virtual_network_name` (String) virtualNetworkName query parameter. Name of the virtual network associated to the multicast configuration.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `fabric_id` (String)
- `id` (String)
- `ip_pool_name` (String)
- `ipv4_ssm_ranges` (List of String)
- `multicast_r_ps` (List of Object) (see [below for nested schema](#nestedobjatt--items--multicast_r_ps))
- `virtual_network_name` (String)

<a id="nestedobjatt--items--multicast_r_ps"></a>
### Nested Schema for `items.multicast_r_ps`

Read-Only:

- `ipv4_address` (String)
- `ipv4_asm_ranges` (List of String)
- `ipv6_address` (String)
- `ipv6_asm_ranges` (List of String)
- `is_default_v4_rp` (String)
- `is_default_v6_rp` (String)
- `network_device_ids` (List of String)
- `rp_device_location` (String)