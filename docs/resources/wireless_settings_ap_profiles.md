---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_settings_ap_profiles Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create and read operations on Wireless.
  This resource allows the user to create a custom AP Profile.
---

# catalystcenter_wireless_settings_ap_profiles (Resource)

It manages create and read operations on Wireless.

- This resource allows the user to create a custom AP Profile.

## Example Usage

```terraform
resource "catalystcenter_wireless_settings_ap_profiles" "example" {
  provider = catalystcenter
 
  parameters {

    ap_power_profile_name  = "string"
    ap_profile_name        = "string"
    awips_enabled          = "false"
    awips_forensic_enabled = "false"
    calendar_power_profiles {

      duration {

        scheduler_date       = ["string"]
        scheduler_day        = ["string"]
        scheduler_end_time   = "string"
        scheduler_start_time = "string"
      }
      power_profile_name = "string"
      scheduler_type     = "string"
    }
    client_limit = 1
    country_code = "string"
    description  = "string"
    management_setting {

      auth_type                  = "string"
      cdp_state                  = "false"
      dot1x_password             = "string"
      dot1x_username             = "string"
      management_enable_password = "string"
      management_password        = "string"
      management_user_name       = "string"
      ssh_enabled                = "false"
      telnet_enabled             = "false"
    }
    mesh_enabled = "false"
    mesh_setting {

      backhaul_client_access    = "false"
      bridge_group_name         = "string"
      ghz24_backhaul_data_rates = "string"
      ghz5_backhaul_data_rates  = "string"
      range                     = 1
      rap_downlink_backhaul     = "string"
    }
    pmf_denial_enabled    = "false"
    remote_worker_enabled = "false"
    rogue_detection_setting {

      rogue_detection                    = "false"
      rogue_detection_min_rssi           = 1
      rogue_detection_report_interval    = 1
      rogue_detection_transient_interval = 1
    }
    time_zone                = "string"
    time_zone_offset_hour    = 1
    time_zone_offset_minutes = 1
  }
}

output "catalystcenter_wireless_settings_ap_profiles_example" {
  value = catalystcenter_wireless_settings_ap_profiles.example
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

- `ap_power_profile_name` (String) Name of the existing AP power profile.
- `ap_profile_name` (String) Name of the Access Point profile. Max length is 32 characters.
- `awips_enabled` (String) Indicates if AWIPS is enabled on the AP.
- `awips_forensic_enabled` (String) Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.
- `calendar_power_profiles` (Block List) (see [below for nested schema](#nestedblock--parameters--calendar_power_profiles))
- `client_limit` (Number) Number of clients. Value should be between 0-1200.
- `country_code` (String) Country Code
- `description` (String) Description of the AP profile. Max length is 241 characters
- `management_setting` (Block List) (see [below for nested schema](#nestedblock--parameters--management_setting))
- `mesh_enabled` (String) This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.
- `mesh_setting` (Block List) (see [below for nested schema](#nestedblock--parameters--mesh_setting))
- `pmf_denial_enabled` (String) Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.
- `remote_worker_enabled` (String) Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.
- `rogue_detection_setting` (Block List) (see [below for nested schema](#nestedblock--parameters--rogue_detection_setting))
- `time_zone` (String) In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.
- `time_zone_offset_hour` (Number) Enter the hour value (HH). The valid range is from -12 through 14.
- `time_zone_offset_minutes` (Number) Enter the minute value (MM). The valid range is from 0 through 59.

<a id="nestedblock--parameters--calendar_power_profiles"></a>
### Nested Schema for `parameters.calendar_power_profiles`

Optional:

- `duration` (Block List) (see [below for nested schema](#nestedblock--parameters--calendar_power_profiles--duration))
- `power_profile_name` (String) Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.
- `scheduler_type` (String) Type of the scheduler.

<a id="nestedblock--parameters--calendar_power_profiles--duration"></a>
### Nested Schema for `parameters.calendar_power_profiles.duration`

Optional:

- `scheduler_date` (List of String) Start and End date of the duration setting, applicable for MONTHLY schedulers. Values must be between 1 and 31.
- `scheduler_day` (List of String) Applies every week on the selected days.  Ex: ["sunday","saturday","tuesday","wednesday","thursday","friday","monday"]
- `scheduler_end_time` (String) End time of the duration setting.
- `scheduler_start_time` (String) Start time of the duration setting.



<a id="nestedblock--parameters--management_setting"></a>
### Nested Schema for `parameters.management_setting`

Optional:

- `auth_type` (String) Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.
- `cdp_state` (String) Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
- `dot1x_password` (String) Password for 802.1X authentication. Length must be 8-120 characters.
- `dot1x_username` (String) Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.
- `management_enable_password` (String) Enable password for managing the AP. Length must be 8-120 characters.
- `management_password` (String) Management password for the AP. Length must be 8-120 characters.
- `management_user_name` (String) Management username must have a minimum of 1 character and a maximum of 32 characters.
- `ssh_enabled` (String) Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.
- `telnet_enabled` (String) Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.


<a id="nestedblock--parameters--mesh_setting"></a>
### Nested Schema for `parameters.mesh_setting`

Optional:

- `backhaul_client_access` (String) Indicates if backhaul client access is enabled on the AP.
- `bridge_group_name` (String) Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.
- `ghz24_backhaul_data_rates` (String) 2.4GHz backhaul data rates.
- `ghz5_backhaul_data_rates` (String) 5GHz backhaul data rates.
- `range` (Number) Range of the mesh network. Value should be between 150-132000
- `rap_downlink_backhaul` (String) Type of downlink backhaul used.


<a id="nestedblock--parameters--rogue_detection_setting"></a>
### Nested Schema for `parameters.rogue_detection_setting`

Optional:

- `rogue_detection` (String) Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters
- `rogue_detection_min_rssi` (Number) Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts
- `rogue_detection_report_interval` (Number) Report interval for rogue detection. Value should be in range 10 and 300.
- `rogue_detection_transient_interval` (Number) Transient interval for rogue detection. Value should be 0 or from 120 to 1800.



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `ap_power_profile_name` (String)
- `ap_profile_name` (String)
- `awips_enabled` (String)
- `awips_forensic_enabled` (String)
- `calendar_power_profiles` (List of Object) (see [below for nested schema](#nestedobjatt--item--calendar_power_profiles))
- `client_limit` (Number)
- `country_code` (String)
- `description` (String)
- `id` (String)
- `management_setting` (List of Object) (see [below for nested schema](#nestedobjatt--item--management_setting))
- `mesh_enabled` (String)
- `mesh_setting` (List of Object) (see [below for nested schema](#nestedobjatt--item--mesh_setting))
- `pmf_denial_enabled` (String)
- `remote_worker_enabled` (String)
- `rogue_detection_setting` (List of Object) (see [below for nested schema](#nestedobjatt--item--rogue_detection_setting))
- `time_zone` (String)
- `time_zone_offset_hour` (Number)
- `time_zone_offset_minutes` (Number)

<a id="nestedobjatt--item--calendar_power_profiles"></a>
### Nested Schema for `item.calendar_power_profiles`

Read-Only:

- `duration` (List of Object) (see [below for nested schema](#nestedobjatt--item--calendar_power_profiles--duration))
- `power_profile_name` (String)
- `scheduler_type` (String)

<a id="nestedobjatt--item--calendar_power_profiles--duration"></a>
### Nested Schema for `item.calendar_power_profiles.duration`

Read-Only:

- `scheduler_date` (List of String)
- `scheduler_day` (List of String)
- `scheduler_end_time` (String)
- `scheduler_start_time` (String)



<a id="nestedobjatt--item--management_setting"></a>
### Nested Schema for `item.management_setting`

Read-Only:

- `auth_type` (String)
- `cdp_state` (String)
- `dot1x_password` (String)
- `dot1x_username` (String)
- `management_enable_password` (String)
- `management_password` (String)
- `management_user_name` (String)
- `ssh_enabled` (String)
- `telnet_enabled` (String)


<a id="nestedobjatt--item--mesh_setting"></a>
### Nested Schema for `item.mesh_setting`

Read-Only:

- `backhaul_client_access` (String)
- `bridge_group_name` (String)
- `ghz24_backhaul_data_rates` (String)
- `ghz5_backhaul_data_rates` (String)
- `range` (Number)
- `rap_downlink_backhaul` (String)


<a id="nestedobjatt--item--rogue_detection_setting"></a>
### Nested Schema for `item.rogue_detection_setting`

Read-Only:

- `rogue_detection` (String)
- `rogue_detection_min_rssi` (Number)
- `rogue_detection_report_interval` (Number)
- `rogue_detection_transient_interval` (Number)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_wireless_settings_ap_profiles.example "id:=string"
```
