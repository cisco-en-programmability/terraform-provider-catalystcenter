---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_controllers_anchor_capable_devices Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  This data source allows the user to get Anchor capable devices
---

# catalystcenter_wireless_controllers_anchor_capable_devices (Data Source)

It performs read operation on Wireless.

- This data source allows the user to get Anchor capable devices

## Example Usage

```terraform
data "catalystcenter_wireless_controllers_anchor_capable_devices" "example" {
  provider = catalystcenter
}

output "catalystcenter_wireless_controllers_anchor_capable_devices_example" {
  value = data.catalystcenter_wireless_controllers_anchor_capable_devices.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `device_ip` (String)
- `device_name` (String)
- `wireless_mgmt_ip` (String)
