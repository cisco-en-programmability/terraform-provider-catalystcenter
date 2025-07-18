---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_device_sync Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs update operation on Devices.
  Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If
  forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result
  can be seen in the child task of each device
---

# catalystcenter_network_device_sync (Resource)

It performs update operation on Devices.

- Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If
forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result
can be seen in the child task of each device


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_network_device_sync" "example" {
  provider   = catalystcenter
  force_sync = "false"
  parameters = ["string"]
}

output "catalystcenter_network_device_sync_example" {
  value = catalystcenter_network_device_sync.example
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

- `force_sync` (Boolean) forceSync query parameter.
- `payload` (List of String) Array of RequestDevicesSyncDevices


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
