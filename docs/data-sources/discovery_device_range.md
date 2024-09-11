---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_discovery_device_range Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Discovery.
  Returns the network devices discovered for the given discovery and for the given range. The maximum number of records
  that can be retrieved is 500. Discovery ID can be obtained using the "Get Discoveries by range" API.
---

# catalystcenter_discovery_device_range (Data Source)

It performs read operation on Discovery.

- Returns the network devices discovered for the given discovery and for the given range. The maximum number of records
that can be retrieved is 500. Discovery ID can be obtained using the "Get Discoveries by range" API.

## Example Usage

```terraform
data "catalystcenter_discovery_device_range" "example" {
  provider          = catalystcenter
  id                = "string"
  records_to_return = 1
  start_index       = 1
  task_id           = "string"
}

output "catalystcenter_discovery_device_range_example" {
  value = data.catalystcenter_discovery_device_range.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. Discovery ID
- `records_to_return` (Number) recordsToReturn path parameter. Number of records to fetch from the start index
- `start_index` (Number) startIndex path parameter. Starting index for the records

### Optional

- `task_id` (String) taskId query parameter.

### Read-Only

- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `anchor_wlc_for_ap` (String)
- `auth_model_id` (String)
- `avg_update_frequency` (Number)
- `boot_date_time` (String)
- `clistatus` (String)
- `duplicate_device_id` (String)
- `error_code` (String)
- `error_description` (String)
- `family` (String)
- `hostname` (String)
- `http_status` (String)
- `id` (String)
- `image_name` (String)
- `ingress_queue_config` (String)
- `interface_count` (String)
- `inventory_collection_status` (String)
- `inventory_reachability_status` (String)
- `last_updated` (String)
- `line_card_count` (String)
- `line_card_id` (String)
- `location` (String)
- `location_name` (String)
- `mac_address` (String)
- `management_ip_address` (String)
- `memory_size` (String)
- `netconf_status` (String)
- `num_updates` (Number)
- `ping_status` (String)
- `platform_id` (String)
- `port_range` (String)
- `qos_status` (String)
- `reachability_failure_reason` (String)
- `reachability_status` (String)
- `role` (String)
- `role_source` (String)
- `serial_number` (String)
- `snmp_contact` (String)
- `snmp_location` (String)
- `snmp_status` (String)
- `software_version` (String)
- `tag` (String)
- `tag_count` (Number)
- `type` (String)
- `up_time` (String)
- `vendor` (String)
- `wlc_ap_device_status` (String)