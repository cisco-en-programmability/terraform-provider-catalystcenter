---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_devices_resync_interval_settings_override Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Devices.
  Overrides the global resync interval on all network devices. This essentially removes device specific intervals if
  set.
---

# catalystcenter_network_devices_resync_interval_settings_override (Resource)

It performs create operation on Devices.

- Overrides the global resync interval on all network devices. This essentially removes device specific intervals if
set.


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_network_devices_resync_interval_settings_override" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_devices_resync_interval_settings_override_example" {
  value = catalystcenter_network_devices_resync_interval_settings_override.example
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


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
