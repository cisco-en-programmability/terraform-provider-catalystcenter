---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_wireless_settings_interfaces Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Wireless.
  This data source allows the user to get all InterfacesThis data source allows the user to get an interface by ID
---

# catalystcenter_wireless_settings_interfaces (Data Source)

It performs read operation on Wireless.

- This data source allows the user to get all Interfaces

- This data source allows the user to get an interface by ID

## Example Usage

```terraform
data "catalystcenter_wireless_settings_interfaces" "example" {
  provider       = catalystcenter
  interface_name = "string"
  limit          = 1
  offset         = 1
  vlan_id        = 1.0
}

output "catalystcenter_wireless_settings_interfaces_example" {
  value = data.catalystcenter_wireless_settings_interfaces.example.items
}

data "catalystcenter_wireless_settings_interfaces" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_settings_interfaces_example" {
  value = data.catalystcenter_wireless_settings_interfaces.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) id path parameter. Interface ID
- `interface_name` (String) interfaceName query parameter. Interface Name
- `limit` (Number) limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
- `offset` (Number) offset query parameter. The first record to show for this page. The first record is numbered 1.
- `vlan_id` (Number) vlanId query parameter. Vlan Id

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `id` (String)
- `interface_name` (String)
- `vlan_id` (Number)


<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `id` (String)
- `interface_name` (String)
- `vlan_id` (Number)
