---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_network_devices_delete_without_cleanup Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Devices.
  This data source action endpoint facilitates the deletion of a network device without performing configuration cleanup
  on the device. To delete a device via API, you must have permission to provision the network device. Although the API
  operation does not change the device configuration, removing a device without cleaning up its configuration could lead
  to a network behaviour that is not consistent with the configurations that are known to the system.
---

# catalystcenter_network_devices_delete_without_cleanup (Resource)

It performs create operation on Devices.

- This data source action endpoint facilitates the deletion of a network device without performing configuration cleanup
on the device. To delete a device via API, you must have permission to provision the network device. Although the API
operation does not change the device configuration, removing a device without cleaning up its configuration could lead
to a network behaviour that is not consistent with the configurations that are known to the system.


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_network_devices_delete_without_cleanup" "example" {
  provider = catalystcenter
  parameters = [{

    id = "string"
  }]
}

output "catalystcenter_network_devices_delete_without_cleanup_example" {
  value = catalystcenter_network_devices_delete_without_cleanup.example
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

- `id` (String) The unique identifier of the network device to be deleted


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
