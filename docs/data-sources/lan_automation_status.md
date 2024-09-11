---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_lan_automation_status Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on LAN Automation.
  Invoke this API to get the LAN Automation session status.Invoke this API to get the LAN Automation session status based on the given Lan Automation session id.
---

# catalystcenter_lan_automation_status (Data Source)

It performs read operation on LAN Automation.

- Invoke this API to get the LAN Automation session status.

- Invoke this API to get the LAN Automation session status based on the given Lan Automation session id.

## Example Usage

```terraform
data "catalystcenter_lan_automation_status" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
}

output "catalystcenter_lan_automation_status_example" {
  value = data.catalystcenter_lan_automation_status.example.items
}

data "catalystcenter_lan_automation_status" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_lan_automation_status_example" {
  value = data.catalystcenter_lan_automation_status.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter. LAN Automation session identifier.
- `limit` (Number) limit query parameter. Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
- `offset` (Number) offset query parameter. Starting index of the LAN Automation session. Minimum value is 1.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `action` (String)
- `creation_time` (String)
- `discovered_device_list` (List of Object) (see [below for nested schema](#nestedobjatt--item--discovered_device_list))
- `discovered_device_site_name_hierarchy` (String)
- `discovery_devices` (List of String)
- `discovery_level` (Number)
- `discovery_timeout` (Number)
- `id` (String)
- `ip_pools` (List of Object) (see [below for nested schema](#nestedobjatt--item--ip_pools))
- `multicast_enabled` (String)
- `peer_device_managment_ipaddress` (String)
- `primary_device_interface_names` (List of String)
- `primary_device_managment_ipaddress` (String)
- `redistribute_isis_to_bgp` (String)
- `status` (String)

<a id="nestedobjatt--item--discovered_device_list"></a>
### Nested Schema for `item.discovered_device_list`

Read-Only:

- `ip_address_in_use_list` (List of String)
- `name` (String)
- `serial_number` (String)
- `state` (String)


<a id="nestedobjatt--item--ip_pools"></a>
### Nested Schema for `item.ip_pools`

Read-Only:

- `ip_pool_name` (String)
- `ip_pool_role` (String)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `action` (String)
- `creation_time` (String)
- `discovered_device_list` (List of Object) (see [below for nested schema](#nestedobjatt--items--discovered_device_list))
- `discovered_device_site_name_hierarchy` (String)
- `discovery_devices` (List of String)
- `discovery_level` (Number)
- `discovery_timeout` (Number)
- `id` (String)
- `ip_pools` (List of Object) (see [below for nested schema](#nestedobjatt--items--ip_pools))
- `multicast_enabled` (String)
- `peer_device_managment_ipaddress` (String)
- `primary_device_interface_names` (List of String)
- `primary_device_managment_ipaddress` (String)
- `redistribute_isis_to_bgp` (String)
- `status` (String)

<a id="nestedobjatt--items--discovered_device_list"></a>
### Nested Schema for `items.discovered_device_list`

Read-Only:

- `ip_address_in_use_list` (List of String)
- `name` (String)
- `serial_number` (String)
- `state` (String)


<a id="nestedobjatt--items--ip_pools"></a>
### Nested Schema for `items.ip_pools`

Read-Only:

- `ip_pool_name` (String)
- `ip_pool_role` (String)