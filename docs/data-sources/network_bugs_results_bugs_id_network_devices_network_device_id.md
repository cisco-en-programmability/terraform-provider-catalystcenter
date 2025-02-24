---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_bugs_results_bugs_id_network_devices_network_device_id Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Compliance.
  Get network bug device for the bug by network device id
---

# catalystcenter_network_bugs_results_bugs_id_network_devices_network_device_id (Data Source)

It performs read operation on Compliance.

- Get network bug device for the bug by network device id

## Example Usage

```terraform
data "catalystcenter_network_bugs_results_bugs_id_network_devices_network_device_id" "example" {
  provider          = catalystcenter
  id                = "string"
  network_device_id = "string"
}

output "catalystcenter_network_bugs_results_bugs_id_network_devices_network_device_id_example" {
  value = data.catalystcenter_network_bugs_results_bugs_id_network_devices_network_device_id.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. Id of the network bug
- `network_device_id` (String) networkDeviceId path parameter. Id of the network device

### Read-Only

- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `response` (List of Object) (see [below for nested schema](#nestedobjatt--items--response))
- `version` (String)

<a id="nestedobjatt--items--response"></a>
### Nested Schema for `items.response`

Read-Only:

- `bug_count` (Number)
- `comments` (String)
- `last_scan_time` (Number)
- `network_device_id` (String)
- `scan_mode` (String)
- `scan_status` (String)
