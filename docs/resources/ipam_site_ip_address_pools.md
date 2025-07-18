---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_ipam_site_ip_address_pools Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create and read operations on Network Settings.
  Reserves (creates) an IP address subpool, which reserves address space from a global pool (or global pools) for a
  particular site (and it's child sites). A subpool must be either an IPv4 or dual-stack pool, with ipV4AddressSpace
  and optionally ipV6AddressSpace properties specified.
---

# catalystcenter_ipam_site_ip_address_pools (Resource)

It manages create and read operations on Network Settings.

- Reserves (creates) an IP address subpool, which reserves address space from a global pool (or global pools) for a
particular site (and it's child sites). A subpool must be either an IPv4 or dual-stack pool, with **ipV4AddressSpace**
and optionally **ipV6AddressSpace** properties specified.

## Example Usage

```terraform
resource "catalystcenter_ipam_site_ip_address_pools" "example" {
  provider = catalystcenter
 
  parameters {

    ip_v4_address_space {

      assigned_addresses         = "string"
      default_assigned_addresses = "string"
      dhcp_servers               = ["string"]
      dns_servers                = ["string"]
      gateway_ip_address         = "string"
      global_pool_id             = "string"
      prefix_length              = 1.0
      slaac_support              = "false"
      subnet                     = "string"
      total_addresses            = "string"
      unassignable_addresses     = "string"
    }
    ip_v6_address_space {

      assigned_addresses         = "string"
      default_assigned_addresses = "string"
      dhcp_servers               = ["string"]
      dns_servers                = ["string"]
      gateway_ip_address         = "string"
      global_pool_id             = "string"
      prefix_length              = 1.0
      slaac_support              = "false"
      subnet                     = "string"
      total_addresses            = "string"
      unassignable_addresses     = "string"
    }
    name      = "string"
    pool_type = "string"
    site_id   = "string"
    site_name = "string"
  }
}

output "catalystcenter_ipam_site_ip_address_pools_example" {
  value = catalystcenter_ipam_site_ip_address_pools.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `ip_v4_address_space` (Block List) (see [below for nested schema](#nestedblock--parameters--ip_v4_address_space))
- `ip_v6_address_space` (Block List) (see [below for nested schema](#nestedblock--parameters--ip_v6_address_space))
- `name` (String) The name for this reserve IP pool. Only letters, numbers, '-' (hyphen), '_' (underscore), '.' (period), and '/' (forward slash) are allowed.
- `pool_type` (String) Once created, a subpool type cannot be changed.  LAN: Assigns IP addresses to LAN interfaces of applicable VNFs and underlay LAN automation.  Management: Assigns IP addresses to management interfaces. A management network is a dedicated network connected to VNFs for VNF management.  Service: Assigns IP addresses to service interfaces. Service networks are used for communication within VNFs.  WAN: Assigns IP addresses to NFVIS for UCS-E provisioning.  Generic: used for all other network types.
- `site_id` (String) The **id** of the site that this subpool belongs to. This must be the **id** of a non-Global site.
- `site_name` (String) The name of the site that this subpool belongs to.

<a id="nestedblock--parameters--ip_v4_address_space"></a>
### Nested Schema for `parameters.ip_v4_address_space`

Optional:

- `assigned_addresses` (String) The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
- `default_assigned_addresses` (String) The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
- `dhcp_servers` (List of String) The DHCP server(s) for this subnet.
- `dns_servers` (List of String) The DNS server(s) for this subnet.
- `gateway_ip_address` (String) The gateway IP address for this subnet.
- `global_pool_id` (String) The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
- `prefix_length` (Number) The network mask component, as a decimal, for the CIDR notation of this subnet.
- `slaac_support` (String) If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.
- `subnet` (String) The IP address component of the CIDR notation for this subnet.
- `total_addresses` (String) The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
- `unassignable_addresses` (String) The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.


<a id="nestedblock--parameters--ip_v6_address_space"></a>
### Nested Schema for `parameters.ip_v6_address_space`

Optional:

- `assigned_addresses` (String) The number of addresses assigned from the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
- `default_assigned_addresses` (String) The number of addresses that are assigned from the pool by default. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
- `dhcp_servers` (List of String) The DHCP server(s) for this subnet.
- `dns_servers` (List of String) The DNS server(s) for this subnet.
- `gateway_ip_address` (String) The gateway IP address for this subnet.
- `global_pool_id` (String) The non-tunnel global pool for this reserve pool (which matches this IP address type). Once added this value cannot be changed.
- `prefix_length` (Number) The network mask component, as a decimal, for the CIDR notation of this subnet.
- `slaac_support` (String) If the prefixLength is 64, this option may be enabled. Stateless Address Auto-configuration (SLAAC) allows network devices to select their IP address without the use of DHCP servers.
- `subnet` (String) The IP address component of the CIDR notation for this subnet.
- `total_addresses` (String) The total number of addresses in the pool. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.
- `unassignable_addresses` (String) The number of addresses in the pool that cannot be assigned. This is a numeric string; since IPv6 address spaces are 128 bits in size, presume this string has a value up to 128 bits for IPv6 address spaces and 32 bits for IPv4 address spaces.



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String)
- `ip_v4_address_space` (List of Object) (see [below for nested schema](#nestedobjatt--item--ip_v4_address_space))
- `ip_v6_address_space` (List of Object) (see [below for nested schema](#nestedobjatt--item--ip_v6_address_space))
- `name` (String)
- `pool_type` (String)
- `site_id` (String)
- `site_name` (String)

<a id="nestedobjatt--item--ip_v4_address_space"></a>
### Nested Schema for `item.ip_v4_address_space`

Read-Only:

- `assigned_addresses` (String)
- `default_assigned_addresses` (String)
- `dhcp_servers` (List of String)
- `dns_servers` (List of String)
- `gateway_ip_address` (String)
- `global_pool_id` (String)
- `prefix_length` (Number)
- `slaac_support` (String)
- `subnet` (String)
- `total_addresses` (String)
- `unassignable_addresses` (String)


<a id="nestedobjatt--item--ip_v6_address_space"></a>
### Nested Schema for `item.ip_v6_address_space`

Read-Only:

- `assigned_addresses` (String)
- `default_assigned_addresses` (String)
- `dhcp_servers` (List of String)
- `dns_servers` (List of String)
- `gateway_ip_address` (String)
- `global_pool_id` (String)
- `prefix_length` (Number)
- `slaac_support` (String)
- `subnet` (String)
- `total_addresses` (String)
- `unassignable_addresses` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_ipam_site_ip_address_pools.example "id:=string"
```
