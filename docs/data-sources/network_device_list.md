---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_device_list Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc.
  You can use the .* in any value to conduct a wildcard search. For example, to find all hostnames beginning with myhost
  in the IP address range 192.25.18.n, issue the following request: GET /dna/intent/api/v1/network-
  device?hostname=myhost.&managementIpAddress=192.25.18..
  If id parameter is provided with comma separated ids, it will return the list of network-devices for the given ids and
  ignores the other request parameters. The API returns a paginated response based on 'limit' and 'offset' parameters,
  allowing up to 500 records per page. 'limit' specifies the number of records, and 'offset' sets the starting point using
  1-based indexing. Use '/dna/intent/api/v1/network-device/count' to get the total record count. For data sets over 500
  records, make multiple calls, adjusting 'limit' and 'offset' to retrieve all records incrementally.
---

# catalystcenter_network_device_list (Data Source)

It performs read operation on Devices.

- Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc.
You can use the .* in any value to conduct a wildcard search. For example, to find all hostnames beginning with myhost
in the IP address range 192.25.18.n, issue the following request: GET /dna/intent/api/v1/network-
device?hostname=myhost.*&managementIpAddress=192.25.18..*
If id parameter is provided with comma separated ids, it will return the list of network-devices for the given ids and
ignores the other request parameters. The API returns a paginated response based on 'limit' and 'offset' parameters,
allowing up to 500 records per page. 'limit' specifies the number of records, and 'offset' sets the starting point using
1-based indexing. Use '/dna/intent/api/v1/network-device/count' to get the total record count. For data sets over 500
records, make multiple calls, adjusting 'limit' and 'offset' to retrieve all records incrementally.

## Example Usage

```terraform
data "catalystcenter_network_device_list" "example" {
  provider                   = catalystcenter
  associated_wlc_ip          = ["string"]
  collection_interval        = ["string"]
  collection_status          = ["string"]
  device_support_level       = "string"
  error_code                 = ["string"]
  error_description          = ["string"]
  family                     = ["string"]
  hostname                   = ["string"]
  id                         = "string"
  license_name               = ["string"]
  license_status             = ["string"]
  license_type               = ["string"]
  limit                      = 1
  location                   = ["string"]
  location_name              = ["string"]
  mac_address                = ["string"]
  management_ip_address      = ["string"]
  module_equpimenttype       = ["string"]
  module_name                = ["string"]
  module_operationstatecode  = ["string"]
  module_partnumber          = ["string"]
  module_servicestate        = ["string"]
  module_vendorequipmenttype = ["string"]
  not_synced_for_minutes     = ["string"]
  offset                     = 1
  platform_id                = ["string"]
  reachability_status        = ["string"]
  role                       = ["string"]
  serial_number              = ["string"]
  series                     = ["string"]
  software_type              = ["string"]
  software_version           = ["string"]
  type                       = ["string"]
  up_time                    = ["string"]
}

output "catalystcenter_network_device_list_example" {
  value = data.catalystcenter_network_device_list.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `associated_wlc_ip` (List of String) associatedWlcIp query parameter.
- `collection_interval` (List of String) collectionInterval query parameter.
- `collection_status` (List of String) collectionStatus query parameter.
- `device_support_level` (String) deviceSupportLevel query parameter.
- `error_code` (List of String) errorCode query parameter.
- `error_description` (List of String) errorDescription query parameter.
- `family` (List of String) family query parameter.
- `hostname` (List of String) hostname query parameter.
- `id` (String) id query parameter. Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
- `license_name` (List of String) license.name query parameter.
- `license_status` (List of String) license.status query parameter.
- `license_type` (List of String) license.type query parameter.
- `limit` (Number) limit query parameter. The number of records to show for this page. Min: 1, Max: 500
- `location` (List of String) location query parameter.
- `location_name` (List of String) locationName query parameter.
- `mac_address` (List of String) macAddress query parameter.
- `management_ip_address` (List of String) managementIpAddress query parameter.
- `module_equpimenttype` (List of String) module+equpimenttype query parameter.
- `module_name` (List of String) module+name query parameter.
- `module_operationstatecode` (List of String) module+operationstatecode query parameter.
- `module_partnumber` (List of String) module+partnumber query parameter.
- `module_servicestate` (List of String) module+servicestate query parameter.
- `module_vendorequipmenttype` (List of String) module+vendorequipmenttype query parameter.
- `not_synced_for_minutes` (List of String) notSyncedForMinutes query parameter.
- `offset` (Number) offset query parameter. offset >= 1 [X gives results from Xth device onwards]
- `platform_id` (List of String) platformId query parameter.
- `reachability_status` (List of String) reachabilityStatus query parameter.
- `role` (List of String) role query parameter.
- `serial_number` (List of String) serialNumber query parameter.
- `series` (List of String) series query parameter.
- `software_type` (List of String) softwareType query parameter.
- `software_version` (List of String) softwareVersion query parameter.
- `type` (List of String) type query parameter.
- `up_time` (List of String) upTime query parameter.

### Read-Only

- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `ap_ethernet_mac_address` (String)
- `ap_manager_interface_ip` (String)
- `associated_wlc_ip` (String)
- `boot_date_time` (String)
- `collection_interval` (String)
- `collection_status` (String)
- `description` (String)
- `device_support_level` (String)
- `dns_resolved_management_address` (String)
- `error_code` (String)
- `error_description` (String)
- `family` (String)
- `hostname` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `interface_count` (String)
- `inventory_status_detail` (String)
- `last_device_resync_start_time` (String)
- `last_managed_resync_reasons` (String)
- `last_update_time` (Number)
- `last_updated` (String)
- `line_card_count` (String)
- `line_card_id` (String)
- `location` (String)
- `location_name` (String)
- `mac_address` (String)
- `managed_atleast_once` (String)
- `management_ip_address` (String)
- `management_state` (String)
- `memory_size` (String)
- `pending_sync_requests_count` (String)
- `platform_id` (String)
- `reachability_failure_reason` (String)
- `reachability_status` (String)
- `reasons_for_device_resync` (String)
- `reasons_for_pending_sync_requests` (String)
- `role` (String)
- `role_source` (String)
- `serial_number` (String)
- `series` (String)
- `snmp_contact` (String)
- `snmp_location` (String)
- `software_type` (String)
- `software_version` (String)
- `sync_requested_by_app` (String)
- `tag_count` (String)
- `tunnel_udp_port` (String)
- `type` (String)
- `up_time` (String)
- `uptime_seconds` (Number)
- `waas_device_mode` (String)
