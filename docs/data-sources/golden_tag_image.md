---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_golden_tag_image Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Software Image Management (SWIM).
  Get golden tag status of an image. Set siteId as -1 for Global site.
---

# catalystcenter_golden_tag_image (Data Source)

It performs read operation on Software Image Management (SWIM).

- Get golden tag status of an image. Set siteId as -1 for Global site.

## Example Usage

```terraform
data "catalystcenter_golden_tag_image" "example" {
  provider                 = catalystcenter
  device_family_identifier = "string"
  device_role              = "string"
  image_id                 = "string"
  site_id                  = "string"
}

output "catalystcenter_golden_tag_image_example" {
  value = data.catalystcenter_golden_tag_image.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `device_family_identifier` (String) deviceFamilyIdentifier path parameter. Device family identifier e.g. : 277696480-283933147, e.g. : 277696480
- `device_role` (String) deviceRole path parameter. Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
- `image_id` (String) imageId path parameter. Image Id in uuid format.
- `site_id` (String) siteId path parameter. Site Id in uuid format. Set siteId as -1 for Global site.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `device_role` (String)
- `inherited_site_id` (String)
- `inherited_site_name` (String)
- `tagged_golden` (String)