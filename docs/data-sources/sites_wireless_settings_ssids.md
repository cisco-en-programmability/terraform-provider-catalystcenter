---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sites_wireless_settings_ssids Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  This data source allows the user to get all SSIDs (Service Set Identifier) at the given siteThis data source allows the user to get an SSID (Service Set Identifier) by ID at the given site
---

# catalystcenter_sites_wireless_settings_ssids (Data Source)

It performs read operation on Wireless.

- This data source allows the user to get all SSIDs (Service Set Identifier) at the given site

- This data source allows the user to get an SSID (Service Set Identifier) by ID at the given site

## Example Usage

```terraform
data "catalystcenter_sites_wireless_settings_ssids" "example" {
  provider    = catalystcenter
  auth_type   = "string"
  l3auth_type = "string"
  limit       = 1
  offset      = 1
  site_id     = "string"
  ssid        = "string"
  wlan_type   = "string"
}

output "catalystcenter_sites_wireless_settings_ssids_example" {
  value = data.catalystcenter_sites_wireless_settings_ssids.example.items
}

data "catalystcenter_sites_wireless_settings_ssids" "example" {
  provider    = catalystcenter
  auth_type   = "string"
  l3auth_type = "string"
  limit       = 1
  offset      = 1
  site_id     = "string"
  ssid        = "string"
  wlan_type   = "string"
}

output "catalystcenter_sites_wireless_settings_ssids_example" {
  value = data.catalystcenter_sites_wireless_settings_ssids.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `auth_type` (String) authType query parameter. Auth Type
- `id` (String) id path parameter. SSID ID
- `l3auth_type` (String) l3authType query parameter. L3 Auth Type
- `limit` (Number) limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
- `offset` (Number) offset query parameter. The first record to show for this page; the first record is numbered 1.
- `site_id` (String) siteId path parameter. Site UUID
- `ssid` (String) ssid query parameter. SSID Name
- `wlan_type` (String) wlanType query parameter. Wlan Type

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `aaa_override` (String)
- `acct_servers` (List of String)
- `acl_name` (String)
- `auth_server` (String)
- `auth_servers` (List of String)
- `auth_type` (String)
- `basic_service_set_client_idle_timeout` (Number)
- `basic_service_set_max_idle_enable` (String)
- `cckm_tsf_tolerance` (Number)
- `client_exclusion_enable` (String)
- `client_exclusion_timeout` (Number)
- `client_rate_limit` (Number)
- `coverage_hole_detection_enable` (String)
- `directed_multicast_service_enable` (String)
- `egress_qos` (String)
- `external_auth_ip_address` (String)
- `fast_transition` (String)
- `fast_transition_over_the_distributed_system_enable` (String)
- `ghz24_policy` (String)
- `ghz6_policy_client_steering` (String)
- `id` (String)
- `ingress_qos` (String)
- `inherited_site_id` (String)
- `inherited_site_name` (String)
- `inherited_site_name_hierarchy` (String)
- `inherited_site_uui_d` (String)
- `is_ap_beacon_protection_enabled` (String)
- `is_auth_key8021x` (String)
- `is_auth_key8021x_plus_ft` (String)
- `is_auth_key8021x_sha256` (String)
- `is_auth_key_easy_psk` (String)
- `is_auth_key_owe` (String)
- `is_auth_key_psk` (String)
- `is_auth_key_psk_plus_ft` (String)
- `is_auth_key_psk_sha256` (String)
- `is_auth_key_sae` (String)
- `is_auth_key_sae_ext` (String)
- `is_auth_key_sae_ext_plus_ft` (String)
- `is_auth_key_sae_plus_ft` (String)
- `is_auth_key_suite_b1921x` (String)
- `is_auth_key_suite_b1x` (String)
- `is_broadcast_ssi_d` (String)
- `is_cckm_enabled` (String)
- `is_custom_nas_id_options` (String)
- `is_enabled` (String)
- `is_fast_lane_enabled` (String)
- `is_hex` (String)
- `is_mac_filtering_enabled` (String)
- `is_posturing_enabled` (String)
- `is_radius_profiling_enabled` (String)
- `is_random_mac_filter_enabled` (String)
- `is_sensor_pnp` (String)
- `l3_auth_type` (String)
- `management_frame_protection_clientprotection` (String)
- `multi_psk_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--multi_psk_settings))
- `nas_options` (List of String)
- `neighbor_list_enable` (String)
- `open_ssid` (String)
- `passphrase` (String)
- `policy_profile_name` (String)
- `profile_name` (String)
- `protected_management_frame` (String)
- `rsn_cipher_suite_ccmp128` (String)
- `rsn_cipher_suite_ccmp256` (String)
- `rsn_cipher_suite_gcmp128` (String)
- `rsn_cipher_suite_gcmp256` (String)
- `session_time_out` (Number)
- `session_time_out_enable` (String)
- `sleeping_client_enable` (String)
- `sleeping_client_timeout` (Number)
- `ssid` (String)
- `ssid_radio_type` (String)
- `web_passthrough` (String)
- `wlan_band_select_enable` (String)
- `wlan_type` (String)

