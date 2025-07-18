---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_device_list Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read and update operations on Devices.
  Adds the device with given credentialUpdate the credentials, management IP address of a given device (or a set of devices) in Catalyst Center and trigger
  an inventory sync.
---

# catalystcenter_network_device_list (Resource)

It manages create, read and update operations on Devices.

- Adds the device with given credential

- Update the credentials, management IP address of a given device (or a set of devices) in Catalyst Center and trigger
an inventory sync.

## Example Usage

```terraform
resource "catalystcenter_network_device_list" "example" {
  provider = catalystcenter
 
  parameters {

    cli_transport           = "string"
    compute_device          = "false"
    enable_password         = "string"
    extended_discovery_info = "string"
    http_password           = "string"
    http_port               = "string"
    http_secure             = "false"
    http_user_name          = "string"
    ip_address              = ["string"]
    meraki_org_id           = ["string"]
    netconf_port            = "string"
    password                = "******"
    serial_number           = "string"
    snmp_auth_passphrase    = "string"
    snmp_auth_protocol      = "string"
    snmp_mode               = "string"
    snmp_priv_passphrase    = "string"
    snmp_priv_protocol      = "string"
    snmp_ro_community       = "string"
    snmp_rw_community       = "string"
    snmp_retry              = 1
    snmp_timeout            = 1
    snmp_user_name          = "string"
    snmp_version            = "string"
    type                    = "string"
    update_mgmt_ipaddress_list {

      exist_mgmt_ip_address = "string"
      new_mgmt_ip_address   = "string"
    }
    user_name = "string"
  }
}

output "catalystcenter_network_device_list_example" {
  value = catalystcenter_network_device_list.example
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

- `cli_transport` (String) CLI transport. Supported values: telnet, ssh. Required if type is NETWORK_DEVICE.
- `compute_device` (String) Compute Device or not. Options are true / false.
- `enable_password` (String) CLI enable password of the device. Required if device is configured to use enable password.
- `extended_discovery_info` (String) This field holds that info as whether to add device with canned data or not. Supported values: DISCOVER_WITH_CANNED_DATA.
- `http_password` (String) HTTP password of the device / API key for catalystcenter Dashboard. Required if type is catalystcenter_DASHBOARD or COMPUTE_DEVICE.
- `http_port` (String) HTTP port of the device. Required if type is COMPUTE_DEVICE.
- `http_secure` (String) Flag to select HTTP / HTTPS protocol. Options are true / false. true for HTTPS and false for HTTP. Default is true.
- `http_user_name` (String) HTTP Username of the device. Required if type is COMPUTE_DEVICE.
- `ip_address` (List of String) IP Address of the device. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.
- `meraki_org_id` (List of String) Selected catalystcenter organization for which the devices needs to be imported. Required if type is catalystcenter_DASHBOARD.
- `netconf_port` (String) Netconf Port of the device. cliTransport must be 'ssh' if netconf is provided. Netconf port is required for eWLC.
- `password` (String, Sensitive) CLI Password of the device. Required if type is NETWORK_DEVICE.
- `serial_number` (String) Serial Number of the Device. Required if extendedDiscoveryInfo is 'DISCOVER_WITH_CANNED_DATA'.
- `snmp_auth_passphrase` (String) SNMPv3 auth passphrase of the device. Required if snmpMode is authNoPriv or authPriv.
- `snmp_auth_protocol` (String) SNMPv3 auth protocol. Supported values: sha, md5. Required if snmpMode is authNoPriv or authPriv.
- `snmp_mode` (String) SNMPv3 mode. Supported values: noAuthnoPriv, authNoPriv, authPriv. Required if snmpVersion is v3.
- `snmp_priv_passphrase` (String) SNMPv3 priv passphrase. Required if snmpMode is authPriv.
- `snmp_priv_protocol` (String) SNMPv3 priv protocol. Supported values: AES128. Required if snmpMode is authPriv.
- `snmp_retry` (Number) SNMP retry count. Max value supported is 3. Default is Global SNMP retry (if exists) or 3.
- `snmp_ro_community` (String) SNMP Read Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.
- `snmp_rw_community` (String) SNMP Write Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.
- `snmp_timeout` (Number) SNMP timeout in seconds. Max value supported is 300. Default is Global SNMP timeout (if exists) or 5.
- `snmp_user_name` (String) SNMPV3 user name of the device. Required if snmpVersion is v3.
- `snmp_version` (String) SNMP version. Values supported: v2, v3. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.
- `type` (String) Type of device being added. Default is NETWORK_DEVICE.
- `update_mgmt_ipaddress_list` (Block List) (see [below for nested schema](#nestedblock--parameters--update_mgmt_ipaddress_list))
- `user_name` (String) CLI user name of the device. Required if type is NETWORK_DEVICE.

<a id="nestedblock--parameters--update_mgmt_ipaddress_list"></a>
### Nested Schema for `parameters.update_mgmt_ipaddress_list`

Optional:

- `exist_mgmt_ip_address` (String) existMgmtIpAddress IP Address of the device.
- `new_mgmt_ip_address` (String) New IP Address to be Updated.



<a id="nestedatt--item"></a>
### Nested Schema for `item`

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

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_network_device_list.example "id:=string"
```
