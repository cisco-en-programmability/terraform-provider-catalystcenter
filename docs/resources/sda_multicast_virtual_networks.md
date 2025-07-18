---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_multicast_virtual_networks Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on SDA.
  Adds multicast for virtual networks based on user input.Updates multicast configurations for virtual networks based on user input.Deletes a multicast configuration for a virtual network based on id.
---

# catalystcenter_sda_multicast_virtual_networks (Resource)

It manages create, read, update and delete operations on SDA.

- Adds multicast for virtual networks based on user input.

- Updates multicast configurations for virtual networks based on user input.

- Deletes a multicast configuration for a virtual network based on id.

## Example Usage

```terraform
resource "catalystcenter_sda_multicast_virtual_networks" "example" {
  provider = catalystcenter
 
  parameters {

    fabric_id       = "string"
    id              = "string"
    ip_pool_name    = "string"
    ipv4_ssm_ranges = ["string"]
    multicast_r_ps {

      ipv4_address       = "string"
      ipv4_asm_ranges    = ["string"]
      ipv6_address       = "string"
      ipv6_asm_ranges    = ["string"]
      is_default_v4_rp   = "false"
      is_default_v6_rp   = "false"
      network_device_ids = ["string"]
      rp_device_location = "string"
    }
    virtual_network_name = "string"
  }
}

output "catalystcenter_sda_multicast_virtual_networks_example" {
  value = catalystcenter_sda_multicast_virtual_networks.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) Array of RequestSdaAddMulticastVirtualNetworks (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `payload` (Block List) Array of RequestSdaAddMulticastVirtualNetworks (see [below for nested schema](#nestedblock--parameters--payload))

<a id="nestedblock--parameters--payload"></a>
### Nested Schema for `parameters.payload`

Optional:

- `fabric_id` (String) ID of the fabric site this multicast configuration is associated with.
- `id` (String) ID of the multicast configuration (updating this field is not allowed).
- `ip_pool_name` (String) Name of the IP Pool associated with the fabric site.
- `ipv4_ssm_ranges` (List of String) IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.
- `multicast_r_ps` (Block List) (see [below for nested schema](#nestedblock--parameters--payload--multicast_r_ps))
- `virtual_network_name` (String) Name of the virtual network associated with the fabric site.

<a id="nestedblock--parameters--payload--multicast_r_ps"></a>
### Nested Schema for `parameters.payload.multicast_r_ps`

Optional:

- `ipv4_address` (String) IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.
- `ipv4_asm_ranges` (List of String) IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
- `ipv6_address` (String) IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.
- `ipv6_asm_ranges` (List of String) IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
- `is_default_v4_rp` (String) Specifies whether it is a default IPv4 RP.
- `is_default_v6_rp` (String) Specifies whether it is a default IPv6 RP.
- `network_device_ids` (List of String) IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.
- `rp_device_location` (String) Device location of the RP.




<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `fabric_id` (String)
- `id` (String)
- `ip_pool_name` (String)
- `ipv4_ssm_ranges` (List of String)
- `multicast_r_ps` (List of Object) (see [below for nested schema](#nestedobjatt--item--multicast_r_ps))
- `virtual_network_name` (String)

<a id="nestedobjatt--item--multicast_r_ps"></a>
### Nested Schema for `item.multicast_r_ps`

Read-Only:

- `ipv4_address` (String)
- `ipv4_asm_ranges` (List of String)
- `ipv6_address` (String)
- `ipv6_asm_ranges` (List of String)
- `is_default_v4_rp` (String)
- `is_default_v6_rp` (String)
- `network_device_ids` (List of String)
- `rp_device_location` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_sda_multicast_virtual_networks.example "id:=string"
```
