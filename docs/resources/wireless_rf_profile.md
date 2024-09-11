---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_rf_profile Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read and delete operations on Wireless.
  Create or Update RF profileDelete RF profile
---

# catalystcenter_wireless_rf_profile (Resource)

It manages create, read and delete operations on Wireless.

- Create or Update RF profile

- Delete RF profile

## Example Usage

```terraform
resource "catalystcenter_wireless_rf_profile" "example" {
  provider = catalystcenter

  parameters {

    channel_width       = "string"
    default_rf_profile  = "false"
    enable_brown_field  = "false"
    enable_custom       = "false"
    enable_radio_type_a = "false"
    enable_radio_type_b = "false"
    enable_radio_type_c = "false"
    name                = "string"
    radio_type_a_properties {

      data_rates           = "string"
      mandatory_data_rates = "string"
      max_power_level      = 1.0
      min_power_level      = 1.0
      parent_profile       = "string"
      power_threshold_v1   = 1.0
      radio_channels       = "string"
      rx_sop_threshold     = "string"
    }
    radio_type_b_properties {

      data_rates           = "string"
      mandatory_data_rates = "string"
      max_power_level      = 1.0
      min_power_level      = 1.0
      parent_profile       = "string"
      power_threshold_v1   = 1.0
      radio_channels       = "string"
      rx_sop_threshold     = "string"
    }
    radio_type_c_properties {

      data_rates           = "string"
      mandatory_data_rates = "string"
      max_power_level      = 1.0
      min_power_level      = 1.0
      parent_profile       = "string"
      power_threshold_v1   = 1.0
      radio_channels       = "string"
      rx_sop_threshold     = "string"
    }
    rf_profile_name = "string"
  }
}

output "catalystcenter_wireless_rf_profile_example" {
  value = catalystcenter_wireless_rf_profile.example
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

Required:

- `rf_profile_name` (String) rfProfileName path parameter. RF profile name to be deleted(required) *non-custom RF profile cannot be deleted

Optional:

- `channel_width` (String) Channel Width
- `default_rf_profile` (String) is Default Rf Profile
- `enable_brown_field` (String) Enable Brown Field
- `enable_custom` (String) Enable Custom
- `enable_radio_type_a` (String) Enable Radio Type A
- `enable_radio_type_b` (String) Enable Radio Type B
- `enable_radio_type_c` (String) Enable Radio Type C (6GHz)
- `name` (String) RF Profile Name
- `radio_type_a_properties` (Block List) (see [below for nested schema](#nestedblock--parameters--radio_type_a_properties))
- `radio_type_b_properties` (Block List) (see [below for nested schema](#nestedblock--parameters--radio_type_b_properties))
- `radio_type_c_properties` (Block List) (see [below for nested schema](#nestedblock--parameters--radio_type_c_properties))

<a id="nestedblock--parameters--radio_type_a_properties"></a>
### Nested Schema for `parameters.radio_type_a_properties`

Optional:

- `data_rates` (String) Data Rates (Default : "6,9,12,18,24,36,48,54")
- `mandatory_data_rates` (String) Mandatory Data Rates (Default: "6,12,24")
- `max_power_level` (Number) Max Power Level  (Default: 30)
- `min_power_level` (Number) Rx Sop Threshold  (Default: -10)
- `parent_profile` (String) Parent Profile (Default : CUSTOM)
- `power_threshold_v1` (Number) Power Threshold V1 ( (Default: -70)
- `radio_channels` (String) Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")
- `rx_sop_threshold` (String) Rx Sop Threshold  (Default: "AUTO")


<a id="nestedblock--parameters--radio_type_b_properties"></a>
### Nested Schema for `parameters.radio_type_b_properties`

Optional:

- `data_rates` (String) Data Rates  (Default: "9,11,12,18,24,36,48,54")
- `mandatory_data_rates` (String) Mandatory Data Rates  (Default: "12")
- `max_power_level` (Number) Max Power Level  (Default: 30)
- `min_power_level` (Number) Min Power Level  (Default: -10)
- `parent_profile` (String) Parent Profile (Default : CUSTOM)
- `power_threshold_v1` (Number) Power Threshold V1  (Default: -70)
- `radio_channels` (String) Radio Channels (Default : "9,11,12,18,24,36,48,54")
- `rx_sop_threshold` (String) Rx Sop Threshold (Default: "AUTO")


<a id="nestedblock--parameters--radio_type_c_properties"></a>
### Nested Schema for `parameters.radio_type_c_properties`

Optional:

- `data_rates` (String) Data Rates  (Default: "6,9,12,18,24,36,48,54")
- `mandatory_data_rates` (String) Mandatory Data Rates  (Default: "6,12,24")
- `max_power_level` (Number) Max Power Level  (Default: 30)
- `min_power_level` (Number) Min Power Level  (Default: -10)
- `parent_profile` (String) Parent Profile (Default : CUSTOM)
- `power_threshold_v1` (Number) Power Threshold V1  (Default: -70)
- `radio_channels` (String) Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")
- `rx_sop_threshold` (String) Rx Sop Threshold  (Default: "AUTO")



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `channel_width` (String)
- `default_rf_profile` (String)
- `enable_brown_field` (String)
- `enable_custom` (String)
- `enable_radio_type_a` (String)
- `enable_radio_type_b` (String)
- `enable_radio_type_c` (String)
- `name` (String)
- `radio_type_a_properties` (List of Object) (see [below for nested schema](#nestedobjatt--item--radio_type_a_properties))
- `radio_type_b_properties` (List of Object) (see [below for nested schema](#nestedobjatt--item--radio_type_b_properties))
- `radio_type_c_properties` (List of Object) (see [below for nested schema](#nestedobjatt--item--radio_type_c_properties))

<a id="nestedobjatt--item--radio_type_a_properties"></a>
### Nested Schema for `item.radio_type_a_properties`

Read-Only:

- `data_rates` (String)
- `mandatory_data_rates` (String)
- `max_power_level` (Number)
- `min_power_level` (Number)
- `parent_profile` (String)
- `power_threshold_v1` (Number)
- `radio_channels` (String)
- `rx_sop_threshold` (String)


<a id="nestedobjatt--item--radio_type_b_properties"></a>
### Nested Schema for `item.radio_type_b_properties`

Read-Only:

- `data_rates` (String)
- `mandatory_data_rates` (String)
- `max_power_level` (Number)
- `min_power_level` (Number)
- `parent_profile` (String)
- `power_threshold_v1` (Number)
- `radio_channels` (String)
- `rx_sop_threshold` (String)


<a id="nestedobjatt--item--radio_type_c_properties"></a>
### Nested Schema for `item.radio_type_c_properties`

Read-Only:

- `data_rates` (String)
- `mandatory_data_rates` (String)
- `max_power_level` (Number)
- `min_power_level` (Number)
- `parent_profile` (String)
- `power_threshold_v1` (Number)
- `radio_channels` (String)
- `rx_sop_threshold` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_wireless_rf_profile.example "id:=string"
```