---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_fabric_sites Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages create, read, update and delete operations on SDA.
  Adds a fabric site based on user input.Updates a fabric site based on user input.Deletes a fabric site based on id.
---

# catalystcenter_sda_fabric_sites (Resource)

It manages create, read, update and delete operations on SDA.

- Adds a fabric site based on user input.

- Updates a fabric site based on user input.

- Deletes a fabric site based on id.

## Example Usage

```terraform
resource "catalystcenter_sda_fabric_sites" "example" {
  provider = catalystcenter
 
  parameters {

    authentication_profile_name = "string"
    id                          = "string"
    is_pub_sub_enabled          = "false"
    site_id                     = "string"
  }
}

output "catalystcenter_sda_fabric_sites_example" {
  value = catalystcenter_sda_fabric_sites.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) Array of RequestSdaAddFabricSite (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Optional:

- `payload` (Block List) Array of RequestApplicationPolicyCreateApplication (see [below for nested schema](#nestedblock--parameters--payload))

<a id="nestedblock--parameters--payload"></a>
### Nested Schema for `parameters.payload`

Optional:

- `authentication_profile_name` (String) Authentication profile used for this fabric.
- `id` (String) ID of the fabric site (updating this field is not allowed).
- `is_pub_sub_enabled` (String) Specifies whether this fabric site will use pub/sub for control nodes.
- `site_id` (String) ID of the network hierarchy.



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `authentication_profile_name` (String)
- `id` (String)
- `is_pub_sub_enabled` (String)
- `site_id` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_sda_fabric_sites.example "id:=string"
```
