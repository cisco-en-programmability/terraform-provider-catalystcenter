---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_images_addon_images_count Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Software Image Management (SWIM).
  Count of add-on images available for the given software image identifier, 'id' can be obtained from the response of
  API [ /dna/intent/api/v1/images?hasAddonImages=true ].
---

# catalystcenter_images_addon_images_count (Data Source)

It performs read operation on Software Image Management (SWIM).

- Count of add-on images available for the given software image identifier, 'id' can be obtained from the response of
API [ /dna/intent/api/v1/images?hasAddonImages=true ].

## Example Usage

```terraform
data "catalystcenter_images_addon_images_count" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_images_addon_images_count_example" {
  value = data.catalystcenter_images_addon_images_count.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) id path parameter. Software image identifier. Check API '/dna/intent/api/v1/images' for id from response.

### Read-Only

- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `count` (Number)