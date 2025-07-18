---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_device_credential Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Network Settings.
  API to get device credential details.
---

# catalystcenter_device_credential (Data Source)

It performs read operation on Network Settings.

- API to get device credential details.

## Example Usage

```terraform
data "catalystcenter_device_credential" "example" {
  provider = catalystcenter
  site_id  = "string"
}

output "catalystcenter_device_credential_example" {
  value = data.catalystcenter_device_credential.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `site_id` (String) siteId query parameter. Site id to retrieve the credential details associated with the site.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `cli` (List of Object) (see [below for nested schema](#nestedobjatt--item--cli))
- `http_read` (List of Object) (see [below for nested schema](#nestedobjatt--item--http_read))
- `http_write` (List of Object) (see [below for nested schema](#nestedobjatt--item--http_write))
- `snmp_v2_read` (List of Object) (see [below for nested schema](#nestedobjatt--item--snmp_v2_read))
- `snmp_v2_write` (List of Object) (see [below for nested schema](#nestedobjatt--item--snmp_v2_write))
- `snmp_v3` (List of Object) (see [below for nested schema](#nestedobjatt--item--snmp_v3))

<a id="nestedobjatt--item--cli"></a>
### Nested Schema for `item.cli`

Read-Only:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `enable_password` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `password` (String)
- `username` (String)


<a id="nestedobjatt--item--http_read"></a>
### Nested Schema for `item.http_read`

Read-Only:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `password` (String)
- `port` (String)
- `secure` (String)
- `username` (String)


<a id="nestedobjatt--item--http_write"></a>
### Nested Schema for `item.http_write`

Read-Only:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `password` (String)
- `port` (String)
- `secure` (String)
- `username` (String)


<a id="nestedobjatt--item--snmp_v2_read"></a>
### Nested Schema for `item.snmp_v2_read`

Read-Only:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `read_community` (String)


<a id="nestedobjatt--item--snmp_v2_write"></a>
### Nested Schema for `item.snmp_v2_write`

Read-Only:

- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `write_community` (String)


<a id="nestedobjatt--item--snmp_v3"></a>
### Nested Schema for `item.snmp_v3`

Read-Only:

- `auth_password` (String)
- `auth_type` (String)
- `comments` (String)
- `credential_type` (String)
- `description` (String)
- `id` (String)
- `instance_tenant_id` (String)
- `instance_uuid` (String)
- `privacy_password` (String)
- `privacy_type` (String)
- `snmp_mode` (String)
- `username` (String)
