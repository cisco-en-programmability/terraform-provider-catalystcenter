---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sites_banner_settings Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Network Settings.
  Retrieve banner settings for a site; null values indicate that the setting will be inherited from the parent site;
  empty objects ({}) indicate that the setting is unset at a site.
---

# catalystcenter_sites_banner_settings (Data Source)

It performs read operation on Network Settings.

- Retrieve banner settings for a site; **null** values indicate that the setting will be inherited from the parent site;
empty objects (**{}**) indicate that the setting is unset at a site.

## Example Usage

```terraform
data "catalystcenter_sites_banner_settings" "example" {
  provider  = catalystcenter
  id        = "string"
  inherited = "false"
}

output "catalystcenter_sites_banner_settings_example" {
  value = data.catalystcenter_sites_banner_settings.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. Site Id

### Optional

- `inherited` (Boolean) _inherited query parameter. Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when **false**, **null** values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

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
