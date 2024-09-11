---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_lan_automation_update_device Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs update operation on LAN Automation.
  Invoke this API to perform a DAY-N update on LAN Automation-related devices. Supported features include Loopback0 IP
  update, hostname update, link addition, and link deletion.
---

# catalystcenter_lan_automation_update_device (Resource)

It performs update operation on LAN Automation.

- Invoke this API to perform a DAY-N update on LAN Automation-related devices. Supported features include Loopback0 IP
update, hostname update, link addition, and link deletion.



~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in DNACenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_lan_automation_update_device" "example" {
  provider = catalystcenter
  feature  = "string"
  parameters {

    hostname_update_devices {

      device_management_ipaddress = "string"
      new_host_name               = "string"
    }
    link_update {

      destination_device_interface_name       = "string"
      destination_device_management_ipaddress = "string"
      ip_pool_name                            = "string"
      source_device_interface_name            = "string"
      source_device_management_ipaddress      = "string"
    }
    loopback_update_device_list {

      device_management_ipaddress = "string"
      new_loopback0_ipaddress     = "string"
    }
  }
}

output "catalystcenter_lan_automation_update_device_example" {
  value = catalystcenter_lan_automation_update_device.example
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

Required:

- `feature` (String) feature query parameter. Feature ID for the update. Supported feature IDs include: LOOPBACK0_IPADDRESS_UPDATE, HOSTNAME_UPDATE, LINK_Add and LINK_DELETE.

Optional:

- `hostname_update_devices` (Block List) (see [below for nested schema](#nestedblock--parameters--hostname_update_devices))
- `link_update` (Block List) (see [below for nested schema](#nestedblock--parameters--link_update))
- `loopback_update_device_list` (Block List) (see [below for nested schema](#nestedblock--parameters--loopback_update_device_list))

<a id="nestedblock--parameters--hostname_update_devices"></a>
### Nested Schema for `parameters.hostname_update_devices`

Optional:

- `device_management_ipaddress` (String) Device Management IP Address
- `new_host_name` (String) New hostname for the device


<a id="nestedblock--parameters--link_update"></a>
### Nested Schema for `parameters.link_update`

Optional:

- `destination_device_interface_name` (String) Destination Device Interface Name
- `destination_device_management_ipaddress` (String) Destination Device Management IP Address
- `ip_pool_name` (String) Name of the IP LAN Pool, required for Link Add should be from discovery site of source and destination device.
- `source_device_interface_name` (String) Source Device Interface Name
- `source_device_management_ipaddress` (String) Source Device Management IP Address


<a id="nestedblock--parameters--loopback_update_device_list"></a>
### Nested Schema for `parameters.loopback_update_device_list`

Optional:

- `device_management_ipaddress` (String) Device Management IP Address
- `new_loopback0_ipaddress` (String) New Loopback0 IP Address from LAN Pool of Device Discovery Site(Shared pool should not be used).



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)