<a id="nestedobjatt--item--multi_psk_settings"></a>
### Nested Schema for `item.multi_psk_settings`

Read-Only:

- `passphrase` (String)
- `passphrase_type` (String)
- `priority` (Number)



<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `aaa_override` (String)
- `acct_servers` (List of String)
- `acl_name` (String)
- `auth_server` (String)
- `auth_servers` (List of String)
- `auth_type` (String)
- `basic_service_set_client_idle_timeout` (Number)
- `basic_service_set_max_idle_enable` (String)
- `cckm_tsf_tolerance` (Number)
- `client_exclusion_enable` (String)
- `client_exclusion_timeout` (Number)
- `client_rate_limit` (Number)
- `coverage_hole_detection_enable` (String)
- `directed_multicast_service_enable` (String)
- `egress_qos` (String)
- `external_auth_ip_address` (String)
- `fast_transition` (String)
- `fast_transition_over_the_distributed_system_enable` (String)
- `ghz24_policy` (String)
- `ghz6_policy_client_steering` (String)
- `id` (String)
- `ingress_qos` (String)
- `inherited_site_id` (String)
- `inherited_site_name` (String)
- `inherited_site_name_hierarchy` (String)
- `is_ap_beacon_protection_enabled` (String)
- `is_auth_key8021x` (String)
- `is_auth_key8021x_plus_ft` (String)
- `is_auth_key8021x_sha256` (String)
- `is_auth_key_easy_psk` (String)
- `is_auth_key_owe` (String)
- `is_auth_key_psk` (String)
- `is_auth_key_psk_plus_ft` (String)
- `is_auth_key_psk_sha256` (String)
- `is_auth_key_sae` (String)
- `is_auth_key_sae_ext` (String)
- `is_auth_key_sae_ext_plus_ft` (String)
- `is_auth_key_sae_plus_ft` (String)
- `is_auth_key_suite_b1921x` (String)
- `is_auth_key_suite_b1x` (String)
- `is_broadcast_ssi_d` (String)
- `is_cckm_enabled` (String)
- `is_custom_nas_id_options` (String)
- `is_enabled` (String)
- `is_fast_lane_enabled` (String)
- `is_hex` (String)
- `is_mac_filtering_enabled` (String)
- `is_posturing_enabled` (String)
- `is_radius_profiling_enabled` (String)
- `is_random_mac_filter_enabled` (String)
- `is_sensor_pnp` (String)
- `l3_auth_type` (String)
- `management_frame_protection_clientprotection` (String)
- `multi_psk_settings` (List of Object) (see [below for nested schema](#nestedobjatt--items--multi_psk_settings))
- `nas_options` (List of String)
- `neighbor_list_enable` (String)
- `open_ssid` (String)
- `passphrase` (String)
- `policy_profile_name` (String)
- `profile_name` (String)
- `protected_management_frame` (String)
- `rsn_cipher_suite_ccmp128` (String)
- `rsn_cipher_suite_ccmp256` (String)
- `rsn_cipher_suite_gcmp128` (String)
- `rsn_cipher_suite_gcmp256` (String)
- `session_time_out` (Number)
- `session_time_out_enable` (String)
- `sleeping_client_enable` (String)
- `sleeping_client_timeout` (Number)
- `ssid` (String)
- `ssid_radio_type` (String)
- `web_passthrough` (String)
- `wlan_band_select_enable` (String)
- `wlan_type` (String)

<a id="nestedobjatt--items--multi_psk_settings"></a>
### Nested Schema for `items.multi_psk_settings`

Read-Only:

- `passphrase` (String)
- `passphrase_type` (String)
- `priority` (Number)
