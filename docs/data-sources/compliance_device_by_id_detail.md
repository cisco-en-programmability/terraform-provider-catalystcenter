---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_compliance_device_by_id_detail Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Compliance.
  Return compliance detailed report for a device.
---

# catalystcenter_compliance_device_by_id_detail (Data Source)

It performs read operation on Compliance.

- Return compliance detailed report for a device.

## Example Usage

```terraform
data "catalystcenter_compliance_device_by_id_detail" "example" {
  provider        = catalystcenter
  category        = "string"
  compliance_type = "string"
  device_uuid     = "string"
  diff_list       = "false"
}

output "catalystcenter_compliance_device_by_id_detail_example" {
  value = data.catalystcenter_compliance_device_by_id_detail.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_uuid` (String) deviceUuid path parameter. Device Id

### Optional

- `category` (String) category query parameter. category can have any value among 'INTENT', 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'DESIGN_OOD' , 'EoX' , 'NETWORK_SETTINGS'
- `compliance_type` (String) complianceType query parameter. Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EoX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
- `diff_list` (Boolean) diffList query parameter. diff list [ pass true to fetch the diff list ]

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `ack_status` (String)
- `compliance_type` (String)
- `device_uuid` (String)
- `last_sync_time` (Number)
- `last_update_time` (Number)
- `source_info_list` (List of Object) (see [below for nested schema](#nestedobjatt--items--source_info_list))
- `state` (String)
- `status` (String)
- `version` (String)

<a id="nestedobjatt--items--source_info_list"></a>
### Nested Schema for `items.source_info_list`

Read-Only:

- `ack_status` (String)
- `app_name` (String)
- `business_key` (List of Object) (see [below for nested schema](#nestedobjatt--items--source_info_list--business_key))
- `count` (Number)
- `diff_list` (List of Object) (see [below for nested schema](#nestedobjatt--items--source_info_list--diff_list))
- `display_name` (String)
- `name` (String)
- `name_with_business_key` (String)
- `source_enum` (String)
- `type` (String)

<a id="nestedobjatt--items--source_info_list--business_key"></a>
### Nested Schema for `items.source_info_list.business_key`

Read-Only:

- `business_key_attributes` (String)
- `other_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--items--source_info_list--business_key--other_attributes))
- `resource_name` (String)

<a id="nestedobjatt--items--source_info_list--business_key--other_attributes"></a>
### Nested Schema for `items.source_info_list.business_key.resource_name`

Read-Only:

- `cfs_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--items--source_info_list--business_key--resource_name--cfs_attributes))
- `name` (String)

<a id="nestedobjatt--items--source_info_list--business_key--resource_name--cfs_attributes"></a>
### Nested Schema for `items.source_info_list.business_key.resource_name.cfs_attributes`

Read-Only:

- `app_name` (String)
- `description` (String)
- `display_name` (String)
- `source` (String)
- `type` (String)




<a id="nestedobjatt--items--source_info_list--diff_list"></a>
### Nested Schema for `items.source_info_list.diff_list`

Read-Only:

- `ack_status` (String)
- `business_key` (String)
- `configured_value` (String)
- `display_name` (String)
- `extended_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--items--source_info_list--diff_list--extended_attributes))
- `instance_uui_d` (String)
- `intended_value` (String)
- `move_from_path` (String)
- `op` (String)
- `path` (String)

<a id="nestedobjatt--items--source_info_list--diff_list--extended_attributes"></a>
### Nested Schema for `items.source_info_list.diff_list.path`

Read-Only:

- `attribute_display_name` (String)
- `data_converter` (String)
- `path` (String)
- `type` (String)
