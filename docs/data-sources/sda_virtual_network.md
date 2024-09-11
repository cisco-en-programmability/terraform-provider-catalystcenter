---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "catalystcenter_sda_virtual_network Data Source - terraform-provider-catalystcenter"
subcategory: ""
description: |-
  It performs read operation on SDA.
  Get virtual network (VN) from SDA Fabric
---

# catalystcenter_sda_virtual_network (Data Source)

It performs read operation on SDA.

- Get virtual network (VN) from SDA Fabric

## Example Usage

```terraform
data "catalystcenter_sda_virtual_network" "example" {
  provider             = catalystcenter
  site_name_hierarchy  = "string"
  virtual_network_name = "string"
}

output "catalystcenter_sda_virtual_network_example" {
  value = data.catalystcenter_sda_virtual_network.example.item
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `site_name_hierarchy` (String) siteNameHierarchy query parameter.
- `virtual_network_name` (String) virtualNetworkName query parameter.

### Read-Only

- `id` (String) The ID of this resource.
- `item` (List of Object) (see [below for nested schema](#nestedatt--item))

<a id="nestedatt--item"></a>
### Nested Schema for `item`

Read-Only:

- `description` (String)
- `execution_id` (String)
- `fabric_name` (String)
- `is_default_vn` (String)
- `is_infra_vn` (String)
- `site_name_hierarchy` (String)
- `status` (String)
- `virtual_network_context_id` (String)
- `virtual_network_id` (String)
- `virtual_network_name` (String)