---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sites_banner_settings Resource - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It manages read and update operations on Network Settings.
  Set banner settings for a site; null values indicate that the setting will be inherited from the parent site;
  empty objects ({}) indicate that the settings is unset.
---

# catalystcenter_sites_banner_settings (Resource)

It manages read and update operations on Network Settings.

- Set banner settings for a site; **null** values indicate that the setting will be inherited from the parent site;
empty objects (**{}**) indicate that the settings is unset.

## Example Usage

```terraform
resource "catalystcenter_sites_banner_settings" "example" {
  provider = catalystcenter
 
  parameters {

    banner {

      message = "string"
      type    = "string"
    }
    id = "string"
  }
}

output "catalystcenter_sites_banner_settings_example" {
  value = catalystcenter_sites_banner_settings.example
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `parameters` (Block List) (see [below for nested schema](#nestedblock--parameters))

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))
- `last_updated` (String)

<a id="nestedblock--parameters"></a>
### Nested Schema for `parameters`

Required:

- `id` (String) id path parameter. Site Id

Optional:

- `banner` (Block List) (see [below for nested schema](#nestedblock--parameters--banner))

<a id="nestedblock--parameters--banner"></a>
### Nested Schema for `parameters.banner`

Optional:

- `message` (String) Custom message that appears when logging into routers, switches, and hubs. Required for custom type.
- `type` (String) Type



<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `banner` (List of Object) (see [below for nested schema](#nestedobjatt--item--banner))

<a id="nestedobjatt--item--banner"></a>
### Nested Schema for `item.banner`

Read-Only:

- `inherited_site_id` (String)
- `inherited_site_name` (String)
- `message` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import catalystcenter_sites_banner_settings.example "id:=string"
```
