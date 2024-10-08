---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_users_external_servers_aaa_attribute Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on User and Roles.
  Get the current value of the custom AAA attribute.
---

# catalystcenter_users_external_servers_aaa_attribute (Data Source)

It performs read operation on User and Roles.

- Get the current value of the custom AAA attribute.

## Example Usage

```terraform
data "catalystcenter_users_external_servers_aaa_attribute" "example" {
  provider = catalystcenter
}

output "catalystcenter_users_external_servers_aaa_attribute_example" {
  value = data.catalystcenter_users_external_servers_aaa_attribute.example.item
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

- `aaa_attributes` (List of Object) (see [below for nested schema](#nestedobjatt--item--aaa_attributes))

<a id="nestedobjatt--item--aaa_attributes"></a>
### Nested Schema for `item.aaa_attributes`

Read-Only:

- `attribute_name` (String)
