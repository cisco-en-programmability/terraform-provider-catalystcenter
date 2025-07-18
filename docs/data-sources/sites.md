---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sites Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on Site Design.
  Get sites.
---

# catalystcenter_sites (Data Source)

It performs read operation on Site Design.

- Get sites.

## Example Usage

```terraform
data "catalystcenter_sites" "example" {
  provider         = catalystcenter
  limit            = 1
  name             = "string"
  name_hierarchy   = "string"
  offset           = 1
  type             = "string"
  units_of_measure = "string"
}

output "catalystcenter_sites_example" {
  value = data.catalystcenter_sites.example.items
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `limit` (Number) limit query parameter. The number of records to show for this page;The minimum is 1, and the maximum is 500.
- `name` (String) name query parameter. Site name.
- `name_hierarchy` (String) nameHierarchy query parameter. Site name hierarchy.
- `offset` (Number) offset query parameter. The first record to show for this page; the first record is numbered 1.
- `type` (String) type query parameter. Site type.
- `units_of_measure` (String) _unitsOfMeasure query parameter. Floor units of measure

### Read-Only

- `id` (String) The ID of this resource.
- `items` (List of Object) (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `address` (String)
- `country` (String)
- `floor_number` (Number)
- `height` (Number)
- `id` (String)
- `latitude` (Number)
- `length` (Number)
- `longitude` (Number)
- `name` (String)
- `name_hierarchy` (String)
- `parent_id` (String)
- `rf_model` (String)
- `site_hierarchy_id` (String)
- `type` (String)
- `units_of_measure` (String)
- `width` (Number)
