---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_settings_anchor_groups Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  This data source allows the user to get AnchorGroups that captured in wireless settings design.
---

# catalystcenter_wireless_settings_anchor_groups (Data Source)

It performs read operation on Wireless.

- This data source allows the user to get AnchorGroups that captured in wireless settings design.

## Example Usage

```terraform
data "catalystcenter_wireless_settings_anchor_groups" "example" {
  provider = catalystcenter
  limit    = "string"
  offset   = "string"
}

output "catalystcenter_wireless_settings_anchor_groups_example" {
  value = data.catalystcenter_wireless_settings_anchor_groups.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `limit` (String) limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
- `offset` (String) offset query parameter. The first record to show for this page, the first record is numbered 1.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `anchor_group_name` (String)
- `id` (String)
- `mobility_anchors` (List of Object) (see [below for nested schema](#nestedobjatt--item--mobility_anchors))

<a id="nestedobjatt--item--mobility_anchors"></a>
### Nested Schema for `item.mobility_anchors`

Read-Only:

- `anchor_priority` (String)
- `device_name` (String)
- `ip_address` (String)
- `mac_address` (String)
- `managed_anchor_wlc` (String)
- `mobility_group_name` (String)
- `peer_device_type` (String)
- `private_ip` (String)
