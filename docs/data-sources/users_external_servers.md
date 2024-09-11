---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_users_external_servers Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on User and Roles.
  Get external users authentication servers.
---

# catalystcenter_users_external_servers (Data Source)

It performs read operation on User and Roles.

- Get external users authentication servers.

## Example Usage

```terraform
data "catalystcenter_users_external_servers" "example" {
  provider      = catalystcenter
  invoke_source = "string"
}

output "catalystcenter_users_external_servers_example" {
  value = data.catalystcenter_users_external_servers.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `invoke_source` (String) invokeSource query parameter. The source that invokes this API. The value of this query parameter must be set to "external".

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `aaa_servers` (List of Object) (see [below for nested schema](#nestedobjatt--item--aaa_servers))

<a id="nestedobjatt--item--aaa_servers"></a>
### Nested Schema for `item.aaa_servers`

Read-Only:

- `aaa_attribute` (String)
- `accounting_port` (Number)
- `authentication_port` (Number)
- `protocol` (String)
- `retries` (Number)
- `role` (String)
- `server_id` (String)
- `server_ip` (String)
- `shared_secret` (String)
- `socket_timeout` (Number)