---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_enterprise_ssid Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on Wireless.
  Creates enterprise SSIDUpdate enterprise SSIDDeletes given enterprise SSID
---

# catalystcenter_wireless_enterprise_ssid (Resource)

It manages create, read, update and delete operations on Wireless.

- Creates enterprise SSID

- Update enterprise SSID

- Deletes given enterprise SSID

## Example Usage

```terraform
resource "catalystcenter_wireless_enterprise_ssid" "example" {
  provider = catalystcenter

  parameters {

    aaa_override                          = "false"
    auth_key_mgmt                         = ["string"]
    basic_service_set_client_idle_timeout = 1
    client_exclusion_timeout              = 1
    client_rate_limit                     = 1.0
    coverage_hole_detection_enable        = "false"
    enable_basic_service_set_max_idle     = "false"
    enable_broadcast_ssi_d                = "false"
    enable_client_exclusion               = "false"
    enable_directed_multicast_service     = "false"
    enable_fast_lane                      = "false"
    enable_mac_filtering                  = "false"
    enable_neighbor_list                  = "false"
    enable_session_time_out               = "false"
    fast_transition                       = "string"
    ghz24_policy                          = "string"
    ghz6_policy_client_steering           = "false"
    mfp_client_protection                 = "string"
    multi_psk_settings {

      passphrase      = "string"
      passphrase_type = "string"
      priority        = 1
    }
    name                       = "string"
    nas_options                = ["string"]
    passphrase                 = "string"
    policy_profile_name        = "string"
    profile_name               = "string"
    protected_management_frame = "string"
    radio_policy               = "string"
    rsn_cipher_suite_ccmp256   = "false"
    rsn_cipher_suite_gcmp128   = "false"
    rsn_cipher_suite_gcmp256   = "false"
    security_level             = "string"
    session_time_out           = 1
    ssid_name                  = "string"
    traffic_type               = "string"
  }
}

output "catalystcenter_wireless_enterprise_ssid_example" {
  value = catalystcenter_wireless_enterprise_ssid.example
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

- `aaa_override` (String) Aaa Override
- `auth_key_mgmt` (List of String) Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft
- `basic_service_set_client_idle_timeout` (Number) Basic Service Set Client Idle Timeout
- `client_exclusion_timeout` (Number) Client Exclusion Timeout
- `client_rate_limit` (Number) Client Rate Limit (in bits per second)
- `coverage_hole_detection_enable` (String) Coverage Hole Detection Enable
- `enable_basic_service_set_max_idle` (String) Enable Basic Service Set Max Idle
- `enable_broadcast_ssi_d` (String) Enable Broadcase SSID
- `enable_client_exclusion` (String) Enable Client Exclusion
- `enable_directed_multicast_service` (String) Enable Directed Multicast Service
- `enable_fast_lane` (String) Enable FastLane
- `enable_mac_filtering` (String) Enable MAC Filtering
- `enable_neighbor_list` (String) Enable Neighbor List
- `enable_session_time_out` (String) Enable Session Timeout
- `fast_transition` (String) Fast Transition
- `ghz24_policy` (String) Ghz24 Policy
- `ghz6_policy_client_steering` (String) Ghz6 Policy Client Steering
- `mfp_client_protection` (String) Management Frame Protection Client
- `multi_psk_settings` (Block List) (see [below for nested schema](#nestedblock--parameters--multi_psk_settings))
- `name` (String) SSID NAME
- `nas_options` (List of String) Nas Options
- `passphrase` (String) Passphrase
- `policy_profile_name` (String) Policy Profile Name
- `profile_name` (String) Profile Name
- `protected_management_frame` (String) (Required applicable for Security Type WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (Optional, Required Applicable for Security Type WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
- `radio_policy` (String) Radio Policy Enum
- `rsn_cipher_suite_ccmp256` (String) Rsn Cipher Suite Ccmp256
- `rsn_cipher_suite_gcmp128` (String) Rsn Cipher Suite Gcmp 128
- `rsn_cipher_suite_gcmp256` (String) Rsn Cipher Suite Gcmp256
- `security_level` (String) Security Level
- `session_time_out` (Number) Session Time Out
- `ssid_name` (String) ssidName path parameter. Enter the SSID name to be deleted
- `traffic_type` (String) Traffic Type Enum (voicedata or data )

<a id="nestedblock--parameters--multi_psk_settings"></a>
### Nested Schema for `parameters.multi_psk_settings`

Optional:

- `passphrase` (String) Passphrase
- `passphrase_type` (String) Passphrase Type
- `priority` (Number) Priority



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `group_uuid` (String)
- `inherited_group_name` (String)
- `inherited_group_uuid` (String)
- `instance_uuid` (String)
- `ssid_details` (List of Object) (see [below for nested schema](#nestedobjatt--item--ssid_details))
- `version` (Number)

<a id="nestedobjatt--item--ssid_details"></a>
### Nested Schema for `item.ssid_details`

Read-Only:

- `aaa_override` (String)
- `auth_server` (String)
- `basic_service_set_client_idle_timeout` (Number)
- `client_exclusion_timeout` (Number)
- `client_rate_limit` (Number)
- `coverage_hole_detection_enable` (String)
- `enable_basic_service_set_max_idle` (String)
- `enable_broadcast_ssi_d` (String)
- `enable_client_exclusion` (String)
- `enable_directed_multicast_service` (String)
- `enable_fast_lane` (String)
- `enable_mac_filtering` (String)
- `enable_neighbor_list` (String)
- `enable_session_time_out` (String)
- `fast_transition` (String)
- `is_enabled` (String)
- `is_fabric` (String)
- `mfp_client_protection` (String)
- `multi_psk_settings` (List of Object) (see [below for nested schema](#nestedobjatt--item--ssid_details--multi_psk_settings))
- `name` (String)
- `nas_options` (List of String)
- `passphrase` (String)
- `protected_management_frame` (String)
- `radio_policy` (String)
- `security_level` (String)
- `session_time_out` (Number)
- `traffic_type` (String)
- `wlan_type` (String)

<a id="nestedobjatt--item--ssid_details--multi_psk_settings"></a>
### Nested Schema for `item.ssid_details.multi_psk_settings`

Read-Only:

- `passphrase` (String)
- `passphrase_type` (String)
- `priority` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_wireless_enterprise_ssid.example "id:=string"
```