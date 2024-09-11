---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_devices_resync_interval_settings Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs update operation on Devices.
  Updates the resync interval (in minutes) globally for devices which do not have custom resync interval. To override
  this setting for all network devices refer to [/networkDevices/resyncIntervalSettings/override]
---

# catalystcenter_network_devices_resync_interval_settings (Resource)

It performs update operation on Devices.

- Updates the resync interval (in minutes) globally for devices which do not have custom resync interval. To override
this setting for all network devices refer to [/networkDevices/resyncIntervalSettings/override]



~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_network_devices_resync_interval_settings" "example" {
  provider = catalystcenter
  parameters {

    interval = 1
  }
}

output "catalystcenter_network_devices_resync_interval_settings_example" {
  value = catalystcenter_network_devices_resync_interval_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `parameters` (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `interval` (Number) Resync Interval should be between 25 to 1440 minutes


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)