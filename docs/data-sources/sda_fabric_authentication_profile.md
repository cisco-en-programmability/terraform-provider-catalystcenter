---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_fabric_authentication_profile Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Get default authentication profile from SDA Fabric
---

# catalystcenter_sda_fabric_authentication_profile (Data Source)

It performs read operation on SDA.

- Get default authentication profile from SDA Fabric

## Example Usage

```terraform
data "catalystcenter_sda_fabric_authentication_profile" "example" {
  provider                   = catalystcenter
  authenticate_template_name = "string"
  site_name_hierarchy        = "string"
}

output "catalystcenter_sda_fabric_authentication_profile_example" {
  value = data.catalystcenter_sda_fabric_authentication_profile.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `site_name_hierarchy` (String) siteNameHierarchy query parameter.

### Optional

- `authenticate_template_name` (String) authenticateTemplateName query parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `authenticate_template_name` (String)
- `authentication_order` (String)
- `description` (String)
- `dot1x_to_mab_fallback_timeout` (String)
- `execution_id` (String)
- `number_of_hosts` (String)
- `site_name_hierarchy` (String)
- `status` (String)
- `wake_on_lan` (String)
