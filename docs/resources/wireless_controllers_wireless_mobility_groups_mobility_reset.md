---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_reset Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs create operation on Wireless.
  This data source action is used to reset wireless mobility which in turn sets mobility group name as 'default'.
---

# catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_reset (Resource)

It performs create operation on Wireless.

- This data source action is used to reset wireless mobility which in turn sets mobility group name as 'default'.


~>**Warning:**
This resource does not represent a real-world entity in Cisco Catalyst Center, therefore changing or deleting this resource on its own has no immediate effect.
Instead, it is a task part of a Cisco Catalyst Center workflow. It is executed in CatalystCenter without any additional verification. It does not check if it was executed before or if a similar configuration or action already existed previously.

## Example Usage

```terraform
resource "catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_reset" "example" {
  provider = catalystcenter
  parameters = [{

    network_device_id = "string"
  }]
}

output "catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_reset_example" {
  value = catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_reset.example
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

- `network_device_id` (String) Network device Id of Cisco wireless controller. Obtain the network device ID value by using the API call GET - /dna/intent/api/v1/network-device/ip-address/${ipAddress}.


<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `task_id` (String)
- `url` (String)
