---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_profiles Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  This data source allows the user to get all Wireless Network ProfilesThis data source allows the user to get a Wireless Network Profile by ID
---

# catalystcenter_wireless_profiles (Data Source)

It performs read operation on Wireless.

- This data source allows the user to get all Wireless Network Profiles

- This data source allows the user to get a Wireless Network Profile by ID

## Example Usage

```terraform
data "catalystcenter_wireless_profiles" "example" {
  provider              = catalystcenter
  limit                 = 1
  offset                = 1
  wireless_profile_name = "string"
}

output "catalystcenter_wireless_profiles_example" {
  value = data.catalystcenter_wireless_profiles.example.items
}

data "catalystcenter_wireless_profiles" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_profiles_example" {
  value = data.catalystcenter_wireless_profiles.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter. Wireless Profile Id
- `limit` (Number) limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500
- `offset` (Number) offset query parameter. The first record to show for this page; the first record is numbered 1
- `wireless_profile_name` (String) wirelessProfileName query parameter. Wireless Profile Name

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `additional_interfaces` (List of String)
- `ap_zones` (List of Object) (see [below for nested schema](#nestedobjatt--item--ap_zones))
- `feature_templates` (List of Object) (see [below for nested schema](#nestedobjatt--item--feature_templates))
- `id` (String)
- `ssid_details` (List of Object) (see [below for nested schema](#nestedobjatt--item--ssid_details))
- `wireless_profile_name` (String)

<a id="nestedobjatt--item--ap_zones"></a>
### Nested Schema for `item.ap_zones`

Read-Only:

- `ap_zone_name` (String)
- `rf_profile_name` (String)
- `ssids` (List of String)


<a id="nestedobjatt--item--feature_templates"></a>
### Nested Schema for `item.feature_templates`

Read-Only:

- `id` (String)
- `ssids` (List of String)


<a id="nestedobjatt--item--ssid_details"></a>
### Nested Schema for `item.ssid_details`

Read-Only:

- `anchor_group_name` (String)
- `dot11be_profile_id` (String)
- `enable_fabric` (String)
- `flex_connect` (List of Object) (see [below for nested schema](#nestedobjatt--item--ssid_details--flex_connect))
- `interface_name` (String)
- `policy_profile_name` (String)
- `ssid_name` (String)
- `vlan_group_name` (String)
- `wlan_profile_name` (String)

<a id="nestedobjatt--item--ssid_details--flex_connect"></a>
### Nested Schema for `item.ssid_details.flex_connect`

Read-Only:

- `enable_flex_connect` (String)
- `local_to_vlan` (Number)




<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `additional_interfaces` (List of String)
- `ap_zones` (List of Object) (see [below for nested schema](#nestedobjatt--items--ap_zones))
- `feature_templates` (List of Object) (see [below for nested schema](#nestedobjatt--items--feature_templates))
- `id` (String)
- `ssid_details` (List of Object) (see [below for nested schema](#nestedobjatt--items--ssid_details))
- `wireless_profile_name` (String)

<a id="nestedobjatt--items--ap_zones"></a>
### Nested Schema for `items.ap_zones`

Read-Only:

- `ap_zone_name` (String)
- `rf_profile_name` (String)
- `ssids` (List of String)


<a id="nestedobjatt--items--feature_templates"></a>
### Nested Schema for `items.feature_templates`

Read-Only:

- `id` (String)
- `ssids` (List of String)


<a id="nestedobjatt--items--ssid_details"></a>
### Nested Schema for `items.ssid_details`

Read-Only:

- `anchor_group_name` (String)
- `dot11be_profile_id` (String)
- `enable_fabric` (String)
- `flex_connect` (List of Object) (see [below for nested schema](#nestedobjatt--items--ssid_details--flex_connect))
- `interface_name` (String)
- `policy_profile_name` (String)
- `ssid_name` (String)
- `vlan_group_name` (String)
- `wlan_profile_name` (String)

<a id="nestedobjatt--items--ssid_details--flex_connect"></a>
### Nested Schema for `items.ssid_details.flex_connect`

Read-Only:

- `enable_flex_connect` (String)
- `local_to_vlan` (Number)
