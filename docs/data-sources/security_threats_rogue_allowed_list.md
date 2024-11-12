---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_security_threats_rogue_allowed_list Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Devices.
  Intent API to fetch all the allowed mac addresses in the system.
---

# catalystcenter_security_threats_rogue_allowed_list (Data Source)

It performs read operation on Devices.

- Intent API to fetch all the allowed mac addresses in the system.

## Example Usage

```terraform
data "catalystcenter_security_threats_rogue_allowed_list" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
}

output "catalystcenter_security_threats_rogue_allowed_list_example" {
  value = data.catalystcenter_security_threats_rogue_allowed_list.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `limit` (Number) limit query parameter. The maximum number of entries to return. If the value exceeds the total count, then the maximum entries will be returned.
- `offset` (Number) offset query parameter. The offset of the first item in the collection to return.

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `category` (Number)
- `last_modified` (Number)
- `mac_address` (String)