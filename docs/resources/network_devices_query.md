---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_devices_query Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Devices.
  Gets the list of Network Devices based on the provided complex filters and aggregation functions. For detailed
  information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
  programmability/catalyst-center-api-specs/blob/main/Assurance/CECatCenter_Org-
  AssuranceNetworkDevices-2.0.1-resolved.yaml
---

# catalystcenter_network_devices_query (Resource)

It performs create operation on Devices.

- Gets the list of Network Devices based on the provided complex filters and aggregation functions. For detailed
information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceNetworkDevices-2.0.1-resolved.yaml


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_network_devices_query" "example" {
  provider = catalystcenter
  parameters = [{

    aggregate_attributes = [{

      function = "string"
      name     = "string"
    }]
    attributes = ["string"]
    end_time   = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = "string"
    }]
    page = [{

      count   = 1
      limit   = 1
      offset  = 1
      sort_by = "string"
    }]
    start_time = 1
    views      = ["string"]
  }]
}

output "catalystcenter_network_devices_query_example" {
  value = catalystcenter_network_devices_query.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `aggregate_attributes` (Block List) (see [below for nested schema](#nestedblock--parameters--aggregate_attributes))
- `attributes` (List of String) Attributes
- `end_time` (Number) End Time
- `filters` (Block List) (see [below for nested schema](#nestedblock--parameters--filters))
- `page` (Block List) (see [below for nested schema](#nestedblock--parameters--page))
- `start_time` (Number) Start Time
- `views` (List of String) Views

Read-Only:

- `items` (List of Object) (see [below for nested schema](#nestedatt--parameters--items))

<a id="nestedblock--parameters--aggregate_attributes"></a>
### Nested Schema for `parameters.aggregate_attributes`

Optional:

- `function` (String) Function
- `name` (String) Name


<a id="nestedblock--parameters--filters"></a>
### Nested Schema for `parameters.filters`

Optional:

- `key` (String) Key
- `operator` (String) Operator
- `value` (String) Value


<a id="nestedblock--parameters--page"></a>
### Nested Schema for `parameters.page`

Optional:

- `count` (Number) Count
- `limit` (Number) Limit
- `offset` (Number) Offset
- `sort_by` (String) Sort By


<a id="nestedatt--parameters--items"></a>
### Nested Schema for `parameters.items`

Read-Only:

- `aggregate_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--aggregate_attributes))
- `ap_details` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--ap_details))
- `client_count` (Number)
- `collection_status` (String)
- `communication_state` (String)
- `device_family` (String)
- `device_group_hierarchy_id` (String)
- `device_role` (String)
- `device_series` (String)
- `device_type` (String)
- `fabric_details` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--fabric_details))
- `fabric_metrics_details` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--fabric_metrics_details))
- `feature_flag_list` (List of String)
- `ha_last_reset_reason` (String)
- `ha_status` (String)
- `id` (String)
- `ipv4_address` (String)
- `ipv6_address` (String)
- `last_boot_time` (Number)
- `mac_address` (String)
- `maintenance_mode_enabled` (String)
- `management_ip_address` (String)
- `metrics_details` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--metrics_details))
- `name` (String)
- `os_type` (String)
- `physical_port_count` (Number)
- `platform_id` (String)
- `port_count` (Number)
- `product_vendor` (String)
- `redundancy_mode` (String)
- `redundancy_peer_state` (String)
- `redundancy_peer_state_derived` (String)
- `redundancy_state` (String)
- `redundancy_state_derived` (String)
- `ring_status` (String)
- `serial_number` (String)
- `site_hierarchy` (String)
- `site_hierarchy_id` (String)
- `site_id` (String)
- `software_version` (String)
- `stack_type` (String)
- `switch_poe_details` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--switch_poe_details))
- `tag_names` (List of String)
- `up_time` (Number)
- `virtual_port_count` (Number)
- `wired_client_count` (Number)
- `wireless_client_count` (Number)

<a id="nestedobjatt--parameters--items--aggregate_attributes"></a>
### Nested Schema for `parameters.items.aggregate_attributes`

Read-Only:

- `function` (String)
- `name` (String)
- `value` (Number)


<a id="nestedobjatt--parameters--items--ap_details"></a>
### Nested Schema for `parameters.items.ap_details`

Read-Only:

- `admin_state` (String)
- `ap_group` (String)
- `ap_operational_state` (String)
- `ap_type` (String)
- `connected_time` (Number)
- `connected_wlc_name` (String)
- `ethernet_mac` (String)
- `flex_group` (String)
- `home_ap_enabled` (String)
- `icap_capability` (String)
- `led_flash_enabled` (String)
- `led_flash_seconds` (Number)
- `operational_mode` (String)
- `policy_tag_name` (String)
- `power_calendar_profile` (String)
- `power_mode` (String)
- `power_profile` (String)
- `power_save_mode` (String)
- `power_save_mode_capable` (String)
- `power_type` (String)
- `protocol` (String)
- `radios` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--ap_details--radios))
- `regulatory_domain` (String)
- `reset_reason` (String)
- `rf_tag_name` (String)
- `site_tag_name` (String)
- `sub_mode` (String)

<a id="nestedobjatt--parameters--items--ap_details--radios"></a>
### Nested Schema for `parameters.items.ap_details.sub_mode`

Read-Only:

- `air_quality` (Number)
- `band` (String)
- `client_count` (Number)
- `id` (String)
- `interference` (Number)
- `noise` (Number)
- `traffic_util` (Number)
- `utilization` (Number)



<a id="nestedobjatt--parameters--items--fabric_details"></a>
### Nested Schema for `parameters.items.fabric_details`

Read-Only:

- `fabric_role` (List of String)
- `fabric_site_id` (String)
- `fabric_site_name` (String)
- `l2_vns` (List of String)
- `l3_vns` (List of String)
- `network_protocol` (String)
- `transit_fabrics` (List of String)


<a id="nestedobjatt--parameters--items--fabric_metrics_details"></a>
### Nested Schema for `parameters.items.fabric_metrics_details`

Read-Only:

- `aaa_status_score` (Number)
- `bgp_bgp_site_score` (Number)
- `bgp_evpn_score` (Number)
- `bgp_peer_infra_vn_score` (Number)
- `bgp_peer_score` (Number)
- `bgp_pubsub_site_score` (Number)
- `bgp_tcp_score` (Number)
- `cts_env_data_download_score` (Number)
- `fabric_site_score` (Number)
- `fabric_transit_score` (Number)
- `fabric_vn_score` (Number)
- `fabsite_fcp_score` (Number)
- `fabsite_fsconn_score` (Number)
- `fabsite_infra_score` (Number)
- `internet_avail_score` (Number)
- `lisp_cp_conn_score` (Number)
- `lisp_transit_conn_score` (Number)
- `mcast_score` (Number)
- `overall_fabric_score` (Number)
- `peer_score` (Number)
- `port_channel_score` (Number)
- `pubsub_infra_vn_score` (Number)
- `pubsub_session_score` (Number)
- `pubsub_transit_conn_score` (Number)
- `remote_internet_avail_score` (Number)
- `tcp_conn_score` (Number)
- `transit_control_plane_score` (Number)
- `transit_services_score` (Number)
- `vn_exit_score` (Number)
- `vn_fcp_score` (Number)
- `vn_service_score` (Number)
- `vn_status_score` (Number)
- `vni_status_score` (Number)


<a id="nestedobjatt--parameters--items--metrics_details"></a>
### Nested Schema for `parameters.items.metrics_details`

Read-Only:

- `air_quality_score` (Number)
- `ap_count` (Number)
- `avg_temperature` (Number)
- `cpu_score` (Number)
- `cpu_utilization` (Number)
- `discard_interfaces` (List of String)
- `discard_score` (Number)
- `error_interfaces` (List of String)
- `error_score` (Number)
- `free_memory_buffer` (Number)
- `free_memory_buffer_score` (Number)
- `free_timer` (Number)
- `free_timer_score` (Number)
- `high_link_utilization_interfaces` (List of String)
- `inter_device_connected_down_interfaces` (List of String)
- `inter_device_link_score` (Number)
- `interference_score` (Number)
- `link_utilization_score` (Number)
- `max_temperature` (Number)
- `memory_score` (Number)
- `memory_utilization` (Number)
- `noise_score` (Number)
- `overall_health_score` (Number)
- `packet_pool` (Number)
- `packet_pool_score` (Number)
- `utilization_score` (Number)
- `wqe_pool` (Number)
- `wqe_pool_score` (Number)


<a id="nestedobjatt--parameters--items--switch_poe_details"></a>
### Nested Schema for `parameters.items.switch_poe_details`

Read-Only:

- `chassis_count` (Number)
- `free_port_count` (Number)
- `module_count` (Number)
- `module_details` (List of Object) (see [below for nested schema](#nestedobjatt--parameters--items--switch_poe_details--module_details))
- `poe_power_allocated` (Number)
- `poe_power_consumed` (Number)
- `poe_version` (String)
- `port_count` (Number)
- `power_budget` (Number)
- `power_consumed` (Number)
- `power_remaining` (Number)
- `system_power_allocated` (Number)
- `system_power_consumed` (Number)
- `used_port_count` (Number)

<a id="nestedobjatt--parameters--items--switch_poe_details--module_details"></a>
### Nested Schema for `parameters.items.switch_poe_details.used_port_count`

Read-Only:

- `chassis_id` (String)
- `interface_power_max` (Number)
- `module_free_port_count` (Number)
- `module_id` (String)
- `module_poe_power_allocated` (Number)
- `module_poe_power_consumed` (Number)
- `module_port_count` (Number)
- `module_power_budget` (Number)
- `module_power_consumed` (Number)
- `module_power_remaining` (Number)
- `module_system_power_allocated` (Number)
- `module_system_power_consumed` (Number)
- `module_used_port_count` (Number)
