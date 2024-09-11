---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_device_config Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Configuration Archive.
  Returns the historical device configurations (running configuration , startup configuration , vlan if applicable) by
  specified criteria
---

# catalystcenter_network_device_config (Data Source)

It performs read operation on Configuration Archive.

- Returns the historical device configurations (running configuration , startup configuration , vlan if applicable) by
specified criteria

## Example Usage

```terraform
data "catalystcenter_network_device_config" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_device_config_example" {
  value = data.catalystcenter_network_device_config.example.items
}

data "catalystcenter_network_device_config" "example" {
  provider          = catalystcenter
  network_device_id = "string"
}

output "catalystcenter_network_device_config_example" {
  value = data.catalystcenter_network_device_config.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `created_by` (String) createdBy query parameter. Comma separated values for createdBy SCHEDULED, USER, CONFIG_CHANGE_EVENT, SCHEDULED_FIRST_TIME, DR_CALL_BACK, PRE_DEPLOY
- `created_time` (String) createdTime query parameter. Supported with logical filters GT,GTE,LT,LTE & BT : time in milliseconds (epoc format)
- `device_id` (String) deviceId query parameter. comma separated device id for example cf35b0a1-407f-412f-b2f4-f0c3156695f9,aaa38191-0c22-4158-befd-779a09d7cec1 . if device id is not provided it will fetch for all devices
- `file_type` (String) fileType query parameter. Config File Type can be RUNNINGCONFIG or STARTUPCONFIG
- `limit` (Number) limit query parameter.
- `offset` (Number) offset query parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `device_id` (String)
- `device_name` (String)
- `ip_address` (String)
- `versions` (List of Object) (see [below for nested schema](#nestedobjatt--items--versions))

<a id="nestedobjatt--items--versions"></a>
### Nested Schema for `items.versions`

Read-Only:

- `config_change_type` (String)
- `created_by` (String)
- `created_time` (Number)
- `files` (List of Object) (see [below for nested schema](#nestedobjatt--items--versions--files))
- `id` (String)
- `last_updated_time` (Number)
- `startup_running_status` (String)
- `syslog_config_event_dto` (List of Object) (see [below for nested schema](#nestedobjatt--items--versions--syslog_config_event_dto))
- `tags` (List of String)

<a id="nestedobjatt--items--versions--files"></a>
### Nested Schema for `items.versions.files`

Read-Only:

- `download_path` (String)
- `file_id` (String)
- `file_type` (String)


<a id="nestedobjatt--items--versions--syslog_config_event_dto"></a>
### Nested Schema for `items.versions.syslog_config_event_dto`

Read-Only:

- `config_method` (String)
- `device_uuid` (String)
- `login_ip_address` (String)
- `out_of_band` (String)
- `process_name` (String)
- `syslog_time` (Number)
- `terminal_name` (String)
- `user_name` (String)